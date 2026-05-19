// Add a display_block interaction returns "Created" response

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
	body := datadogV2.LLMObsAnnotationQueueInteractionsRequest{
		Data: datadogV2.LLMObsAnnotationQueueInteractionsDataRequest{
			Attributes: datadogV2.LLMObsAnnotationQueueInteractionsDataAttributesRequest{
				Interactions: []datadogV2.LLMObsAnnotationQueueInteractionItem{
					datadogV2.LLMObsAnnotationQueueInteractionItem{
						LLMObsDisplayBlockInteractionItem: &datadogV2.LLMObsDisplayBlockInteractionItem{
							Type: datadogV2.LLMOBSDISPLAYBLOCKINTERACTIONTYPE_DISPLAY_BLOCK,
							DisplayBlock: []datadogV2.LLMObsContentBlock{
								{
									Type:    datadogV2.LLMOBSCONTENTBLOCKTYPE_MARKDOWN,
									Content: "## Triage Instructions",
								},
							},
						}},
				},
			},
			Type: datadogV2.LLMOBSANNOTATIONQUEUEINTERACTIONSTYPE_INTERACTIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateLLMObsAnnotationQueueInteractions", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.CreateLLMObsAnnotationQueueInteractions(ctx, "queue_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.CreateLLMObsAnnotationQueueInteractions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.CreateLLMObsAnnotationQueueInteractions`:\n%s\n", responseContent)
}
