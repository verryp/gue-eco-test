package helper

import (
	"fmt"
)

func CacheKeyTokenBlacklisted(jwtID string) string {
	return fmt.Sprintf("jwt-id-blacklisted:%s", jwtID)
}
