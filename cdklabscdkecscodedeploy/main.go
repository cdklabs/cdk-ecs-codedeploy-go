// CDK Constructs for performing ECS Deployments with CodeDeploy
package cdklabscdkecscodedeploy

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdklabs/cdk-ecs-codedeploy.ApiCanary",
		reflect.TypeOf((*ApiCanary)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTestStep", GoMethod: "AddTestStep"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "artifactsBucket", GoGetter: "ArtifactsBucket"},
			_jsii_.MemberProperty{JsiiProperty: "canaryId", GoGetter: "CanaryId"},
			_jsii_.MemberProperty{JsiiProperty: "canaryName", GoGetter: "CanaryName"},
			_jsii_.MemberProperty{JsiiProperty: "canaryState", GoGetter: "CanaryState"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "durationAlarm", GoGetter: "DurationAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "metricDuration", GoMethod: "MetricDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricSuccessPercent", GoMethod: "MetricSuccessPercent"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "successAlarm", GoGetter: "SuccessAlarm"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ApiCanary{}
			_jsii_.InitJsiiProxy(&j.Type__awssyntheticsCanary)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-ecs-codedeploy.ApiCanaryProps",
		reflect.TypeOf((*ApiCanaryProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-ecs-codedeploy.ApiTestStep",
		reflect.TypeOf((*ApiTestStep)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-ecs-codedeploy.AppSpecHooks",
		reflect.TypeOf((*AppSpecHooks)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-ecs-codedeploy.ApplicationLoadBalancedCodeDeployedFargateService",
		reflect.TypeOf((*ApplicationLoadBalancedCodeDeployedFargateService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessLogBucket", GoGetter: "AccessLogBucket"},
			_jsii_.MemberMethod{JsiiMethod: "addServiceAsTarget", GoMethod: "AddServiceAsTarget"},
			_jsii_.MemberProperty{JsiiProperty: "apiCanary", GoGetter: "ApiCanary"},
			_jsii_.MemberProperty{JsiiProperty: "application", GoGetter: "Application"},
			_jsii_.MemberProperty{JsiiProperty: "assignPublicIp", GoGetter: "AssignPublicIp"},
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberMethod{JsiiMethod: "createAWSLogDriver", GoMethod: "CreateAWSLogDriver"},
			_jsii_.MemberProperty{JsiiProperty: "deployment", GoGetter: "Deployment"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentGroup", GoGetter: "DeploymentGroup"},
			_jsii_.MemberMethod{JsiiMethod: "getDefaultCluster", GoMethod: "GetDefaultCluster"},
			_jsii_.MemberProperty{JsiiProperty: "greenTargetGroup", GoGetter: "GreenTargetGroup"},
			_jsii_.MemberProperty{JsiiProperty: "healthAlarm", GoGetter: "HealthAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "internalDesiredCount", GoGetter: "InternalDesiredCount"},
			_jsii_.MemberProperty{JsiiProperty: "listener", GoGetter: "Listener"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancer", GoGetter: "LoadBalancer"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "redirectListener", GoGetter: "RedirectListener"},
			_jsii_.MemberProperty{JsiiProperty: "service", GoGetter: "Service"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroup", GoGetter: "TargetGroup"},
			_jsii_.MemberProperty{JsiiProperty: "taskDefinition", GoGetter: "TaskDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "testListener", GoGetter: "TestListener"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService{}
			_jsii_.InitJsiiProxy(&j.Type__awsecspatternsApplicationLoadBalancedFargateService)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-ecs-codedeploy.ApplicationLoadBalancedCodeDeployedFargateServiceProps",
		reflect.TypeOf((*ApplicationLoadBalancedCodeDeployedFargateServiceProps)(nil)).Elem(),
	)
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
