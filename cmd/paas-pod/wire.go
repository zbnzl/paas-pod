//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/zbnzl/paas-pod/internal/biz"
	"github.com/zbnzl/paas-pod/internal/conf"
	"github.com/zbnzl/paas-pod/internal/dao"
	"github.com/zbnzl/paas-pod/internal/data"
	"github.com/zbnzl/paas-pod/internal/server"
	"github.com/zbnzl/paas-pod/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, dao.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
