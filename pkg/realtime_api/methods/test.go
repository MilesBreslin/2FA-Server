package methods

func init() {
	//m.description = "Test if methods work"
	add("test",test_method)
}

func test_method(in []interface{}) ([]interface{}, uint16) {
	return in, 200
}