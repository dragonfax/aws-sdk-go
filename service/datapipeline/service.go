// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package datapipeline

import (
	"github.com/dragonfax/aws-sdk-go/aws"
	"github.com/dragonfax/aws-sdk-go/aws/defaults"
	"github.com/dragonfax/aws-sdk-go/aws/request"
	"github.com/dragonfax/aws-sdk-go/aws/service"
	"github.com/dragonfax/aws-sdk-go/aws/service/serviceinfo"
	"github.com/dragonfax/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/dragonfax/aws-sdk-go/internal/signer/v4"
)

// AWS Data Pipeline configures and manages a data-driven workflow called a
// pipeline. AWS Data Pipeline handles the details of scheduling and ensuring
// that data dependencies are met so that your application can focus on processing
// the data.
//
// AWS Data Pipeline provides a JAR implementation of a task runner called
// AWS Data Pipeline Task Runner. AWS Data Pipeline Task Runner provides logic
// for common data management scenarios, such as performing database queries
// and running data analysis using Amazon Elastic MapReduce (Amazon EMR). You
// can use AWS Data Pipeline Task Runner as your task runner, or you can write
// your own task runner to provide custom data management.
//
// AWS Data Pipeline implements two main sets of functionality. Use the first
// set to create a pipeline and define data sources, schedules, dependencies,
// and the transforms to be performed on the data. Use the second set in your
// task runner application to receive the next task ready for processing. The
// logic for performing the task, such as querying the data, running data analysis,
// or converting the data from one format to another, is contained within the
// task runner. The task runner performs the task assigned to it by the web
// service, reporting progress to the web service as it does so. When the task
// is done, the task runner reports the final success or failure of the task
// to the web service.
type DataPipeline struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new DataPipeline client.
func New(config *aws.Config) *DataPipeline {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:       defaults.DefaultConfig.Merge(config),
			ServiceName:  "datapipeline",
			APIVersion:   "2012-10-29",
			JSONVersion:  "1.1",
			TargetPrefix: "DataPipeline",
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

	return &DataPipeline{service}
}

// newRequest creates a new request for a DataPipeline operation and runs any
// custom request initialization.
func (c *DataPipeline) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
