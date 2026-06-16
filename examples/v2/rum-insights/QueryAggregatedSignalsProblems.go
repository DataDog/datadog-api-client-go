// Query aggregated signals and problems returns "Successful response" response

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
	body := datadogV2.AggregatedSignalsProblemsRequest{
		Data: datadogV2.AggregatedSignalsProblemsRequestData{
			Attributes: datadogV2.AggregatedSignalsProblemsRequestAttributes{
				ApplicationId: "ccbc53b1-74f2-496b-bdd7-9a8fa7b7376b",
				Criteria: &datadogV2.AggregatedWaterfallPerformanceCriteria{
					Max:    datadog.PtrFloat64(5.0),
					Metric: datadogV2.AGGREGATEDWATERFALLPERFORMANCECRITERIAMETRIC_LARGEST_CONTENTFUL_PAINT,
					Min:    datadog.PtrFloat64(2.5),
				},
				DetectionTypes: []string{
					"high_script_evaluations",
					"uncompressed_resources",
				},
				Filter:     datadog.PtrString("@session.type:user"),
				From:       1762437564,
				SampleSize: 30,
				To:         1762523964,
				ViewName:   "/account/login(/:type)",
			},
			Type: datadogV2.AGGREGATEDSIGNALSPROBLEMSREQUESTTYPE_AGGREGATED_SIGNALS_PROBLEMS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.QueryAggregatedSignalsProblems", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMInsightsApi(apiClient)
	resp, r, err := api.QueryAggregatedSignalsProblems(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMInsightsApi.QueryAggregatedSignalsProblems`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMInsightsApi.QueryAggregatedSignalsProblems`:\n%s\n", responseContent)
}
