package gobuy

import (
	"encoding/base64"
)

func formatBasicAuthorization(token string) string {
	tokenData := []byte(token)
	return "Basic " + base64.StdEncoding.EncodeToString(tokenData)
}
