// Get an AWS cloud authentication persona mapping returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetAWSCloudAuthPersonaMapping", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudAuthenticationApi(apiClient)
	resp, r, err := api.GetAWSCloudAuthPersonaMapping(ctx, "persona_mapping_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAuthenticationApi.GetAWSCloudAuthPersonaMapping`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudAuthenticationApi.GetAWSCloudAuthPersonaMapping`:\n%s\n", responseContent)
}
