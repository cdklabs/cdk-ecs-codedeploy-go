package cdklabscdkecscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// The properties for the ApplicationLoadBalancedCodeDeployedFargateService service.
// Experimental.
type ApplicationLoadBalancedCodeDeployedFargateServiceProps struct {
	// A list of Capacity Provider strategies used to place a service.
	// Default: - undefined.
	//
	// Experimental.
	CapacityProviderStrategies *[]*awsecs.CapacityProviderStrategy `field:"optional" json:"capacityProviderStrategies" yaml:"capacityProviderStrategies"`
	// Certificate Manager certificate to associate with the load balancer.
	//
	// Setting this option will set the load balancer protocol to HTTPS.
	// Default: - No certificate associated with the load balancer, if using
	// the HTTP protocol. For HTTPS, a DNS-validated certificate will be
	// created for the load balancer's specified domain name if a domain name
	// and domain zone are specified.
	//
	// Experimental.
	Certificate awscertificatemanager.ICertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	// Default: - disabled.
	//
	// Experimental.
	CircuitBreaker *awsecs.DeploymentCircuitBreaker `field:"optional" json:"circuitBreaker" yaml:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	// Default: - AWS Cloud Map service discovery is not enabled.
	//
	// Experimental.
	CloudMapOptions *awsecs.CloudMapOptions `field:"optional" json:"cloudMapOptions" yaml:"cloudMapOptions"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Default: - create a new cluster; if both cluster and vpc are omitted, a new VPC will be created for you.
	//
	// Experimental.
	Cluster awsecs.ICluster `field:"optional" json:"cluster" yaml:"cluster"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	// Default: - Rolling update (ECS).
	//
	// Experimental.
	DeploymentController *awsecs.DeploymentController `field:"optional" json:"deploymentController" yaml:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The minimum value is 1.
	// Default: - If the feature flag, ECS_REMOVE_DEFAULT_DESIRED_COUNT is false, the default is 1;
	// if true, the default is 1 for all new services and uses the existing services desired count
	// when updating an existing service.
	//
	// Experimental.
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// The domain name for the service, e.g. "api.example.com.".
	// Default: - No domain name.
	//
	// Experimental.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The Route53 hosted zone for the domain, e.g. "example.com.".
	// Default: - No Route53 hosted domain zone.
	//
	// Experimental.
	DomainZone awsroute53.IHostedZone `field:"optional" json:"domainZone" yaml:"domainZone"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Default: false.
	//
	// Experimental.
	EnableECSManagedTags *bool `field:"optional" json:"enableECSManagedTags" yaml:"enableECSManagedTags"`
	// Whether ECS Exec should be enabled.
	// Default: - false.
	//
	// Experimental.
	EnableExecuteCommand *bool `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	// Default: - defaults to 60 seconds if at least one load balancer is in-use and it is not already set.
	//
	// Experimental.
	HealthCheckGracePeriod awscdk.Duration `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// The load balancer idle timeout, in seconds.
	//
	// Can be between 1 and 4000 seconds.
	// Default: - CloudFormation sets idle timeout to 60 seconds.
	//
	// Experimental.
	IdleTimeout awscdk.Duration `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// Listener port of the application load balancer that will serve traffic to the service.
	// Default: - The default listener port is determined from the protocol (port 80 for HTTP,
	// port 443 for HTTPS). A domain name and zone must be also be specified if using HTTPS.
	//
	// Experimental.
	ListenerPort *float64 `field:"optional" json:"listenerPort" yaml:"listenerPort"`
	// The application load balancer that will serve traffic to the service.
	//
	// The VPC attribute of a load balancer must be specified for it to be used
	// to create a new service with this pattern.
	//
	// [disable-awslint:ref-via-interface].
	// Default: - a new load balancer will be created.
	//
	// Experimental.
	LoadBalancer awselasticloadbalancingv2.IApplicationLoadBalancer `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// Name of the load balancer.
	// Default: - Automatically generated name.
	//
	// Experimental.
	LoadBalancerName *string `field:"optional" json:"loadBalancerName" yaml:"loadBalancerName"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	// Default: - 100 if daemon, otherwise 200.
	//
	// Experimental.
	MaxHealthyPercent *float64 `field:"optional" json:"maxHealthyPercent" yaml:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	// Default: - 0 if daemon, otherwise 50.
	//
	// Experimental.
	MinHealthyPercent *float64 `field:"optional" json:"minHealthyPercent" yaml:"minHealthyPercent"`
	// Determines whether or not the Security Group for the Load Balancer's Listener will be open to all traffic by default.
	// Default: true -- The security group allows ingress from all IP addresses.
	//
	// Experimental.
	OpenListener *bool `field:"optional" json:"openListener" yaml:"openListener"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Default: - none.
	//
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The protocol for connections from clients to the load balancer.
	//
	// The load balancer port is determined from the protocol (port 80 for
	// HTTP, port 443 for HTTPS).  If HTTPS, either a certificate or domain
	// name and domain zone must also be specified.
	// Default: HTTP. If a certificate is specified, the protocol will be
	// set by default to HTTPS.
	//
	// Experimental.
	Protocol awselasticloadbalancingv2.ApplicationProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The protocol version to use.
	// Default: ApplicationProtocolVersion.HTTP1
	//
	// Experimental.
	ProtocolVersion awselasticloadbalancingv2.ApplicationProtocolVersion `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
	// Determines whether the Load Balancer will be internet-facing.
	// Default: true.
	//
	// Experimental.
	PublicLoadBalancer *bool `field:"optional" json:"publicLoadBalancer" yaml:"publicLoadBalancer"`
	// Specifies whether the Route53 record should be a CNAME, an A record using the Alias feature or no record at all.
	//
	// This is useful if you need to work with DNS systems that do not support alias records.
	// Default: ApplicationLoadBalancedServiceRecordType.ALIAS
	//
	// Experimental.
	RecordType awsecspatterns.ApplicationLoadBalancedServiceRecordType `field:"optional" json:"recordType" yaml:"recordType"`
	// Specifies whether the load balancer should redirect traffic on port 80 to port 443 to support HTTP->HTTPS redirects This is only valid if the protocol of the ALB is HTTPS.
	// Default: false.
	//
	// Experimental.
	RedirectHTTP *bool `field:"optional" json:"redirectHTTP" yaml:"redirectHTTP"`
	// The name of the service.
	// Default: - CloudFormation-generated name.
	//
	// Experimental.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The security policy that defines which ciphers and protocols are supported by the ALB Listener.
	// Default: - The recommended elastic load balancing security policy.
	//
	// Experimental.
	SslPolicy awselasticloadbalancingv2.SslPolicy `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
	// The protocol for connections from the load balancer to the ECS tasks.
	//
	// The default target port is determined from the protocol (port 80 for
	// HTTP, port 443 for HTTPS).
	// Default: HTTP.
	//
	// Experimental.
	TargetProtocol awselasticloadbalancingv2.ApplicationProtocol `field:"optional" json:"targetProtocol" yaml:"targetProtocol"`
	// The properties required to create a new task definition.
	//
	// TaskDefinition or TaskImageOptions must be specified, but not both.
	// Default: none.
	//
	// Experimental.
	TaskImageOptions *awsecspatterns.ApplicationLoadBalancedTaskImageOptions `field:"optional" json:"taskImageOptions" yaml:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Default: - uses the VPC defined in the cluster or creates a new VPC.
	//
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The number of cpu units used by the task.
	//
	// Valid values, which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 0.5GB, 1GB, 2GB
	//
	// 512 (.5 vCPU) - Available memory values: 1GB, 2GB, 3GB, 4GB
	//
	// 1024 (1 vCPU) - Available memory values: 2GB, 3GB, 4GB, 5GB, 6GB, 7GB, 8GB
	//
	// 2048 (2 vCPU) - Available memory values: Between 4GB and 16GB in 1GB increments
	//
	// 4096 (4 vCPU) - Available memory values: Between 8GB and 30GB in 1GB increments
	//
	// 8192 (8 vCPU) - Available memory values: Between 16GB and 60GB in 4GB increments
	//
	// 16384 (16 vCPU) - Available memory values: Between 32GB and 120GB in 8GB increments
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Default: 256.
	//
	// Experimental.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// The amount (in MiB) of memory used by the task.
	//
	// This field is required and you must use one of the following values, which determines your range of valid values
	// for the cpu parameter:
	//
	// 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available cpu values: 256 (.25 vCPU)
	//
	// 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available cpu values: 512 (.5 vCPU)
	//
	// 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB) - Available cpu values: 1024 (1 vCPU)
	//
	// Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) - Available cpu values: 2048 (2 vCPU)
	//
	// Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available cpu values: 4096 (4 vCPU)
	//
	// Between 16384 (16 GB) and 61440 (60 GB) in increments of 4096 (4 GB) - Available cpu values: 8192 (8 vCPU)
	//
	// Between 32768 (32 GB) and 122880 (120 GB) in increments of 8192 (8 GB) - Available cpu values: 16384 (16 vCPU)
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Default: 512.
	//
	// Experimental.
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// The platform version on which to run your service.
	//
	// If one is not specified, the LATEST platform version is used by default. For more information, see
	// [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	// Default: Latest.
	//
	// Experimental.
	PlatformVersion awsecs.FargatePlatformVersion `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// The runtime platform of the task definition.
	// Default: - If the property is undefined, `operatingSystemFamily` is LINUX and `cpuArchitecture` is X86_64.
	//
	// Experimental.
	RuntimePlatform *awsecs.RuntimePlatform `field:"optional" json:"runtimePlatform" yaml:"runtimePlatform"`
	// The task definition to use for tasks in the service. TaskDefinition or TaskImageOptions must be specified, but not both.
	//
	// [disable-awslint:ref-via-interface].
	// Default: - none.
	//
	// Experimental.
	TaskDefinition awsecs.FargateTaskDefinition `field:"optional" json:"taskDefinition" yaml:"taskDefinition"`
	// Determines whether the service will be assigned a public IP address.
	// Default: false.
	//
	// Experimental.
	AssignPublicIp *bool `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// The health check command and associated configuration parameters for the container.
	// Default: - Health check configuration from container.
	//
	// Experimental.
	HealthCheck *awsecs.HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The security groups to associate with the service.
	//
	// If you do not specify a security group, a new security group is created.
	// Default: - A new security group is created.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The subnets to associate with the service.
	// Default: - Public subnets if `assignPublicIp` is set, otherwise the first available one of Private, Isolated, Public, in that order.
	//
	// Experimental.
	TaskSubnets *awsec2.SubnetSelection `field:"optional" json:"taskSubnets" yaml:"taskSubnets"`
	// The bucket to use for access logs from the Application Load Balancer.
	// Default: - a new S3 bucket will be created.
	//
	// Experimental.
	AccessLogBucket awss3.IBucket `field:"optional" json:"accessLogBucket" yaml:"accessLogBucket"`
	// The prefix to use for access logs from the Application Load Balancer.
	// Default: - none.
	//
	// Experimental.
	AccessLogPrefix *string `field:"optional" json:"accessLogPrefix" yaml:"accessLogPrefix"`
	// The frequency for running the api canaries.
	// Default: - 5 minutes.
	//
	// Experimental.
	ApiCanarySchedule awscdk.Duration `field:"optional" json:"apiCanarySchedule" yaml:"apiCanarySchedule"`
	// The number of threads to run concurrently for the synthetic test.
	// Default: - 20.
	//
	// Experimental.
	ApiCanaryThreadCount *float64 `field:"optional" json:"apiCanaryThreadCount" yaml:"apiCanaryThreadCount"`
	// The threshold for how long a api canary can take to run.
	// Default: - no alarm is created for test duration.
	//
	// Experimental.
	ApiCanaryTimeout awscdk.Duration `field:"optional" json:"apiCanaryTimeout" yaml:"apiCanaryTimeout"`
	// The steps to run in the canary.
	// Default: - no synthetic test will be created.
	//
	// Experimental.
	ApiTestSteps *[]*ApiTestStep `field:"optional" json:"apiTestSteps" yaml:"apiTestSteps"`
	// The physical, human-readable name of the CodeDeploy Application.
	// Default: an auto-generated name will be used.
	//
	// Experimental.
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// The deployment configuration to use for the deployment group.
	// Default: - EcsDeploymentConfig.ALL_AT_ONCE
	//
	// Experimental.
	DeploymentConfig awscodedeploy.IEcsDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
	// The physical, human-readable name of the CodeDeploy Deployment Group.
	// Default: An auto-generated name will be used.
	//
	// Experimental.
	DeploymentGroupName *string `field:"optional" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// The timeout for a CodeDeploy deployment.
	// Default: - 60 minutes.
	//
	// Experimental.
	DeploymentTimeout awscdk.Duration `field:"optional" json:"deploymentTimeout" yaml:"deploymentTimeout"`
	// The amount of time for ELB to wait before changing the state of a deregistering target from 'draining' to 'unused'.
	// Default: - 300 seconds.
	//
	// Experimental.
	DeregistrationDelay awscdk.Duration `field:"optional" json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// Optional lifecycle hooks.
	// Default: - no lifecycle hooks.
	//
	// Experimental.
	Hooks *AppSpecHooks `field:"optional" json:"hooks" yaml:"hooks"`
	// The threshold for response time alarm.
	// Default: - no alarm will be created.
	//
	// Experimental.
	ResponseTimeAlarmThreshold awscdk.Duration `field:"optional" json:"responseTimeAlarmThreshold" yaml:"responseTimeAlarmThreshold"`
	// The healthcheck to configure on the Application Load Balancer target groups.
	// Default: - no health check is configured.
	//
	// Experimental.
	TargetHealthCheck *awselasticloadbalancingv2.HealthCheck `field:"optional" json:"targetHealthCheck" yaml:"targetHealthCheck"`
	// The time to wait before terminating the original (blue) task set.
	// Default: - 10 minutes.
	//
	// Experimental.
	TerminationWaitTime awscdk.Duration `field:"optional" json:"terminationWaitTime" yaml:"terminationWaitTime"`
	// The port to use for test traffic on the listener.
	// Default: - listenerPort + 1.
	//
	// Experimental.
	TestPort *float64 `field:"optional" json:"testPort" yaml:"testPort"`
}

