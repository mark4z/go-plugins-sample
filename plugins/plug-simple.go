package main

import (
	"fmt"
	"github.com/dubbogo/dubbo-go-proxy/pkg/context"
)

type Filter interface {

	// Do run filter, use c.next() to next filter, before is pre logic, after is post logic.
	Do() context.FilterFunc
}

func T() {
	fmt.Println("hello world")
}
