// List Model Lab project facet keys returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.ListModelLabProjectFacetKeys", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewModelLabAPIApi(apiClient)
	resp, r, err := api.ListModelLabProjectFacetKeys(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelLabAPIApi.ListModelLabProjectFacetKeys`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ModelLabAPIApi.ListModelLabProjectFacetKeys`:\n%s\n", responseContent)
}
