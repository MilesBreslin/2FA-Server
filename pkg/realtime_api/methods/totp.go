package methods

import (
    "../../totp"
)

func init() {
    // Actually add the method to the global list
    Add("GetTOTPToken",getTOTPToken_method)
}

type getTOTPToken_reponse struct {
    Token           string              `json:"token"`
}

func getTOTPToken_method(in []interface{}) ([]interface{}, uint16) {
    out := make([]interface{},0)
    for _, obj := range in {
        if secret, ok := obj.(map[string]interface{})["secret"]; ok {
            switch secret.(type) {
            case string:
                token := totp.GetTOTPToken(secret.(string))
                out = append(
                    out,
                    getTOTPToken_reponse{
                        Token: token,
                    },
                )
            default:
                return nil, 400
            }
        } else {
            return nil, 400
        }
    }
    return out, 200
}