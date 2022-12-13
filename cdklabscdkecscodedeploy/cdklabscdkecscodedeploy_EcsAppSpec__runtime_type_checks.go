//go:build !no_runtime_type_checking

// CDK Constructs for performing ECS Deployments with CodeDeploy
package cdklabscdkecscodedeploy

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateNewEcsAppSpecParameters(targetService *TargetService) error {
	if targetService == nil {
		return fmt.Errorf("parameter targetService is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(targetService, func() string { return "parameter targetService" }); err != nil {
		return err
	}

	return nil
}

