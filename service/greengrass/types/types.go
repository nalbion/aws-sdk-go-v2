// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// Information about a bulk deployment. You cannot start a new bulk deployment
// while another one is still running or in a non-terminal state.
type BulkDeployment struct {

	// The ARN of the bulk deployment.
	BulkDeploymentArn *string

	// The ID of the bulk deployment.
	BulkDeploymentId *string

	// The time, in ISO format, when the deployment was created.
	CreatedAt *string

	noSmithyDocumentSerde
}

// Relevant metrics on input records processed during bulk deployment.
type BulkDeploymentMetrics struct {

	// The total number of records that returned a non-retryable error. For example,
	// this can occur if a group record from the input file uses an invalid format or
	// specifies a nonexistent group version, or if the execution role doesn't grant
	// permission to deploy a group or group version.
	InvalidInputRecords int32

	// The total number of group records from the input file that have been processed
	// so far, or attempted.
	RecordsProcessed int32

	// The total number of deployment attempts that returned a retryable error. For
	// example, a retry is triggered if the attempt to deploy a group returns a
	// throttling error. ''StartBulkDeployment'' retries a group deployment up to five
	// times.
	RetryAttempts int32

	noSmithyDocumentSerde
}

// Information about an individual group deployment in a bulk deployment operation.
type BulkDeploymentResult struct {

	// The time, in ISO format, when the deployment was created.
	CreatedAt *string

	// The ARN of the group deployment.
	DeploymentArn *string

	// The ID of the group deployment.
	DeploymentId *string

	// The current status of the group deployment: ''InProgress'', ''Building'',
	// ''Success'', or ''Failure''.
	DeploymentStatus *string

	// The type of the deployment.
	DeploymentType DeploymentType

	// Details about the error.
	ErrorDetails []ErrorDetail

	// The error message for a failed deployment
	ErrorMessage *string

	// The ARN of the Greengrass group.
	GroupArn *string

	noSmithyDocumentSerde
}

// Information about a Greengrass core's connectivity.
type ConnectivityInfo struct {

	// The endpoint for the Greengrass core. Can be an IP address or DNS.
	HostAddress *string

	// The ID of the connectivity information.
	Id *string

	// Metadata for this endpoint.
	Metadata *string

	// The port of the Greengrass core. Usually 8883.
	PortNumber int32

	noSmithyDocumentSerde
}

// Information about a connector. Connectors run on the Greengrass core and contain
// built-in integration with local infrastructure, device protocols, AWS, and other
// cloud services.
type Connector struct {

	// The ARN of the connector.
	//
	// This member is required.
	ConnectorArn *string

	// A descriptive or arbitrary ID for the connector. This value must be unique
	// within the connector definition version. Max length is 128 characters with
	// pattern [a-zA-Z0-9:_-]+.
	//
	// This member is required.
	Id *string

	// The parameters or configuration that the connector uses.
	Parameters map[string]string

	noSmithyDocumentSerde
}

// Information about the connector definition version, which is a container for
// connectors.
type ConnectorDefinitionVersion struct {

	// A list of references to connectors in this version, with their corresponding
	// configuration settings.
	Connectors []Connector

	noSmithyDocumentSerde
}

// Information about a core.
type Core struct {

	// The ARN of the certificate associated with the core.
	//
	// This member is required.
	CertificateArn *string

	// A descriptive or arbitrary ID for the core. This value must be unique within the
	// core definition version. Max length is 128 characters with pattern
	// ''[a-zA-Z0-9:_-]+''.
	//
	// This member is required.
	Id *string

	// The ARN of the thing which is the core.
	//
	// This member is required.
	ThingArn *string

	// If true, the core's local shadow is automatically synced with the cloud.
	SyncShadow bool

	noSmithyDocumentSerde
}

// Information about a core definition version.
type CoreDefinitionVersion struct {

	// A list of cores in the core definition version.
	Cores []Core

	noSmithyDocumentSerde
}

// Information about a definition.
type DefinitionInformation struct {

	// The ARN of the definition.
	Arn *string

	// The time, in milliseconds since the epoch, when the definition was created.
	CreationTimestamp *string

	// The ID of the definition.
	Id *string

	// The time, in milliseconds since the epoch, when the definition was last updated.
	LastUpdatedTimestamp *string

	// The ID of the latest version associated with the definition.
	LatestVersion *string

	// The ARN of the latest version associated with the definition.
	LatestVersionArn *string

	// The name of the definition.
	Name *string

	// Tag(s) attached to the resource arn.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Information about a deployment.
type Deployment struct {

	// The time, in milliseconds since the epoch, when the deployment was created.
	CreatedAt *string

	// The ARN of the deployment.
	DeploymentArn *string

	// The ID of the deployment.
	DeploymentId *string

	// The type of the deployment.
	DeploymentType DeploymentType

	// The ARN of the group for this deployment.
	GroupArn *string

	noSmithyDocumentSerde
}

// Information about a device.
type Device struct {

	// The ARN of the certificate associated with the device.
	//
	// This member is required.
	CertificateArn *string

	// A descriptive or arbitrary ID for the device. This value must be unique within
	// the device definition version. Max length is 128 characters with pattern
	// ''[a-zA-Z0-9:_-]+''.
	//
	// This member is required.
	Id *string

	// The thing ARN of the device.
	//
	// This member is required.
	ThingArn *string

	// If true, the device's local shadow will be automatically synced with the cloud.
	SyncShadow bool

	noSmithyDocumentSerde
}

// Information about a device definition version.
type DeviceDefinitionVersion struct {

	// A list of devices in the definition version.
	Devices []Device

	noSmithyDocumentSerde
}

// Details about the error.
type ErrorDetail struct {

	// A detailed error code.
	DetailedErrorCode *string

	// A detailed error message.
	DetailedErrorMessage *string

	noSmithyDocumentSerde
}

// Information about a Lambda function.
type Function struct {

	// A descriptive or arbitrary ID for the function. This value must be unique within
	// the function definition version. Max length is 128 characters with pattern
	// ''[a-zA-Z0-9:_-]+''.
	//
	// This member is required.
	Id *string

	// The ARN of the Lambda function.
	FunctionArn *string

	// The configuration of the Lambda function.
	FunctionConfiguration *FunctionConfiguration

	noSmithyDocumentSerde
}

// The configuration of the Lambda function.
type FunctionConfiguration struct {

	// The expected encoding type of the input payload for the function. The default is
	// ''json''.
	EncodingType EncodingType

	// The environment configuration of the function.
	Environment *FunctionConfigurationEnvironment

	// The execution arguments.
	ExecArgs *string

	// The name of the function executable.
	Executable *string

	// The memory size, in KB, which the function requires. This setting is not
	// applicable and should be cleared when you run the Lambda function without
	// containerization.
	MemorySize int32

	// True if the function is pinned. Pinned means the function is long-lived and
	// starts when the core starts.
	Pinned bool

	// The allowed function execution time, after which Lambda should terminate the
	// function. This timeout still applies to pinned Lambda functions for each
	// request.
	Timeout int32

	noSmithyDocumentSerde
}

// The environment configuration of the function.
type FunctionConfigurationEnvironment struct {

	// If true, the Lambda function is allowed to access the host's /sys folder. Use
	// this when the Lambda function needs to read device information from /sys. This
	// setting applies only when you run the Lambda function in a Greengrass container.
	AccessSysfs bool

	// Configuration related to executing the Lambda function
	Execution *FunctionExecutionConfig

	// A list of the resources, with their permissions, to which the Lambda function
	// will be granted access. A Lambda function can have at most 10 resources.
	// ResourceAccessPolicies apply only when you run the Lambda function in a
	// Greengrass container.
	ResourceAccessPolicies []ResourceAccessPolicy

	// Environment variables for the Lambda function's configuration.
	Variables map[string]string

	noSmithyDocumentSerde
}

// The default configuration that applies to all Lambda functions in the group.
// Individual Lambda functions can override these settings.
type FunctionDefaultConfig struct {

	// Configuration information that specifies how a Lambda function runs.
	Execution *FunctionDefaultExecutionConfig

	noSmithyDocumentSerde
}

// Configuration information that specifies how a Lambda function runs.
type FunctionDefaultExecutionConfig struct {

	// Specifies whether the Lambda function runs in a Greengrass container (default)
	// or without containerization. Unless your scenario requires that you run without
	// containerization, we recommend that you run in a Greengrass container. Omit this
	// value to run the Lambda function with the default containerization for the
	// group.
	IsolationMode FunctionIsolationMode

	// Specifies the user and group whose permissions are used when running the Lambda
	// function. You can specify one or both values to override the default values. We
	// recommend that you avoid running as root unless absolutely necessary to minimize
	// the risk of unintended changes or malicious attacks. To run as root, you must
	// set ''IsolationMode'' to ''NoContainer'' and update config.json in
	// ''greengrass-root/config'' to set ''allowFunctionsToRunAsRoot'' to ''yes''.
	RunAs *FunctionRunAsConfig

	noSmithyDocumentSerde
}

// Information about a function definition version.
type FunctionDefinitionVersion struct {

	// The default configuration that applies to all Lambda functions in this function
	// definition version. Individual Lambda functions can override these settings.
	DefaultConfig *FunctionDefaultConfig

	// A list of Lambda functions in this function definition version.
	Functions []Function

	noSmithyDocumentSerde
}

// Configuration information that specifies how a Lambda function runs.
type FunctionExecutionConfig struct {

	// Specifies whether the Lambda function runs in a Greengrass container (default)
	// or without containerization. Unless your scenario requires that you run without
	// containerization, we recommend that you run in a Greengrass container. Omit this
	// value to run the Lambda function with the default containerization for the
	// group.
	IsolationMode FunctionIsolationMode

	// Specifies the user and group whose permissions are used when running the Lambda
	// function. You can specify one or both values to override the default values. We
	// recommend that you avoid running as root unless absolutely necessary to minimize
	// the risk of unintended changes or malicious attacks. To run as root, you must
	// set ''IsolationMode'' to ''NoContainer'' and update config.json in
	// ''greengrass-root/config'' to set ''allowFunctionsToRunAsRoot'' to ''yes''.
	RunAs *FunctionRunAsConfig

	noSmithyDocumentSerde
}

// Specifies the user and group whose permissions are used when running the Lambda
// function. You can specify one or both values to override the default values. We
// recommend that you avoid running as root unless absolutely necessary to minimize
// the risk of unintended changes or malicious attacks. To run as root, you must
// set ''IsolationMode'' to ''NoContainer'' and update config.json in
// ''greengrass-root/config'' to set ''allowFunctionsToRunAsRoot'' to ''yes''.
type FunctionRunAsConfig struct {

	// The group ID whose permissions are used to run a Lambda function.
	Gid int32

	// The user ID whose permissions are used to run a Lambda function.
	Uid int32

	noSmithyDocumentSerde
}

// Information about a certificate authority for a group.
type GroupCertificateAuthorityProperties struct {

	// The ARN of the certificate authority for the group.
	GroupCertificateAuthorityArn *string

	// The ID of the certificate authority for the group.
	GroupCertificateAuthorityId *string

	noSmithyDocumentSerde
}

// Information about a group.
type GroupInformation struct {

	// The ARN of the group.
	Arn *string

	// The time, in milliseconds since the epoch, when the group was created.
	CreationTimestamp *string

	// The ID of the group.
	Id *string

	// The time, in milliseconds since the epoch, when the group was last updated.
	LastUpdatedTimestamp *string

	// The ID of the latest version associated with the group.
	LatestVersion *string

	// The ARN of the latest version associated with the group.
	LatestVersionArn *string

	// The name of the group.
	Name *string

	noSmithyDocumentSerde
}

// Group owner related settings for local resources.
type GroupOwnerSetting struct {

	// If true, AWS IoT Greengrass automatically adds the specified Linux OS group
	// owner of the resource to the Lambda process privileges. Thus the Lambda process
	// will have the file access permissions of the added Linux group.
	AutoAddGroupOwner bool

	// The name of the Linux OS group whose privileges will be added to the Lambda
	// process. This field is optional.
	GroupOwner *string

	noSmithyDocumentSerde
}

// Information about a group version.
type GroupVersion struct {

	// The ARN of the connector definition version for this group.
	ConnectorDefinitionVersionArn *string

	// The ARN of the core definition version for this group.
	CoreDefinitionVersionArn *string

	// The ARN of the device definition version for this group.
	DeviceDefinitionVersionArn *string

	// The ARN of the function definition version for this group.
	FunctionDefinitionVersionArn *string

	// The ARN of the logger definition version for this group.
	LoggerDefinitionVersionArn *string

	// The ARN of the resource definition version for this group.
	ResourceDefinitionVersionArn *string

	// The ARN of the subscription definition version for this group.
	SubscriptionDefinitionVersionArn *string

	noSmithyDocumentSerde
}

// Attributes that define a local device resource.
type LocalDeviceResourceData struct {

	// Group/owner related settings for local resources.
	GroupOwnerSetting *GroupOwnerSetting

	// The local absolute path of the device resource. The source path for a device
	// resource can refer only to a character device or block device under ''/dev''.
	SourcePath *string

	noSmithyDocumentSerde
}

// Attributes that define a local volume resource.
type LocalVolumeResourceData struct {

	// The absolute local path of the resource inside the Lambda environment.
	DestinationPath *string

	// Allows you to configure additional group privileges for the Lambda process. This
	// field is optional.
	GroupOwnerSetting *GroupOwnerSetting

	// The local absolute path of the volume resource on the host. The source path for
	// a volume resource type cannot start with ''/sys''.
	SourcePath *string

	noSmithyDocumentSerde
}

// Information about a logger
type Logger struct {

	// The component that will be subject to logging.
	//
	// This member is required.
	Component LoggerComponent

	// A descriptive or arbitrary ID for the logger. This value must be unique within
	// the logger definition version. Max length is 128 characters with pattern
	// ''[a-zA-Z0-9:_-]+''.
	//
	// This member is required.
	Id *string

	// The level of the logs.
	//
	// This member is required.
	Level LoggerLevel

	// The type of log output which will be used.
	//
	// This member is required.
	Type LoggerType

	// The amount of file space, in KB, to use if the local file system is used for
	// logging purposes.
	Space int32

	noSmithyDocumentSerde
}

// Information about a logger definition version.
type LoggerDefinitionVersion struct {

	// A list of loggers.
	Loggers []Logger

	noSmithyDocumentSerde
}

// Information about a resource.
type Resource struct {

	// The resource ID, used to refer to a resource in the Lambda function
	// configuration. Max length is 128 characters with pattern ''[a-zA-Z0-9:_-]+''.
	// This must be unique within a Greengrass group.
	//
	// This member is required.
	Id *string

	// The descriptive resource name, which is displayed on the AWS IoT Greengrass
	// console. Max length 128 characters with pattern ''[a-zA-Z0-9:_-]+''. This must
	// be unique within a Greengrass group.
	//
	// This member is required.
	Name *string

	// A container of data for all resource types.
	//
	// This member is required.
	ResourceDataContainer *ResourceDataContainer

	noSmithyDocumentSerde
}

// A policy used by the function to access a resource.
type ResourceAccessPolicy struct {

	// The ID of the resource. (This ID is assigned to the resource when you create the
	// resource definiton.)
	//
	// This member is required.
	ResourceId *string

	// The permissions that the Lambda function has to the resource. Can be one of
	// ''rw'' (read/write) or ''ro'' (read-only).
	Permission Permission

	noSmithyDocumentSerde
}

// A container for resource data. The container takes only one of the following
// supported resource data types: ''LocalDeviceResourceData'',
// ''LocalVolumeResourceData'', ''SageMakerMachineLearningModelResourceData'',
// ''S3MachineLearningModelResourceData'', ''SecretsManagerSecretResourceData''.
type ResourceDataContainer struct {

	// Attributes that define the local device resource.
	LocalDeviceResourceData *LocalDeviceResourceData

	// Attributes that define the local volume resource.
	LocalVolumeResourceData *LocalVolumeResourceData

	// Attributes that define an Amazon S3 machine learning resource.
	S3MachineLearningModelResourceData *S3MachineLearningModelResourceData

	// Attributes that define an Amazon SageMaker machine learning resource.
	SageMakerMachineLearningModelResourceData *SageMakerMachineLearningModelResourceData

	// Attributes that define a secret resource, which references a secret from AWS
	// Secrets Manager.
	SecretsManagerSecretResourceData *SecretsManagerSecretResourceData

	noSmithyDocumentSerde
}

// Information about a resource definition version.
type ResourceDefinitionVersion struct {

	// A list of resources.
	Resources []Resource

	noSmithyDocumentSerde
}

// The owner setting for downloaded machine learning resources.
type ResourceDownloadOwnerSetting struct {

	// The group owner of the resource. This is the name of an existing Linux OS group
	// on the system or a GID. The group's permissions are added to the Lambda process.
	//
	// This member is required.
	GroupOwner *string

	// The permissions that the group owner has to the resource. Valid values are
	// ''rw'' (read/write) or ''ro'' (read-only).
	//
	// This member is required.
	GroupPermission Permission

	noSmithyDocumentSerde
}

// Runtime configuration for a thing.
type RuntimeConfiguration struct {

	// Configuration for telemetry service.
	TelemetryConfiguration *TelemetryConfiguration

	noSmithyDocumentSerde
}

// Attributes that define an Amazon S3 machine learning resource.
type S3MachineLearningModelResourceData struct {

	// The absolute local path of the resource inside the Lambda environment.
	DestinationPath *string

	// The owner setting for downloaded machine learning resources.
	OwnerSetting *ResourceDownloadOwnerSetting

	// The URI of the source model in an S3 bucket. The model package must be in tar.gz
	// or .zip format.
	S3Uri *string

	noSmithyDocumentSerde
}

// Attributes that define an Amazon SageMaker machine learning resource.
type SageMakerMachineLearningModelResourceData struct {

	// The absolute local path of the resource inside the Lambda environment.
	DestinationPath *string

	// The owner setting for downloaded machine learning resources.
	OwnerSetting *ResourceDownloadOwnerSetting

	// The ARN of the Amazon SageMaker training job that represents the source model.
	SageMakerJobArn *string

	noSmithyDocumentSerde
}

// Attributes that define a secret resource, which references a secret from AWS
// Secrets Manager. AWS IoT Greengrass stores a local, encrypted copy of the secret
// on the Greengrass core, where it can be securely accessed by connectors and
// Lambda functions.
type SecretsManagerSecretResourceData struct {

	// The ARN of the Secrets Manager secret to make available on the core. The value
	// of the secret's latest version (represented by the ''AWSCURRENT'' staging label)
	// is included by default.
	ARN *string

	// Optional. The staging labels whose values you want to make available on the
	// core, in addition to ''AWSCURRENT''.
	AdditionalStagingLabelsToDownload []string

	noSmithyDocumentSerde
}

// Information about a subscription.
type Subscription struct {

	// A descriptive or arbitrary ID for the subscription. This value must be unique
	// within the subscription definition version. Max length is 128 characters with
	// pattern ''[a-zA-Z0-9:_-]+''.
	//
	// This member is required.
	Id *string

	// The source of the subscription. Can be a thing ARN, a Lambda function ARN, a
	// connector ARN, 'cloud' (which represents the AWS IoT cloud), or
	// 'GGShadowService'.
	//
	// This member is required.
	Source *string

	// The MQTT topic used to route the message.
	//
	// This member is required.
	Subject *string

	// Where the message is sent to. Can be a thing ARN, a Lambda function ARN, a
	// connector ARN, 'cloud' (which represents the AWS IoT cloud), or
	// 'GGShadowService'.
	//
	// This member is required.
	Target *string

	noSmithyDocumentSerde
}

// Information about a subscription definition version.
type SubscriptionDefinitionVersion struct {

	// A list of subscriptions.
	Subscriptions []Subscription

	noSmithyDocumentSerde
}

// Configuration settings for running telemetry.
type TelemetryConfiguration struct {

	// Configure telemetry to be on or off.
	//
	// This member is required.
	Telemetry Telemetry

	// Synchronization status of the device reported configuration with the desired
	// configuration.
	ConfigurationSyncStatus ConfigurationSyncStatus

	noSmithyDocumentSerde
}

// Configuration settings for running telemetry.
type TelemetryConfigurationUpdate struct {

	// Configure telemetry to be on or off.
	//
	// This member is required.
	Telemetry Telemetry

	noSmithyDocumentSerde
}

// Information about a version.
type VersionInformation struct {

	// The ARN of the version.
	Arn *string

	// The time, in milliseconds since the epoch, when the version was created.
	CreationTimestamp *string

	// The ID of the parent definition that the version is associated with.
	Id *string

	// The ID of the version.
	Version *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde