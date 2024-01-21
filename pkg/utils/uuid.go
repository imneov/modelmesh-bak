package utils

import (
	"crypto/rand"
	"fmt"
)

const (
	defaultRequestPrefix = "req-"
)

// ReqUUID ReqID returns a request id.
func ReqUUID() string {
	return UUID(defaultRequestPrefix)
}

// UUID generate an uuid.
func UUID(prefix string) string {
	uuid := make([]byte, 16)
	if _, err := rand.Read(uuid); err != nil {
		return ""
	}
	// see section 4.1.1.
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// see section 4.1.3.
	uuid[6] = uuid[6]&^0xf0 | 0x40
	uid := fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])

	prefixLen := len(prefix)
	if prefixLen > 0 && prefixLen < 15 {
		uid = prefix + uid[prefixLen:]
	}
	return uid
}
