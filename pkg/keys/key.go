package keys

import (
    "github.com/opensource2fa/server/pkg/totp"
)

type Key struct {
    Secret          string              `json:"secret"`
    Id              uint64              `json:"id"`
}

func (k *Key) GetCode() (string, error) {
    return totp.GetTOTPToken(k.Secret) , nil
}

func MapToKey(obj map[string]interface{}) (*Key) {
	return &Key{
		Secret: obj["secret"].(string),
		Id: uint64(obj["id"].(float64)),
	}
}