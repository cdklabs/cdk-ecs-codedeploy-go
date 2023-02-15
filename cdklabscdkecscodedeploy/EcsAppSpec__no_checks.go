//go:build no_runtime_type_checking

// CDK Constructs for performing ECS Deployments with CodeDeploy
package cdklabscdkecscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func validateNewEcsAppSpecParameters(targetService *TargetService) error {
	return nil
}

