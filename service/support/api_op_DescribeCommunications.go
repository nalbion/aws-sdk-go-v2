// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns communications and attachments for one or more support cases. Use the
// afterTime and beforeTime parameters to filter by date. You can use the caseId
// parameter to restrict the results to a specific case. Case data is available for
// 12 months after creation. If a case was created more than 12 months ago, a
// request for data might cause an error. You can use the maxResults and nextToken
// parameters to control the pagination of the results. Set maxResults to the
// number of cases that you want to display on each page, and use nextToken to
// specify the resumption of pagination.
//
// * You must have a Business, Enterprise
// On-Ramp, or Enterprise Support plan to use the Amazon Web Services Support
// API.
//
// * If you call the Amazon Web Services Support API from an account that
// does not have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see Amazon Web Services Support
// (http://aws.amazon.com/premiumsupport/).
func (c *Client) DescribeCommunications(ctx context.Context, params *DescribeCommunicationsInput, optFns ...func(*Options)) (*DescribeCommunicationsOutput, error) {
	if params == nil {
		params = &DescribeCommunicationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCommunications", params, optFns, c.addOperationDescribeCommunicationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCommunicationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCommunicationsInput struct {

	// The support case ID requested or returned in the call. The case ID is an
	// alphanumeric string formatted as shown in this example:
	// case-12345678910-2013-c4c1d2bf33c5cf47
	//
	// This member is required.
	CaseId *string

	// The start date for a filtered date search on support case communications. Case
	// communications are available for 12 months after creation.
	AfterTime *string

	// The end date for a filtered date search on support case communications. Case
	// communications are available for 12 months after creation.
	BeforeTime *string

	// The maximum number of results to return before paginating.
	MaxResults *int32

	// A resumption point for pagination.
	NextToken *string

	noSmithyDocumentSerde
}

// The communications returned by the DescribeCommunications operation.
type DescribeCommunicationsOutput struct {

	// The communications for the case.
	Communications []types.Communication

	// A resumption point for pagination.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCommunicationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCommunications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCommunications{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeCommunicationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCommunications(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeCommunicationsAPIClient is a client that implements the
// DescribeCommunications operation.
type DescribeCommunicationsAPIClient interface {
	DescribeCommunications(context.Context, *DescribeCommunicationsInput, ...func(*Options)) (*DescribeCommunicationsOutput, error)
}

var _ DescribeCommunicationsAPIClient = (*Client)(nil)

// DescribeCommunicationsPaginatorOptions is the paginator options for
// DescribeCommunications
type DescribeCommunicationsPaginatorOptions struct {
	// The maximum number of results to return before paginating.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeCommunicationsPaginator is a paginator for DescribeCommunications
type DescribeCommunicationsPaginator struct {
	options   DescribeCommunicationsPaginatorOptions
	client    DescribeCommunicationsAPIClient
	params    *DescribeCommunicationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeCommunicationsPaginator returns a new DescribeCommunicationsPaginator
func NewDescribeCommunicationsPaginator(client DescribeCommunicationsAPIClient, params *DescribeCommunicationsInput, optFns ...func(*DescribeCommunicationsPaginatorOptions)) *DescribeCommunicationsPaginator {
	if params == nil {
		params = &DescribeCommunicationsInput{}
	}

	options := DescribeCommunicationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeCommunicationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeCommunicationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeCommunications page.
func (p *DescribeCommunicationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeCommunicationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeCommunications(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeCommunications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "DescribeCommunications",
	}
}