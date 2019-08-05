package client

import (
    "../../keys"
    "../../status_codes"
)

func (c *Client) GetKey(id uint64) (*keys.Key, error) {
    obj, response := c.runMethod("GetKey",[]interface{}{id})
    if response == 200 {
        return keys.MapToKey(obj[0].(map[string]interface{})), nil
    }
    return nil, status_codes.StatusToError(response)
}

func (c *Client) AddKey(secret string) (uint64, error) {
    send := make([]interface{},1)
    send[0] = struct{Secret string `json:"secret"`}{
        Secret: secret,
    }
    obj, response := c.runMethod("AddKey",send)
    if response == 200 {
        return uint64(obj[0].(float64)), nil
    }
    return 0, status_codes.StatusToError(response)
}