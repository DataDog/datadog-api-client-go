// Get all hosts with metadata for your organization returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewHostsApi(apiClient)
	resp, r, err := api.ListHosts(ctx, *datadogV1.NewListHostsOptionalParameters().WithIncludeHostsMetadata(true))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ListHosts`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `HostsApi.ListHosts`:\n%s\n", responseContent)
}
