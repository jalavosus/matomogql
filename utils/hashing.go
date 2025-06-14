package utils

import (
	"crypto/sha256"
)

type BytesLike interface{ string | []byte }

func Sha256Sum[T BytesLike](data T) []byte {
	hasher := sha256.New()
	hasher.Write([]byte(data))

	return hasher.Sum(nil)
}
