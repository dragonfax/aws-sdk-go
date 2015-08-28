// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudfront

import (
	"github.com/dragonfax/aws-sdk-go/aws"
	"github.com/dragonfax/aws-sdk-go/aws/defaults"
	"github.com/dragonfax/aws-sdk-go/aws/request"
	"github.com/dragonfax/aws-sdk-go/aws/service"
	"github.com/dragonfax/aws-sdk-go/aws/service/serviceinfo"
	"github.com/dragonfax/aws-sdk-go/internal/protocol/restxml"
	"github.com/dragonfax/aws-sdk-go/internal/signer/v4"
)

// CloudFront is a client for CloudFront.
type CloudFront struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new CloudFront client.
func New(config *aws.Config) *CloudFront {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "cloudfront",
			APIVersion:  "2015-04-17",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &CloudFront{service}
}

// newRequest creates a new request for a CloudFront operation and runs any
// custom request initialization.
func (c *CloudFront) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
