// Get the RUM configuration returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetRumConfig", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumConfigApi(apiClient)
	resp, r, err := api.GetRumConfig(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumConfigApi.GetRumConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumConfigApi.GetRumConfig`:\n%s\n", responseContent)
}
