package mi_test

import (
	"testing"

	"github.com/freemiumvpn/fpn-openvpn-server/internal/openvpn/mi"
	"github.com/stretchr/testify/assert"
)

func Test_ParseLine(t *testing.T) {
	input := []byte("foo")
	output := mi.ParseLine(input)

	assert.Equal(t, output, "foo")
}

func Test_ParseLines(t *testing.T) {
	input := make([][]byte, 0)
	input = append(input, []byte("foo"))
	input = append(input, []byte("baz"))

	output := mi.ParseLines(input)

	assert.Equal(t, output, "foo\nbaz\n")
}
