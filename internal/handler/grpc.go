package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/seed95/product-service/internal"
	"github.com/seed95/product-service/internal/api"
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/service"
	"github.com/seed95/product-service/pkg/logger"
	"github.com/seed95/product-service/pkg/logger/keyval"
	"github.com/seed95/product-service/pkg/proto/micro"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"strings"
	"time"
)

// OpCodes
const (
	NewProductOpCode = 1
)

type (
	gRPCHandler struct {
		config  *internal.Config
		service service.ProductService
		logger  logger.Logger
	}

	Setting struct {
		Config  *internal.Config
		Service service.ProductService
		Logger  logger.Logger
	}
)

func New(s *Setting) (*grpc.Server, error) {

	handler := &gRPCHandler{
		config:  s.Config,
		service: s.Service,
		logger:  s.Logger,
	}

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(logInterceptor(s.Logger)))
	micro.RegisterMicroServiceServer(grpcServer, handler)

	return grpcServer, nil
}

func (h *gRPCHandler) GeneralCall(ctx context.Context, req *micro.RequestMessage) (res *micro.ResponseMessage, err error) {

	// Recover panic
	defer func() {
		if r := recover(); r != nil {
			res = nil
			err = derror.New(derror.InternalServer, fmt.Sprintf("%+v", r))
		}
	}()

	var payload interface{}

	serviceContext, cancel := context.WithTimeout(ctx, h.config.ServiceTimeout)
	_ = serviceContext
	defer cancel()

	// ServiceRequest Common
	common := api.Common{}

	reqBytes, _ := json.Marshal(req)
	if err = json.Unmarshal(reqBytes, &common); err != nil {
		return nil, derror.New(derror.BadRequest, err.Error())
	}

	fmt.Println(common)

	switch req.GetOpCode() {

	case NewProductOpCode:
		// Make service request
		serviceRequest := api.CreateNewProductRequest{Common: &common}
		if err = json.Unmarshal([]byte(req.GetPayload()), &serviceRequest); err != nil {
			err = derror.New(derror.BadRequest, err.Error())
			break
		}

		fmt.Println(serviceRequest)
		// Call service
		payload, err = h.service.CreateNewProduct(ctx, serviceRequest)

	default:
		err = derror.NotImplemented

	}

	res = makeResponse(payload, err)
	return res, nil
}

func logInterceptor(l logger.Logger) grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		start := time.Now()
		resp, err = handler(ctx, req)

		commonKeyVal := []keyval.Pair{
			keyval.String("duration", time.Since(start).String()),
			keyval.String("req", fmt.Sprintf("%+v", req)),
			keyval.String("res", fmt.Sprintf("%+v", resp)),
		}
		logger.LogReqRes(l, "grpc."+strings.Split(info.FullMethod, "/")[2], err, commonKeyVal...)

		return resp, err

	}
}

func makeResponse(payload interface{}, err error) *micro.ResponseMessage {
	res := &micro.ResponseMessage{}
	if err != nil {
		res.StatusMessage = derror.StatusText(err)
		res.StatusCode = int32(derror.StatusCode(err))
		res.Payload = "{}"
	} else {
		res.StatusMessage = "Ok"
		res.StatusCode = int32(codes.OK)
		payloadBytes, _ := json.Marshal(payload)
		res.Payload = string(payloadBytes)
	}
	return res
}
