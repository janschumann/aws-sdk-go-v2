// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package appstreamiface provides an interface to enable mocking the Amazon AppStream service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package appstreamiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
)

// AppStreamAPI provides an interface to enable mocking the
// appstream.AppStream service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon AppStream.
//    func myFunc(svc appstreamiface.AppStreamAPI) bool {
//        // Make svc.AssociateFleet request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := appstream.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockAppStreamClient struct {
//        appstreamiface.AppStreamAPI
//    }
//    func (m *mockAppStreamClient) AssociateFleet(input *appstream.AssociateFleetInput) (*appstream.AssociateFleetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockAppStreamClient{}
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
type AppStreamAPI interface {
	AssociateFleetRequest(*appstream.AssociateFleetInput) appstream.AssociateFleetRequest

	CreateDirectoryConfigRequest(*appstream.CreateDirectoryConfigInput) appstream.CreateDirectoryConfigRequest

	CreateFleetRequest(*appstream.CreateFleetInput) appstream.CreateFleetRequest

	CreateImageBuilderRequest(*appstream.CreateImageBuilderInput) appstream.CreateImageBuilderRequest

	CreateImageBuilderStreamingURLRequest(*appstream.CreateImageBuilderStreamingURLInput) appstream.CreateImageBuilderStreamingURLRequest

	CreateStackRequest(*appstream.CreateStackInput) appstream.CreateStackRequest

	CreateStreamingURLRequest(*appstream.CreateStreamingURLInput) appstream.CreateStreamingURLRequest

	DeleteDirectoryConfigRequest(*appstream.DeleteDirectoryConfigInput) appstream.DeleteDirectoryConfigRequest

	DeleteFleetRequest(*appstream.DeleteFleetInput) appstream.DeleteFleetRequest

	DeleteImageRequest(*appstream.DeleteImageInput) appstream.DeleteImageRequest

	DeleteImageBuilderRequest(*appstream.DeleteImageBuilderInput) appstream.DeleteImageBuilderRequest

	DeleteStackRequest(*appstream.DeleteStackInput) appstream.DeleteStackRequest

	DescribeDirectoryConfigsRequest(*appstream.DescribeDirectoryConfigsInput) appstream.DescribeDirectoryConfigsRequest

	DescribeFleetsRequest(*appstream.DescribeFleetsInput) appstream.DescribeFleetsRequest

	DescribeImageBuildersRequest(*appstream.DescribeImageBuildersInput) appstream.DescribeImageBuildersRequest

	DescribeImagesRequest(*appstream.DescribeImagesInput) appstream.DescribeImagesRequest

	DescribeSessionsRequest(*appstream.DescribeSessionsInput) appstream.DescribeSessionsRequest

	DescribeStacksRequest(*appstream.DescribeStacksInput) appstream.DescribeStacksRequest

	DisassociateFleetRequest(*appstream.DisassociateFleetInput) appstream.DisassociateFleetRequest

	ExpireSessionRequest(*appstream.ExpireSessionInput) appstream.ExpireSessionRequest

	ListAssociatedFleetsRequest(*appstream.ListAssociatedFleetsInput) appstream.ListAssociatedFleetsRequest

	ListAssociatedStacksRequest(*appstream.ListAssociatedStacksInput) appstream.ListAssociatedStacksRequest

	ListTagsForResourceRequest(*appstream.ListTagsForResourceInput) appstream.ListTagsForResourceRequest

	StartFleetRequest(*appstream.StartFleetInput) appstream.StartFleetRequest

	StartImageBuilderRequest(*appstream.StartImageBuilderInput) appstream.StartImageBuilderRequest

	StopFleetRequest(*appstream.StopFleetInput) appstream.StopFleetRequest

	StopImageBuilderRequest(*appstream.StopImageBuilderInput) appstream.StopImageBuilderRequest

	TagResourceRequest(*appstream.TagResourceInput) appstream.TagResourceRequest

	UntagResourceRequest(*appstream.UntagResourceInput) appstream.UntagResourceRequest

	UpdateDirectoryConfigRequest(*appstream.UpdateDirectoryConfigInput) appstream.UpdateDirectoryConfigRequest

	UpdateFleetRequest(*appstream.UpdateFleetInput) appstream.UpdateFleetRequest

	UpdateStackRequest(*appstream.UpdateStackInput) appstream.UpdateStackRequest

	WaitUntilFleetStarted(*appstream.DescribeFleetsInput) error
	WaitUntilFleetStartedWithContext(aws.Context, *appstream.DescribeFleetsInput, ...aws.WaiterOption) error

	WaitUntilFleetStopped(*appstream.DescribeFleetsInput) error
	WaitUntilFleetStoppedWithContext(aws.Context, *appstream.DescribeFleetsInput, ...aws.WaiterOption) error
}

var _ AppStreamAPI = (*appstream.AppStream)(nil)
