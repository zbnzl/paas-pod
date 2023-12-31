// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/zbnzl/paas-pod/internal/biz"
	"github.com/zbnzl/paas-pod/internal/conf"
	"github.com/zbnzl/paas-pod/internal/dao"
	"github.com/zbnzl/paas-pod/internal/data"
	"github.com/zbnzl/paas-pod/internal/server"
	"github.com/zbnzl/paas-pod/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	db, err := dao.NewMysqlDb(confData)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	iPodRepository := data.NewPodRepository(db)
	clientset := dao.NewK8sClient(confData)
	podUsecase := biz.NewPodUsecase(iPodRepository, clientset, logger)
	podService := service.NewPodService(podUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, greeterService, podService, logger)
	httpServer := server.NewHTTPServer(confServer, greeterService, podService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
