package handshake

const (
	ResponseEOF = 0xfe
	ResponseOK  = 0x00
	ResponseERR = 0xff
)

// Capability flags
const (
	CapabilityServerMySQL uint32 = 1 << iota
	CapabilityServerPostgres
)

// Supported auth forms.
const (
	// CFAuthToken transmits the cf auth token in the clear.
	CFAuthToken = "cf_auth_token"
)

const (
	// MaxPacketSize is the maximum payload length of a packet
	// the server supports.
	MaxPacketSize = (1 << 24) - 1

	// protocolVersion is the current version of the protocol.
	protocolVersion = 1
)
const (
	serverVersion = "1"
)
