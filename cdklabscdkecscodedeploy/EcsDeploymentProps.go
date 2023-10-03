package cdklabscdkecscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy"
)

// Construction properties of EcsDeployment.
// Experimental.
type EcsDeploymentProps struct {
	// The deployment group to target for this deployment.
	// Experimental.
	DeploymentGroup awscodedeploy.IEcsDeploymentGroup `field:"required" json:"deploymentGroup" yaml:"deploymentGroup"`
	// The ECS service to target for the deployment.
	//
	// see: https://docs.aws.amazon.com/codedeploy/latest/userguide/reference-appspec-file-structure-resources.html#reference-appspec-file-structure-resources-ecs
	// Experimental.
	TargetService *TargetService `field:"required" json:"targetService" yaml:"targetService"`
	// The configuration for rollback in the event that a deployment fails.
	// Default: : no automatic rollback triggered.
	//
	// Experimental.
	AutoRollback *awscodedeploy.AutoRollbackConfig `field:"optional" json:"autoRollback" yaml:"autoRollback"`
	// The description for the deployment.
	// Default: no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional lifecycle hooks.
	// Default: - no lifecycle hooks.
	//
	// Experimental.
	Hooks *AppSpecHooks `field:"optional" json:"hooks" yaml:"hooks"`
	// The timeout for the deployment.
	//
	// If the timeout is reached, it will trigger a rollback of the stack.
	// Default: 30 minutes.
	//
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

