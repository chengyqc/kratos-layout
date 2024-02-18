//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos-layout/internal/server"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(config *conf.Config, engine *gin.Engine) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		//data.ProviderSet,
		//biz.ProviderSet,
		//service.ProviderSet,
		newApp,
	))
}
