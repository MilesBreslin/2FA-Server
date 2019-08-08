package client

import (
    "../../keys"
    "../../status_codes"
)

func (c *Client) GetKey(id uint64) (*keys.Key, error) {
    // Send the request and wait for the return object and response code
    obj, response := c.runMethod("GetKey",[]interface{}{id})
    // If the response code is OK, return the object
    if response == status_codes.OK {
        return keys.MapToKey(obj[0].(map[string]interface{})), nil
    }
    // Else return nothing and the error for the status code
    return nil, status_codes.StatusToError(response)
}

func (c *Client) AddKey(secret string) (uint64, error) {
    send := make([]interface{},1)
    send[0] = struct{Secret string `json:"secret"`}{
        Secret: secret,
    }
    obj, response := c.runMethod("AddKey",send)
    if response == status_codes.OK {
        return uint64(obj[0].(float64)), nil
    }
    return 0, status_codes.StatusToError(response)
}

func (c *Client) ListKeys() ([]uint64, error) {
    obj, response := c.runMethod("ListKeys",nil)
    if response == status_codes.OK {
        ret := make([]uint64, len(obj))
        for index, val := range obj {
            ret[index] = uint64(val.(float64))
        }
        return ret, nil
    }
    return nil, status_codes.StatusToError(response)
}

func (c *Client) GetKeyToken(id uint64) (string, error) {
    // Send the request and wait for the return object and response code
    obj, response := c.runMethod("GetKeyToken",[]interface{}{id})
    // If the response code is OK, return the object
    if response == status_codes.OK {
        return keys.MapToKey(obj[0].(string)), nil
    }
    // Else return nothing and the error for the status code
    return "", status_codes.StatusToError(response)
}