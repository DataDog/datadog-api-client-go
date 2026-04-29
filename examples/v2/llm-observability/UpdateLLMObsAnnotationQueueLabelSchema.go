// Update annotation queue label schema returns "OK" response

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
	body := datadogV2.LLMObsAnnotationQueueLabelSchemaUpdateRequest{
		Data: datadogV2.LLMObsAnnotationQueueLabelSchemaUpdateData{
			Attributes: datadogV2.LLMObsAnnotationQueueLabelSchemaUpdateAttributes{
				AnnotationSchema: datadogV2.LLMObsAnnotationSchema{
					LabelSchemas: []datadogV2.LLMObsLabelSchema{
						{
							Description:   datadog.PtrString("Rating of the response quality."),
							HasAssessment: datadog.PtrBool(false),
							HasReasoning:  datadog.PtrBool(false),
							Id:            datadog.PtrString("ab12cd34"),
							IsAssessment:  datadog.PtrBool(false),
							IsInteger:     datadog.PtrBool(false),
							IsRequired:    datadog.PtrBool(true),
							Max:           datadog.PtrFloat64(5.0),
							Min:           datadog.PtrFloat64(0.0),
							Name:          "quality",
							Type:          datadogV2.LLMOBSLABELSCHEMATYPE_SCORE,
							Values: []string{
								"good",
								"bad",
								"neutral",
							},
						},
					},
				},
			},
			Type: datadogV2.LLMOBSANNOTATIONQUEUETYPE_QUEUES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateLLMObsAnnotationQueueLabelSchema", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.UpdateLLMObsAnnotationQueueLabelSchema(ctx, "queue_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.UpdateLLMObsAnnotationQueueLabelSchema`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.UpdateLLMObsAnnotationQueueLabelSchema`:\n%s\n", responseContent)
}
