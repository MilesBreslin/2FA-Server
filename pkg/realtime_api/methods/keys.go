package methods

import (
	"../../keys"
)

func init() {
    // Actually add the method to the global list
    Add("AddKey",addKey_method)
    Add("GetKey",getKey_method)
}

func addKey_method(in []interface{}) ([]interface{}, uint16) {
    for _, obj := range in {
        if secret, ok := obj.(map[string]interface{})["secret"]; ok {
            switch secret.(type) {
            case string:
                keys.AddKey(secret.(string))
            default:
                return nil, 400
            }
        } else {
            return nil, 400
        }
    }
    return nil, 200
}

func getKey_method(in []interface{}) ([]interface{}, uint16) {
	out := make([]interface{}, 0)
    for _, obj := range in {
        if secret, ok := obj.(map[string]interface{})["id"]; ok {
            switch secret.(type) {
            case float64:
                k, err := keys.GetKey(id.(string))
                if err != nil {
                	return nil, 404
                }
                append(out, *k)
            default:
                return nil, 400
            }
        } else {
            return nil, 400
        }
    }
    return nil, 200
}