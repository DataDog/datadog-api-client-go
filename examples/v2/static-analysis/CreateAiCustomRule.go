// Create an AI custom rule returns "Successfully created" response

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
	body := datadogV2.AiCustomRuleRequest{
		Data: &datadogV2.AiCustomRuleRequestData{
			Attributes: &datadogV2.AiCustomRuleRequestAttributes{
				Name: datadog.PtrString("my-ai-rule"),
			},
			Id:   datadog.PtrString("my-ai-rule"),
			Type: datadogV2.AICUSTOMRULEDATATYPE_AI_RULE.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateAiCustomRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStaticAnalysisApi(apiClient)
	resp, r, err := api.CreateAiCustomRule(ctx, "my-ai-ruleset", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticAnalysisApi.CreateAiCustomRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StaticAnalysisApi.CreateAiCustomRule`:\n%s\n", responseContent)
}
