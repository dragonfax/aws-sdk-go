// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package machinelearning

import (
	"github.com/dragonfax/aws-sdk-go/aws"
	"github.com/dragonfax/aws-sdk-go/aws/defaults"
	"github.com/dragonfax/aws-sdk-go/aws/request"
	"github.com/dragonfax/aws-sdk-go/aws/service"
	"github.com/dragonfax/aws-sdk-go/aws/service/serviceinfo"
	"github.com/dragonfax/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/dragonfax/aws-sdk-go/internal/signer/v4"
)

// Definition of the public APIs exposed by Amazon Machine Learning
type MachineLearning struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new MachineLearning client.
func New(config *aws.Config) *MachineLearning {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:       defaults.DefaultConfig.Merge(config),
			ServiceName:  "machinelearning",
			APIVersion:   "2014-12-12",
			JSONVersion:  "1.1",
			TargetPrefix: "AmazonML_20141212",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &MachineLearning{service}
}

// newRequest creates a new request for a MachineLearning operation and runs any
// custom request initialization.
func (c *MachineLearning) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
