package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

type md5Hash struct{}

func NewMd5Hasher() *md5Hash {
	return &md5Hash{}
}

func (m *md5Hash) Hash(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
