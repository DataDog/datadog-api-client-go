// Create or update annotations returns "OK — annotations created or updated. Per-item errors are listed in `errors`."
// response

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
	body := datadogV2.LLMObsAnnotationsRequest{
		Data: datadogV2.LLMObsAnnotationsDataRequest{
			Attributes: datadogV2.LLMObsAnnotationsDataAttributesRequest{
				Annotations: []datadogV2.LLMObsUpsertAnnotationItem{
					{
						InteractionId: "00000000-0000-0000-0000-000000000001",
						LabelValues: []datadogV2.LLMObsAnnotationLabelValue{
							{
								LabelSchemaId: "abc-123",
								Value: datadogV2.LLMObsAnnotationLabelValueValue{
									AnyValueString: datadog.PtrString("good")},
							},
							{
								LabelSchemaId: "ef56gh78",
								Value: datadogV2.LLMObsAnnotationLabelValueValue{
									AnyValueString: datadog.PtrString("positive")},
							},
						},
					},
				},
			},
			Type: datadogV2.LLMOBSANNOTATIONSTYPE_ANNOTATIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpsertLLMObsAnnotations", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.UpsertLLMObsAnnotations(ctx, "queue_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.UpsertLLMObsAnnotations`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.UpsertLLMObsAnnotations`:\n%s\n", responseContent)
}
