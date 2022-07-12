// Generate a new external ID returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.AWSAccount{
		AccountId: common.PtrString("1234567"),
		AccountSpecificNamespaceRules: map[string]bool{
			"auto_scaling": false,
			"opswork":      false,
		},
		CspmResourceCollectionEnabled: common.PtrBool(true),
		ExcludedRegions: []string{
			"us-east-1",
			"us-west-2",
		},
		FilterTags: []string{
			"$KEY:$VALUE",
		},
		HostTags: []string{
			"$KEY:$VALUE",
		},
		MetricsCollectionEnabled:  common.PtrBool(false),
		ResourceCollectionEnabled: common.PtrBool(true),
		RoleName:                  common.PtrString("DatadogAWSIntegrationRole"),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewAWSIntegrationApi(apiClient)
	resp, r, err := api.CreateNewAWSExternalID(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.CreateNewAWSExternalID`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AWSIntegrationApi.CreateNewAWSExternalID`:\n%s\n", responseContent)
}
