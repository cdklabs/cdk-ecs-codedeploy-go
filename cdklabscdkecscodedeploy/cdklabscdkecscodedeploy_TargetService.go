// CDK Constructs for performing ECS Deployments with CodeDeploy
package cdklabscdkecscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Describe the target for CodeDeploy to use when creating a deployment for a {@link ecs.EcsDeploymentGroup}.
// Experimental.
type TargetService struct {
	// The name of the Amazon ECS container that contains your Amazon ECS application.
	//
	// It must be a container specified in your Amazon ECS task definition.
	// Experimental.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// The port on the container where traffic will be routed to.
	// Experimental.
	ContainerPort *float64 `field:"required" json:"containerPort" yaml:"containerPort"`
	// The TaskDefintion to deploy to the target services.
	// Experimental.
	TaskDefinition awsecs.ITaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
	// Network configuration for ECS services that have a network type of `awsvpc`.
	// Experimental.
	AwsvpcConfiguration *AwsvpcConfiguration `field:"optional" json:"awsvpcConfiguration" yaml:"awsvpcConfiguration"`
	// A list of Amazon ECS capacity providers to use for the deployment.
	// Experimental.
	CapacityProviderStrategy *[]*awsecs.CapacityProviderStrategy `field:"optional" json:"capacityProviderStrategy" yaml:"capacityProviderStrategy"`
	// The platform version of the Fargate tasks in the deployed Amazon ECS service.
	//
	// {@link https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html}
	// Experimental.
	PlatformVersion awsecs.FargatePlatformVersion `field:"optional" json:"platformVersion" yaml:"platformVersion"`
}

