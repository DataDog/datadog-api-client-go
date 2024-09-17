// Update a restriction policy returns "OK" response

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
	// there is a valid "user" in the system

	body := datadogV2.RestrictionPolicyUpdateRequest{
		Data: datadogV2.RestrictionPolicy{
			Id:   "dashboard:test-update",
			Type: datadogV2.RESTRICTIONPOLICYTYPE_RESTRICTION_POLICY,
			Attributes: datadogV2.RestrictionPolicyAttributes{
				Bindings: []datadogV2.RestrictionPolicyBinding{
					{
						Relation: "editor",
						Principals: []string{
							"org:00000000-0000-beef-0000-000000000000",
						},
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRestrictionPoliciesApi(apiClient)
	resp, r, err := api.UpdateRestrictionPolicy(ctx, "dashboard:test-update", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestrictionPoliciesApi.UpdateRestrictionPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RestrictionPoliciesApi.UpdateRestrictionPolicy`:\n%s\n", responseContent)
}
