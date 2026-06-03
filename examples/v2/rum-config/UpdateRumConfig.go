// Update the RUM configuration returns "OK" response

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
	body := datadogV2.RumConfigUpdateRequest{
		Data: datadogV2.RumConfigUpdateData{
			Attributes: datadogV2.RumConfigUpdateAttributes{
				EnforcedApplicationTags: true,
			},
			Type: datadogV2.RUMCONFIGTYPE_RUM_CONFIG,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateRumConfig", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumConfigApi(apiClient)
	resp, r, err := api.UpdateRumConfig(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumConfigApi.UpdateRumConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumConfigApi.UpdateRumConfig`:\n%s\n", responseContent)
}
