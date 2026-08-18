// Delete LLM Observability data returns "Accepted" response

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
	body := datadogV2.LLMObsDataDeletionRequest{
		Data: datadogV2.LLMObsDataDeletionRequestData{
			Attributes: datadogV2.LLMObsDataDeletionRequestAttributes{
				Delay: datadog.PtrInt64(0),
				From:  1705314600000,
				Query: map[string]string{
					"query": "@trace_id:abc123def456",
				},
				To: 1705315200000,
			},
			Type: datadogV2.LLMOBSDATADELETIONREQUESTTYPE_CREATE_DELETION_REQ,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteLLMObsData", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.DeleteLLMObsData(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.DeleteLLMObsData`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.DeleteLLMObsData`:\n%s\n", responseContent)
}
