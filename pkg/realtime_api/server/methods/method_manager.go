package methods

// A method is a function with specific inputs and outputs
type method func([]interface{}) ([]interface{}, uint16)

// The global list of methods is a map with a name to look it up by
var methods map[string]method

// Add the method to the methods map
func Add(n string, m method) {
	// Ensure methods has been initialized
	// This will initialize it when run for the first time
	// This elimantes race conditions
    if methods == nil {
        methods = make(map[string]method)
    }
    // Assign method to name in global map
    methods[n] = m
}

// Basically a wrapper to ensure that "ok" gets returned
// Also ensures readonly access when using this function
func Get(n string) (method, bool) {
    m, ok := methods[n]
    return m, ok
}