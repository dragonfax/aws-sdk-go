// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package elasticbeanstalkiface provides an interface for the AWS Elastic Beanstalk.
package elasticbeanstalkiface

import (
	"github.com/dragonfax/aws-sdk-go/aws/request"
	"github.com/dragonfax/aws-sdk-go/service/elasticbeanstalk"
)

// ElasticBeanstalkAPI is the interface type for elasticbeanstalk.ElasticBeanstalk.
type ElasticBeanstalkAPI interface {
	AbortEnvironmentUpdateRequest(*elasticbeanstalk.AbortEnvironmentUpdateInput) (*request.Request, *elasticbeanstalk.AbortEnvironmentUpdateOutput)

	AbortEnvironmentUpdate(*elasticbeanstalk.AbortEnvironmentUpdateInput) (*elasticbeanstalk.AbortEnvironmentUpdateOutput, error)

	CheckDNSAvailabilityRequest(*elasticbeanstalk.CheckDNSAvailabilityInput) (*request.Request, *elasticbeanstalk.CheckDNSAvailabilityOutput)

	CheckDNSAvailability(*elasticbeanstalk.CheckDNSAvailabilityInput) (*elasticbeanstalk.CheckDNSAvailabilityOutput, error)

	CreateApplicationRequest(*elasticbeanstalk.CreateApplicationInput) (*request.Request, *elasticbeanstalk.ApplicationDescriptionMessage)

	CreateApplication(*elasticbeanstalk.CreateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error)

	CreateApplicationVersionRequest(*elasticbeanstalk.CreateApplicationVersionInput) (*request.Request, *elasticbeanstalk.ApplicationVersionDescriptionMessage)

	CreateApplicationVersion(*elasticbeanstalk.CreateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error)

	CreateConfigurationTemplateRequest(*elasticbeanstalk.CreateConfigurationTemplateInput) (*request.Request, *elasticbeanstalk.ConfigurationSettingsDescription)

	CreateConfigurationTemplate(*elasticbeanstalk.CreateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error)

	CreateEnvironmentRequest(*elasticbeanstalk.CreateEnvironmentInput) (*request.Request, *elasticbeanstalk.EnvironmentDescription)

	CreateEnvironment(*elasticbeanstalk.CreateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error)

	CreateStorageLocationRequest(*elasticbeanstalk.CreateStorageLocationInput) (*request.Request, *elasticbeanstalk.CreateStorageLocationOutput)

	CreateStorageLocation(*elasticbeanstalk.CreateStorageLocationInput) (*elasticbeanstalk.CreateStorageLocationOutput, error)

	DeleteApplicationRequest(*elasticbeanstalk.DeleteApplicationInput) (*request.Request, *elasticbeanstalk.DeleteApplicationOutput)

	DeleteApplication(*elasticbeanstalk.DeleteApplicationInput) (*elasticbeanstalk.DeleteApplicationOutput, error)

	DeleteApplicationVersionRequest(*elasticbeanstalk.DeleteApplicationVersionInput) (*request.Request, *elasticbeanstalk.DeleteApplicationVersionOutput)

	DeleteApplicationVersion(*elasticbeanstalk.DeleteApplicationVersionInput) (*elasticbeanstalk.DeleteApplicationVersionOutput, error)

	DeleteConfigurationTemplateRequest(*elasticbeanstalk.DeleteConfigurationTemplateInput) (*request.Request, *elasticbeanstalk.DeleteConfigurationTemplateOutput)

	DeleteConfigurationTemplate(*elasticbeanstalk.DeleteConfigurationTemplateInput) (*elasticbeanstalk.DeleteConfigurationTemplateOutput, error)

	DeleteEnvironmentConfigurationRequest(*elasticbeanstalk.DeleteEnvironmentConfigurationInput) (*request.Request, *elasticbeanstalk.DeleteEnvironmentConfigurationOutput)

	DeleteEnvironmentConfiguration(*elasticbeanstalk.DeleteEnvironmentConfigurationInput) (*elasticbeanstalk.DeleteEnvironmentConfigurationOutput, error)

	DescribeApplicationVersionsRequest(*elasticbeanstalk.DescribeApplicationVersionsInput) (*request.Request, *elasticbeanstalk.DescribeApplicationVersionsOutput)

	DescribeApplicationVersions(*elasticbeanstalk.DescribeApplicationVersionsInput) (*elasticbeanstalk.DescribeApplicationVersionsOutput, error)

	DescribeApplicationsRequest(*elasticbeanstalk.DescribeApplicationsInput) (*request.Request, *elasticbeanstalk.DescribeApplicationsOutput)

	DescribeApplications(*elasticbeanstalk.DescribeApplicationsInput) (*elasticbeanstalk.DescribeApplicationsOutput, error)

	DescribeConfigurationOptionsRequest(*elasticbeanstalk.DescribeConfigurationOptionsInput) (*request.Request, *elasticbeanstalk.DescribeConfigurationOptionsOutput)

	DescribeConfigurationOptions(*elasticbeanstalk.DescribeConfigurationOptionsInput) (*elasticbeanstalk.DescribeConfigurationOptionsOutput, error)

	DescribeConfigurationSettingsRequest(*elasticbeanstalk.DescribeConfigurationSettingsInput) (*request.Request, *elasticbeanstalk.DescribeConfigurationSettingsOutput)

	DescribeConfigurationSettings(*elasticbeanstalk.DescribeConfigurationSettingsInput) (*elasticbeanstalk.DescribeConfigurationSettingsOutput, error)

	DescribeEnvironmentHealthRequest(*elasticbeanstalk.DescribeEnvironmentHealthInput) (*request.Request, *elasticbeanstalk.DescribeEnvironmentHealthOutput)

	DescribeEnvironmentHealth(*elasticbeanstalk.DescribeEnvironmentHealthInput) (*elasticbeanstalk.DescribeEnvironmentHealthOutput, error)

	DescribeEnvironmentResourcesRequest(*elasticbeanstalk.DescribeEnvironmentResourcesInput) (*request.Request, *elasticbeanstalk.DescribeEnvironmentResourcesOutput)

	DescribeEnvironmentResources(*elasticbeanstalk.DescribeEnvironmentResourcesInput) (*elasticbeanstalk.DescribeEnvironmentResourcesOutput, error)

	DescribeEnvironmentsRequest(*elasticbeanstalk.DescribeEnvironmentsInput) (*request.Request, *elasticbeanstalk.DescribeEnvironmentsOutput)

	DescribeEnvironments(*elasticbeanstalk.DescribeEnvironmentsInput) (*elasticbeanstalk.DescribeEnvironmentsOutput, error)

	DescribeEventsRequest(*elasticbeanstalk.DescribeEventsInput) (*request.Request, *elasticbeanstalk.DescribeEventsOutput)

	DescribeEvents(*elasticbeanstalk.DescribeEventsInput) (*elasticbeanstalk.DescribeEventsOutput, error)

	DescribeEventsPages(*elasticbeanstalk.DescribeEventsInput, func(*elasticbeanstalk.DescribeEventsOutput, bool) bool) error

	DescribeInstancesHealthRequest(*elasticbeanstalk.DescribeInstancesHealthInput) (*request.Request, *elasticbeanstalk.DescribeInstancesHealthOutput)

	DescribeInstancesHealth(*elasticbeanstalk.DescribeInstancesHealthInput) (*elasticbeanstalk.DescribeInstancesHealthOutput, error)

	ListAvailableSolutionStacksRequest(*elasticbeanstalk.ListAvailableSolutionStacksInput) (*request.Request, *elasticbeanstalk.ListAvailableSolutionStacksOutput)

	ListAvailableSolutionStacks(*elasticbeanstalk.ListAvailableSolutionStacksInput) (*elasticbeanstalk.ListAvailableSolutionStacksOutput, error)

	RebuildEnvironmentRequest(*elasticbeanstalk.RebuildEnvironmentInput) (*request.Request, *elasticbeanstalk.RebuildEnvironmentOutput)

	RebuildEnvironment(*elasticbeanstalk.RebuildEnvironmentInput) (*elasticbeanstalk.RebuildEnvironmentOutput, error)

	RequestEnvironmentInfoRequest(*elasticbeanstalk.RequestEnvironmentInfoInput) (*request.Request, *elasticbeanstalk.RequestEnvironmentInfoOutput)

	RequestEnvironmentInfo(*elasticbeanstalk.RequestEnvironmentInfoInput) (*elasticbeanstalk.RequestEnvironmentInfoOutput, error)

	RestartAppServerRequest(*elasticbeanstalk.RestartAppServerInput) (*request.Request, *elasticbeanstalk.RestartAppServerOutput)

	RestartAppServer(*elasticbeanstalk.RestartAppServerInput) (*elasticbeanstalk.RestartAppServerOutput, error)

	RetrieveEnvironmentInfoRequest(*elasticbeanstalk.RetrieveEnvironmentInfoInput) (*request.Request, *elasticbeanstalk.RetrieveEnvironmentInfoOutput)

	RetrieveEnvironmentInfo(*elasticbeanstalk.RetrieveEnvironmentInfoInput) (*elasticbeanstalk.RetrieveEnvironmentInfoOutput, error)

	SwapEnvironmentCNAMEsRequest(*elasticbeanstalk.SwapEnvironmentCNAMEsInput) (*request.Request, *elasticbeanstalk.SwapEnvironmentCNAMEsOutput)

	SwapEnvironmentCNAMEs(*elasticbeanstalk.SwapEnvironmentCNAMEsInput) (*elasticbeanstalk.SwapEnvironmentCNAMEsOutput, error)

	TerminateEnvironmentRequest(*elasticbeanstalk.TerminateEnvironmentInput) (*request.Request, *elasticbeanstalk.EnvironmentDescription)

	TerminateEnvironment(*elasticbeanstalk.TerminateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error)

	UpdateApplicationRequest(*elasticbeanstalk.UpdateApplicationInput) (*request.Request, *elasticbeanstalk.ApplicationDescriptionMessage)

	UpdateApplication(*elasticbeanstalk.UpdateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error)

	UpdateApplicationVersionRequest(*elasticbeanstalk.UpdateApplicationVersionInput) (*request.Request, *elasticbeanstalk.ApplicationVersionDescriptionMessage)

	UpdateApplicationVersion(*elasticbeanstalk.UpdateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error)

	UpdateConfigurationTemplateRequest(*elasticbeanstalk.UpdateConfigurationTemplateInput) (*request.Request, *elasticbeanstalk.ConfigurationSettingsDescription)

	UpdateConfigurationTemplate(*elasticbeanstalk.UpdateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error)

	UpdateEnvironmentRequest(*elasticbeanstalk.UpdateEnvironmentInput) (*request.Request, *elasticbeanstalk.EnvironmentDescription)

	UpdateEnvironment(*elasticbeanstalk.UpdateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error)

	ValidateConfigurationSettingsRequest(*elasticbeanstalk.ValidateConfigurationSettingsInput) (*request.Request, *elasticbeanstalk.ValidateConfigurationSettingsOutput)

	ValidateConfigurationSettings(*elasticbeanstalk.ValidateConfigurationSettingsInput) (*elasticbeanstalk.ValidateConfigurationSettingsOutput, error)
}
