// List tag configurations with a tag filter returns "Success" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsApi.ListTagConfigurations(ctx, *datadog.NewListTagConfigurationsOptionalParameters().WithFilterTags("ExampleListtagconfigurationswithatagfilterreturnsSuccessresponse"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListTagConfigurations`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.ListTagConfigurations`:\n%s\n", responseContent)
}
