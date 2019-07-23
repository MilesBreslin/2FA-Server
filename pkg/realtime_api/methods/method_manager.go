package methods

import "log"

type method func([]interface{}) ([]interface{}, uint16)

var methods map[string]method

func Add(n string, m method) {
    if methods == nil {
        methods = make(map[string]method)
    }
    methods[n] = m
}

func Get(n string) (method, bool) {
    log.Println("get"+n)
    log.Println(methods)
    m, ok := methods[n]
    return m, ok
}