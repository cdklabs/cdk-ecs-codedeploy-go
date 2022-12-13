// CDK Constructs for performing ECS Deployments with CodeDeploy
package cdklabscdkecscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Network configuration for ECS services that have a network type of `awsvpc`.
// Experimental.
type AwsvpcConfiguration struct {
	// Assign a public IP address to the task.
	// Experimental.
	AssignPublicIp *bool `field:"required" json:"assignPublicIp" yaml:"assignPublicIp"`
	// The Security Groups to use for the task.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// The VPC to use for the task.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// The Subnets to use for the task.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"required" json:"vpcSubnets" yaml:"vpcSubnets"`
}

