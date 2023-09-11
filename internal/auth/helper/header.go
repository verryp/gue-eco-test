package helper

import (
	"strings"
)

func SplitHeaderBearerToken(token string) string {
	apiKey := strings.Split(token, " ")
	return apiKey[1]
}
