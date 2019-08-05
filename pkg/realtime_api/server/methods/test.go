package methods

import (
    "../../../status_codes"
)

func init() {
	// Actually add the method to the global list
	Add("test",test_method)
}

func test_method(in []interface{}) ([]interface{}, uint16) {
	return in, status_codes.OK
}