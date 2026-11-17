// Get funnel step suggestions returns "OK" response

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
	body := datadogV2.ProductAnalyticsFunnelStepSuggestionsRequest{
		Data: datadogV2.ProductAnalyticsFunnelStepSuggestionsRequestData{
			Attributes: datadogV2.ProductAnalyticsFunnelStepSuggestionsRequestAttributes{
				ApplicationId: datadog.PtrString("40d8ca4b-2ac0-481f-b16e-f110699ecff6"),
				Expression:    "node_01->node_02",
				Steps: map[string]string{
					"node_01": "@type:view @view.name:/checkout",
					"node_02": "@type:action @action.name:pay",
				},
			},
			Type: datadogV2.PRODUCTANALYTICSFUNNELSTEPSUGGESTIONSREQUESTTYPE_FUNNEL_SUGGESTIONS_STEPS_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.QueryProductAnalyticsFunnelStepSuggestions", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewProductAnalyticsApi(apiClient)
	resp, r, err := api.QueryProductAnalyticsFunnelStepSuggestions(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAnalyticsApi.QueryProductAnalyticsFunnelStepSuggestions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ProductAnalyticsApi.QueryProductAnalyticsFunnelStepSuggestions`:\n%s\n", responseContent)
}
