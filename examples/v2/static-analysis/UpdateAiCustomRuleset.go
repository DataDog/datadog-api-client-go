// Update an AI custom ruleset returns "Successfully updated" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.AiCustomRulesetUpdateRequest{
		Data: &datadogV2.AiCustomRulesetUpdateData{
			Attributes: &datadogV2.AiCustomRulesetUpdateAttributes{
				Description:      datadog.PtrString("Ruleset description"),
				Name:             datadog.PtrString("my-ai-ruleset"),
				ShortDescription: datadog.PtrString("Ruleset short description"),
			},
			Id:   datadog.PtrString("my-ai-ruleset"),
			Type: datadogV2.AICUSTOMRULESETDATATYPE_AI_RULESET.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateAiCustomRuleset", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStaticAnalysisApi(apiClient)
	r, err := api.UpdateAiCustomRuleset(ctx, "my-ai-ruleset", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticAnalysisApi.UpdateAiCustomRuleset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
