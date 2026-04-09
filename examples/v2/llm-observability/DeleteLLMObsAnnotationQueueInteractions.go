// Delete annotation queue interactions returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.LLMObsDeleteAnnotationQueueInteractionsRequest{
		Data: datadogV2.LLMObsDeleteAnnotationQueueInteractionsDataRequest{
			Attributes: datadogV2.LLMObsDeleteAnnotationQueueInteractionsDataAttributesRequest{
				InteractionIds: []string{
					"00000000-0000-0000-0000-000000000000",
					"00000000-0000-0000-0000-000000000001",
				},
			},
			Type: datadogV2.LLMOBSANNOTATIONQUEUEINTERACTIONSTYPE_INTERACTIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteLLMObsAnnotationQueueInteractions", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	r, err := api.DeleteLLMObsAnnotationQueueInteractions(ctx, "queue_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.DeleteLLMObsAnnotationQueueInteractions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
