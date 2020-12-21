package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("/go/src/out/plugins.so")
	if err != nil {
		fmt.Println(err)
	}
	lookup, err := p.Lookup("F")
	fmt.Println(err)

	lookup.(func())()

	lookup, err = p.Lookup("F")
	lookup.(func())()
	fmt.Println(err)
}
