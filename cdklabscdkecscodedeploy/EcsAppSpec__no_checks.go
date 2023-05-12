//go:build no_runtime_type_checking

package cdklabscdkecscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func validateNewEcsAppSpecParameters(targetService *TargetService, hooks *AppSpecHooks) error {
	return nil
}

