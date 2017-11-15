package singularity_test

import (
	singularity "github.com/lenfree/go-singularity"
)

func ExampleClient_CreateRequest() {
	config := singularity.Config{
		Host: "localhost/singularity",
	}
	client := singularity.New(config)
	onDemandTypeReq := singularity.NewOnDemandRequest()
	client.CreateRequest(onDemandTypeReq)

	// Output:
	// {"request":
	//	{
	//		"id":"lenfree-test",
	//		"requestType":"ON_DEMAND",
	//		"numRetriesOnFailure":0,
	//		"rackSensitive":false,
	//		"loadBalanced":false,
	//		"killOldNonLongRunningTasksAfterMillis":0,
	//		"scheduledExpectedRuntimeMillis":0,
	//		"bounceAfterScale":false,
	//		"skipHealthchecks":false,
	//		"taskLogErrorRegex": "",
	//		"taskLogErrorRegexCaseSensitive":false
	//	},
	//	"state":"ACTIVE"}"
	//)
}
