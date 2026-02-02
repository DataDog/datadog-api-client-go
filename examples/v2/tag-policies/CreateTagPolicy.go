// Create a tag policy returns "Created" response

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
	body := datadogV2.TagPolicyCreateRequest{
		Data: datadogV2.TagPolicyDataRequest{
			Attributes: datadogV2.TagPolicyAttributesRequest{
				Enabled:    true,
				Negated:    false,
				PolicyName: "production-tags-policy",
				Required:   true,
				Scope:      "env",
				Source:     "logs",
				TagKey:     "service",
				TagValuePatterns: []string{
					"api",
					"web",
				},
			},
			Type: datadogV2.TAGPOLICYTYPE_TAG_POLICY,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateTagPolicy", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTagPoliciesApi(apiClient)
	resp, r, err := api.CreateTagPolicy(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagPoliciesApi.CreateTagPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TagPoliciesApi.CreateTagPolicy`:\n%s\n", responseContent)
}
