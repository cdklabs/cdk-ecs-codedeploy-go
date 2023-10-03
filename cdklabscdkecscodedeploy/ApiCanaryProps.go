package cdklabscdkecscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssynthetics"
)

// Experimental.
type ApiCanaryProps struct {
	// The base URL to use for tests.
	// Experimental.
	BaseUrl *string `field:"required" json:"baseUrl" yaml:"baseUrl"`
	// The s3 location that stores the data of the canary runs.
	// Default: - A new s3 bucket will be created without a prefix.
	//
	// Experimental.
	ArtifactsBucketLocation *awssynthetics.ArtifactsBucketLocation `field:"optional" json:"artifactsBucketLocation" yaml:"artifactsBucketLocation"`
	// The name of the canary.
	//
	// Be sure to give it a descriptive name that distinguishes it from
	// other canaries in your account.
	//
	// Do not include secrets or proprietary information in your canary name. The canary name
	// makes up part of the canary ARN, which is included in outbound calls over the internet.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/servicelens_canaries_security.html
	//
	// Default: - A unique name will be generated from the construct ID.
	//
	// Experimental.
	CanaryName *string `field:"optional" json:"canaryName" yaml:"canaryName"`
	// The threshold for triggering an alarm on the test duration.
	// Default: - no alarm is created for duration.
	//
	// Experimental.
	DurationAlarmThreshold awscdk.Duration `field:"optional" json:"durationAlarmThreshold" yaml:"durationAlarmThreshold"`
	// How many days should failed runs be retained.
	// Default: Duration.days(31)
	//
	// Experimental.
	FailureRetentionPeriod awscdk.Duration `field:"optional" json:"failureRetentionPeriod" yaml:"failureRetentionPeriod"`
	// Canary execution role.
	//
	// This is the role that will be assumed by the canary upon execution.
	// It controls the permissions that the canary will have. The role must
	// be assumable by the AWS Lambda service principal.
	//
	// If not supplied, a role will be created with all the required permissions.
	// If you provide a Role, you must add the required permissions.
	// See: required permissions: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-executionrolearn
	//
	// Default: - A unique role will be generated for this canary.
	// You can add permissions to roles by calling 'addToRolePolicy'.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Specify the schedule for how often the canary runs.
	//
	// For example, if you set `schedule` to `rate(10 minutes)`, then the canary will run every 10 minutes.
	// You can set the schedule with `Schedule.rate(Duration)` (recommended) or you can specify an expression using `Schedule.expression()`.
	// Default: 'rate(5 minutes)'.
	//
	// Experimental.
	Schedule awssynthetics.Schedule `field:"optional" json:"schedule" yaml:"schedule"`
	// The list of security groups to associate with the canary's network interfaces.
	//
	// You must provide `vpc` when using this prop.
	// Default: - If the canary is placed within a VPC and a security group is
	// not specified a dedicated security group will be created for this canary.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Whether or not the canary should start after creation.
	// Default: true.
	//
	// Experimental.
	StartAfterCreation *bool `field:"optional" json:"startAfterCreation" yaml:"startAfterCreation"`
	// The steps to perform in the synthetic test.
	// Experimental.
	Steps *[]*ApiTestStep `field:"optional" json:"steps" yaml:"steps"`
	// How many days should successful runs be retained.
	// Default: Duration.days(31)
	//
	// Experimental.
	SuccessRetentionPeriod awscdk.Duration `field:"optional" json:"successRetentionPeriod" yaml:"successRetentionPeriod"`
	// The number of threads to run concurrently for the synthetic test.
	// Default: - 20.
	//
	// Experimental.
	ThreadCount *float64 `field:"optional" json:"threadCount" yaml:"threadCount"`
	// How long the canary will be in a 'RUNNING' state.
	//
	// For example, if you set `timeToLive` to be 1 hour and `schedule` to be `rate(10 minutes)`,
	// your canary will run at 10 minute intervals for an hour, for a total of 6 times.
	// Default: - no limit.
	//
	// Experimental.
	TimeToLive awscdk.Duration `field:"optional" json:"timeToLive" yaml:"timeToLive"`
	// The VPC where this canary is run.
	//
	// Specify this if the canary needs to access resources in a VPC.
	// Default: - Not in VPC.
	//
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place the network interfaces within the VPC.
	//
	// You must provide `vpc` when using this prop.
	// Default: - the Vpc default strategy if not specified.
	//
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

