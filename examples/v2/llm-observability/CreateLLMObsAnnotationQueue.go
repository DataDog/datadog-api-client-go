// Create an LLM Observability annotation queue returns "Created" response

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
	body := datadogV2.LLMObsAnnotationQueueRequest{
		Data: datadogV2.LLMObsAnnotationQueueDataRequest{
			Attributes: datadogV2.LLMObsAnnotationQueueDataAttributesRequest{
				Description: datadog.PtrString("Queue for annotating customer support traces"),
				Name:        "My annotation queue",
				ProjectId:   "a33671aa-24fd-4dcd-9b33-a8ec7dde7751",
			},
			Type: datadogV2.LLMOBSANNOTATIONQUEUETYPE_QUEUES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateLLMObsAnnotationQueue", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.CreateLLMObsAnnotationQueue(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.CreateLLMObsAnnotationQueue`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.CreateLLMObsAnnotationQueue`:\n%s\n", responseContent)
}
