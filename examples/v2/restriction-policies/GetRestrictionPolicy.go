// Get a restriction policy returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRestrictionPoliciesApi(apiClient)
	resp, r, err := api.GetRestrictionPolicy(ctx, "dashboard:test-get")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestrictionPoliciesApi.GetRestrictionPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RestrictionPoliciesApi.GetRestrictionPolicy`:\n%s\n", responseContent)
}
