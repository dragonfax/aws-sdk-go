// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package workspacesiface provides an interface for the Amazon WorkSpaces.
package workspacesiface

import (
	"github.com/dragonfax/aws-sdk-go/aws/request"
	"github.com/dragonfax/aws-sdk-go/service/workspaces"
)

// WorkSpacesAPI is the interface type for workspaces.WorkSpaces.
type WorkSpacesAPI interface {
	CreateWorkspacesRequest(*workspaces.CreateWorkspacesInput) (*request.Request, *workspaces.CreateWorkspacesOutput)

	CreateWorkspaces(*workspaces.CreateWorkspacesInput) (*workspaces.CreateWorkspacesOutput, error)

	DescribeWorkspaceBundlesRequest(*workspaces.DescribeWorkspaceBundlesInput) (*request.Request, *workspaces.DescribeWorkspaceBundlesOutput)

	DescribeWorkspaceBundles(*workspaces.DescribeWorkspaceBundlesInput) (*workspaces.DescribeWorkspaceBundlesOutput, error)

	DescribeWorkspaceBundlesPages(*workspaces.DescribeWorkspaceBundlesInput, func(*workspaces.DescribeWorkspaceBundlesOutput, bool) bool) error

	DescribeWorkspaceDirectoriesRequest(*workspaces.DescribeWorkspaceDirectoriesInput) (*request.Request, *workspaces.DescribeWorkspaceDirectoriesOutput)

	DescribeWorkspaceDirectories(*workspaces.DescribeWorkspaceDirectoriesInput) (*workspaces.DescribeWorkspaceDirectoriesOutput, error)

	DescribeWorkspaceDirectoriesPages(*workspaces.DescribeWorkspaceDirectoriesInput, func(*workspaces.DescribeWorkspaceDirectoriesOutput, bool) bool) error

	DescribeWorkspacesRequest(*workspaces.DescribeWorkspacesInput) (*request.Request, *workspaces.DescribeWorkspacesOutput)

	DescribeWorkspaces(*workspaces.DescribeWorkspacesInput) (*workspaces.DescribeWorkspacesOutput, error)

	DescribeWorkspacesPages(*workspaces.DescribeWorkspacesInput, func(*workspaces.DescribeWorkspacesOutput, bool) bool) error

	RebootWorkspacesRequest(*workspaces.RebootWorkspacesInput) (*request.Request, *workspaces.RebootWorkspacesOutput)

	RebootWorkspaces(*workspaces.RebootWorkspacesInput) (*workspaces.RebootWorkspacesOutput, error)

	RebuildWorkspacesRequest(*workspaces.RebuildWorkspacesInput) (*request.Request, *workspaces.RebuildWorkspacesOutput)

	RebuildWorkspaces(*workspaces.RebuildWorkspacesInput) (*workspaces.RebuildWorkspacesOutput, error)

	TerminateWorkspacesRequest(*workspaces.TerminateWorkspacesInput) (*request.Request, *workspaces.TerminateWorkspacesOutput)

	TerminateWorkspaces(*workspaces.TerminateWorkspacesInput) (*workspaces.TerminateWorkspacesOutput, error)
}
