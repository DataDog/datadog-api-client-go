// Get all Container Image groups returns "OK" response

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
	api := datadogV2.NewContainerImagesApi(apiClient)
	resp, r, err := api.ListContainerImages(ctx, *datadogV2.NewListContainerImagesOptionalParameters().WithGroupBy("short_image"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerImagesApi.ListContainerImages`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ContainerImagesApi.ListContainerImages`:\n%s\n", responseContent)
}
