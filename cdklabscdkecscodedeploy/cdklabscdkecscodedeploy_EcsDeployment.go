// CDK Constructs for performing ECS Deployments with CodeDeploy
package cdklabscdkecscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecs-codedeploy-go/cdklabscdkecscodedeploy/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-ecs-codedeploy-go/cdklabscdkecscodedeploy/internal"
)

// A CodeDeploy Deployment for a Amazon ECS service DeploymentGroup.
//
// An `EcsDeploymentGroup`
// must only have 1 EcsDeployment. This limit is enforced by making the constructor protected
// and requiring the use of a static method such as `forDeploymentGroup` to initialize.
// The `scope` will always be set to the `EcsDeploymentGroup` and the `id` will always
// be set to the string 'Deployment' to force an error if mulitiple EcsDeployment constructs
// are created for a single EcsDeploymentGroup.
// Experimental.
type EcsDeployment interface {
	constructs.Construct
	// The id of the deployment that was created.
	// Experimental.
	DeploymentId() *string
	// Experimental.
	SetDeploymentId(val *string)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsDeployment
type jsiiProxy_EcsDeployment struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_EcsDeployment) DeploymentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewEcsDeployment(scope constructs.Construct, id *string, props *EcsDeploymentProps) EcsDeployment {
	_init_.Initialize()

	if err := validateNewEcsDeploymentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsDeployment{}

	_jsii_.Create(
		"@cdklabs/cdk-ecs-codedeploy.EcsDeployment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEcsDeployment_Override(e EcsDeployment, scope constructs.Construct, id *string, props *EcsDeploymentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-ecs-codedeploy.EcsDeployment",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_EcsDeployment)SetDeploymentId(val *string) {
	if err := j.validateSetDeploymentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentId",
		val,
	)
}

// Create a new deployment for a given `EcsDeploymentGroup`.
// Experimental.
func EcsDeployment_ForDeploymentGroup(props *EcsDeploymentProps) EcsDeployment {
	_init_.Initialize()

	if err := validateEcsDeployment_ForDeploymentGroupParameters(props); err != nil {
		panic(err)
	}
	var returns EcsDeployment

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-ecs-codedeploy.EcsDeployment",
		"forDeploymentGroup",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func EcsDeployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcsDeployment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-ecs-codedeploy.EcsDeployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDeployment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

