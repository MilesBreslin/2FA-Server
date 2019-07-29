package client

import (
    "github.com/opensource2fa/pkg/keys"
    "github.com/opensource2fa/pkg/status_codes"
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