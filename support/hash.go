package support

import (
	"crypto/md5"
)
func Hash(key, s string) string{
	h := md5.New()
	h.Write([]byte(key))
	h.Write([]byte(s))
	v := h.Sum(nil)
	return string(v[:])
}