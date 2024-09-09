package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

type BytesLike interface{ string | []byte }

func Sha256Sum[T BytesLike](data T) []byte {
	hasher := sha256.New()
	hasher.Write([]byte(data))

	return hasher.Sum(nil)
}

func Sha256SumHex[T BytesLike](data T) string {
	return hex.EncodeToString(Sha256Sum(data))
}
