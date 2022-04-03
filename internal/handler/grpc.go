package handler

import (
	"context"
	"encoding/json"
	"github.com/seed95/product-service/internal"
	"github.com/seed95/product-service/internal/api"
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/service"
	"github.com/seed95/product-service/pkg/logger"
	"github.com/seed95/product-service/pkg/logger/keyval"
	"github.com/seed95/product-service/pkg/proto/micro"
	"google.golang.org/grpc/codes"
	"strings"
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
		logger  logger.Logger
	}

	Setting struct {
		Config  *internal.Config
		Service service.ProductService
		Logger  logger.Logger
	}
)

func New(s *Setting) (micro.MicroServiceServer, error) {
	return &gRPCHandler{
		config:  s.Config,
		service: s.Service,
		logger:  s.Logger,
	}, nil
}

func (h *gRPCHandler) GeneralCall(ctx context.Context, req *micro.RequestMessage) (res *micro.ResponseMessage, err error) {

	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("req", req.String()),
			keyval.String("res", res.String()),
		}
		logger.LogReqRes(h.logger, "http.general", err, commonKeyVal...)
	}()

	var payload interface{}

	serviceContext, cancel := context.WithTimeout(ctx, h.config.ServiceTimeout)
	_ = serviceContext
	defer cancel()

	// ServiceRequest Common
	common := api.Common{
		Language:    strings.ToLower(req.GetLanguage()),
		Username:    req.GetUsername(),
		CompanyName: req.GetCompanyName(),
	}

	switch req.GetOpCode() {

	case NewProductOpCode:
		// Make service request
		serviceRequest := &api.CreateNewProductRequest{Common: &common}
		if err = json.Unmarshal([]byte(req.GetPayload()), serviceRequest); err != nil {
			err = derror.New(derror.BadRequest, err.Error())
			break
		}

		// Call service
		payload, err = h.service.CreateNewProduct(ctx, serviceRequest)

	default:
		err = derror.NotImplemented
	}

	res = makeResponse(payload, err)
	return res, nil
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
