package main

import (
	"fmt"
	pkg "github.com/mark4z/go-plugins-server/pkg"
)

type AA struct {
}

func (a *AA) Do() pkg.A {
	a2 := pkg.A{Name: "Alice"}
	fmt.Println("a")
	return a2
}
