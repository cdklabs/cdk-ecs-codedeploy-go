package cdklabscdkecscodedeploy

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-ecs-codedeploy.AwsvpcConfiguration",
		reflect.TypeOf((*AwsvpcConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-ecs-codedeploy.EcsAppSpec",
		reflect.TypeOf((*EcsAppSpec)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_EcsAppSpec{}
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-ecs-codedeploy.EcsDeployment",
		reflect.TypeOf((*EcsDeployment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deploymentId", GoGetter: "DeploymentId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDeployment{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-ecs-codedeploy.EcsDeploymentProps",
		reflect.TypeOf((*EcsDeploymentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-ecs-codedeploy.TargetService",
		reflect.TypeOf((*TargetService)(nil)).Elem(),
	)
}
