//go:build !no_runtime_type_checking

package cdklabscdkecscodedeploy

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateNewEcsAppSpecParameters(targetService *TargetService, hooks *AppSpecHooks) error {
	if targetService == nil {
		return fmt.Errorf("parameter targetService is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(targetService, func() string { return "parameter targetService" }); err != nil {
		return err
	}

	if err := _jsii_.ValidateStruct(hooks, func() string { return "parameter hooks" }); err != nil {
		return err
	}

	return nil
}

