package main

import (
	"github.com/dubbogo/dubbo-go-proxy/pkg/context"
	"github.com/dubbogo/dubbo-go-proxy/pkg/logger"
)

func O() func(t context.Context) {
	logger.Info("-->")
	return func(t context.Context) {
		logger.Info("O()-->")
		t.Next()
		t.Next()
		t.Next()
		t.Next()
		logger.Info("O()<--")
	}
}
