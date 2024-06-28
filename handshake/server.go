package handshake

import (
	"fmt"
	"net"

	accessv1alpha1 "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	_ "github.com/go-sql-driver/mysql"

	"go.commonfate.io/entities/cf/grantoutput"
)

type Server struct {
	conn           *Conn
	tokenValidator TokenValidator
	grantValidator GrantValidator
}

type TokenValidator interface {
	Validate(token string) (*accessv1alpha1.GetCallerIdentityResponse, error)
}

type GrantValidator interface {
	Validate(principal *entityv1alpha1.EID, grantID string) (*accessv1alpha1.Grant, *grantoutput.AWSRDS, error)
}

func NewHandshakeServer(conn net.Conn, connectionID uint32, tokenValidator TokenValidator, grantValidator GrantValidator) *Server {
	return &Server{
		tokenValidator: tokenValidator,
		grantValidator: grantValidator,
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

type HandshakeResponse struct {
	CallerIdentity *accessv1alpha1.GetCallerIdentityResponse
	Grant          *accessv1alpha1.Grant
	GrantOutput    *grantoutput.AWSRDS
}

func (s *Server) Handshake() (*HandshakeResponse, error) {
	err := s.conn.writeHandshake(serverVersion)
	if err != nil {
		return nil, err
	}
	data, err := s.conn.readOnePacket()
	if err != nil {
		return nil, err
	}

	handshakeResponse, err := parseClientHandshakePacket(data)
	if err != nil {
		return nil, err
	}

	// validate the token
	callerID, err := s.tokenValidator.Validate(handshakeResponse.token)
	if err != nil {
		return nil, s.conn.writeErrorPacketFromError(err)
	}

	// validate the grant
	grant, grantOutput, err := s.grantValidator.Validate(callerID.Principal.Eid, handshakeResponse.grantID)
	if err != nil {
		return nil, s.conn.writeErrorPacketFromError(err)
	}

	// Negotiation worked, send OK packet.
	err = s.conn.writePacket([]byte{ResponseOK})
	if err != nil {
		return nil, err
	}

	return &HandshakeResponse{
		CallerIdentity: callerID,
		Grant:          grant,
		GrantOutput:    grantOutput,
	}, nil
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

type handshakeResponse struct {
	token   string
	grantID string
}

func parseClientHandshakePacket(data []byte) (*handshakeResponse, error) {
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

	return &handshakeResponse{
		token:   token,
		grantID: grantID,
	}, nil
}
