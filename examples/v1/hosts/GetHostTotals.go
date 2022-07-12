// Get the total number of active hosts returns "OK" response

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
	api := datadog.NewHostsApi(apiClient)
	resp, r, err := api.GetHostTotals(ctx, *datadog.NewGetHostTotalsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.GetHostTotals`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `HostsApi.GetHostTotals`:\n%s\n", responseContent)
}
