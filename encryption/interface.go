package encryption

import (
	"crypto/rand"
	"io"
)

const (
	BlockBIt_128 = 128
	BlockBIt_192 = 192
	BlockBIt_256 = 256
	InitVector_16 = 16
)

func MakeIV(len int)[]byte  {
	iv := make([]byte,len)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	return iv
}
