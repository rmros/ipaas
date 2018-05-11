package uuid

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/nu7hatch/gouuid"
)

func md5_32(Seed ...string) string {
	var buf []byte
	if len(Seed) > 0 {
		buf = []byte(Seed[0])
	} else {
		u4, _ := uuid.NewV4()
		buf = u4[:]
	}
	m5 := md5.New()
	m5.Write(buf)
	m5str := hex.EncodeToString(m5.Sum(nil))
	return m5str
}

func md5_16(Seed ...string) string {
	m5str := md5_32(Seed...)
	return m5str[8:24]
}

// UU return the 16 len string
func UU(Seed ...string) string {
	return md5_16(Seed...)
}
