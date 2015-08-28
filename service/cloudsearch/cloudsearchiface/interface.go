// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudsearchiface provides an interface for the Amazon CloudSearch.
package cloudsearchiface

import (
	"github.com/dragonfax/aws-sdk-go/aws/request"
	"github.com/dragonfax/aws-sdk-go/service/cloudsearch"
)

// CloudSearchAPI is the interface type for cloudsearch.CloudSearch.
type CloudSearchAPI interface {
	BuildSuggestersRequest(*cloudsearch.BuildSuggestersInput) (*request.Request, *cloudsearch.BuildSuggestersOutput)

	BuildSuggesters(*cloudsearch.BuildSuggestersInput) (*cloudsearch.BuildSuggestersOutput, error)

	CreateDomainRequest(*cloudsearch.CreateDomainInput) (*request.Request, *cloudsearch.CreateDomainOutput)

	CreateDomain(*cloudsearch.CreateDomainInput) (*cloudsearch.CreateDomainOutput, error)

	DefineAnalysisSchemeRequest(*cloudsearch.DefineAnalysisSchemeInput) (*request.Request, *cloudsearch.DefineAnalysisSchemeOutput)

	DefineAnalysisScheme(*cloudsearch.DefineAnalysisSchemeInput) (*cloudsearch.DefineAnalysisSchemeOutput, error)

	DefineExpressionRequest(*cloudsearch.DefineExpressionInput) (*request.Request, *cloudsearch.DefineExpressionOutput)

	DefineExpression(*cloudsearch.DefineExpressionInput) (*cloudsearch.DefineExpressionOutput, error)

	DefineIndexFieldRequest(*cloudsearch.DefineIndexFieldInput) (*request.Request, *cloudsearch.DefineIndexFieldOutput)

	DefineIndexField(*cloudsearch.DefineIndexFieldInput) (*cloudsearch.DefineIndexFieldOutput, error)

	DefineSuggesterRequest(*cloudsearch.DefineSuggesterInput) (*request.Request, *cloudsearch.DefineSuggesterOutput)

	DefineSuggester(*cloudsearch.DefineSuggesterInput) (*cloudsearch.DefineSuggesterOutput, error)

	DeleteAnalysisSchemeRequest(*cloudsearch.DeleteAnalysisSchemeInput) (*request.Request, *cloudsearch.DeleteAnalysisSchemeOutput)

	DeleteAnalysisScheme(*cloudsearch.DeleteAnalysisSchemeInput) (*cloudsearch.DeleteAnalysisSchemeOutput, error)

	DeleteDomainRequest(*cloudsearch.DeleteDomainInput) (*request.Request, *cloudsearch.DeleteDomainOutput)

	DeleteDomain(*cloudsearch.DeleteDomainInput) (*cloudsearch.DeleteDomainOutput, error)

	DeleteExpressionRequest(*cloudsearch.DeleteExpressionInput) (*request.Request, *cloudsearch.DeleteExpressionOutput)

	DeleteExpression(*cloudsearch.DeleteExpressionInput) (*cloudsearch.DeleteExpressionOutput, error)

	DeleteIndexFieldRequest(*cloudsearch.DeleteIndexFieldInput) (*request.Request, *cloudsearch.DeleteIndexFieldOutput)

	DeleteIndexField(*cloudsearch.DeleteIndexFieldInput) (*cloudsearch.DeleteIndexFieldOutput, error)

	DeleteSuggesterRequest(*cloudsearch.DeleteSuggesterInput) (*request.Request, *cloudsearch.DeleteSuggesterOutput)

	DeleteSuggester(*cloudsearch.DeleteSuggesterInput) (*cloudsearch.DeleteSuggesterOutput, error)

	DescribeAnalysisSchemesRequest(*cloudsearch.DescribeAnalysisSchemesInput) (*request.Request, *cloudsearch.DescribeAnalysisSchemesOutput)

	DescribeAnalysisSchemes(*cloudsearch.DescribeAnalysisSchemesInput) (*cloudsearch.DescribeAnalysisSchemesOutput, error)

	DescribeAvailabilityOptionsRequest(*cloudsearch.DescribeAvailabilityOptionsInput) (*request.Request, *cloudsearch.DescribeAvailabilityOptionsOutput)

	DescribeAvailabilityOptions(*cloudsearch.DescribeAvailabilityOptionsInput) (*cloudsearch.DescribeAvailabilityOptionsOutput, error)

	DescribeDomainsRequest(*cloudsearch.DescribeDomainsInput) (*request.Request, *cloudsearch.DescribeDomainsOutput)

	DescribeDomains(*cloudsearch.DescribeDomainsInput) (*cloudsearch.DescribeDomainsOutput, error)

	DescribeExpressionsRequest(*cloudsearch.DescribeExpressionsInput) (*request.Request, *cloudsearch.DescribeExpressionsOutput)

	DescribeExpressions(*cloudsearch.DescribeExpressionsInput) (*cloudsearch.DescribeExpressionsOutput, error)

	DescribeIndexFieldsRequest(*cloudsearch.DescribeIndexFieldsInput) (*request.Request, *cloudsearch.DescribeIndexFieldsOutput)

	DescribeIndexFields(*cloudsearch.DescribeIndexFieldsInput) (*cloudsearch.DescribeIndexFieldsOutput, error)

	DescribeScalingParametersRequest(*cloudsearch.DescribeScalingParametersInput) (*request.Request, *cloudsearch.DescribeScalingParametersOutput)

	DescribeScalingParameters(*cloudsearch.DescribeScalingParametersInput) (*cloudsearch.DescribeScalingParametersOutput, error)

	DescribeServiceAccessPoliciesRequest(*cloudsearch.DescribeServiceAccessPoliciesInput) (*request.Request, *cloudsearch.DescribeServiceAccessPoliciesOutput)

	DescribeServiceAccessPolicies(*cloudsearch.DescribeServiceAccessPoliciesInput) (*cloudsearch.DescribeServiceAccessPoliciesOutput, error)

	DescribeSuggestersRequest(*cloudsearch.DescribeSuggestersInput) (*request.Request, *cloudsearch.DescribeSuggestersOutput)

	DescribeSuggesters(*cloudsearch.DescribeSuggestersInput) (*cloudsearch.DescribeSuggestersOutput, error)

	IndexDocumentsRequest(*cloudsearch.IndexDocumentsInput) (*request.Request, *cloudsearch.IndexDocumentsOutput)

	IndexDocuments(*cloudsearch.IndexDocumentsInput) (*cloudsearch.IndexDocumentsOutput, error)

	ListDomainNamesRequest(*cloudsearch.ListDomainNamesInput) (*request.Request, *cloudsearch.ListDomainNamesOutput)

	ListDomainNames(*cloudsearch.ListDomainNamesInput) (*cloudsearch.ListDomainNamesOutput, error)

	UpdateAvailabilityOptionsRequest(*cloudsearch.UpdateAvailabilityOptionsInput) (*request.Request, *cloudsearch.UpdateAvailabilityOptionsOutput)

	UpdateAvailabilityOptions(*cloudsearch.UpdateAvailabilityOptionsInput) (*cloudsearch.UpdateAvailabilityOptionsOutput, error)

	UpdateScalingParametersRequest(*cloudsearch.UpdateScalingParametersInput) (*request.Request, *cloudsearch.UpdateScalingParametersOutput)

	UpdateScalingParameters(*cloudsearch.UpdateScalingParametersInput) (*cloudsearch.UpdateScalingParametersOutput, error)

	UpdateServiceAccessPoliciesRequest(*cloudsearch.UpdateServiceAccessPoliciesInput) (*request.Request, *cloudsearch.UpdateServiceAccessPoliciesOutput)

	UpdateServiceAccessPolicies(*cloudsearch.UpdateServiceAccessPoliciesInput) (*cloudsearch.UpdateServiceAccessPoliciesOutput, error)
}
