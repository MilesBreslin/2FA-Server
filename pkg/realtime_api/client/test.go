package client

func (c *Client) Test(params []interface{}) ([]interface{}) {
	obj, response := c.runMethod("test",params)
	if response == 200 {
		return obj
	}
	return nil
}