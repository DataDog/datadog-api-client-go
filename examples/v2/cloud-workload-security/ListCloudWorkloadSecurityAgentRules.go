// Get all Cloud Workload Security Agent rules returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.CloudWorkloadSecurityApi(apiClient)
	resp, r, err := api.ListCloudWorkloadSecurityAgentRules(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkloadSecurityApi.ListCloudWorkloadSecurityAgentRules`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudWorkloadSecurityApi.ListCloudWorkloadSecurityAgentRules`:\n%s\n", responseContent)
}
