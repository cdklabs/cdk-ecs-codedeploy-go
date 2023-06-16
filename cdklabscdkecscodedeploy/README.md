# CDK ECS CodeDeploy

[![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg)](https://constructs.dev/packages/@cdklabs/cdk-ecs-codedeploy)
[![npm version](https://badge.fury.io/js/@cdklabs%2Fcdk-ecs-codedeploy.svg)](https://badge.fury.io/js/@cdklabs%2Fcdk-ecs-codedeploy)
[![Maven Central](https://maven-badges.herokuapp.com/maven-central/io.github.cdklabs/cdk-ecs-codedeploy/badge.svg)](https://maven-badges.herokuapp.com/maven-central/io.github.cdklabs/cdk-ecs-codedeploy)
[![PyPI version](https://badge.fury.io/py/cdklabs.ecs-codedeploy.svg)](https://badge.fury.io/py/cdklabs.ecs-codedeploy)
[![NuGet version](https://badge.fury.io/nu/Cdklabs.CdkEcsCodeDeploy.svg)](https://badge.fury.io/nu/Cdklabs.CdkEcsCodeDeploy)
[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/cdklabs/cdk-ecs-codedeploy)
[![Mergify](https://img.shields.io/endpoint.svg?url=https://api.mergify.com/badges/cdklabs/cdk-ecs-codedeploy&style=flat)](https://mergify.io)

This project contains CDK constructs to create CodeDeploy ECS deployments.

## Installation

<details><summary><strong>TypeScript</strong></summary>

```bash
yarn add @cdklabs/cdk-ecs-codedeploy
```

</details><details><summary><strong>Java</strong></summary>

See https://mvnrepository.com/artifact/io.github.cdklabs/cdk-ecs-codedeploy

</details><details><summary><strong>Python</strong></summary>

See https://pypi.org/project/cdklabs.ecs-codedeploy/

</details><details><summary><strong>C#</strong></summary>

See https://www.nuget.org/packages/Cdklabs.CdkEcsCodeDeploy/

</details>

### Deployments

CodeDeploy for ECS can manage the deployment of new task definitions to ECS services.  Only 1 deployment construct can be defined for a given EcsDeploymentGroup.

```go
var deploymentGroup iEcsDeploymentGroup
var taskDefinition iTaskDefinition


cdklabscdkecscodedeploy.NewEcsDeployment(&EcsDeploymentProps{
	DeploymentGroup: DeploymentGroup,
	TargetService: &TargetService{
		TaskDefinition: *TaskDefinition,
		ContainerName: jsii.String("mycontainer"),
		ContainerPort: jsii.Number(80),
	},
})
```

The deployment will use the AutoRollbackConfig for the EcsDeploymentGroup unless it is overridden in the deployment:

```go
var deploymentGroup iEcsDeploymentGroup
var taskDefinition iTaskDefinition


cdklabscdkecscodedeploy.NewEcsDeployment(&EcsDeploymentProps{
	DeploymentGroup: DeploymentGroup,
	TargetService: &TargetService{
		TaskDefinition: *TaskDefinition,
		ContainerName: jsii.String("mycontainer"),
		ContainerPort: jsii.Number(80),
	},
	AutoRollback: &AutoRollbackConfig{
		FailedDeployment: jsii.Boolean(true),
		DeploymentInAlarm: jsii.Boolean(true),
		StoppedDeployment: jsii.Boolean(false),
	},
})
```

By default, the deployment will timeout after 30 minutes. The timeout value can be overridden:

```go
var deploymentGroup iEcsDeploymentGroup
var taskDefinition iTaskDefinition


cdklabscdkecscodedeploy.NewEcsDeployment(&EcsDeploymentProps{
	DeploymentGroup: DeploymentGroup,
	TargetService: &TargetService{
		TaskDefinition: *TaskDefinition,
		ContainerName: jsii.String("mycontainer"),
		ContainerPort: jsii.Number(80),
	},
	Timeout: awscdk.Duration_Minutes(jsii.Number(60)),
})
```

### API Canaries

CodeDeploy can leverage Cloudwatch Alarms to trigger automatic rollbacks. The `ApiCanary` construct simplifies the process for creating CloudWatch Synthetics Canaries to monitor APIs. The following code demonstrates a canary that monitors https://xkcd.com/908/info.0.json and checks the JSON response to assert that `safe_title` has the value of `'The Cloud'`.

```go
canary := cdklabscdkecscodedeploy.NewApiCanary(stack, jsii.String("Canary"), &ApiCanaryProps{
	BaseUrl: jsii.String("https://xkcd.com"),
	DurationAlarmThreshold: awscdk.Duration_Seconds(jsii.Number(5)),
	ThreadCount: jsii.Number(5),
	Steps: []apiTestStep{
		&apiTestStep{
			Name: jsii.String("info"),
			Path: jsii.String("/908/info.0.json"),
			JmesPath: jsii.String("safe_title"),
			ExpectedValue: jsii.String("The Cloud"),
		},
	},
})
```

### Application Load Balanced CodeDeployed Fargate Service

An L3 construct named `ApplicationLoadBalancedCodeDeployedFargateService` extends [ApplicationLoadBalancedFargateService](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_ecs_patterns.ApplicationLoadBalancedFargateService.html) and adds support for deploying new versions of the service with AWS CodeDeploy. Additionally, an Amazon CloudWatch Synthetic canary is created via the `ApiCanary` construct and is monitored by the CodeDeploy deployment to trigger rollback if the canary begins to alarm.

```go
var cluster iCluster
var image containerImage

service := cdklabscdkecscodedeploy.NewApplicationLoadBalancedCodeDeployedFargateService(stack, jsii.String("Service"), &ApplicationLoadBalancedCodeDeployedFargateServiceProps{
	Cluster: Cluster,
	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
		Image: *Image,
	},
	ApiTestSteps: []apiTestStep{
		&apiTestStep{
			Name: jsii.String("health"),
			Path: jsii.String("/health"),
			JmesPath: jsii.String("status"),
			ExpectedValue: jsii.String("ok"),
		},
	},
})
```

## Local Development

```bash
yarn install
yarn build
yarn test
```

To run an integration test and update the snapshot, run:

```bash
yarn integ:ecs-deployment:deploy
```

To recreate snapshots for integration tests, run:

```bash
yarn integ:snapshot-all
```

## Security

See [CONTRIBUTING](CONTRIBUTING.md#security-issue-notifications) for more information.

## License

This project is licensed under the Apache-2.0 License.
