// Delete an AWS cloud authentication persona mapping returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteAWSCloudAuthPersonaMapping", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudAuthenticationApi(apiClient)
	r, err := api.DeleteAWSCloudAuthPersonaMapping(ctx, "persona_mapping_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAuthenticationApi.DeleteAWSCloudAuthPersonaMapping`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
