//go:build no_runtime_type_checking

// CDK Constructs for performing ECS Deployments with CodeDeploy
package cdklabscdkecscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func validateEcsDeployment_ForDeploymentGroupParameters(props *EcsDeploymentProps) error {
	return nil
}

func validateEcsDeployment_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsDeployment) validateSetDeploymentIdParameters(val *string) error {
	return nil
}

func validateNewEcsDeploymentParameters(scope constructs.Construct, id *string, props *EcsDeploymentProps) error {
	return nil
}

