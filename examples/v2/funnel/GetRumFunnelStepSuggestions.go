// Get rum funnel step suggestions returns "Successful response with funnel step suggestions" response

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
	body := datadogV2.FunnelSuggestionRequest{
		Data: &datadogV2.FunnelSuggestionRequestData{
			Attributes: &datadogV2.FunnelSuggestionRequestDataAttributes{
				DataSource: datadog.PtrString(""),
				Search: &datadogV2.FunnelSuggestionRequestDataAttributesSearch{
					CrossSessionFilter: datadog.PtrString(""),
					QueryString:        datadog.PtrString("@type:view"),
					Steps: []datadogV2.FunnelSuggestionRequestDataAttributesSearchStepsItems{
						{
							Facet:      datadog.PtrString("@view.name"),
							StepFilter: datadog.PtrString(""),
							Value:      datadog.PtrString("/apm/home"),
						},
					},
					SubqueryId: datadog.PtrString(""),
				},
				TermSearch: &datadogV2.FunnelSuggestionRequestDataAttributesTermSearch{
					Query: datadog.PtrString("apm"),
				},
				Time: &datadogV2.FunnelSuggestionRequestDataAttributesTime{
					From: datadog.PtrInt64(1756425600000),
					To:   datadog.PtrInt64(1756857600000),
				},
			},
			Id:   datadog.PtrString("funnel_suggestion_request"),
			Type: datadogV2.FUNNELSUGGESTIONREQUESTDATATYPE_FUNNEL_SUGGESTION_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetRumFunnelStepSuggestions", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFunnelApi(apiClient)
	resp, r, err := api.GetRumFunnelStepSuggestions(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunnelApi.GetRumFunnelStepSuggestions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FunnelApi.GetRumFunnelStepSuggestions`:\n%s\n", responseContent)
}
