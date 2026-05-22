// Aggregate LLM Observability experimentation returns "OK" response

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
	body := datadogV2.LLMObsExperimentationAnalyticsRequest{
		Data: datadogV2.LLMObsExperimentationAnalyticsDataRequest{
			Attributes: datadogV2.LLMObsExperimentationAnalyticsDataAttributesRequest{
				Aggregate: datadogV2.LLMObsExperimentationAnalyticsAggregate{
					Compute: []datadogV2.LLMObsExperimentationAnalyticsCompute{
						{
							Metric: "score_value",
							Name:   datadog.PtrString("avg_faithfulness"),
						},
					},
					DatasetVersion: *datadog.NewNullableInt64(nil),
					GroupBy: []datadogV2.LLMObsExperimentationAnalyticsGroupBy{
						{
							Field: "span_id",
						},
					},
					Indexes: []string{
						"experiment-evals",
					},
					Limit: *datadog.NewNullableInt32(datadog.PtrInt32(1000)),
					Search: datadogV2.LLMObsExperimentationAnalyticsSearch{
						Query: "@experiment_id:3fd6b5e0-8910-4b1c-a7d0-5b84de329012",
					},
					Time: &datadogV2.LLMObsExperimentationAnalyticsTimeRange{
						From: 1705312200000,
						To:   1705315800000,
					},
				},
			},
			Type: datadogV2.LLMOBSEXPERIMENTATIONTYPE_EXPERIMENTATION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.AggregateLLMObsExperimentation", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.AggregateLLMObsExperimentation(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.AggregateLLMObsExperimentation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.AggregateLLMObsExperimentation`:\n%s\n", responseContent)
}
