package client

import (
    "github.com/opensource2fa/server/pkg/status_codes"
)

func (c *Client) Test(params []interface{}) ([]interface{}) {
	obj, response := c.runMethod("test",params)
	if response == status_codes.OK {
		return obj
	}
	return nil
}