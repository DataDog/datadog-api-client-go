// Get All Container groups returns "OK" response

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
	api := datadogV2.NewContainersApi(apiClient)
	resp, r, err := api.ListContainers(ctx, *datadogV2.NewListContainersOptionalParameters().WithGroupBy("short_image"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.ListContainers`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ContainersApi.ListContainers`:\n%s\n", responseContent)
}
