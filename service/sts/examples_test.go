// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package sts_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/dragonfax/aws-sdk-go/aws"
	"github.com/dragonfax/aws-sdk-go/service/sts"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleSTS_AssumeRole() {
	svc := sts.New(nil)

	params := &sts.AssumeRoleInput{
		RoleArn:         aws.String("arnType"),      // Required
		RoleSessionName: aws.String("userNameType"), // Required
		DurationSeconds: aws.Int64(1),
		ExternalId:      aws.String("externalIdType"),
		Policy:          aws.String("sessionPolicyDocumentType"),
		SerialNumber:    aws.String("serialNumberType"),
		TokenCode:       aws.String("tokenCodeType"),
	}
	resp, err := svc.AssumeRole(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSTS_AssumeRoleWithSAML() {
	svc := sts.New(nil)

	params := &sts.AssumeRoleWithSAMLInput{
		PrincipalArn:    aws.String("arnType"),           // Required
		RoleArn:         aws.String("arnType"),           // Required
		SAMLAssertion:   aws.String("SAMLAssertionType"), // Required
		DurationSeconds: aws.Int64(1),
		Policy:          aws.String("sessionPolicyDocumentType"),
	}
	resp, err := svc.AssumeRoleWithSAML(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSTS_AssumeRoleWithWebIdentity() {
	svc := sts.New(nil)

	params := &sts.AssumeRoleWithWebIdentityInput{
		RoleArn:          aws.String("arnType"),         // Required
		RoleSessionName:  aws.String("userNameType"),    // Required
		WebIdentityToken: aws.String("clientTokenType"), // Required
		DurationSeconds:  aws.Int64(1),
		Policy:           aws.String("sessionPolicyDocumentType"),
		ProviderId:       aws.String("urlType"),
	}
	resp, err := svc.AssumeRoleWithWebIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSTS_DecodeAuthorizationMessage() {
	svc := sts.New(nil)

	params := &sts.DecodeAuthorizationMessageInput{
		EncodedMessage: aws.String("encodedMessageType"), // Required
	}
	resp, err := svc.DecodeAuthorizationMessage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSTS_GetFederationToken() {
	svc := sts.New(nil)

	params := &sts.GetFederationTokenInput{
		Name:            aws.String("userNameType"), // Required
		DurationSeconds: aws.Int64(1),
		Policy:          aws.String("sessionPolicyDocumentType"),
	}
	resp, err := svc.GetFederationToken(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSTS_GetSessionToken() {
	svc := sts.New(nil)

	params := &sts.GetSessionTokenInput{
		DurationSeconds: aws.Int64(1),
		SerialNumber:    aws.String("serialNumberType"),
		TokenCode:       aws.String("tokenCodeType"),
	}
	resp, err := svc.GetSessionToken(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
