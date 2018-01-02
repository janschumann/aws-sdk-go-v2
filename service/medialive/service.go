// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// MediaLive provides the API operation methods for making requests to
// AWS Elemental MediaLive. See this package's package overview docs
// for details on the service.
//
// MediaLive methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type MediaLive struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*MediaLive)

// Used for custom request initialization logic
var initRequest func(*MediaLive, *aws.Request)

// Service information constants
const (
	ServiceName = "medialive" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the MediaLive client with a config.
// If additional configuration is needed for the client instance use the
// optional aws.Config parameter to add your extra config.
//
// Example:
//     // Create a MediaLive client from just a config.
//     svc := medialive.New(myConfig)
//
//     // Create a MediaLive client with additional configuration
//     svc := medialive.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func New(config aws.Config) *MediaLive {
	var signingName string
	signingName = "medialive"
	signingRegion := config.Region

	svc := &MediaLive{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2017-10-14",
				JSONVersion:   "1.1",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restjson.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a MediaLive operation and runs any
// custom request initialization.
func (c *MediaLive) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
