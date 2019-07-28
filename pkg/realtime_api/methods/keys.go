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
    for _, obj := range in {
        if secret, ok := obj.(map[string]interface{})["secret"]; ok {
            switch secret.(type) {
            case string:
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
    for _, obj := range in {
        if id, ok := obj.(map[string]interface{})["id"]; ok {
            switch id.(type) {
            case float64:
                k, err := keys.GetKey(uint64(id.(float64)))
                if err != nil {
                	return nil, 404
                }
                out = append(out, *k)
            default:
                return nil, 400
            }
        } else {
            return nil, 400
        }
    }
    return out, 200
}

func listKeys_method(in []interface{}) ([]interface{}, uint16) {
    list := keys.GetList()
    out := make([]interface{},len(list))
    for index, val := range list {
        out[index] = val
    }
    return out, 200
}