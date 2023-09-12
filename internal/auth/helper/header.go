package helper

import (
	"strings"
)

func SplitHeaderBearerToken(token string) string {
	apiKey := strings.Split(token, " ")
	return apiKey[1]
}

func TrimHeaderIPAddress(xForwardedFor string) string {
	ipAddress := strings.Split(xForwardedFor, ",")
	userIPAddress := ipAddress[len(ipAddress)-1]

	return strings.TrimSpace(userIPAddress)
}
