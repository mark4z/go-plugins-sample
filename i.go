package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("/go/src/out/plug-simple.so")
	if err != nil {
		fmt.Println(err)
	}
	lookup, err := p.Lookup("T")
	fmt.Println(err)

	lookup.(func())()

	p, err = plugin.Open("/go/src/out/hello_world.so")
	lookup, err = p.Lookup("T")
	lookup.(func())()
	fmt.Println(err)
}
