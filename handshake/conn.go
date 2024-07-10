package handshake

import (
	"fmt"
	"io"
	"strings"
)

// writePacket writes a packet, for this protocol we don't worry about breaking packets up into smaller chunks
func (c *Conn) writePacket(data []byte) error {
	index := 0
	packetLength := len(data)

	// Compute and write the header.
	var header [4]byte
	header[0] = byte(packetLength)
	header[1] = byte(packetLength >> 8)
	header[2] = byte(packetLength >> 16)
	header[3] = c.sequence
	if n, err := c.netConn.Write(header[:]); err != nil {
		return fmt.Errorf("Write(header) failed: %w", err)
	} else if n != 4 {
		return fmt.Errorf("Write(header) returned a short write: %v < 4", n)
	}

	// Write the body.
	if n, err := c.netConn.Write(data[index : index+packetLength]); err != nil {
		return fmt.Errorf("Write(packet) failed: %w", err)
	} else if n != packetLength {
		return fmt.Errorf("Write(packet) returned a short write: %v < 4", n)
	}

	// Update our state.
	c.sequence++

	return nil
}

func (c *Conn) readHeader() (int, error) {
	var header [4]byte
	if _, err := io.ReadFull(c.netConn, header[:]); err != nil {
		if err == io.EOF {
			return 0, err
		}
		if strings.HasSuffix(err.Error(), "read: connection reset by peer") {
			return 0, io.EOF
		}
		return 0, fmt.Errorf("io.ReadFull(header size) failed: %w", err)
	}

	sequence := uint8(header[3])
	if sequence != c.sequence {
		return 0, fmt.Errorf("invalid sequence, expected %v got %v", c.sequence, sequence)
	}

	c.sequence++

	return int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16), nil
}

func (c *Conn) writeErrorPacketFromError(err error) error {
	if se, ok := err.(*ServerError); ok {
		return c.writeErrorPacket(uint16(se.Num), "%v", se.Message)
	}
	return c.writeErrorPacket(ErrUnknownError, "unknown error: %v", err)
}

// writeErrorPacket writes an error packet.
// Server -> Client.
func (c *Conn) writeErrorPacket(errorCode uint16, format string, args ...interface{}) error {
	errorMessage := fmt.Sprintf(format, args...)
	length := 1 + // message type
		2 + // error code
		len(errorMessage) // message
	data := make([]byte, length)
	pos := 0
	pos = writeByte(data, pos, ResponseERR)
	pos = writeUint16(data, pos, errorCode)

	_ = writeEOFString(data, pos, errorMessage)

	return c.writePacket(data)
}

// readOnePacket reads a single packet into a newly allocated buffer.
func (c *Conn) readOnePacket() ([]byte, error) {
	length, err := c.readHeader()
	if err != nil {
		return nil, err
	}
	if length > MaxPacketSize {
		// for simplicity we don't support data spread over multiple packets, and we shouldn't need to
		return nil, fmt.Errorf("reading packets larger than max packet size is not supported")
	}
	if length == 0 {
		// This can be caused by the packet after a packet of
		// exactly size MaxPacketSize.
		return nil, nil
	}

	data := make([]byte, length)
	if _, err := io.ReadFull(c.netConn, data); err != nil {
		return nil, fmt.Errorf("io.ReadFull(packet body of length %v) failed: %w", length, err)
	}
	return data, nil
}
