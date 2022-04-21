// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all license usage records for a license configuration, displaying license
// consumption details by resource at a selected point in time. Use this action to
// audit the current license consumption for any license inventory and
// configuration.
func (c *Client) ListUsageForLicenseConfiguration(ctx context.Context, params *ListUsageForLicenseConfigurationInput, optFns ...func(*Options)) (*ListUsageForLicenseConfigurationOutput, error) {
	if params == nil {
		params = &ListUsageForLicenseConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUsageForLicenseConfiguration", params, optFns, c.addOperationListUsageForLicenseConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUsageForLicenseConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUsageForLicenseConfigurationInput struct {

	// Amazon Resource Name (ARN) of the license configuration.
	//
	// This member is required.
	LicenseConfigurationArn *string

	// Filters to scope the results. The following filters and logical operators are
	// supported:
	//
	// * resourceArn - The ARN of the license configuration resource.
	// Logical operators are EQUALS | NOT_EQUALS.
	//
	// * resourceType - The resource type
	// (EC2_INSTANCE | EC2_HOST | EC2_AMI | SYSTEMS_MANAGER_MANAGED_INSTANCE). Logical
	// operators are EQUALS | NOT_EQUALS.
	//
	// * resourceAccount - The ID of the account
	// that owns the resource. Logical operators are EQUALS | NOT_EQUALS.
	Filters []types.Filter

	// Maximum number of results to return in a single call.
	MaxResults *int32

	// Token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListUsageForLicenseConfigurationOutput struct {

	// Information about the license configurations.
	LicenseConfigurationUsageList []types.LicenseConfigurationUsage

	// Token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListUsageForLicenseConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListUsageForLicenseConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListUsageForLicenseConfiguration{}, middleware.After)
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
	if err = addOpListUsageForLicenseConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUsageForLicenseConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListUsageForLicenseConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "ListUsageForLicenseConfiguration",
	}
}