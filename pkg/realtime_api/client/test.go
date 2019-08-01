package client

func (c *client) Test(params []interface{}) ([]interface{}) {
	obj, response := c.runMethod("test",params)
	if response == 200 {
		return obj
	}
	return nil
}