// Generate data transformation code returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.DataTransformationRequest{
		ChatHistory: []datadogV2.ChatHistoryItem{
			{
				Content: "Please add error handling",
				Role:    datadogV2.CHATHISTORYITEMROLE_USER,
			},
		},
		Context: &datadogV2.DataTransformationContext{
			ContextVariables: `{ "timestamp": 1234567890 }`,
			CurrentScript:    "return data.timestamp;",
		},
		Language:   datadogV2.DATATRANSFORMATIONLANGUAGE_JAVASCRIPT.Ptr(),
		Stream:     datadog.PtrBool(true),
		UserPrompt: "Convert timestamp to ISO format",
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateDataTransformation", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewWorkflowAutomationApi(apiClient)
	resp, r, err := api.CreateDataTransformation(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowAutomationApi.CreateDataTransformation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
