package cdklabscdkecscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecs-codedeploy-go/cdklabscdkecscodedeploy/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-ecs-codedeploy-go/cdklabscdkecscodedeploy/internal"
)

// A Fargate service running on an ECS cluster fronted by an application load balancer and deployed by CodeDeploy.
// Experimental.
type ApplicationLoadBalancedCodeDeployedFargateService interface {
	awsecspatterns.ApplicationLoadBalancedFargateService
	// S3 Bucket used for access logs.
	// Experimental.
	AccessLogBucket() awss3.IBucket
	// Experimental.
	SetAccessLogBucket(val awss3.IBucket)
	// API Canary for the service.
	// Experimental.
	ApiCanary() ApiCanary
	// Experimental.
	SetApiCanary(val ApiCanary)
	// CodeDeploy application for this service.
	// Experimental.
	Application() awscodedeploy.EcsApplication
	// Experimental.
	SetApplication(val awscodedeploy.EcsApplication)
	// Determines whether the service will be assigned a public IP address.
	// Experimental.
	AssignPublicIp() *bool
	// Certificate Manager certificate to associate with the load balancer.
	// Experimental.
	Certificate() awscertificatemanager.ICertificate
	// The cluster that hosts the service.
	// Experimental.
	Cluster() awsecs.ICluster
	// CodeDeploy deployment for this service.
	// Experimental.
	Deployment() EcsDeployment
	// Experimental.
	SetDeployment(val EcsDeployment)
	// CodeDeploy deployment group for this service.
	// Experimental.
	DeploymentGroup() awscodedeploy.EcsDeploymentGroup
	// Experimental.
	SetDeploymentGroup(val awscodedeploy.EcsDeploymentGroup)
	// Test target group to use for CodeDeploy deployments.
	// Experimental.
	GreenTargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup
	// Experimental.
	SetGreenTargetGroup(val awselasticloadbalancingv2.ApplicationTargetGroup)
	// Composite alarm for monitoring health of service.
	// Experimental.
	HealthAlarm() awscloudwatch.IAlarm
	// Experimental.
	SetHealthAlarm(val awscloudwatch.IAlarm)
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The default is 1 for all new services and uses the existing services desired count
	// when updating an existing service if one is not provided.
	// Experimental.
	InternalDesiredCount() *float64
	// The listener for the service.
	// Experimental.
	Listener() awselasticloadbalancingv2.ApplicationListener
	// The Application Load Balancer for the service.
	// Experimental.
	LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The redirect listener for the service if redirectHTTP is enabled.
	// Experimental.
	RedirectListener() awselasticloadbalancingv2.ApplicationListener
	// The Fargate service in this construct.
	// Experimental.
	Service() awsecs.FargateService
	// The target group for the service.
	// Experimental.
	TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup
	// The Fargate task definition in this construct.
	// Experimental.
	TaskDefinition() awsecs.FargateTaskDefinition
	// Test listener to use for CodeDeploy deployments.
	// Experimental.
	TestListener() awselasticloadbalancingv2.ApplicationListener
	// Experimental.
	SetTestListener(val awselasticloadbalancingv2.ApplicationListener)
	// Adds service as a target of the target group.
	// Experimental.
	AddServiceAsTarget(service awsecs.BaseService)
	// Experimental.
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	// Returns the default cluster.
	// Experimental.
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationLoadBalancedCodeDeployedFargateService
type jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService struct {
	internal.Type__awsecspatternsApplicationLoadBalancedFargateService
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) AccessLogBucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"accessLogBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) ApiCanary() ApiCanary {
	var returns ApiCanary
	_jsii_.Get(
		j,
		"apiCanary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) Application() awscodedeploy.EcsApplication {
	var returns awscodedeploy.EcsApplication
	_jsii_.Get(
		j,
		"application",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) AssignPublicIp() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) Certificate() awscertificatemanager.ICertificate {
	var returns awscertificatemanager.ICertificate
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) Deployment() EcsDeployment {
	var returns EcsDeployment
	_jsii_.Get(
		j,
		"deployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) DeploymentGroup() awscodedeploy.EcsDeploymentGroup {
	var returns awscodedeploy.EcsDeploymentGroup
	_jsii_.Get(
		j,
		"deploymentGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) GreenTargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns awselasticloadbalancingv2.ApplicationTargetGroup
	_jsii_.Get(
		j,
		"greenTargetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) HealthAlarm() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"healthAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) InternalDesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalDesiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) Listener() awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer {
	var returns awselasticloadbalancingv2.ApplicationLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) RedirectListener() awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"redirectListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) Service() awsecs.FargateService {
	var returns awsecs.FargateService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns awselasticloadbalancingv2.ApplicationTargetGroup
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) TaskDefinition() awsecs.FargateTaskDefinition {
	var returns awsecs.FargateTaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) TestListener() awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"testListener",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ApplicationLoadBalancedCodeDeployedFargateService class.
// Experimental.
func NewApplicationLoadBalancedCodeDeployedFargateService(scope constructs.Construct, id *string, props *ApplicationLoadBalancedCodeDeployedFargateServiceProps) ApplicationLoadBalancedCodeDeployedFargateService {
	_init_.Initialize()

	if err := validateNewApplicationLoadBalancedCodeDeployedFargateServiceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService{}

	_jsii_.Create(
		"@cdklabs/cdk-ecs-codedeploy.ApplicationLoadBalancedCodeDeployedFargateService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ApplicationLoadBalancedCodeDeployedFargateService class.
// Experimental.
func NewApplicationLoadBalancedCodeDeployedFargateService_Override(a ApplicationLoadBalancedCodeDeployedFargateService, scope constructs.Construct, id *string, props *ApplicationLoadBalancedCodeDeployedFargateServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-ecs-codedeploy.ApplicationLoadBalancedCodeDeployedFargateService",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService)SetAccessLogBucket(val awss3.IBucket) {
	if err := j.validateSetAccessLogBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessLogBucket",
		val,
	)
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService)SetApiCanary(val ApiCanary) {
	_jsii_.Set(
		j,
		"apiCanary",
		val,
	)
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService)SetApplication(val awscodedeploy.EcsApplication) {
	if err := j.validateSetApplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"application",
		val,
	)
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService)SetDeployment(val EcsDeployment) {
	if err := j.validateSetDeploymentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deployment",
		val,
	)
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService)SetDeploymentGroup(val awscodedeploy.EcsDeploymentGroup) {
	if err := j.validateSetDeploymentGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentGroup",
		val,
	)
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService)SetGreenTargetGroup(val awselasticloadbalancingv2.ApplicationTargetGroup) {
	if err := j.validateSetGreenTargetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"greenTargetGroup",
		val,
	)
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService)SetHealthAlarm(val awscloudwatch.IAlarm) {
	_jsii_.Set(
		j,
		"healthAlarm",
		val,
	)
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService)SetTestListener(val awselasticloadbalancingv2.ApplicationListener) {
	if err := j.validateSetTestListenerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"testListener",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func ApplicationLoadBalancedCodeDeployedFargateService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationLoadBalancedCodeDeployedFargateService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-ecs-codedeploy.ApplicationLoadBalancedCodeDeployedFargateService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) AddServiceAsTarget(service awsecs.BaseService) {
	if err := a.validateAddServiceAsTargetParameters(service); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addServiceAsTarget",
		[]interface{}{service},
	)
}

func (a *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	if err := a.validateCreateAWSLogDriverParameters(prefix); err != nil {
		panic(err)
	}
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		a,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	if err := a.validateGetDefaultClusterParameters(scope); err != nil {
		panic(err)
	}
	var returns awsecs.Cluster

	_jsii_.Invoke(
		a,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

