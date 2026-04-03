// Create allocation for a flag in an environment returns "Created" response

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
	// there is a valid "feature_flag" in the system
	FeatureFlagDataAttributesVariants0ID := uuid.MustParse(os.Getenv("FEATURE_FLAG_DATA_ATTRIBUTES_VARIANTS_0_ID"))
	FeatureFlagDataID := uuid.MustParse(os.Getenv("FEATURE_FLAG_DATA_ID"))

	// there is a valid "environment" in the system
	EnvironmentDataID := uuid.MustParse(os.Getenv("ENVIRONMENT_DATA_ID"))

	body := datadogV2.CreateAllocationsRequest{
		Data: datadogV2.AllocationDataRequest{
			Type: datadogV2.ALLOCATIONDATATYPE_ALLOCATIONS,
			Attributes: datadogV2.UpsertAllocationRequest{
				Name:           "New targeting rule Example-Feature-Flag",
				Key:            "new-targeting-rule-example-feature-flag",
				TargetingRules: []datadogV2.TargetingRuleRequest{},
				VariantWeights: []datadogV2.VariantWeightRequest{
					{
						VariantId: datadog.PtrUUID(FeatureFlagDataAttributesVariants0ID),
						Value:     100,
					},
				},
				GuardrailMetrics: []datadogV2.GuardrailMetricRequest{},
				Type:             datadogV2.ALLOCATIONTYPE_CANARY,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFeatureFlagsApi(apiClient)
	resp, r, err := api.CreateAllocationsForFeatureFlagInEnvironment(ctx, FeatureFlagDataID, EnvironmentDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.CreateAllocationsForFeatureFlagInEnvironment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.CreateAllocationsForFeatureFlagInEnvironment`:\n%s\n", responseContent)
}
