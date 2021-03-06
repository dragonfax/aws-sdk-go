// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package ses_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/dragonfax/aws-sdk-go/aws"
	"github.com/dragonfax/aws-sdk-go/service/ses"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleSES_DeleteIdentity() {
	svc := ses.New(nil)

	params := &ses.DeleteIdentityInput{
		Identity: aws.String("Identity"), // Required
	}
	resp, err := svc.DeleteIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_DeleteIdentityPolicy() {
	svc := ses.New(nil)

	params := &ses.DeleteIdentityPolicyInput{
		Identity:   aws.String("Identity"),   // Required
		PolicyName: aws.String("PolicyName"), // Required
	}
	resp, err := svc.DeleteIdentityPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_DeleteVerifiedEmailAddress() {
	svc := ses.New(nil)

	params := &ses.DeleteVerifiedEmailAddressInput{
		EmailAddress: aws.String("Address"), // Required
	}
	resp, err := svc.DeleteVerifiedEmailAddress(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_GetIdentityDkimAttributes() {
	svc := ses.New(nil)

	params := &ses.GetIdentityDkimAttributesInput{
		Identities: []*string{ // Required
			aws.String("Identity"), // Required
			// More values...
		},
	}
	resp, err := svc.GetIdentityDkimAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_GetIdentityNotificationAttributes() {
	svc := ses.New(nil)

	params := &ses.GetIdentityNotificationAttributesInput{
		Identities: []*string{ // Required
			aws.String("Identity"), // Required
			// More values...
		},
	}
	resp, err := svc.GetIdentityNotificationAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_GetIdentityPolicies() {
	svc := ses.New(nil)

	params := &ses.GetIdentityPoliciesInput{
		Identity: aws.String("Identity"), // Required
		PolicyNames: []*string{ // Required
			aws.String("PolicyName"), // Required
			// More values...
		},
	}
	resp, err := svc.GetIdentityPolicies(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_GetIdentityVerificationAttributes() {
	svc := ses.New(nil)

	params := &ses.GetIdentityVerificationAttributesInput{
		Identities: []*string{ // Required
			aws.String("Identity"), // Required
			// More values...
		},
	}
	resp, err := svc.GetIdentityVerificationAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_GetSendQuota() {
	svc := ses.New(nil)

	var params *ses.GetSendQuotaInput
	resp, err := svc.GetSendQuota(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_GetSendStatistics() {
	svc := ses.New(nil)

	var params *ses.GetSendStatisticsInput
	resp, err := svc.GetSendStatistics(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_ListIdentities() {
	svc := ses.New(nil)

	params := &ses.ListIdentitiesInput{
		IdentityType: aws.String("IdentityType"),
		MaxItems:     aws.Int64(1),
		NextToken:    aws.String("NextToken"),
	}
	resp, err := svc.ListIdentities(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_ListIdentityPolicies() {
	svc := ses.New(nil)

	params := &ses.ListIdentityPoliciesInput{
		Identity: aws.String("Identity"), // Required
	}
	resp, err := svc.ListIdentityPolicies(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_ListVerifiedEmailAddresses() {
	svc := ses.New(nil)

	var params *ses.ListVerifiedEmailAddressesInput
	resp, err := svc.ListVerifiedEmailAddresses(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_PutIdentityPolicy() {
	svc := ses.New(nil)

	params := &ses.PutIdentityPolicyInput{
		Identity:   aws.String("Identity"),   // Required
		Policy:     aws.String("Policy"),     // Required
		PolicyName: aws.String("PolicyName"), // Required
	}
	resp, err := svc.PutIdentityPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_SendEmail() {
	svc := ses.New(nil)

	params := &ses.SendEmailInput{
		Destination: &ses.Destination{ // Required
			BccAddresses: []*string{
				aws.String("Address"), // Required
				// More values...
			},
			CcAddresses: []*string{
				aws.String("Address"), // Required
				// More values...
			},
			ToAddresses: []*string{
				aws.String("Address"), // Required
				// More values...
			},
		},
		Message: &ses.Message{ // Required
			Body: &ses.Body{ // Required
				Html: &ses.Content{
					Data:    aws.String("MessageData"), // Required
					Charset: aws.String("Charset"),
				},
				Text: &ses.Content{
					Data:    aws.String("MessageData"), // Required
					Charset: aws.String("Charset"),
				},
			},
			Subject: &ses.Content{ // Required
				Data:    aws.String("MessageData"), // Required
				Charset: aws.String("Charset"),
			},
		},
		Source: aws.String("Address"), // Required
		ReplyToAddresses: []*string{
			aws.String("Address"), // Required
			// More values...
		},
		ReturnPath:    aws.String("Address"),
		ReturnPathArn: aws.String("AmazonResourceName"),
		SourceArn:     aws.String("AmazonResourceName"),
	}
	resp, err := svc.SendEmail(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_SendRawEmail() {
	svc := ses.New(nil)

	params := &ses.SendRawEmailInput{
		RawMessage: &ses.RawMessage{ // Required
			Data: []byte("PAYLOAD"), // Required
		},
		Destinations: []*string{
			aws.String("Address"), // Required
			// More values...
		},
		FromArn:       aws.String("AmazonResourceName"),
		ReturnPathArn: aws.String("AmazonResourceName"),
		Source:        aws.String("Address"),
		SourceArn:     aws.String("AmazonResourceName"),
	}
	resp, err := svc.SendRawEmail(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_SetIdentityDkimEnabled() {
	svc := ses.New(nil)

	params := &ses.SetIdentityDkimEnabledInput{
		DkimEnabled: aws.Bool(true),         // Required
		Identity:    aws.String("Identity"), // Required
	}
	resp, err := svc.SetIdentityDkimEnabled(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_SetIdentityFeedbackForwardingEnabled() {
	svc := ses.New(nil)

	params := &ses.SetIdentityFeedbackForwardingEnabledInput{
		ForwardingEnabled: aws.Bool(true),         // Required
		Identity:          aws.String("Identity"), // Required
	}
	resp, err := svc.SetIdentityFeedbackForwardingEnabled(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_SetIdentityNotificationTopic() {
	svc := ses.New(nil)

	params := &ses.SetIdentityNotificationTopicInput{
		Identity:         aws.String("Identity"),         // Required
		NotificationType: aws.String("NotificationType"), // Required
		SnsTopic:         aws.String("NotificationTopic"),
	}
	resp, err := svc.SetIdentityNotificationTopic(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_VerifyDomainDkim() {
	svc := ses.New(nil)

	params := &ses.VerifyDomainDkimInput{
		Domain: aws.String("Domain"), // Required
	}
	resp, err := svc.VerifyDomainDkim(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_VerifyDomainIdentity() {
	svc := ses.New(nil)

	params := &ses.VerifyDomainIdentityInput{
		Domain: aws.String("Domain"), // Required
	}
	resp, err := svc.VerifyDomainIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_VerifyEmailAddress() {
	svc := ses.New(nil)

	params := &ses.VerifyEmailAddressInput{
		EmailAddress: aws.String("Address"), // Required
	}
	resp, err := svc.VerifyEmailAddress(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSES_VerifyEmailIdentity() {
	svc := ses.New(nil)

	params := &ses.VerifyEmailIdentityInput{
		EmailAddress: aws.String("Address"), // Required
	}
	resp, err := svc.VerifyEmailIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
