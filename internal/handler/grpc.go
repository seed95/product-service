package handler

import (
	"context"
	"fmt"
	"github.com/seed95/product-service/internal"
	"github.com/seed95/product-service/internal/service"
	"github.com/seed95/product-service/pkg/proto/micro"
	"time"
)

// OpCodes
const (
	NewProductOpCode = 1
)

const (
	_serviceTimeout = 10 * time.Second
)

type (
	gRPCHandler struct {
		config  *internal.Config
		service service.ProductService
	}

	Setting struct {
		Config  *internal.Config
		Service service.ProductService
	}
)

func New(s *Setting) (micro.MicroServiceServer, error) {
	return &gRPCHandler{
		config:  s.Config,
		service: s.Service,
	}, nil
}

func (g *gRPCHandler) GeneralCall(ctx context.Context, req *micro.RequestMessage) (res *micro.ResponseMessage, err error) {

	serviceContext, cancel := context.WithTimeout(ctx, g.config.ServiceTimeout)
	_ = serviceContext
	defer cancel()

	switch req.GetOpCode() {
	case NewProductOpCode:
		fmt.Print("new product")
	default:
		panic("implement me")
	}
	return nil, err
}
