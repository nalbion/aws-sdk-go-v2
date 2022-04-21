// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers the user in the specified user pool and creates a user name, password,
// and user attributes. This action might generate an SMS text message. Starting
// June 1, 2021, US telecom carriers require you to register an origination phone
// number before you can send SMS messages to US phone numbers. If you use SMS text
// messages in Amazon Cognito, you must register a phone number with Amazon
// Pinpoint (https://console.aws.amazon.com/pinpoint/home/). Amazon Cognito uses
// the registered number automatically. Otherwise, Amazon Cognito users who must
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
func (c *Client) SignUp(ctx context.Context, params *SignUpInput, optFns ...func(*Options)) (*SignUpOutput, error) {
	if params == nil {
		params = &SignUpInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SignUp", params, optFns, c.addOperationSignUpMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SignUpOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to register a user.
type SignUpInput struct {

	// The ID of the client associated with the user pool.
	//
	// This member is required.
	ClientId *string

	// The password of the user you want to register.
	//
	// This member is required.
	Password *string

	// The user name of the user you want to register.
	//
	// This member is required.
	Username *string

	// The Amazon Pinpoint analytics metadata for collecting metrics for SignUp calls.
	AnalyticsMetadata *types.AnalyticsMetadataType

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers. You create custom workflows by assigning
	// Lambda functions to user pool triggers. When you use the SignUp API action,
	// Amazon Cognito invokes any functions that are assigned to the following
	// triggers: pre sign-up, custom message, and post confirmation. When Amazon
	// Cognito invokes any of these functions, it passes a JSON payload, which the
	// function receives as input. This payload contains a clientMetadata attribute,
	// which provides the data that you assigned to the ClientMetadata parameter in
	// your SignUp request. In your function code in Lambda, you can process the
	// clientMetadata value to enhance your workflow for your specific needs. For more
	// information, see  Customizing user pool Workflows with Lambda Triggers
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

	// A keyed-hash message authentication code (HMAC) calculated using the secret key
	// of a user pool client and username plus the client ID in the message.
	SecretHash *string

	// An array of name-value pairs representing user attributes. For custom
	// attributes, you must prepend the custom: prefix to the attribute name.
	UserAttributes []types.AttributeType

	// Contextual data such as the user's device fingerprint, IP address, or location
	// used for evaluating the risk of an unexpected event by Amazon Cognito advanced
	// security.
	UserContextData *types.UserContextDataType

	// The validation data in the request to register a user.
	ValidationData []types.AttributeType

	noSmithyDocumentSerde
}

// The response from the server for a registration request.
type SignUpOutput struct {

	// A response from the server indicating that a user registration has been
	// confirmed.
	//
	// This member is required.
	UserConfirmed bool

	// The UUID of the authenticated user. This isn't the same as username.
	//
	// This member is required.
	UserSub *string

	// The code delivery details returned by the server response to the user
	// registration request.
	CodeDeliveryDetails *types.CodeDeliveryDetailsType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSignUpMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSignUp{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSignUp{}, middleware.After)
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
	if err = addOpSignUpValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSignUp(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSignUp(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SignUp",
	}
}