// Create the RUM configuration returns "Created" response

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
	body := datadogV2.RUMConfigCreateRequest{
		Data: datadogV2.RUMConfigCreateData{
			Attributes: datadogV2.RUMConfigCreateAttributes{
				EnforcedApplicationTags: true,
			},
			Type: datadogV2.RUMCONFIGTYPE_RUM_CONFIG,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateRUMConfig", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMConfigApi(apiClient)
	resp, r, err := api.CreateRUMConfig(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMConfigApi.CreateRUMConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMConfigApi.CreateRUMConfig`:\n%s\n", responseContent)
}
