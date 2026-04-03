// Update targeting rules for a flag in an environment returns "OK" response

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

	body := datadogV2.OverwriteAllocationsRequest{
		Data: []datadogV2.AllocationDataRequest{
			{
				Type: datadogV2.ALLOCATIONDATATYPE_ALLOCATIONS,
				Attributes: datadogV2.UpsertAllocationRequest{
					Key:            "overwrite-allocation-example-feature-flag",
					Name:           "New targeting rule Example-Feature-Flag",
					TargetingRules: []datadogV2.TargetingRuleRequest{},
					VariantWeights: []datadogV2.VariantWeightRequest{
						{
							VariantId: datadog.PtrUUID(FeatureFlagDataAttributesVariants0ID),
							Value:     100,
						},
					},
					ExposureSchedule: &datadogV2.ExposureScheduleRequest{
						RolloutOptions: datadogV2.RolloutOptionsRequest{
							Strategy:            datadogV2.ROLLOUTSTRATEGY_UNIFORM_INTERVALS,
							Autostart:           *datadog.NewNullableBool(datadog.PtrBool(false)),
							SelectionIntervalMs: datadog.PtrInt64(86400000),
						},
						RolloutSteps: []datadogV2.ExposureRolloutStepRequest{
							{
								ExposureRatio:    0.05,
								IntervalMs:       *datadog.NewNullableInt64(nil),
								IsPauseRecord:    false,
								GroupedStepIndex: 0,
							},
							{
								ExposureRatio:    0.25,
								IntervalMs:       *datadog.NewNullableInt64(nil),
								IsPauseRecord:    false,
								GroupedStepIndex: 1,
							},
							{
								ExposureRatio:    1,
								IntervalMs:       *datadog.NewNullableInt64(nil),
								IsPauseRecord:    false,
								GroupedStepIndex: 2,
							},
						},
					},
					GuardrailMetrics: []datadogV2.GuardrailMetricRequest{},
					Type:             datadogV2.ALLOCATIONTYPE_CANARY,
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFeatureFlagsApi(apiClient)
	resp, r, err := api.UpdateAllocationsForFeatureFlagInEnvironment(ctx, FeatureFlagDataID, EnvironmentDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.UpdateAllocationsForFeatureFlagInEnvironment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.UpdateAllocationsForFeatureFlagInEnvironment`:\n%s\n", responseContent)
}
