// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package mobileanalytics_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/dragonfax/aws-sdk-go/aws"
	"github.com/dragonfax/aws-sdk-go/service/mobileanalytics"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleMobileAnalytics_PutEvents() {
	svc := mobileanalytics.New(nil)

	params := &mobileanalytics.PutEventsInput{
		ClientContext: aws.String("String"), // Required
		Events: []*mobileanalytics.Event{ // Required
			{ // Required
				EventType: aws.String("String50Chars"),    // Required
				Timestamp: aws.String("ISO8601Timestamp"), // Required
				Attributes: map[string]*string{
					"Key": aws.String("String0to1000Chars"), // Required
					// More values...
				},
				Metrics: map[string]*float64{
					"Key": aws.Float64(1.0), // Required
					// More values...
				},
				Session: &mobileanalytics.Session{
					Duration:       aws.Int64(1),
					Id:             aws.String("String50Chars"),
					StartTimestamp: aws.String("ISO8601Timestamp"),
					StopTimestamp:  aws.String("ISO8601Timestamp"),
				},
				Version: aws.String("String10Chars"),
			},
			// More values...
		},
		ClientContextEncoding: aws.String("String"),
	}
	resp, err := svc.PutEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
