// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package textractiface provides an interface to enable mocking the Amazon Textract service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package textractiface

import (
	"github.com/aws/aws-sdk-go-v2/service/textract"
)

// TextractAPI provides an interface to enable mocking the
// textract.Textract service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Textract.
//    func myFunc(svc textractiface.TextractAPI) bool {
//        // Make svc.AnalyzeDocument request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := textract.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockTextractClient struct {
//        textractiface.TextractAPI
//    }
//    func (m *mockTextractClient) AnalyzeDocument(input *textract.AnalyzeDocumentInput) (*textract.AnalyzeDocumentOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockTextractClient{}
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
type TextractAPI interface {
	AnalyzeDocumentRequest(*textract.AnalyzeDocumentInput) textract.AnalyzeDocumentRequest

	DetectDocumentTextRequest(*textract.DetectDocumentTextInput) textract.DetectDocumentTextRequest

	GetDocumentAnalysisRequest(*textract.GetDocumentAnalysisInput) textract.GetDocumentAnalysisRequest

	GetDocumentTextDetectionRequest(*textract.GetDocumentTextDetectionInput) textract.GetDocumentTextDetectionRequest

	StartDocumentAnalysisRequest(*textract.StartDocumentAnalysisInput) textract.StartDocumentAnalysisRequest

	StartDocumentTextDetectionRequest(*textract.StartDocumentTextDetectionInput) textract.StartDocumentTextDetectionRequest
}

var _ TextractAPI = (*textract.Textract)(nil)
