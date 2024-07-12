package handshake

import (
	"fmt"
	"net"
)

type Server struct {
	conn             *Conn
	sessionValidator SessionValidator
}

type ValidateResult struct {
	SessionID string
}
type SessionValidator interface {
	Validate(handshake ClientHandshake) (*ValidateResult, error)
}

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as SessionValidators
type SessionValidatorFunc func(handshake ClientHandshake) (*ValidateResult, error)

func (f SessionValidatorFunc) Validate(handshake ClientHandshake) (*ValidateResult, error) {
	return f(handshake)
}

func NewHandshakeServer(conn net.Conn, connectionID uint32, sessionValidator SessionValidator) *Server {
	return &Server{
		sessionValidator: sessionValidator,
		conn: &Conn{
			connectionID: connectionID,
			netConn:      conn,
		},
	}
}

type Conn struct {
	sequence     uint8
	connectionID uint32
	netConn      net.Conn
}

func (s *Server) Handshake() error {
	err := s.conn.writeHandshake(serverVersion)
	if err != nil {
		return err
	}
	data, err := s.conn.readOnePacket()
	if err != nil {
		return err
	}

	handshakeResponse, err := parseClientHandshakePacket(data)
	if err != nil {
		return err
	}

	// validate the token
	result, err := s.sessionValidator.Validate(*handshakeResponse)
	if err != nil {
		wErr := s.conn.writeErrorPacketFromError(err)
		if wErr != nil {
			return wErr
		}
		return fmt.Errorf("session validation failed: %w", err)
	}

	// Negotiation worked, send OK packet with the session ID.
	err = s.conn.writeOKPacket(result.SessionID)
	if err != nil {
		return err
	}

	return nil
}

func (c *Conn) writeOKPacket(sessionID string) error {
	length := 1 + // message type
		len(sessionID) // sessionID
	data := make([]byte, length)
	pos := 0
	pos = writeByte(data, pos, ResponseOK)
	pos = writeEOFString(data, pos, sessionID)
	// Sanity check.
	if pos != len(data) {
		return fmt.Errorf("error building Handshake packet: got %v bytes expected %v", pos, len(data))
	}

	return c.writePacket(data)
}

type ClientHandshake struct {
	Token   string
	GrantID string
	// Optional Session ID
	SessionID string
}

// writeHandshake writes the Initial Handshake Packet, server side
func (c *Conn) writeHandshake(serverVersion string) error {
	capabilities := CapabilityServerMySQL |
		CapabilityServerPostgres

	length :=
		1 + // protocol version
			lenNullString(serverVersion) +
			4 + // connection ID
			4 + // capability flags
			lenNullString(CFAuthToken) // auth-plugin-name

	data := make([]byte, length)
	pos := 0

	// Protocol version.
	pos = writeByte(data, pos, protocolVersion)

	// Copy server version.
	pos = writeNullString(data, pos, serverVersion)

	// Add connectionID in.
	pos = writeUint32(data, pos, c.connectionID)

	// Lower part of the capability flags.
	pos = writeUint32(data, pos, capabilities)

	// Copy authPluginName. We always start with cf_auth_token.
	pos = writeNullString(data, pos, CFAuthToken)

	// Sanity check.
	if pos != len(data) {
		return fmt.Errorf("error building Handshake packet: got %v bytes expected %v", pos, len(data))
	}

	return c.writePacket(data)
}

func parseClientHandshakePacket(data []byte) (*ClientHandshake, error) {
	pos := 0

	// Client flags, 4 bytes. not currently used for anything
	_, pos, ok := readUint32(data, pos)
	if !ok {
		return nil, fmt.Errorf("parseClientHandshakePacket: can't read client flags")
	}

	// token
	token, pos, ok := readNullString(data, pos)
	if !ok {
		return nil, fmt.Errorf("parseClientHandshakePacket: can't read token")
	}
	grantID, _, ok := readNullString(data, pos)
	if !ok {
		return nil, fmt.Errorf("parseClientHandshakePacket: can't read grantID")
	}
	sessionID, _, ok := readNullString(data, pos)
	if !ok {
		return nil, fmt.Errorf("parseClientHandshakePacket: can't read sessionID")
	}

	return &ClientHandshake{
		Token:     token,
		GrantID:   grantID,
		SessionID: sessionID,
	}, nil
}
