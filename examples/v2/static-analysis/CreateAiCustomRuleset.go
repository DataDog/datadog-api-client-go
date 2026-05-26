// Create an AI custom ruleset returns "Successfully created" response

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
	body := datadogV2.AiCustomRulesetRequest{
		Data: &datadogV2.AiCustomRulesetRequestData{
			Attributes: &datadogV2.AiCustomRulesetRequestAttributes{
				Description:      "Ruleset description",
				Name:             "my-ai-ruleset",
				ShortDescription: "Ruleset short description",
			},
			Id:   datadog.PtrString("my-ai-ruleset"),
			Type: datadogV2.AICUSTOMRULESETDATATYPE_AI_RULESET.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateAiCustomRuleset", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStaticAnalysisApi(apiClient)
	resp, r, err := api.CreateAiCustomRuleset(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticAnalysisApi.CreateAiCustomRuleset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StaticAnalysisApi.CreateAiCustomRuleset`:\n%s\n", responseContent)
}
