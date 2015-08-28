// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package dynamodbstreamsiface provides an interface for the Amazon DynamoDB Streams.
package dynamodbstreamsiface

import (
	"github.com/dragonfax/aws-sdk-go/aws/request"
	"github.com/dragonfax/aws-sdk-go/service/dynamodbstreams"
)

// DynamoDBStreamsAPI is the interface type for dynamodbstreams.DynamoDBStreams.
type DynamoDBStreamsAPI interface {
	DescribeStreamRequest(*dynamodbstreams.DescribeStreamInput) (*request.Request, *dynamodbstreams.DescribeStreamOutput)

	DescribeStream(*dynamodbstreams.DescribeStreamInput) (*dynamodbstreams.DescribeStreamOutput, error)

	GetRecordsRequest(*dynamodbstreams.GetRecordsInput) (*request.Request, *dynamodbstreams.GetRecordsOutput)

	GetRecords(*dynamodbstreams.GetRecordsInput) (*dynamodbstreams.GetRecordsOutput, error)

	GetShardIteratorRequest(*dynamodbstreams.GetShardIteratorInput) (*request.Request, *dynamodbstreams.GetShardIteratorOutput)

	GetShardIterator(*dynamodbstreams.GetShardIteratorInput) (*dynamodbstreams.GetShardIteratorOutput, error)

	ListStreamsRequest(*dynamodbstreams.ListStreamsInput) (*request.Request, *dynamodbstreams.ListStreamsOutput)

	ListStreams(*dynamodbstreams.ListStreamsInput) (*dynamodbstreams.ListStreamsOutput, error)
}
