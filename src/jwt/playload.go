package jwt

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

type JwtPlayload struct {
	Iss string `json:"iss"`
	Iat int64  `json:"iat"`
}

func (p JwtPlayload) Base64Content() (string, error) {

	payloadByte, err := json.Marshal(p)

	if nil != err {
		return "", err
	}

	return strings.TrimRight(base64.URLEncoding.EncodeToString(payloadByte), "="), nil
}
