package cdklabscdkecscodedeploy


// Lifecycle hooks configuration.
// Experimental.
type AppSpecHooks struct {
	// Lambda or ARN of a lambda to run tasks after the test listener serves traffic to the replacement task set.
	// Experimental.
	AfterAllowTestTraffic interface{} `field:"optional" json:"afterAllowTestTraffic" yaml:"afterAllowTestTraffic"`
	// Lambda or ARN of a lambda to run tasks after the second target group serves traffic to the replacement task set.
	// Experimental.
	AfterAllowTraffic interface{} `field:"optional" json:"afterAllowTraffic" yaml:"afterAllowTraffic"`
	// Lambda or ARN of a lambda to run tasks after the replacement task set is created and one of the target groups is associated with it.
	// Experimental.
	AfterInstall interface{} `field:"optional" json:"afterInstall" yaml:"afterInstall"`
	// Lambda or ARN of a lambda to run tasks after the second target group is associated with the replacement task set, but before traffic is shifted to the replacement task set.
	// Experimental.
	BeforeAllowTraffic interface{} `field:"optional" json:"beforeAllowTraffic" yaml:"beforeAllowTraffic"`
	// Lambda or ARN of a lambda to run tasks before the replacement task set is created.
	// Experimental.
	BeforeInstall interface{} `field:"optional" json:"beforeInstall" yaml:"beforeInstall"`
}

