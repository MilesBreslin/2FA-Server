package methods

import (
    "github.com/opensource2fa/pkg/totp"
)

func init() {
    // Actually add the method to the global list
    Add("GetTOTPToken",getTOTPToken_method)
}

func getTOTPToken_method(in []interface{}) ([]interface{}, uint16) {
    out := make([]interface{},0)
    // For each object in the array, ensure it has "secret" as type string
    for _, secret := range in {
        switch secret.(type) {
        case string:
            // Generate the TOTP token and append as a special datatype to the reply
            token := totp.GetTOTPToken(secret.(string))
            out = append(out,token)
        default:
            return nil, 400
        }
    }
    return out, 200
}
