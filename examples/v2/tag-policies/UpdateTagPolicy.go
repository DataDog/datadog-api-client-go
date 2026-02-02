// Update a tag policy returns "OK" response

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
	body := datadogV2.TagPolicyUpdateRequest{
		Data: datadogV2.TagPolicyDataUpdateRequest{
			Attributes: datadogV2.TagPolicyAttributesUpdateRequest{
				Enabled:    datadog.PtrBool(true),
				Negated:    datadog.PtrBool(false),
				PolicyName: datadog.PtrString("production-tags-policy"),
				Required:   datadog.PtrBool(true),
				Scope:      datadog.PtrString("env"),
				Source:     datadog.PtrString("logs"),
				TagKey:     datadog.PtrString("service"),
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
	configuration.SetUnstableOperationEnabled("v2.UpdateTagPolicy", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTagPoliciesApi(apiClient)
	resp, r, err := api.UpdateTagPolicy(ctx, "123", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagPoliciesApi.UpdateTagPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TagPoliciesApi.UpdateTagPolicy`:\n%s\n", responseContent)
}
