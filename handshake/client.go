package handshake

import (
	"context"
	"fmt"
	"net"

	"github.com/common-fate/clio"
	"github.com/common-fate/sdk/config"
)

type Client struct {
	conn    *Conn
	grantID string
}

func NewHandshakeClient(conn net.Conn, grantID string) *Client {
	return &Client{
		grantID: grantID,
		conn: &Conn{
			netConn: conn,
		},
	}
}

func (c *Client) Handshake() error {
	data, err := c.conn.readOnePacket()
	if err != nil {
		return err
	}

	capabilities, authMethod, err := c.parseInitialHandshakePacket(data)
	if err != nil {
		return err
	}
	if capabilities&CapabilityServerMySQL > 0 {
		clio.Info("server supports mysql database connections")
	}
	if capabilities&CapabilityServerPostgres > 0 {
		clio.Info("server supports postgres database connections")
	}

	if authMethod != CFAuthToken {
		return fmt.Errorf("only cf auth token is supported")
	}

	// get the auth token
	cfg, err := config.LoadDefault(context.Background())
	if err != nil {
		return err
	}

	tok, err := cfg.TokenStore.Token()
	if err != nil {
		return err
	}

	// send response
	err = c.writeClientHandshakeResponse(tok.AccessToken, c.grantID)
	if err != nil {
		return err
	}
	// read ok
	data, err = c.conn.readOnePacket()
	if err != nil {
		return err
	}
	switch data[0] {
	case ResponseOK:
		clio.Success("handshake success")
		return nil
	case ResponseERR:
		return parseErrorPacket(data)
	}
	return nil
}

// ParseErrorPacket parses the error packet and returns a SQLError.
func parseErrorPacket(data []byte) error {
	// We already read the type.
	pos := 1

	// Error code is 2 bytes.
	code, pos, ok := readUint16(data, pos)
	if !ok {
		return fmt.Errorf("invalid error packet code: %v", data)
	}

	// Human readable error message is the rest.
	msg := string(data[pos:])
	return fmt.Errorf("error code: %v, msg: %s", code, msg)
}

// writeHandshake writes the Initial Handshake Packet, server side
func (c *Client) writeClientHandshakeResponse(authToken string, grantID string) error {
	var capabilities uint32 = 0 // not using this for anything yet

	length :=
		4 + // capability flags
			lenNullString(authToken) + // auth token
			lenNullString(grantID) // grant id

	data := make([]byte, length)
	pos := 0

	// The capability flags.
	pos = writeUint32(data, pos, capabilities)

	// Copy the auth token
	pos = writeNullString(data, pos, authToken)

	// Copy the grantID
	pos = writeNullString(data, pos, grantID)

	// Sanity check.
	if pos != len(data) {
		return fmt.Errorf("error building Client Handshake packet: got %v bytes expected %v", pos, len(data))
	}

	return c.conn.writePacket(data)
}

// parseInitialHandshakePacket parses the initial handshake from the server.
func (c *Client) parseInitialHandshakePacket(data []byte) (uint32, string, error) {
	pos := 0

	// Protocol version.
	pver, pos, ok := readByte(data, pos)
	if !ok {
		return 0, "", fmt.Errorf("parseInitialHandshakePacket: packet has no protocol version")
	}

	// Server is allowed to immediately send ERR packet
	if pver == ResponseERR {
		errorCode, pos, _ := readUint16(data, pos)
		// Normally there would be a 1-byte sql_state_marker field and a 5-byte
		// sql_state field here, but docs say these will not be present in this case.
		errorMsg, _, _ := readEOFString(data, pos)
		return 0, "", fmt.Errorf("immediate error from server errorCode=%v errorMsg=%v", errorCode, errorMsg)
	}

	if pver != protocolVersion {
		return 0, "", fmt.Errorf("bad protocol version: %v", pver)
	}

	// Read the server version.
	serverVersion, pos, ok := readNullString(data, pos)
	if !ok {
		return 0, "", fmt.Errorf("parseInitialHandshakePacket: packet has no server version")
	}
	clio.Infof("server version %s", serverVersion)

	// Read the connection id.
	c.conn.connectionID, pos, ok = readUint32(data, pos)
	if !ok {
		return 0, "", fmt.Errorf("parseInitialHandshakePacket: packet has no connection id")
	}

	// Lower 2 bytes of the capability flags.
	capabilities, pos, ok := readUint32(data, pos)
	if !ok {
		return 0, "", fmt.Errorf("parseInitialHandshakePacket: packet has no capability flags ")
	}

	authMethod, _, ok := readNullString(data, pos)
	if !ok {
		return 0, "", fmt.Errorf("parseInitialHandshakePacket: packet has no auth method")
	}
	return capabilities, authMethod, nil
}
