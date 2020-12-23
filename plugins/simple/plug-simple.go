package main

import (
	"github.com/dubbogo/dubbo-go-proxy/pkg/context"
	"github.com/dubbogo/dubbo-go-proxy/pkg/filter"
	"github.com/dubbogo/dubbo-go-proxy/pkg/logger"
)

func New() filter.Filter {
	return &simpleFilter{}
}

type simpleFilter struct {
}

func (s simpleFilter) Do() context.FilterFunc {
	return func(c context.Context) {
		logger.Info("--------> hello world")
	}
}
