// Create a tag rule returns "Created" response

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
	body := datadogV2.TagRuleCreateRequest{
		Data: datadogV2.TagRuleCreateData{
			Attributes: datadogV2.TagRuleCreateAttributes{
				Enabled:    datadog.PtrBool(true),
				Negated:    datadog.PtrBool(false),
				PolicyName: "Service tag must be one of api or web",
				PolicyType: datadogV2.TAGRULECREATETYPE_SURFACING,
				Required:   datadog.PtrBool(true),
				Scope:      "env",
				Source:     datadogV2.TAGRULESOURCE_LOGS,
				TagKey:     "service",
				TagValuePatterns: []string{
					"api",
					"web",
				},
			},
			Type: datadogV2.TAGRULERESOURCETYPE_TAG_POLICY,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateTagRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTagRulesApi(apiClient)
	resp, r, err := api.CreateTagRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagRulesApi.CreateTagRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TagRulesApi.CreateTagRule`:\n%s\n", responseContent)
}
