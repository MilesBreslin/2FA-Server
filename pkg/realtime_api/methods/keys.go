package methods

import (
	"../../keys"
)

func init() {
    // Actually add the method to the global list
    Add("AddKey",addKey_method)
    Add("GetKey",getKey_method)
    Add("ListKeys",listKeys_method)
}

func addKey_method(in []interface{}) ([]interface{}, uint16) {
    out := make([]interface{}, 0)
    // For each object in the array, ensure it has "secret" as type string
    for _, obj := range in {
        if secret, ok := obj.(map[string]interface{})["secret"]; ok {
            switch secret.(type) {
            case string:
                // Add the Key and append its Id to the output array
                k := keys.AddKey(secret.(string))
                out = append(out, k.Id)
            default:
                return out, 400
            }
        } else {
            return out, 400
        }
    }
    return out, 200
}

func getKey_method(in []interface{}) ([]interface{}, uint16) {
	out := make([]interface{}, 0)
    // For each object in the array, ensure it is type float64 (JSON default numeric type)
    for _, id := range in {
        switch id.(type) {
        case float64:
            // Retrieve key, return 404 if no exist, and append key to output array
            k, err := keys.GetKey(uint64(id.(float64)))
            if err != nil {
            	return nil, 404
            }
            out = append(out, *k)
        default:
            return nil, 400
        }
    }
    return out, 200
}

func listKeys_method(in []interface{}) ([]interface{}, uint16) {
    // No input parameters, just return the list
    list := keys.GetList()

    // Convert []uint64 to []interface{}
    out := make([]interface{},len(list))
    for index, val := range list {
        out[index] = val
    }
    return out, 200
}