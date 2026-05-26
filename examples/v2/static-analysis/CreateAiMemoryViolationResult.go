// Create an AI memory violation result returns "Successfully created" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.AiMemoryViolationResultRequest{
		Data: &datadogV2.AiMemoryViolationResultRequestData{
			Attributes: &datadogV2.AiMemoryViolationResultRequestAttributes{
				Line:         10,
				Message:      "This is a false positive because the input is sanitized.",
				Name:         "src/main.py",
				RepositoryId: "my-repo",
				Rule:         "my-ai-ruleset/my-ai-rule",
				Sha:          "abc123def456789012345678901234567890abcd",
				Type:         datadogV2.AIMEMORYVIOLATIONTYPE_FP,
			},
			Id:   datadog.PtrString("violation-abc"),
			Type: datadogV2.AIMEMORYVIOLATIONRESULTDATATYPE_AI_MEMORY_VIOLATION_RESULT.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateAiMemoryViolationResult", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStaticAnalysisApi(apiClient)
	r, err := api.CreateAiMemoryViolationResult(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticAnalysisApi.CreateAiMemoryViolationResult`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
