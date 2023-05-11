package cdklabscdkecscodedeploy


// Experimental.
type ApiTestStep struct {
	// Name of test.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Path of HTTP request, relative to baseUrl.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Optional body to include in HTTP request.
	// Experimental.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// Expected value to compare against the jmesPath.
	// Experimental.
	ExpectedValue interface{} `field:"optional" json:"expectedValue" yaml:"expectedValue"`
	// Optional headers to include in HTTP request.
	// Experimental.
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// JMESPath to apply against the response from the HTTP request and compare against expected value.
	// Experimental.
	JmesPath *string `field:"optional" json:"jmesPath" yaml:"jmesPath"`
	// Optional method to for HTTP request.
	// Experimental.
	Method *string `field:"optional" json:"method" yaml:"method"`
}

