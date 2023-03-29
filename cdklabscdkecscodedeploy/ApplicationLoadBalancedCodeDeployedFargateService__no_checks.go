//go:build no_runtime_type_checking

// CDK Constructs for performing ECS Deployments with CodeDeploy
package cdklabscdkecscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) validateAddServiceAsTargetParameters(service awsecs.BaseService) error {
	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) validateCreateAWSLogDriverParameters(prefix *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	return nil
}

func validateApplicationLoadBalancedCodeDeployedFargateService_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) validateSetAccessLogBucketParameters(val awss3.IBucket) error {
	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) validateSetApplicationParameters(val awscodedeploy.EcsApplication) error {
	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) validateSetDeploymentParameters(val EcsDeployment) error {
	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) validateSetDeploymentGroupParameters(val awscodedeploy.EcsDeploymentGroup) error {
	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) validateSetGreenTargetGroupParameters(val awselasticloadbalancingv2.ApplicationTargetGroup) error {
	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) validateSetTestListenerParameters(val awselasticloadbalancingv2.ApplicationListener) error {
	return nil
}

func validateNewApplicationLoadBalancedCodeDeployedFargateServiceParameters(scope constructs.Construct, id *string, props *ApplicationLoadBalancedCodeDeployedFargateServiceProps) error {
	return nil
}

