// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get information about a parameter. Request results are returned on a best-effort
// basis. If you specify MaxResults in the request, the response includes
// information up to the limit specified. The number of items returned, however,
// can be between zero and the value of MaxResults. If the service reaches an
// internal limit while processing the results, it stops the operation and returns
// the matching values up to that point and a NextToken. You can specify the
// NextToken in a subsequent call to get the next set of results. If you change the
// KMS key alias for the KMS key used to encrypt a parameter, then you must also
// update the key alias the parameter uses to reference KMS. Otherwise,
// DescribeParameters retrieves whatever the original key alias was referencing.
func (c *Client) DescribeParameters(ctx context.Context, params *DescribeParametersInput, optFns ...func(*Options)) (*DescribeParametersOutput, error) {
	if params == nil {
		params = &DescribeParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeParameters", params, optFns, c.addOperationDescribeParametersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeParametersInput struct {

	// This data type is deprecated. Instead, use ParameterFilters.
	Filters []types.ParametersFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// Filters to limit the request results.
	ParameterFilters []types.ParameterStringFilter

	noSmithyDocumentSerde
}

type DescribeParametersOutput struct {

	// The token to use when requesting the next set of items.
	NextToken *string

	// Parameters returned by the request.
	Parameters []types.ParameterMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeParametersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeParameters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeParameters{}, middleware.After)
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
	if err = addOpDescribeParametersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeParameters(options.Region), middleware.Before); err != nil {
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

// DescribeParametersAPIClient is a client that implements the DescribeParameters
// operation.
type DescribeParametersAPIClient interface {
	DescribeParameters(context.Context, *DescribeParametersInput, ...func(*Options)) (*DescribeParametersOutput, error)
}

var _ DescribeParametersAPIClient = (*Client)(nil)

// DescribeParametersPaginatorOptions is the paginator options for
// DescribeParameters
type DescribeParametersPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeParametersPaginator is a paginator for DescribeParameters
type DescribeParametersPaginator struct {
	options   DescribeParametersPaginatorOptions
	client    DescribeParametersAPIClient
	params    *DescribeParametersInput
	nextToken *string
	firstPage bool
}

// NewDescribeParametersPaginator returns a new DescribeParametersPaginator
func NewDescribeParametersPaginator(client DescribeParametersAPIClient, params *DescribeParametersInput, optFns ...func(*DescribeParametersPaginatorOptions)) *DescribeParametersPaginator {
	if params == nil {
		params = &DescribeParametersInput{}
	}

	options := DescribeParametersPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeParametersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeParametersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeParameters page.
func (p *DescribeParametersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeParametersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeParameters(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeParameters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeParameters",
	}
}