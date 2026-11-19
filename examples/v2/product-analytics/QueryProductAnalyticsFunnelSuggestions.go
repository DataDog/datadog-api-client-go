// Get feature suggestions returns "OK" response

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
	body := datadogV2.ProductAnalyticsFunnelSuggestionsRequest{
		Data: datadogV2.ProductAnalyticsFunnelSuggestionsRequestData{
			Attributes: &datadogV2.ProductAnalyticsFunnelSuggestionsRequestAttributes{
				ApplicationId: datadog.PtrString("40d8ca4b-2ac0-481f-b16e-f110699ecff6"),
			},
			Type: datadogV2.PRODUCTANALYTICSFUNNELSUGGESTIONSREQUESTTYPE_FUNNEL_SUGGESTIONS_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.QueryProductAnalyticsFunnelSuggestions", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewProductAnalyticsApi(apiClient)
	resp, r, err := api.QueryProductAnalyticsFunnelSuggestions(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAnalyticsApi.QueryProductAnalyticsFunnelSuggestions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ProductAnalyticsApi.QueryProductAnalyticsFunnelSuggestions`:\n%s\n", responseContent)
}
