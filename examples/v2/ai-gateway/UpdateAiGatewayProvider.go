// Update an AI Gateway provider returns "OK" response

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
	body := datadogV2.AiGatewayProviderUpdateRequest{
		Data: datadogV2.AiGatewayProviderUpdateData{
			Attributes: datadogV2.AiGatewayProviderUpdateAttributes{
				CostMultiplier: datadog.PtrFloat64(1.0),
				Name:           datadog.PtrString("production-us-v2"),
			},
			Id:   "42",
			Type: datadogV2.AIGATEWAYPROVIDERTYPE_PROVIDERS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateAiGatewayProvider", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	resp, r, err := api.UpdateAiGatewayProvider(ctx, 9223372036854775807, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.UpdateAiGatewayProvider`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGatewayApi.UpdateAiGatewayProvider`:\n%s\n", responseContent)
}
