package packet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var confs = []map[string]interface{}{
	{
		"pack":  Packet{Code: []byte{0, 1, 2, 3}, Type: 3, Arguments: [][]byte{{4}, {5}, {6}}},
		"bytes": []byte{0, 1, 2, 3, 0, 0, 0, 3, 0, 0, 0, 5, 4, 0, 5, 0, 6},
	},
	{
		"pack":  Packet{Code: []byte{0, 1, 2, 3}, Type: 3, Arguments: [][]byte{}},
		"bytes": []byte{0, 1, 2, 3, 0, 0, 0, 3, 0, 0, 0, 0},
	},
}

func TestBytes(t *testing.T) {
	for _, conf := range confs {
		pack := conf["pack"].(Packet)
		bytes, err := pack.Bytes()
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, bytes, conf["bytes"])
	}
}

func TestConstructor(t *testing.T) {
	for _, conf := range confs {
		pack, err := New(conf["bytes"].([]byte))
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, *pack, conf["pack"])
	}
}

func TestHandle(t *testing.T) {
	expected := "The Handle"
	pack := Packet{Arguments: [][]byte{[]byte(expected)}}
	assert.Equal(t, pack.Handle(), expected)
}
