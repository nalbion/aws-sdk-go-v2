// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the user attribute verification code for the specified attribute name. This
// action might generate an SMS text message. Starting June 1, 2021, US telecom
// carriers require you to register an origination phone number before you can send
// SMS messages to US phone numbers. If you use SMS text messages in Amazon
// Cognito, you must register a phone number with Amazon Pinpoint
// (https://console.aws.amazon.com/pinpoint/home/). Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in. If you have never used SMS text messages with Amazon Cognito or any
// other Amazon Web Service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In sandbox mode
// (https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html) , you can send
// messages only to verified phone numbers. After you test your app while in the
// sandbox environment, you can move out of the sandbox and into production. For
// more information, see  SMS message settings for Amazon Cognito user pools
// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-sms-userpool-settings.html)
// in the Amazon Cognito Developer Guide.
func (c *Client) GetUserAttributeVerificationCode(ctx context.Context, params *GetUserAttributeVerificationCodeInput, optFns ...func(*Options)) (*GetUserAttributeVerificationCodeOutput, error) {
	if params == nil {
		params = &GetUserAttributeVerificationCodeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUserAttributeVerificationCode", params, optFns, c.addOperationGetUserAttributeVerificationCodeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUserAttributeVerificationCodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to get user attribute verification.
type GetUserAttributeVerificationCodeInput struct {

	// The access token returned by the server response to get the user attribute
	// verification code.
	//
	// This member is required.
	AccessToken *string

	// The attribute name returned by the server response to get the user attribute
	// verification code.
	//
	// This member is required.
	AttributeName *string

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers. You create custom workflows by assigning
	// Lambda functions to user pool triggers. When you use the
	// GetUserAttributeVerificationCode API action, Amazon Cognito invokes the function
	// that is assigned to the custom message trigger. When Amazon Cognito invokes this
	// function, it passes a JSON payload, which the function receives as input. This
	// payload contains a clientMetadata attribute, which provides the data that you
	// assigned to the ClientMetadata parameter in your
	// GetUserAttributeVerificationCode request. In your function code in Lambda, you
	// can process the clientMetadata value to enhance your workflow for your specific
	// needs. For more information, see  Customizing user pool Workflows with Lambda
	// Triggers
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide. When you use the ClientMetadata
	// parameter, remember that Amazon Cognito won't do the following:
	//
	// * Store the
	// ClientMetadata value. This data is available only to Lambda triggers that are
	// assigned to a user pool to support custom workflows. If your user pool
	// configuration doesn't include triggers, the ClientMetadata parameter serves no
	// purpose.
	//
	// * Validate the ClientMetadata value.
	//
	// * Encrypt the ClientMetadata
	// value. Don't use Amazon Cognito to provide sensitive information.
	ClientMetadata map[string]string

	noSmithyDocumentSerde
}

// The verification code response returned by the server response to get the user
// attribute verification code.
type GetUserAttributeVerificationCodeOutput struct {

	// The code delivery details returned by the server in response to the request to
	// get the user attribute verification code.
	CodeDeliveryDetails *types.CodeDeliveryDetailsType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUserAttributeVerificationCodeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetUserAttributeVerificationCode{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetUserAttributeVerificationCode{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = addOpGetUserAttributeVerificationCodeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUserAttributeVerificationCode(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetUserAttributeVerificationCode(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetUserAttributeVerificationCode",
	}
}