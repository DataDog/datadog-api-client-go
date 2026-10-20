// Patch a saved filter accepts the current timestamp for rules

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.PatchSavedFilterRequest{
		Data: datadogV2.PatchSavedFilterData{
			Attributes: datadogV2.PatchSavedFilterAttributes{
				LastUpdatedAt: *datadog.NewNullableTime(datadog.PtrTime(time.Date(2026, 10, 16, 20, 20, 30, 0, time.UTC))),
				TargetingRules: []datadogV2.SavedFilterTargetingRuleRequest{
					{
						Conditions: []datadogV2.SavedFilterConditionRequest{
							{
								Attribute: "user_id",
								Operator:  datadogV2.SAVEDFILTERCONDITIONOPERATOR_ONE_OF,
								Value: []string{
									"test-user-2",
								},
							},
						},
					},
				},
			},
			Type: datadogV2.SAVEDFILTERDATATYPE_SAVED_FILTERS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFeatureFlagsApi(apiClient)
	resp, r, err := api.UpdateSavedFilterAttributes(ctx, uuid.MustParse("a8c393b4-3a9c-48fd-aa48-402c97e8d9c9"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.UpdateSavedFilterAttributes`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.UpdateSavedFilterAttributes`:\n%s\n", responseContent)
}
