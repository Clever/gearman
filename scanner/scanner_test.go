package scanner

import (
	"bytes"
	"github.com/Clever/gearman/packet"
	"github.com/stretchr/testify/assert"
	"testing"
)

func packetWithArgs(args [][]byte) *packet.Packet {
	return &packet.Packet{
		Code:      []byte{0, 1, 2, 3},
		Type:      1,
		Arguments: args,
	}
}

func TestScanner(t *testing.T) {
	arg := []byte("arg")
	packetNoArgs := packetWithArgs([][]byte{})
	packetOneArg := packetWithArgs([][]byte{arg})
	packetMultArgs := packetWithArgs([][]byte{arg, arg, arg})
	tmp := []byte{}
	noArgB, err := packetNoArgs.Bytes()
	assert.Nil(t, err, nil)
	tmp = append(tmp, noArgB...)
	oneArgB, err := packetOneArg.Bytes()
	assert.Nil(t, err, nil)
	tmp = append(tmp, oneArgB...)
	multArgB, err := packetMultArgs.Bytes()
	assert.Nil(t, err, nil)
	tmp = append(tmp, multArgB...)
	buf := bytes.NewBuffer(tmp)

	scanner := New(buf)
	assert.Equal(t, scanner.Scan(), true)
	assert.Equal(t, scanner.Bytes(), noArgB)
	assert.Equal(t, scanner.Scan(), true)
	assert.Equal(t, scanner.Bytes(), oneArgB)
	assert.Equal(t, scanner.Scan(), true)
	assert.Equal(t, scanner.Bytes(), multArgB)
}
