// Push events for an LLM Observability experiment returns "Accepted" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.LLMObsExperimentEventsRequest{
		Data: datadogV2.LLMObsExperimentEventsDataRequest{
			Attributes: datadogV2.LLMObsExperimentEventsDataAttributesRequest{
				Metrics: []datadogV2.LLMObsExperimentMetric{
					{
						Assessment:  datadogV2.LLMOBSMETRICASSESSMENT_PASS.Ptr(),
						Error:       &datadogV2.LLMObsExperimentMetricError{},
						Label:       "faithfulness",
						MetricType:  datadogV2.LLMOBSMETRICSCORETYPE_SCORE,
						SpanId:      "span-7a1b2c3d",
						Tags:        []string{},
						TimestampMs: 1705314600000,
					},
				},
				Spans: []datadogV2.LLMObsExperimentSpan{
					{
						DatasetId: "9f64e5c7-dc5a-45c8-a17c-1b85f0bec97d",
						Duration:  1500000000,
						Meta: &datadogV2.LLMObsExperimentSpanMeta{
							Error: &datadogV2.LLMObsExperimentSpanError{
								Message: datadog.PtrString("Model response timed out"),
								Stack: datadog.PtrString(`Traceback (most recent call last):
  File "main.py", line 10, in <module>
    response = model.generate(input)
  File "model.py", line 45, in generate
    raise TimeoutError("Model response timed out")
TimeoutError: Model response timed out`),
								Type: datadog.PtrString("TimeoutError"),
							},
							Input:  *datadogV2.NewNullableAnyValue(nil),
							Output: *datadogV2.NewNullableAnyValue(nil),
						},
						Name:      "llm_call",
						ProjectId: "a33671aa-24fd-4dcd-9b33-a8ec7dde7751",
						SpanId:    "span-7a1b2c3d",
						StartNs:   1705314600000000000,
						Status:    datadogV2.LLMOBSEXPERIMENTSPANSTATUS_OK,
						Tags:      []string{},
						TraceId:   "abc123def456",
					},
				},
			},
			Type: datadogV2.LLMOBSEVENTTYPE_EVENTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateLLMObsExperimentEvents", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	r, err := api.CreateLLMObsExperimentEvents(ctx, "experiment_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.CreateLLMObsExperimentEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
