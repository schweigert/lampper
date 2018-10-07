package bytes

import "errors"

func FromUint32(v uint32) []byte {
	b := make([]byte, 4)
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	return b
}

func ToUint32(b []byte) (uint32, error) {
	if len(b) < 4 {
		return 0, errors.New("buffer size is less than 4")
	}
	v := uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
	return v, nil
}
