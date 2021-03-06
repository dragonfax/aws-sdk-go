// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package datapipeline_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/dragonfax/aws-sdk-go/aws"
	"github.com/dragonfax/aws-sdk-go/service/datapipeline"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleDataPipeline_ActivatePipeline() {
	svc := datapipeline.New(nil)

	params := &datapipeline.ActivatePipelineInput{
		PipelineId: aws.String("id"), // Required
		ParameterValues: []*datapipeline.ParameterValue{
			{ // Required
				Id:          aws.String("fieldNameString"),  // Required
				StringValue: aws.String("fieldStringValue"), // Required
			},
			// More values...
		},
		StartTimestamp: aws.Time(time.Now()),
	}
	resp, err := svc.ActivatePipeline(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDataPipeline_AddTags() {
	svc := datapipeline.New(nil)

	params := &datapipeline.AddTagsInput{
		PipelineId: aws.String("id"), // Required
		Tags: []*datapipeline.Tag{ // Required
			{ // Required
				Key:   aws.String("tagKey"),   // Required
				Value: aws.String("tagValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.AddTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDataPipeline_CreatePipeline() {
	svc := datapipeline.New(nil)

	params := &datapipeline.CreatePipelineInput{
		Name:        aws.String("id"), // Required
		UniqueId:    aws.String("id"), // Required
		Description: aws.String("string"),
		Tags: []*datapipeline.Tag{
			{ // Required
				Key:   aws.String("tagKey"),   // Required
				Value: aws.String("tagValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.CreatePipeline(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDataPipeline_DeactivatePipeline() {
	svc := datapipeline.New(nil)

	params := &datapipeline.DeactivatePipelineInput{
		PipelineId:   aws.String("id"), // Required
		CancelActive: aws.Bool(true),
	}
	resp, err := svc.DeactivatePipeline(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDataPipeline_DeletePipeline() {
	svc := datapipeline.New(nil)

	params := &datapipeline.DeletePipelineInput{
		PipelineId: aws.String("id"), // Required
	}
	resp, err := svc.DeletePipeline(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDataPipeline_DescribeObjects() {
	svc := datapipeline.New(nil)

	params := &datapipeline.DescribeObjectsInput{
		ObjectIds: []*string{ // Required
			aws.String("id"), // Required
			// More values...
		},
		PipelineId:          aws.String("id"), // Required
		EvaluateExpressions: aws.Bool(true),
		Marker:              aws.String("string"),
	}
	resp, err := svc.DescribeObjects(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDataPipeline_DescribePipelines() {
	svc := datapipeline.New(nil)

	params := &datapipeline.DescribePipelinesInput{
		PipelineIds: []*string{ // Required
			aws.String("id"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribePipelines(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDataPipeline_EvaluateExpression() {
	svc := datapipeline.New(nil)

	params := &datapipeline.EvaluateExpressionInput{
		Expression: aws.String("longString"), // Required
		ObjectId:   aws.String("id"),         // Required
		PipelineId: aws.String("id"),         // Required
	}
	resp, err := svc.EvaluateExpression(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDataPipeline_GetPipelineDefinition() {
	svc := datapipeline.New(nil)

	params := &datapipeline.GetPipelineDefinitionInput{
		PipelineId: aws.String("id"), // Required
		Version:    aws.String("string"),
	}
	resp, err := svc.GetPipelineDefinition(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDataPipeline_ListPipelines() {
	svc := datapipeline.New(nil)

	params := &datapipeline.ListPipelinesInput{
		Marker: aws.String("string"),
	}
	resp, err := svc.ListPipelines(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDataPipeline_PollForTask() {
	svc := datapipeline.New(nil)

	params := &datapipeline.PollForTaskInput{
		WorkerGroup: aws.String("string"), // Required
		Hostname:    aws.String("id"),
		InstanceIdentity: &datapipeline.InstanceIdentity{
			Document:  aws.String("string"),
			Signature: aws.String("string"),
		},
	}
	resp, err := svc.PollForTask(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDataPipeline_PutPipelineDefinition() {
	svc := datapipeline.New(nil)

	params := &datapipeline.PutPipelineDefinitionInput{
		PipelineId: aws.String("id"), // Required
		PipelineObjects: []*datapipeline.PipelineObject{ // Required
			{ // Required
				Fields: []*datapipeline.Field{ // Required
					{ // Required
						Key:         aws.String("fieldNameString"), // Required
						RefValue:    aws.String("fieldNameString"),
						StringValue: aws.String("fieldStringValue"),
					},
					// More values...
				},
				Id:   aws.String("id"), // Required
				Name: aws.String("id"), // Required
			},
			// More values...
		},
		ParameterObjects: []*datapipeline.ParameterObject{
			{ // Required
				Attributes: []*datapipeline.ParameterAttribute{ // Required
					{ // Required
						Key:         aws.String("attributeNameString"),  // Required
						StringValue: aws.String("attributeValueString"), // Required
					},
					// More values...
				},
				Id: aws.String("fieldNameString"), // Required
			},
			// More values...
		},
		ParameterValues: []*datapipeline.ParameterValue{
			{ // Required
				Id:          aws.String("fieldNameString"),  // Required
				StringValue: aws.String("fieldStringValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.PutPipelineDefinition(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDataPipeline_QueryObjects() {
	svc := datapipeline.New(nil)

	params := &datapipeline.QueryObjectsInput{
		PipelineId: aws.String("id"),     // Required
		Sphere:     aws.String("string"), // Required
		Limit:      aws.Int64(1),
		Marker:     aws.String("string"),
		Query: &datapipeline.Query{
			Selectors: []*datapipeline.Selector{
				{ // Required
					FieldName: aws.String("string"),
					Operator: &datapipeline.Operator{
						Type: aws.String("OperatorType"),
						Values: []*string{
							aws.String("string"), // Required
							// More values...
						},
					},
				},
				// More values...
			},
		},
	}
	resp, err := svc.QueryObjects(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDataPipeline_RemoveTags() {
	svc := datapipeline.New(nil)

	params := &datapipeline.RemoveTagsInput{
		PipelineId: aws.String("id"), // Required
		TagKeys: []*string{ // Required
			aws.String("string"), // Required
			// More values...
		},
	}
	resp, err := svc.RemoveTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDataPipeline_ReportTaskProgress() {
	svc := datapipeline.New(nil)

	params := &datapipeline.ReportTaskProgressInput{
		TaskId: aws.String("taskId"), // Required
		Fields: []*datapipeline.Field{
			{ // Required
				Key:         aws.String("fieldNameString"), // Required
				RefValue:    aws.String("fieldNameString"),
				StringValue: aws.String("fieldStringValue"),
			},
			// More values...
		},
	}
	resp, err := svc.ReportTaskProgress(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDataPipeline_ReportTaskRunnerHeartbeat() {
	svc := datapipeline.New(nil)

	params := &datapipeline.ReportTaskRunnerHeartbeatInput{
		TaskrunnerId: aws.String("id"), // Required
		Hostname:     aws.String("id"),
		WorkerGroup:  aws.String("string"),
	}
	resp, err := svc.ReportTaskRunnerHeartbeat(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDataPipeline_SetStatus() {
	svc := datapipeline.New(nil)

	params := &datapipeline.SetStatusInput{
		ObjectIds: []*string{ // Required
			aws.String("id"), // Required
			// More values...
		},
		PipelineId: aws.String("id"),     // Required
		Status:     aws.String("string"), // Required
	}
	resp, err := svc.SetStatus(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDataPipeline_SetTaskStatus() {
	svc := datapipeline.New(nil)

	params := &datapipeline.SetTaskStatusInput{
		TaskId:          aws.String("taskId"),     // Required
		TaskStatus:      aws.String("TaskStatus"), // Required
		ErrorId:         aws.String("string"),
		ErrorMessage:    aws.String("errorMessage"),
		ErrorStackTrace: aws.String("string"),
	}
	resp, err := svc.SetTaskStatus(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDataPipeline_ValidatePipelineDefinition() {
	svc := datapipeline.New(nil)

	params := &datapipeline.ValidatePipelineDefinitionInput{
		PipelineId: aws.String("id"), // Required
		PipelineObjects: []*datapipeline.PipelineObject{ // Required
			{ // Required
				Fields: []*datapipeline.Field{ // Required
					{ // Required
						Key:         aws.String("fieldNameString"), // Required
						RefValue:    aws.String("fieldNameString"),
						StringValue: aws.String("fieldStringValue"),
					},
					// More values...
				},
				Id:   aws.String("id"), // Required
				Name: aws.String("id"), // Required
			},
			// More values...
		},
		ParameterObjects: []*datapipeline.ParameterObject{
			{ // Required
				Attributes: []*datapipeline.ParameterAttribute{ // Required
					{ // Required
						Key:         aws.String("attributeNameString"),  // Required
						StringValue: aws.String("attributeValueString"), // Required
					},
					// More values...
				},
				Id: aws.String("fieldNameString"), // Required
			},
			// More values...
		},
		ParameterValues: []*datapipeline.ParameterValue{
			{ // Required
				Id:          aws.String("fieldNameString"),  // Required
				StringValue: aws.String("fieldStringValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.ValidatePipelineDefinition(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
