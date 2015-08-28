// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package codecommit provides a client for AWS CodeCommit.
package codecommit

import (
	"time"

	"github.com/dragonfax/aws-sdk-go/aws/awsutil"
	"github.com/dragonfax/aws-sdk-go/aws/request"
)

const opBatchGetRepositories = "BatchGetRepositories"

// BatchGetRepositoriesRequest generates a request for the BatchGetRepositories operation.
func (c *CodeCommit) BatchGetRepositoriesRequest(input *BatchGetRepositoriesInput) (req *request.Request, output *BatchGetRepositoriesOutput) {
	op := &request.Operation{
		Name:       opBatchGetRepositories,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchGetRepositoriesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &BatchGetRepositoriesOutput{}
	req.Data = output
	return
}

// Gets information about one or more repositories.
//
// The description field for a repository accepts all HTML characters and all
// valid Unicode characters. Applications that do not HTML-encode the description
// and display it in a web page could expose users to potentially malicious
// code. Make sure that you HTML-encode the description field in any application
// that uses this API to display the repository description on a web page.
func (c *CodeCommit) BatchGetRepositories(input *BatchGetRepositoriesInput) (*BatchGetRepositoriesOutput, error) {
	req, out := c.BatchGetRepositoriesRequest(input)
	err := req.Send()
	return out, err
}

const opCreateBranch = "CreateBranch"

// CreateBranchRequest generates a request for the CreateBranch operation.
func (c *CodeCommit) CreateBranchRequest(input *CreateBranchInput) (req *request.Request, output *CreateBranchOutput) {
	op := &request.Operation{
		Name:       opCreateBranch,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateBranchInput{}
	}

	req = c.newRequest(op, input, output)
	output = &CreateBranchOutput{}
	req.Data = output
	return
}

// Creates a new branch in a repository and points the branch to a commit.
//
// Calling the create branch operation does not set a repository's default
// branch. To do this, call the update default branch operation.
func (c *CodeCommit) CreateBranch(input *CreateBranchInput) (*CreateBranchOutput, error) {
	req, out := c.CreateBranchRequest(input)
	err := req.Send()
	return out, err
}

const opCreateRepository = "CreateRepository"

// CreateRepositoryRequest generates a request for the CreateRepository operation.
func (c *CodeCommit) CreateRepositoryRequest(input *CreateRepositoryInput) (req *request.Request, output *CreateRepositoryOutput) {
	op := &request.Operation{
		Name:       opCreateRepository,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateRepositoryInput{}
	}

	req = c.newRequest(op, input, output)
	output = &CreateRepositoryOutput{}
	req.Data = output
	return
}

// Creates a new, empty repository.
func (c *CodeCommit) CreateRepository(input *CreateRepositoryInput) (*CreateRepositoryOutput, error) {
	req, out := c.CreateRepositoryRequest(input)
	err := req.Send()
	return out, err
}

const opDeleteRepository = "DeleteRepository"

// DeleteRepositoryRequest generates a request for the DeleteRepository operation.
func (c *CodeCommit) DeleteRepositoryRequest(input *DeleteRepositoryInput) (req *request.Request, output *DeleteRepositoryOutput) {
	op := &request.Operation{
		Name:       opDeleteRepository,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRepositoryInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DeleteRepositoryOutput{}
	req.Data = output
	return
}

// Deletes a repository. If a specified repository was already deleted, a null
// repository ID will be returned.
//
// Deleting a repository also deletes all associated objects and metadata.
// After a repository is deleted, all future push calls to the deleted repository
// will fail.
func (c *CodeCommit) DeleteRepository(input *DeleteRepositoryInput) (*DeleteRepositoryOutput, error) {
	req, out := c.DeleteRepositoryRequest(input)
	err := req.Send()
	return out, err
}

const opGetBranch = "GetBranch"

// GetBranchRequest generates a request for the GetBranch operation.
func (c *CodeCommit) GetBranchRequest(input *GetBranchInput) (req *request.Request, output *GetBranchOutput) {
	op := &request.Operation{
		Name:       opGetBranch,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetBranchInput{}
	}

	req = c.newRequest(op, input, output)
	output = &GetBranchOutput{}
	req.Data = output
	return
}

// Retrieves information about a repository branch, including its name and the
// last commit ID.
func (c *CodeCommit) GetBranch(input *GetBranchInput) (*GetBranchOutput, error) {
	req, out := c.GetBranchRequest(input)
	err := req.Send()
	return out, err
}

const opGetRepository = "GetRepository"

// GetRepositoryRequest generates a request for the GetRepository operation.
func (c *CodeCommit) GetRepositoryRequest(input *GetRepositoryInput) (req *request.Request, output *GetRepositoryOutput) {
	op := &request.Operation{
		Name:       opGetRepository,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRepositoryInput{}
	}

	req = c.newRequest(op, input, output)
	output = &GetRepositoryOutput{}
	req.Data = output
	return
}

// Gets information about a repository.
//
// The description field for a repository accepts all HTML characters and all
// valid Unicode characters. Applications that do not HTML-encode the description
// and display it in a web page could expose users to potentially malicious
// code. Make sure that you HTML-encode the description field in any application
// that uses this API to display the repository description on a web page.
func (c *CodeCommit) GetRepository(input *GetRepositoryInput) (*GetRepositoryOutput, error) {
	req, out := c.GetRepositoryRequest(input)
	err := req.Send()
	return out, err
}

const opListBranches = "ListBranches"

// ListBranchesRequest generates a request for the ListBranches operation.
func (c *CodeCommit) ListBranchesRequest(input *ListBranchesInput) (req *request.Request, output *ListBranchesOutput) {
	op := &request.Operation{
		Name:       opListBranches,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListBranchesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ListBranchesOutput{}
	req.Data = output
	return
}

// Gets information about one or more branches in a repository.
func (c *CodeCommit) ListBranches(input *ListBranchesInput) (*ListBranchesOutput, error) {
	req, out := c.ListBranchesRequest(input)
	err := req.Send()
	return out, err
}

const opListRepositories = "ListRepositories"

// ListRepositoriesRequest generates a request for the ListRepositories operation.
func (c *CodeCommit) ListRepositoriesRequest(input *ListRepositoriesInput) (req *request.Request, output *ListRepositoriesOutput) {
	op := &request.Operation{
		Name:       opListRepositories,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListRepositoriesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ListRepositoriesOutput{}
	req.Data = output
	return
}

// Gets information about one or more repositories.
func (c *CodeCommit) ListRepositories(input *ListRepositoriesInput) (*ListRepositoriesOutput, error) {
	req, out := c.ListRepositoriesRequest(input)
	err := req.Send()
	return out, err
}

const opUpdateDefaultBranch = "UpdateDefaultBranch"

// UpdateDefaultBranchRequest generates a request for the UpdateDefaultBranch operation.
func (c *CodeCommit) UpdateDefaultBranchRequest(input *UpdateDefaultBranchInput) (req *request.Request, output *UpdateDefaultBranchOutput) {
	op := &request.Operation{
		Name:       opUpdateDefaultBranch,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateDefaultBranchInput{}
	}

	req = c.newRequest(op, input, output)
	output = &UpdateDefaultBranchOutput{}
	req.Data = output
	return
}

// Sets or changes the default branch name for the specified repository.
//
// If you use this operation to change the default branch name to the current
// default branch name, a success message is returned even though the default
// branch did not change.
func (c *CodeCommit) UpdateDefaultBranch(input *UpdateDefaultBranchInput) (*UpdateDefaultBranchOutput, error) {
	req, out := c.UpdateDefaultBranchRequest(input)
	err := req.Send()
	return out, err
}

const opUpdateRepositoryDescription = "UpdateRepositoryDescription"

// UpdateRepositoryDescriptionRequest generates a request for the UpdateRepositoryDescription operation.
func (c *CodeCommit) UpdateRepositoryDescriptionRequest(input *UpdateRepositoryDescriptionInput) (req *request.Request, output *UpdateRepositoryDescriptionOutput) {
	op := &request.Operation{
		Name:       opUpdateRepositoryDescription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateRepositoryDescriptionInput{}
	}

	req = c.newRequest(op, input, output)
	output = &UpdateRepositoryDescriptionOutput{}
	req.Data = output
	return
}

// Sets or changes the comment or description for a repository.
//
// The description field for a repository accepts all HTML characters and all
// valid Unicode characters. Applications that do not HTML-encode the description
// and display it in a web page could expose users to potentially malicious
// code. Make sure that you HTML-encode the description field in any application
// that uses this API to display the repository description on a web page.
func (c *CodeCommit) UpdateRepositoryDescription(input *UpdateRepositoryDescriptionInput) (*UpdateRepositoryDescriptionOutput, error) {
	req, out := c.UpdateRepositoryDescriptionRequest(input)
	err := req.Send()
	return out, err
}

const opUpdateRepositoryName = "UpdateRepositoryName"

// UpdateRepositoryNameRequest generates a request for the UpdateRepositoryName operation.
func (c *CodeCommit) UpdateRepositoryNameRequest(input *UpdateRepositoryNameInput) (req *request.Request, output *UpdateRepositoryNameOutput) {
	op := &request.Operation{
		Name:       opUpdateRepositoryName,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateRepositoryNameInput{}
	}

	req = c.newRequest(op, input, output)
	output = &UpdateRepositoryNameOutput{}
	req.Data = output
	return
}

// Renames a repository.
func (c *CodeCommit) UpdateRepositoryName(input *UpdateRepositoryNameInput) (*UpdateRepositoryNameOutput, error) {
	req, out := c.UpdateRepositoryNameRequest(input)
	err := req.Send()
	return out, err
}

// Represents the input of a batch get repositories operation.
type BatchGetRepositoriesInput struct {
	// The names of the repositories to get information about.
	RepositoryNames []*string `locationName:"repositoryNames" type:"list" required:"true"`

	metadataBatchGetRepositoriesInput `json:"-" xml:"-"`
}

type metadataBatchGetRepositoriesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s BatchGetRepositoriesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s BatchGetRepositoriesInput) GoString() string {
	return s.String()
}

// Represents the output of a batch get repositories operation.
type BatchGetRepositoriesOutput struct {
	// A list of repositories returned by the batch get repositories operation.
	Repositories []*RepositoryMetadata `locationName:"repositories" type:"list"`

	// Returns a list of repository names for which information could not be found.
	RepositoriesNotFound []*string `locationName:"repositoriesNotFound" type:"list"`

	metadataBatchGetRepositoriesOutput `json:"-" xml:"-"`
}

type metadataBatchGetRepositoriesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s BatchGetRepositoriesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s BatchGetRepositoriesOutput) GoString() string {
	return s.String()
}

// Returns information about a branch.
type BranchInfo struct {
	// The name of the branch.
	BranchName *string `locationName:"branchName" type:"string"`

	// The ID of the last commit made to the branch.
	CommitId *string `locationName:"commitId" type:"string"`

	metadataBranchInfo `json:"-" xml:"-"`
}

type metadataBranchInfo struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s BranchInfo) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s BranchInfo) GoString() string {
	return s.String()
}

// Represents the input of a create branch operation.
type CreateBranchInput struct {
	// The name of the new branch to create.
	BranchName *string `locationName:"branchName" type:"string" required:"true"`

	// The ID of the commit to point the new branch to.
	//
	// If this commit ID is not specified, the new branch will point to the commit
	// that is pointed to by the repository's default branch.
	CommitId *string `locationName:"commitId" type:"string" required:"true"`

	// The name of the repository in which you want to create the new branch.
	RepositoryName *string `locationName:"repositoryName" type:"string" required:"true"`

	metadataCreateBranchInput `json:"-" xml:"-"`
}

type metadataCreateBranchInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s CreateBranchInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateBranchInput) GoString() string {
	return s.String()
}

type CreateBranchOutput struct {
	metadataCreateBranchOutput `json:"-" xml:"-"`
}

type metadataCreateBranchOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s CreateBranchOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateBranchOutput) GoString() string {
	return s.String()
}

// Represents the input of a create repository operation.
type CreateRepositoryInput struct {
	// A comment or description about the new repository.
	RepositoryDescription *string `locationName:"repositoryDescription" type:"string"`

	// The name of the new repository to be created.
	//
	// The repository name must be unique across the calling AWS account. In addition,
	// repository names are restricted to alphanumeric characters. The suffix ".git"
	// is prohibited.
	RepositoryName *string `locationName:"repositoryName" type:"string" required:"true"`

	metadataCreateRepositoryInput `json:"-" xml:"-"`
}

type metadataCreateRepositoryInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s CreateRepositoryInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateRepositoryInput) GoString() string {
	return s.String()
}

// Represents the output of a create repository operation.
type CreateRepositoryOutput struct {
	// Information about the newly created repository.
	RepositoryMetadata *RepositoryMetadata `locationName:"repositoryMetadata" type:"structure"`

	metadataCreateRepositoryOutput `json:"-" xml:"-"`
}

type metadataCreateRepositoryOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s CreateRepositoryOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateRepositoryOutput) GoString() string {
	return s.String()
}

// Represents the input of a delete repository operation.
type DeleteRepositoryInput struct {
	// The name of the repository to delete.
	RepositoryName *string `locationName:"repositoryName" type:"string" required:"true"`

	metadataDeleteRepositoryInput `json:"-" xml:"-"`
}

type metadataDeleteRepositoryInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DeleteRepositoryInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteRepositoryInput) GoString() string {
	return s.String()
}

// Represents the output of a delete repository operation.
type DeleteRepositoryOutput struct {
	// The ID of the repository that was deleted.
	RepositoryId *string `locationName:"repositoryId" type:"string"`

	metadataDeleteRepositoryOutput `json:"-" xml:"-"`
}

type metadataDeleteRepositoryOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DeleteRepositoryOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteRepositoryOutput) GoString() string {
	return s.String()
}

// Represents the input of a get branch operation.
type GetBranchInput struct {
	// The name of the branch for which you want to retrieve information.
	BranchName *string `locationName:"branchName" type:"string"`

	// Repository name is restricted to alphanumeric characters (a-z, A-Z, 0-9),
	// ".", "_", and "-". Additionally, the suffix ".git" is prohibited in a repository
	// name.
	RepositoryName *string `locationName:"repositoryName" type:"string"`

	metadataGetBranchInput `json:"-" xml:"-"`
}

type metadataGetBranchInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s GetBranchInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetBranchInput) GoString() string {
	return s.String()
}

// Represents the output of a get branch operation.
type GetBranchOutput struct {
	// The name of the branch.
	Branch *BranchInfo `locationName:"branch" type:"structure"`

	metadataGetBranchOutput `json:"-" xml:"-"`
}

type metadataGetBranchOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s GetBranchOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetBranchOutput) GoString() string {
	return s.String()
}

// Represents the input of a get repository operation.
type GetRepositoryInput struct {
	// The name of the repository to get information about.
	RepositoryName *string `locationName:"repositoryName" type:"string" required:"true"`

	metadataGetRepositoryInput `json:"-" xml:"-"`
}

type metadataGetRepositoryInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s GetRepositoryInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRepositoryInput) GoString() string {
	return s.String()
}

// Represents the output of a get repository operation.
type GetRepositoryOutput struct {
	// Information about the repository.
	RepositoryMetadata *RepositoryMetadata `locationName:"repositoryMetadata" type:"structure"`

	metadataGetRepositoryOutput `json:"-" xml:"-"`
}

type metadataGetRepositoryOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s GetRepositoryOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRepositoryOutput) GoString() string {
	return s.String()
}

// Represents the input of a list branches operation.
type ListBranchesInput struct {
	// An enumeration token that allows the operation to batch the results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The name of the repository that contains the branches.
	RepositoryName *string `locationName:"repositoryName" type:"string" required:"true"`

	metadataListBranchesInput `json:"-" xml:"-"`
}

type metadataListBranchesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListBranchesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListBranchesInput) GoString() string {
	return s.String()
}

// Represents the output of a list branches operation.
type ListBranchesOutput struct {
	// The list of branch names.
	Branches []*string `locationName:"branches" type:"list"`

	// An enumeration token that returns the batch of the results.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataListBranchesOutput `json:"-" xml:"-"`
}

type metadataListBranchesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListBranchesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListBranchesOutput) GoString() string {
	return s.String()
}

// Represents the input of a list repositories operation.
type ListRepositoriesInput struct {
	// An enumeration token that allows the operation to batch the results of the
	// operation. Batch sizes are 1,000 for list repository operations. When the
	// client sends the token back to AWS CodeCommit, another page of 1,000 records
	// is retrieved.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The order in which to sort the results of a list repositories operation.
	Order *string `locationName:"order" type:"string" enum:"OrderEnum"`

	// The criteria used to sort the results of a list repositories operation.
	SortBy *string `locationName:"sortBy" type:"string" enum:"SortByEnum"`

	metadataListRepositoriesInput `json:"-" xml:"-"`
}

type metadataListRepositoriesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListRepositoriesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListRepositoriesInput) GoString() string {
	return s.String()
}

// Represents the output of a list repositories operation.
type ListRepositoriesOutput struct {
	// An enumeration token that allows the operation to batch the results of the
	// operation. Batch sizes are 1,000 for list repository operations. When the
	// client sends the token back to AWS CodeCommit, another page of 1,000 records
	// is retrieved.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Lists the repositories called by the list repositories operation.
	Repositories []*RepositoryNameIdPair `locationName:"repositories" type:"list"`

	metadataListRepositoriesOutput `json:"-" xml:"-"`
}

type metadataListRepositoriesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListRepositoriesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListRepositoriesOutput) GoString() string {
	return s.String()
}

// Information about a repository.
type RepositoryMetadata struct {
	// The ID of the AWS account associated with the repository.
	AccountId *string `locationName:"accountId" type:"string"`

	// The Amazon Resource Name (ARN) of the repository.
	Arn *string `type:"string"`

	// The URL to use for cloning the repository over HTTPS.
	CloneUrlHttp *string `locationName:"cloneUrlHttp" type:"string"`

	// The URL to use for cloning the repository over SSH.
	CloneUrlSsh *string `locationName:"cloneUrlSsh" type:"string"`

	// The date and time the repository was created, in timestamp format.
	CreationDate *time.Time `locationName:"creationDate" type:"timestamp" timestampFormat:"unix"`

	// The repository's default branch name.
	DefaultBranch *string `locationName:"defaultBranch" type:"string"`

	// The date and time the repository was last modified, in timestamp format.
	LastModifiedDate *time.Time `locationName:"lastModifiedDate" type:"timestamp" timestampFormat:"unix"`

	// A comment or description about the repository.
	RepositoryDescription *string `locationName:"repositoryDescription" type:"string"`

	// The ID of the repository.
	RepositoryId *string `locationName:"repositoryId" type:"string"`

	// The repository's name.
	RepositoryName *string `locationName:"repositoryName" type:"string"`

	metadataRepositoryMetadata `json:"-" xml:"-"`
}

type metadataRepositoryMetadata struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s RepositoryMetadata) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RepositoryMetadata) GoString() string {
	return s.String()
}

// Information about a repository name and ID.
type RepositoryNameIdPair struct {
	// The ID associated with the repository name.
	RepositoryId *string `locationName:"repositoryId" type:"string"`

	// Repository name is restricted to alphanumeric characters (a-z, A-Z, 0-9),
	// ".", "_", and "-". Additionally, the suffix ".git" is prohibited in a repository
	// name.
	RepositoryName *string `locationName:"repositoryName" type:"string"`

	metadataRepositoryNameIdPair `json:"-" xml:"-"`
}

type metadataRepositoryNameIdPair struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s RepositoryNameIdPair) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RepositoryNameIdPair) GoString() string {
	return s.String()
}

// Represents the input of an update default branch operation.
type UpdateDefaultBranchInput struct {
	// The name of the branch to set as the default.
	DefaultBranchName *string `locationName:"defaultBranchName" type:"string" required:"true"`

	// The name of the repository to set or change the default branch for.
	RepositoryName *string `locationName:"repositoryName" type:"string" required:"true"`

	metadataUpdateDefaultBranchInput `json:"-" xml:"-"`
}

type metadataUpdateDefaultBranchInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s UpdateDefaultBranchInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateDefaultBranchInput) GoString() string {
	return s.String()
}

type UpdateDefaultBranchOutput struct {
	metadataUpdateDefaultBranchOutput `json:"-" xml:"-"`
}

type metadataUpdateDefaultBranchOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s UpdateDefaultBranchOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateDefaultBranchOutput) GoString() string {
	return s.String()
}

// Represents the input of an update repository description operation.
type UpdateRepositoryDescriptionInput struct {
	// The new comment or description for the specified repository.
	RepositoryDescription *string `locationName:"repositoryDescription" type:"string"`

	// The name of the repository to set or change the comment or description for.
	RepositoryName *string `locationName:"repositoryName" type:"string" required:"true"`

	metadataUpdateRepositoryDescriptionInput `json:"-" xml:"-"`
}

type metadataUpdateRepositoryDescriptionInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s UpdateRepositoryDescriptionInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateRepositoryDescriptionInput) GoString() string {
	return s.String()
}

type UpdateRepositoryDescriptionOutput struct {
	metadataUpdateRepositoryDescriptionOutput `json:"-" xml:"-"`
}

type metadataUpdateRepositoryDescriptionOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s UpdateRepositoryDescriptionOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateRepositoryDescriptionOutput) GoString() string {
	return s.String()
}

// Represents the input of an update repository description operation.
type UpdateRepositoryNameInput struct {
	// Repository name is restricted to alphanumeric characters (a-z, A-Z, 0-9),
	// ".", "_", and "-". Additionally, the suffix ".git" is prohibited in a repository
	// name.
	NewName *string `locationName:"newName" type:"string" required:"true"`

	// Repository name is restricted to alphanumeric characters (a-z, A-Z, 0-9),
	// ".", "_", and "-". Additionally, the suffix ".git" is prohibited in a repository
	// name.
	OldName *string `locationName:"oldName" type:"string" required:"true"`

	metadataUpdateRepositoryNameInput `json:"-" xml:"-"`
}

type metadataUpdateRepositoryNameInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s UpdateRepositoryNameInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateRepositoryNameInput) GoString() string {
	return s.String()
}

type UpdateRepositoryNameOutput struct {
	metadataUpdateRepositoryNameOutput `json:"-" xml:"-"`
}

type metadataUpdateRepositoryNameOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s UpdateRepositoryNameOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateRepositoryNameOutput) GoString() string {
	return s.String()
}

const (
	// @enum OrderEnum
	OrderEnumAscending = "ascending"
	// @enum OrderEnum
	OrderEnumDescending = "descending"
)

const (
	// @enum SortByEnum
	SortByEnumRepositoryName = "repositoryName"
	// @enum SortByEnum
	SortByEnumLastModifiedDate = "lastModifiedDate"
)
