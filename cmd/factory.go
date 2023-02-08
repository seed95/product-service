package main

import (
	"errors"
	"fmt"
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/handler"
	"google.golang.org/grpc"
	nativeLog "log"
	"os"

	"github.com/seed95/product-service/internal"
	"github.com/seed95/product-service/internal/repo/product"
	"github.com/seed95/product-service/internal/service"
	"github.com/seed95/product-service/pkg/logger"
	kitzap "github.com/seed95/product-service/pkg/logger/zap"

	"go.uber.org/zap/zapcore"
)

type ServerFactory struct {
	Config     *internal.Config
	Logger     logger.Logger
	Service    service.ProductService
	GRPRServer *grpc.Server
}

func NewServerFactory() (*ServerFactory, error) {

	configPrefix := os.Getenv("CONFIG_PREFIX")
	config := internal.NewConfig(configPrefix)

	// Setup logger
	zapLogger, err := newLogger(&config.Log)
	if err != nil {
		return nil, err
	}

	productRepo, err := product.New(&product.Setting{Config: &config.ProductRepo, Logger: zapLogger})
	if err != nil {
		return nil, err
	}

	productService, err := service.New(&service.Setting{ProductRepo: productRepo, Logger: zapLogger})
	if err != nil {
		return nil, err
	}

	grpcServer, err := handler.New(&handler.Setting{
		Config:  config,
		Service: productService,
		Logger:  zapLogger,
	})
	if err != nil {
		nativeLog.Fatal(err)
	}
	//grpcServer := grpc.NewServer(grpc.UnaryInterceptor(inteceptor))
	//micro.RegisterMicroServiceServer(grpcServer, gRPCHandler)

	return &ServerFactory{
		Config:     config,
		Logger:     zapLogger,
		Service:    productService,
		GRPRServer: grpcServer,
	}, nil
}

func newLogger(config *internal.LogConfig) (logger.Logger, error) {
	var cores []zapcore.Core

	// std log
	stdCore, err := kitzap.NewStandardCore(true, kitzap.Level(config.StdLevel))
	if err != nil {
		return nil, errors.New(fmt.Sprintf(derror.CreateStdLogErrorFormat, err))
	}

	cores = append(cores, stdCore)
	loggerCores := kitzap.NewZapLoggerWithCores(cores...)
	return loggerCores, nil
}
