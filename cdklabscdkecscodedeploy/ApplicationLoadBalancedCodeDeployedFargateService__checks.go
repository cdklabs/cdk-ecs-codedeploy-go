//go:build !no_runtime_type_checking

package cdklabscdkecscodedeploy

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) validateAddServiceAsTargetParameters(service awsecs.BaseService) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) validateCreateAWSLogDriverParameters(prefix *string) error {
	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateApplicationLoadBalancedCodeDeployedFargateService_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) validateSetAccessLogBucketParameters(val awss3.IBucket) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) validateSetApplicationParameters(val awscodedeploy.EcsApplication) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) validateSetDeploymentParameters(val EcsDeployment) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) validateSetDeploymentGroupParameters(val awscodedeploy.EcsDeploymentGroup) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) validateSetGreenTargetGroupParameters(val awselasticloadbalancingv2.ApplicationTargetGroup) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancedCodeDeployedFargateService) validateSetTestListenerParameters(val awselasticloadbalancingv2.ApplicationListener) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewApplicationLoadBalancedCodeDeployedFargateServiceParameters(scope constructs.Construct, id *string, props *ApplicationLoadBalancedCodeDeployedFargateServiceProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

