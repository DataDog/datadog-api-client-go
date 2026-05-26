// List Model Lab run artifacts returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.ListModelLabRunArtifacts", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewModelLabAPIApi(apiClient)
	resp, r, err := api.ListModelLabRunArtifacts(ctx, 70158, *datadogV2.NewListModelLabRunArtifactsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelLabAPIApi.ListModelLabRunArtifacts`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ModelLabAPIApi.ListModelLabRunArtifacts`:\n%s\n", responseContent)
}
