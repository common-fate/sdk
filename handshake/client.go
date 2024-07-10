package handshake

import (
	"fmt"
	"net"

	"github.com/common-fate/clio"
	"golang.org/x/oauth2"
)

type Client struct {
	conn        *Conn
	grantID     string
	tokenSource oauth2.TokenSource
}

func NewHandshakeClient(conn net.Conn, grantID string, tokenSource oauth2.TokenSource) *Client {
	return &Client{
		grantID: grantID,
		conn: &Conn{
			netConn: conn,
		},
		tokenSource: tokenSource,
	}
}

type ServerHandshake struct {
	ConnectionID       uint32
	ServerCapabilities uint32
	AuthMethod         string
	ServerVersion      string
	ProtocolVersion    int
}

func (c *Client) Handshake() (*ServerHandshake, error) {
	clio.Debugw("beginning handshake for grant", "grantId", c.grantID)
	data, err := c.conn.readOnePacket()
	if err != nil {
		return nil, err
	}
	clio.Debugw("server handshake initial packet", "packet", data)

	handshake, err := c.parseInitialHandshakePacket(data)
	if err != nil {
		return nil, err
	}

	clio.Debugw("parsed initial handshake packet", "connectionId", c.conn.connectionID)

	if handshake.AuthMethod != CFAuthToken {
		return nil, fmt.Errorf("only %s auth method is supported", CFAuthToken)
	}

	token, err := c.tokenSource.Token()
	if err != nil {
		return nil, err
	}

	clio.Debug("writing handshake response")
	// send response
	err = c.writeClientHandshakeResponse(token.AccessToken, c.grantID)
	if err != nil {
		return nil, err
	}
	// read ok
	data, err = c.conn.readOnePacket()
	if err != nil {
		return nil, err
	}
	clio.Debug("server handshake response packet", "packet", data)
	switch data[0] {
	case ResponseOK:
		clio.Debug("handshake successful")
		return handshake, nil
	case ResponseERR:
		return nil, parseErrorPacket(data)
	default:
		return nil, fmt.Errorf("handshake failed: unexpected response from server")
	}
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
func (c *Client) parseInitialHandshakePacket(data []byte) (*ServerHandshake, error) {
	pos := 0

	// Protocol version.
	pver, pos, ok := readByte(data, pos)
	if !ok {
		return nil, fmt.Errorf("parseInitialHandshakePacket: packet has no protocol version")
	}

	// Server is allowed to immediately send ERR packet
	if pver == ResponseERR {
		errorCode, pos, _ := readUint16(data, pos)
		// Normally there would be a 1-byte sql_state_marker field and a 5-byte
		// sql_state field here, but docs say these will not be present in this case.
		errorMsg, _, _ := readEOFString(data, pos)
		return nil, fmt.Errorf("immediate error from server errorCode=%v errorMsg=%v", errorCode, errorMsg)
	}

	if pver != protocolVersion {
		return nil, fmt.Errorf("bad protocol version: %v", pver)
	}

	// Read the server version.
	serverVersion, pos, ok := readNullString(data, pos)
	if !ok {
		return nil, fmt.Errorf("parseInitialHandshakePacket: packet has no server version")
	}

	// Read the connection id.
	c.conn.connectionID, pos, ok = readUint32(data, pos)
	if !ok {
		return nil, fmt.Errorf("parseInitialHandshakePacket: packet has no connection id")
	}

	// Lower 2 bytes of the capability flags.
	capabilities, pos, ok := readUint32(data, pos)
	if !ok {
		return nil, fmt.Errorf("parseInitialHandshakePacket: packet has no capability flags ")
	}

	authMethod, _, ok := readNullString(data, pos)
	if !ok {
		return nil, fmt.Errorf("parseInitialHandshakePacket: packet has no auth method")
	}
	return &ServerHandshake{
		ConnectionID:       c.conn.connectionID,
		ServerCapabilities: capabilities,
		AuthMethod:         authMethod,
		ServerVersion:      serverVersion,
		ProtocolVersion:    int(pver),
	}, nil
}
