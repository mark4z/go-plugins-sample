package main

import (
	"fmt"
	"github.com/dubbogo/dubbo-go-proxy-filter/pkg/context"
	"github.com/dubbogo/dubbo-go-proxy-filter/pkg/filter"
)

type F struct {
	v string
}

func (f *F) Do() context.FilterFunc {
	return func(c context.Context) {
		fmt.Println("F()", f.v)
	}
}

func NewF() filter.Filter {
	return &F{"0.0.2"}
}
