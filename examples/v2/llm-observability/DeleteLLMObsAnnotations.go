// Delete annotations returns "OK — annotations deleted. Errors for annotations that could not be deleted are listed in
// `errors`." response

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
	body := datadogV2.LLMObsDeleteAnnotationsRequest{
		Data: datadogV2.LLMObsDeleteAnnotationsDataRequest{
			Attributes: datadogV2.LLMObsDeleteAnnotationsDataAttributesRequest{
				AnnotationIds: []string{
					"00000000-0000-0000-0000-000000000000",
					"00000000-0000-0000-0000-000000000001",
				},
			},
			Type: datadogV2.LLMOBSANNOTATIONSTYPE_ANNOTATIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteLLMObsAnnotations", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.DeleteLLMObsAnnotations(ctx, "queue_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.DeleteLLMObsAnnotations`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.DeleteLLMObsAnnotations`:\n%s\n", responseContent)
}
