// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package guarddutyiface provides an interface to enable mocking the Amazon GuardDuty service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package guarddutyiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
)

// GuardDutyAPI provides an interface to enable mocking the
// guardduty.GuardDuty service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon GuardDuty.
//    func myFunc(svc guarddutyiface.GuardDutyAPI) bool {
//        // Make svc.AcceptInvitation request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := guardduty.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockGuardDutyClient struct {
//        guarddutyiface.GuardDutyAPI
//    }
//    func (m *mockGuardDutyClient) AcceptInvitation(input *guardduty.AcceptInvitationInput) (*guardduty.AcceptInvitationOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockGuardDutyClient{}
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
type GuardDutyAPI interface {
	AcceptInvitationRequest(*guardduty.AcceptInvitationInput) guardduty.AcceptInvitationRequest

	ArchiveFindingsRequest(*guardduty.ArchiveFindingsInput) guardduty.ArchiveFindingsRequest

	CreateDetectorRequest(*guardduty.CreateDetectorInput) guardduty.CreateDetectorRequest

	CreateIPSetRequest(*guardduty.CreateIPSetInput) guardduty.CreateIPSetRequest

	CreateMembersRequest(*guardduty.CreateMembersInput) guardduty.CreateMembersRequest

	CreateSampleFindingsRequest(*guardduty.CreateSampleFindingsInput) guardduty.CreateSampleFindingsRequest

	CreateThreatIntelSetRequest(*guardduty.CreateThreatIntelSetInput) guardduty.CreateThreatIntelSetRequest

	DeclineInvitationsRequest(*guardduty.DeclineInvitationsInput) guardduty.DeclineInvitationsRequest

	DeleteDetectorRequest(*guardduty.DeleteDetectorInput) guardduty.DeleteDetectorRequest

	DeleteIPSetRequest(*guardduty.DeleteIPSetInput) guardduty.DeleteIPSetRequest

	DeleteInvitationsRequest(*guardduty.DeleteInvitationsInput) guardduty.DeleteInvitationsRequest

	DeleteMembersRequest(*guardduty.DeleteMembersInput) guardduty.DeleteMembersRequest

	DeleteThreatIntelSetRequest(*guardduty.DeleteThreatIntelSetInput) guardduty.DeleteThreatIntelSetRequest

	DisassociateFromMasterAccountRequest(*guardduty.DisassociateFromMasterAccountInput) guardduty.DisassociateFromMasterAccountRequest

	DisassociateMembersRequest(*guardduty.DisassociateMembersInput) guardduty.DisassociateMembersRequest

	GetDetectorRequest(*guardduty.GetDetectorInput) guardduty.GetDetectorRequest

	GetFindingsRequest(*guardduty.GetFindingsInput) guardduty.GetFindingsRequest

	GetFindingsStatisticsRequest(*guardduty.GetFindingsStatisticsInput) guardduty.GetFindingsStatisticsRequest

	GetIPSetRequest(*guardduty.GetIPSetInput) guardduty.GetIPSetRequest

	GetInvitationsCountRequest(*guardduty.GetInvitationsCountInput) guardduty.GetInvitationsCountRequest

	GetMasterAccountRequest(*guardduty.GetMasterAccountInput) guardduty.GetMasterAccountRequest

	GetMembersRequest(*guardduty.GetMembersInput) guardduty.GetMembersRequest

	GetThreatIntelSetRequest(*guardduty.GetThreatIntelSetInput) guardduty.GetThreatIntelSetRequest

	InviteMembersRequest(*guardduty.InviteMembersInput) guardduty.InviteMembersRequest

	ListDetectorsRequest(*guardduty.ListDetectorsInput) guardduty.ListDetectorsRequest

	ListDetectorsPages(*guardduty.ListDetectorsInput, func(*guardduty.ListDetectorsOutput, bool) bool) error
	ListDetectorsPagesWithContext(aws.Context, *guardduty.ListDetectorsInput, func(*guardduty.ListDetectorsOutput, bool) bool, ...aws.Option) error

	ListFindingsRequest(*guardduty.ListFindingsInput) guardduty.ListFindingsRequest

	ListFindingsPages(*guardduty.ListFindingsInput, func(*guardduty.ListFindingsOutput, bool) bool) error
	ListFindingsPagesWithContext(aws.Context, *guardduty.ListFindingsInput, func(*guardduty.ListFindingsOutput, bool) bool, ...aws.Option) error

	ListIPSetsRequest(*guardduty.ListIPSetsInput) guardduty.ListIPSetsRequest

	ListIPSetsPages(*guardduty.ListIPSetsInput, func(*guardduty.ListIPSetsOutput, bool) bool) error
	ListIPSetsPagesWithContext(aws.Context, *guardduty.ListIPSetsInput, func(*guardduty.ListIPSetsOutput, bool) bool, ...aws.Option) error

	ListInvitationsRequest(*guardduty.ListInvitationsInput) guardduty.ListInvitationsRequest

	ListInvitationsPages(*guardduty.ListInvitationsInput, func(*guardduty.ListInvitationsOutput, bool) bool) error
	ListInvitationsPagesWithContext(aws.Context, *guardduty.ListInvitationsInput, func(*guardduty.ListInvitationsOutput, bool) bool, ...aws.Option) error

	ListMembersRequest(*guardduty.ListMembersInput) guardduty.ListMembersRequest

	ListMembersPages(*guardduty.ListMembersInput, func(*guardduty.ListMembersOutput, bool) bool) error
	ListMembersPagesWithContext(aws.Context, *guardduty.ListMembersInput, func(*guardduty.ListMembersOutput, bool) bool, ...aws.Option) error

	ListThreatIntelSetsRequest(*guardduty.ListThreatIntelSetsInput) guardduty.ListThreatIntelSetsRequest

	ListThreatIntelSetsPages(*guardduty.ListThreatIntelSetsInput, func(*guardduty.ListThreatIntelSetsOutput, bool) bool) error
	ListThreatIntelSetsPagesWithContext(aws.Context, *guardduty.ListThreatIntelSetsInput, func(*guardduty.ListThreatIntelSetsOutput, bool) bool, ...aws.Option) error

	StartMonitoringMembersRequest(*guardduty.StartMonitoringMembersInput) guardduty.StartMonitoringMembersRequest

	StopMonitoringMembersRequest(*guardduty.StopMonitoringMembersInput) guardduty.StopMonitoringMembersRequest

	UnarchiveFindingsRequest(*guardduty.UnarchiveFindingsInput) guardduty.UnarchiveFindingsRequest

	UpdateDetectorRequest(*guardduty.UpdateDetectorInput) guardduty.UpdateDetectorRequest

	UpdateFindingsFeedbackRequest(*guardduty.UpdateFindingsFeedbackInput) guardduty.UpdateFindingsFeedbackRequest

	UpdateIPSetRequest(*guardduty.UpdateIPSetInput) guardduty.UpdateIPSetRequest

	UpdateThreatIntelSetRequest(*guardduty.UpdateThreatIntelSetInput) guardduty.UpdateThreatIntelSetRequest
}

var _ GuardDutyAPI = (*guardduty.GuardDuty)(nil)
