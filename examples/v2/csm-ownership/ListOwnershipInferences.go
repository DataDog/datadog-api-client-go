// List ownership inferences for a resource returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.ListOwnershipInferences", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCSMOwnershipApi(apiClient)
	resp, r, err := api.ListOwnershipInferences(ctx, "test-resource")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CSMOwnershipApi.ListOwnershipInferences`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CSMOwnershipApi.ListOwnershipInferences`:\n%s\n", responseContent)
}
