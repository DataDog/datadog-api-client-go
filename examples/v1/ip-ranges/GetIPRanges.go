// List IP Ranges returns "OK" response

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
	api := datadog.NewIPRangesApi(apiClient)
	resp, r, err := api.GetIPRanges(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPRangesApi.GetIPRanges`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IPRangesApi.GetIPRanges`:\n%s\n", responseContent)
}
