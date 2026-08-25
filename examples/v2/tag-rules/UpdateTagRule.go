// Update a tag rule returns "OK" response

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
	body := datadogV2.TagRuleUpdateRequest{
		Data: datadogV2.TagRuleUpdateData{
			Attributes: &datadogV2.TagRuleUpdateAttributes{
				RuleType:         datadogV2.TAGRULETYPE_SURFACING.Ptr(),
				TagValuePatterns: []string{},
			},
			Id:   "123",
			Type: datadogV2.TAGRULERESOURCETYPE_TAG_RULE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateTagRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTagRulesApi(apiClient)
	resp, r, err := api.UpdateTagRule(ctx, "rule_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagRulesApi.UpdateTagRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TagRulesApi.UpdateTagRule`:\n%s\n", responseContent)
}
