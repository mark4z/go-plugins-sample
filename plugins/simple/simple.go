package main

import (
	"github.com/dubbogo/dubbo-go-proxy/pkg/context"
	"github.com/dubbogo/dubbo-go-proxy/pkg/logger"
	"github.com/mark4z/go-plugins-server/foo"
)

func O(t context.Context) {
	logger.Info("-->", t)
}

func I(t foo.Foo) {
	logger.Info("-->", t)
}
