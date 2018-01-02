// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudformationiface provides an interface to enable mocking the AWS CloudFormation service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudformationiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

// CloudFormationAPI provides an interface to enable mocking the
// cloudformation.CloudFormation service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS CloudFormation.
//    func myFunc(svc cloudformationiface.CloudFormationAPI) bool {
//        // Make svc.CancelUpdateStack request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := cloudformation.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCloudFormationClient struct {
//        cloudformationiface.CloudFormationAPI
//    }
//    func (m *mockCloudFormationClient) CancelUpdateStack(input *cloudformation.CancelUpdateStackInput) (*cloudformation.CancelUpdateStackOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCloudFormationClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type CloudFormationAPI interface {
	CancelUpdateStackRequest(*cloudformation.CancelUpdateStackInput) cloudformation.CancelUpdateStackRequest

	ContinueUpdateRollbackRequest(*cloudformation.ContinueUpdateRollbackInput) cloudformation.ContinueUpdateRollbackRequest

	CreateChangeSetRequest(*cloudformation.CreateChangeSetInput) cloudformation.CreateChangeSetRequest

	CreateStackRequest(*cloudformation.CreateStackInput) cloudformation.CreateStackRequest

	CreateStackInstancesRequest(*cloudformation.CreateStackInstancesInput) cloudformation.CreateStackInstancesRequest

	CreateStackSetRequest(*cloudformation.CreateStackSetInput) cloudformation.CreateStackSetRequest

	DeleteChangeSetRequest(*cloudformation.DeleteChangeSetInput) cloudformation.DeleteChangeSetRequest

	DeleteStackRequest(*cloudformation.DeleteStackInput) cloudformation.DeleteStackRequest

	DeleteStackInstancesRequest(*cloudformation.DeleteStackInstancesInput) cloudformation.DeleteStackInstancesRequest

	DeleteStackSetRequest(*cloudformation.DeleteStackSetInput) cloudformation.DeleteStackSetRequest

	DescribeAccountLimitsRequest(*cloudformation.DescribeAccountLimitsInput) cloudformation.DescribeAccountLimitsRequest

	DescribeChangeSetRequest(*cloudformation.DescribeChangeSetInput) cloudformation.DescribeChangeSetRequest

	DescribeStackEventsRequest(*cloudformation.DescribeStackEventsInput) cloudformation.DescribeStackEventsRequest

	DescribeStackEventsPages(*cloudformation.DescribeStackEventsInput, func(*cloudformation.DescribeStackEventsOutput, bool) bool) error
	DescribeStackEventsPagesWithContext(aws.Context, *cloudformation.DescribeStackEventsInput, func(*cloudformation.DescribeStackEventsOutput, bool) bool, ...aws.Option) error

	DescribeStackInstanceRequest(*cloudformation.DescribeStackInstanceInput) cloudformation.DescribeStackInstanceRequest

	DescribeStackResourceRequest(*cloudformation.DescribeStackResourceInput) cloudformation.DescribeStackResourceRequest

	DescribeStackResourcesRequest(*cloudformation.DescribeStackResourcesInput) cloudformation.DescribeStackResourcesRequest

	DescribeStackSetRequest(*cloudformation.DescribeStackSetInput) cloudformation.DescribeStackSetRequest

	DescribeStackSetOperationRequest(*cloudformation.DescribeStackSetOperationInput) cloudformation.DescribeStackSetOperationRequest

	DescribeStacksRequest(*cloudformation.DescribeStacksInput) cloudformation.DescribeStacksRequest

	DescribeStacksPages(*cloudformation.DescribeStacksInput, func(*cloudformation.DescribeStacksOutput, bool) bool) error
	DescribeStacksPagesWithContext(aws.Context, *cloudformation.DescribeStacksInput, func(*cloudformation.DescribeStacksOutput, bool) bool, ...aws.Option) error

	EstimateTemplateCostRequest(*cloudformation.EstimateTemplateCostInput) cloudformation.EstimateTemplateCostRequest

	ExecuteChangeSetRequest(*cloudformation.ExecuteChangeSetInput) cloudformation.ExecuteChangeSetRequest

	GetStackPolicyRequest(*cloudformation.GetStackPolicyInput) cloudformation.GetStackPolicyRequest

	GetTemplateRequest(*cloudformation.GetTemplateInput) cloudformation.GetTemplateRequest

	GetTemplateSummaryRequest(*cloudformation.GetTemplateSummaryInput) cloudformation.GetTemplateSummaryRequest

	ListChangeSetsRequest(*cloudformation.ListChangeSetsInput) cloudformation.ListChangeSetsRequest

	ListExportsRequest(*cloudformation.ListExportsInput) cloudformation.ListExportsRequest

	ListExportsPages(*cloudformation.ListExportsInput, func(*cloudformation.ListExportsOutput, bool) bool) error
	ListExportsPagesWithContext(aws.Context, *cloudformation.ListExportsInput, func(*cloudformation.ListExportsOutput, bool) bool, ...aws.Option) error

	ListImportsRequest(*cloudformation.ListImportsInput) cloudformation.ListImportsRequest

	ListImportsPages(*cloudformation.ListImportsInput, func(*cloudformation.ListImportsOutput, bool) bool) error
	ListImportsPagesWithContext(aws.Context, *cloudformation.ListImportsInput, func(*cloudformation.ListImportsOutput, bool) bool, ...aws.Option) error

	ListStackInstancesRequest(*cloudformation.ListStackInstancesInput) cloudformation.ListStackInstancesRequest

	ListStackResourcesRequest(*cloudformation.ListStackResourcesInput) cloudformation.ListStackResourcesRequest

	ListStackResourcesPages(*cloudformation.ListStackResourcesInput, func(*cloudformation.ListStackResourcesOutput, bool) bool) error
	ListStackResourcesPagesWithContext(aws.Context, *cloudformation.ListStackResourcesInput, func(*cloudformation.ListStackResourcesOutput, bool) bool, ...aws.Option) error

	ListStackSetOperationResultsRequest(*cloudformation.ListStackSetOperationResultsInput) cloudformation.ListStackSetOperationResultsRequest

	ListStackSetOperationsRequest(*cloudformation.ListStackSetOperationsInput) cloudformation.ListStackSetOperationsRequest

	ListStackSetsRequest(*cloudformation.ListStackSetsInput) cloudformation.ListStackSetsRequest

	ListStacksRequest(*cloudformation.ListStacksInput) cloudformation.ListStacksRequest

	ListStacksPages(*cloudformation.ListStacksInput, func(*cloudformation.ListStacksOutput, bool) bool) error
	ListStacksPagesWithContext(aws.Context, *cloudformation.ListStacksInput, func(*cloudformation.ListStacksOutput, bool) bool, ...aws.Option) error

	SetStackPolicyRequest(*cloudformation.SetStackPolicyInput) cloudformation.SetStackPolicyRequest

	SignalResourceRequest(*cloudformation.SignalResourceInput) cloudformation.SignalResourceRequest

	StopStackSetOperationRequest(*cloudformation.StopStackSetOperationInput) cloudformation.StopStackSetOperationRequest

	UpdateStackRequest(*cloudformation.UpdateStackInput) cloudformation.UpdateStackRequest

	UpdateStackInstancesRequest(*cloudformation.UpdateStackInstancesInput) cloudformation.UpdateStackInstancesRequest

	UpdateStackSetRequest(*cloudformation.UpdateStackSetInput) cloudformation.UpdateStackSetRequest

	UpdateTerminationProtectionRequest(*cloudformation.UpdateTerminationProtectionInput) cloudformation.UpdateTerminationProtectionRequest

	ValidateTemplateRequest(*cloudformation.ValidateTemplateInput) cloudformation.ValidateTemplateRequest

	WaitUntilChangeSetCreateComplete(*cloudformation.DescribeChangeSetInput) error
	WaitUntilChangeSetCreateCompleteWithContext(aws.Context, *cloudformation.DescribeChangeSetInput, ...aws.WaiterOption) error

	WaitUntilStackCreateComplete(*cloudformation.DescribeStacksInput) error
	WaitUntilStackCreateCompleteWithContext(aws.Context, *cloudformation.DescribeStacksInput, ...aws.WaiterOption) error

	WaitUntilStackDeleteComplete(*cloudformation.DescribeStacksInput) error
	WaitUntilStackDeleteCompleteWithContext(aws.Context, *cloudformation.DescribeStacksInput, ...aws.WaiterOption) error

	WaitUntilStackExists(*cloudformation.DescribeStacksInput) error
	WaitUntilStackExistsWithContext(aws.Context, *cloudformation.DescribeStacksInput, ...aws.WaiterOption) error

	WaitUntilStackUpdateComplete(*cloudformation.DescribeStacksInput) error
	WaitUntilStackUpdateCompleteWithContext(aws.Context, *cloudformation.DescribeStacksInput, ...aws.WaiterOption) error
}

var _ CloudFormationAPI = (*cloudformation.CloudFormation)(nil)
