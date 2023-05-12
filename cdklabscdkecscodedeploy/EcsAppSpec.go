package cdklabscdkecscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecs-codedeploy-go/cdklabscdkecscodedeploy/jsii"
)

// Represents an AppSpec to be used for ECS services.
//
// see: https://docs.aws.amazon.com/codedeploy/latest/userguide/reference-appspec-file-structure-resources.html#reference-appspec-file-structure-resources-ecs
// Experimental.
type EcsAppSpec interface {
	// Render JSON string for this AppSpec to be used.
	//
	// Returns: string representation of this AppSpec.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsAppSpec
type jsiiProxy_EcsAppSpec struct {
	_ byte // padding
}

// Experimental.
func NewEcsAppSpec(targetService *TargetService, hooks *AppSpecHooks) EcsAppSpec {
	_init_.Initialize()

	if err := validateNewEcsAppSpecParameters(targetService, hooks); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsAppSpec{}

	_jsii_.Create(
		"@cdklabs/cdk-ecs-codedeploy.EcsAppSpec",
		[]interface{}{targetService, hooks},
		&j,
	)

	return &j
}

// Experimental.
func NewEcsAppSpec_Override(e EcsAppSpec, targetService *TargetService, hooks *AppSpecHooks) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-ecs-codedeploy.EcsAppSpec",
		[]interface{}{targetService, hooks},
		e,
	)
}

func (e *jsiiProxy_EcsAppSpec) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

