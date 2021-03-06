// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package mobileanalytics

import (
	"github.com/dragonfax/aws-sdk-go/aws"
	"github.com/dragonfax/aws-sdk-go/aws/defaults"
	"github.com/dragonfax/aws-sdk-go/aws/request"
	"github.com/dragonfax/aws-sdk-go/aws/service"
	"github.com/dragonfax/aws-sdk-go/aws/service/serviceinfo"
	"github.com/dragonfax/aws-sdk-go/internal/protocol/restjson"
	"github.com/dragonfax/aws-sdk-go/internal/signer/v4"
)

// Amazon Mobile Analytics is a service for collecting, visualizing, and understanding
// app usage data at scale.
type MobileAnalytics struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new MobileAnalytics client.
func New(config *aws.Config) *MobileAnalytics {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "mobileanalytics",
			APIVersion:  "2014-06-05",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &MobileAnalytics{service}
}

// newRequest creates a new request for a MobileAnalytics operation and runs any
// custom request initialization.
func (c *MobileAnalytics) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
