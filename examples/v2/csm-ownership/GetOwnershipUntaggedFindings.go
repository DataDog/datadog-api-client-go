// Count untagged findings by ownership confidence returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetOwnershipUntaggedFindings", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCSMOwnershipApi(apiClient)
	resp, r, err := api.GetOwnershipUntaggedFindings(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CSMOwnershipApi.GetOwnershipUntaggedFindings`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CSMOwnershipApi.GetOwnershipUntaggedFindings`:\n%s\n", responseContent)
}
