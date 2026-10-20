// Patch a saved filter replaces the rules of a saved list

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.PatchSavedFilterRequest{
		Data: datadogV2.PatchSavedFilterData{
			Type: datadogV2.SAVEDFILTERDATATYPE_SAVED_FILTERS,
			Attributes: datadogV2.PatchSavedFilterAttributes{
				TargetingRules: []datadogV2.SavedFilterTargetingRuleRequest{
					{
						Conditions: []datadogV2.SavedFilterConditionRequest{
							{
								Operator:  datadogV2.SAVEDFILTERCONDITIONOPERATOR_ONE_OF,
								Attribute: "user_id",
								Value: []string{
									"test-user-3",
								},
							},
						},
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFeatureFlagsApi(apiClient)
	resp, r, err := api.UpdateSavedFilterAttributes(ctx, uuid.MustParse("d8e3a81b-1505-4600-99c6-c563d01d46b1"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.UpdateSavedFilterAttributes`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.UpdateSavedFilterAttributes`:\n%s\n", responseContent)
}
