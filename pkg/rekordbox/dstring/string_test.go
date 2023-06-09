package dstring_test

import (
	`testing`

	`github.com/ambientsound/rex/pkg/rekordbox/dstring`
	"github.com/stretchr/testify/assert"
)

func TestShortAsciiString_EmptyString_MarshalBinary(t *testing.T) {
	s := dstring.New("")

	data, err := s.MarshalBinary()

	assert.NoError(t, err)
	assert.Equal(t, []byte{0x03}, data)
}

func TestShortAsciiString_MarshalBinary(t *testing.T) {
	s := dstring.New("/PIONEER/USBANLZ/P05B/0001069F/ANLZ0000.DAT")

	data, err := s.MarshalBinary()

	assert.NoError(t, err)
	assert.Equal(t, []byte{
		0x59, 0x2f, 0x50, 0x49, 0x4f, 0x4e, 0x45, 0x45, 0x52, 0x2f, 0x55, 0x53,
		0x42, 0x41, 0x4e, 0x4c, 0x5a, 0x2f, 0x50, 0x30, 0x35, 0x42, 0x2f, 0x30,
		0x30, 0x30, 0x31, 0x30, 0x36, 0x39, 0x46, 0x2f, 0x41, 0x4e, 0x4c, 0x5a,
		0x30, 0x30, 0x30, 0x30, 0x2e, 0x44, 0x41, 0x54,
	}, data)
}

// FIXME: not verified with actual production data!
func TestLongAsciiString_MarshalBinary(t *testing.T) {
	s := dstring.LongAsciiString("/PIONEER/USBANLZ/P05B/0001069F/ANLZ0000.DAT")

	data, err := s.MarshalBinary()

	assert.NoError(t, err)
	assert.Equal(t, []byte{
		0x40, 0x2f, 0x00, 0x00,
		0x2f, 0x50, 0x49, 0x4f, 0x4e, 0x45, 0x45, 0x52, 0x2f, 0x55, 0x53,
		0x42, 0x41, 0x4e, 0x4c, 0x5a, 0x2f, 0x50, 0x30, 0x35, 0x42, 0x2f, 0x30,
		0x30, 0x30, 0x31, 0x30, 0x36, 0x39, 0x46, 0x2f, 0x41, 0x4e, 0x4c, 0x5a,
		0x30, 0x30, 0x30, 0x30, 0x2e, 0x44, 0x41, 0x54,
	}, data)
}

func TestUnicodeString_MarshalBinary(t *testing.T) {
	s := dstring.New("Rødhåd feat. Vril")
	data, err := s.MarshalBinary()

	assert.NoError(t, err)
	assert.Equal(t, []byte{
		0x90, 0x26, 0x00, 0x00, 0x52, 0x00, 0xf8, 0x00, 0x64, 0x00, 0x68, 0x00,
		0xe5, 0x00, 0x64, 0x00, 0x20, 0x00, 0x66, 0x00, 0x65, 0x00, 0x61, 0x00,
		0x74, 0x00, 0x2e, 0x00, 0x20, 0x00, 0x56, 0x00, 0x72, 0x00, 0x69, 0x00,
		0x6c, 0x00,
	}, data)
}

func TestIsrcString_MarshalBinary(t *testing.T) {
	s := dstring.IsrcString("GBJX38209003")
	data, err := s.MarshalBinary()

	assert.NoError(t, err)
	assert.Equal(t, []byte{
		0x90, 0x12, 0x00, 0x00, 0x03, 0x47, 0x42, 0x4a, 0x58, 0x33, 0x38, 0x32,
		0x30, 0x39, 0x30, 0x30, 0x33, 0x00,
	}, data)
}
