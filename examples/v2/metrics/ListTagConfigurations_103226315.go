// List tag configurations with configured filter returns "Success" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v2/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewMetricsApi(apiClient)
	resp, r, err := api.ListTagConfigurations(ctx, *datadog.NewListTagConfigurationsOptionalParameters().WithFilterConfigured(true))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListTagConfigurations`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.ListTagConfigurations`:\n%s\n", responseContent)
}
