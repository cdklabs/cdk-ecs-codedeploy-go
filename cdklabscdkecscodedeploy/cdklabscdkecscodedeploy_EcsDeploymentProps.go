// CDK Constructs for performing ECS Deployments with CodeDeploy
package cdklabscdkecscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy"
)

// Construction properties of EcsDeployment.
// Experimental.
type EcsDeploymentProps struct {
	// The AppSpec to use for the deployment.
	//
	// see: https://docs.aws.amazon.com/codedeploy/latest/userguide/reference-appspec-file-structure-resources.html#reference-appspec-file-structure-resources-ecs
	// Experimental.
	Appspec EcsAppSpec `field:"required" json:"appspec" yaml:"appspec"`
	// The deployment group to target for this deployment.
	// Experimental.
	DeploymentGroup awscodedeploy.IEcsDeploymentGroup `field:"required" json:"deploymentGroup" yaml:"deploymentGroup"`
	// The configuration for rollback in the event that a deployment fails.
	// Experimental.
	AutoRollback *awscodedeploy.AutoRollbackConfig `field:"optional" json:"autoRollback" yaml:"autoRollback"`
	// The description for the deployment.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The timeout for the deployment.
	//
	// If the timeout is reached, it will trigger a rollback of the stack.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

