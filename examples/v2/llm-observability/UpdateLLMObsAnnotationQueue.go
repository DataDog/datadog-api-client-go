// Update an LLM Observability annotation queue returns "OK" response

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
	body := datadogV2.LLMObsAnnotationQueueUpdateRequest{
		Data: datadogV2.LLMObsAnnotationQueueUpdateDataRequest{
			Attributes: datadogV2.LLMObsAnnotationQueueUpdateDataAttributesRequest{
				Description: datadog.PtrString("Updated description"),
				Name:        datadog.PtrString("Updated queue name"),
			},
			Type: datadogV2.LLMOBSANNOTATIONQUEUETYPE_QUEUES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateLLMObsAnnotationQueue", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.UpdateLLMObsAnnotationQueue(ctx, "queue_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.UpdateLLMObsAnnotationQueue`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.UpdateLLMObsAnnotationQueue`:\n%s\n", responseContent)
}
