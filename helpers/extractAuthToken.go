package helpers

import "strings"

func ExtractBearerToken(header string) string {
	token := strings.TrimPrefix(header, "Bearer ")
	return token
}