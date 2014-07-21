package packet

import (
	"bytes"
	"encoding/binary"
)

// Packet contains a Gearman packet. See http://gearman.org/protocol/
type Packet struct {
	// The Code for the packet: either \0REQ or \0RES
	Code []byte
	// The Type of the packet, e.g. WorkStatus
	Type int
	// The Arguments of the packet
	Arguments [][]byte
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface
func (packet *Packet) UnmarshalBinary(data []byte) error {
	packetType := int32(0)
	if err := binary.Read(bytes.NewBuffer(data[4:8]), binary.BigEndian, &packetType); err != nil {
		return err
	}
	packet.Type = int(packetType)
	arguments := [][]byte{}
	if len(data) > 12 {
		arguments = bytes.Split(data[12:len(data)], []byte{0})
	}
	packet.Arguments = arguments

	code := make([]byte, 4)
	copy(code, data[0:4])
	packet.Code = code

	return nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface
func (packet *Packet) MarshalBinary() ([]byte, error) {
	buf := bytes.NewBuffer(packet.Code)
	if err := binary.Write(buf, binary.BigEndian, int32(packet.Type)); err != nil {
		return nil, err
	}
	size := len(packet.Arguments) - 1 // One for each null-byte separator
	for _, argument := range packet.Arguments {
		size += len(argument)
	}
	if size < 0 {
		size = 0
	}
	if err := binary.Write(buf, binary.BigEndian, int32(size)); err != nil {
		return nil, err
	}
	if len(packet.Arguments) > 0 {
		// Need special handling for last argument (don't write null byte)
		for _, argument := range packet.Arguments[0 : len(packet.Arguments)-1] {
			buf.Write(argument)
			buf.Write([]byte{0})
		}
		buf.Write(packet.Arguments[len(packet.Arguments)-1])
	}
	return buf.Bytes(), nil
}

// Handle assumes that the first argument of the packet is the job handle, returns it as a string
func (packet *Packet) Handle() string {
	return string(packet.Arguments[0])
}
