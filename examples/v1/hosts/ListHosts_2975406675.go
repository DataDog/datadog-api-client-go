// Get all hosts with metadata for your organization returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.HostsApi(apiClient)
	resp, r, err := api.ListHosts(ctx, *datadog.NewListHostsOptionalParameters().WithIncludeHostsMetadata(true))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ListHosts`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `HostsApi.ListHosts`:\n%s\n", responseContent)
}
