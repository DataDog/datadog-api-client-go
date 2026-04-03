// Create targeting rules for a flag env returns "Created" response

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
	body := datadogV2.CreateAllocationsRequest{
		Data: datadogV2.AllocationDataRequest{
			Attributes: datadogV2.UpsertAllocationRequest{
				ExperimentId: *datadog.NewNullableString(datadog.PtrString("550e8400-e29b-41d4-a716-446655440030")),
				ExposureSchedule: &datadogV2.ExposureScheduleRequest{
					AbsoluteStartTime: *datadog.NewNullableTime(datadog.PtrTime(time.Date(2025, 6, 13, 12, 0, 0, 0, time.UTC))),
					ControlVariantId:  *datadog.NewNullableString(datadog.PtrString("550e8400-e29b-41d4-a716-446655440012")),
					ControlVariantKey: *datadog.NewNullableString(datadog.PtrString("control")),
					Id:                datadog.PtrUUID(uuid.MustParse("550e8400-e29b-41d4-a716-446655440010")),
					RolloutOptions: datadogV2.RolloutOptionsRequest{
						Autostart:           *datadog.NewNullableBool(datadog.PtrBool(false)),
						SelectionIntervalMs: datadog.PtrInt64(3600000),
						Strategy:            datadogV2.ROLLOUTSTRATEGY_UNIFORM_INTERVALS,
					},
					RolloutSteps: []datadogV2.ExposureRolloutStepRequest{
						{
							ExposureRatio:    0.5,
							GroupedStepIndex: 1,
							Id:               datadog.PtrUUID(uuid.MustParse("550e8400-e29b-41d4-a716-446655440040")),
							IntervalMs:       *datadog.NewNullableInt64(datadog.PtrInt64(3600000)),
							IsPauseRecord:    false,
						},
					},
				},
				GuardrailMetrics: []datadogV2.GuardrailMetricRequest{
					{
						MetricId:      "metric-error-rate",
						TriggerAction: datadogV2.GUARDRAILTRIGGERACTION_PAUSE,
					},
				},
				Id:   datadog.PtrUUID(uuid.MustParse("550e8400-e29b-41d4-a716-446655440020")),
				Key:  "prod-rollout",
				Name: "Production Rollout",
				TargetingRules: []datadogV2.TargetingRuleRequest{
					{
						Conditions: []datadogV2.ConditionRequest{
							{
								Attribute: "user_tier",
								Operator:  datadogV2.CONDITIONOPERATOR_ONE_OF,
								Value: []string{
									"premium",
									"enterprise",
								},
							},
						},
					},
				},
				Type: datadogV2.ALLOCATIONTYPE_FEATURE_GATE,
				VariantWeights: []datadogV2.VariantWeightRequest{
					{
						Value:      50,
						VariantId:  datadog.PtrUUID(uuid.MustParse("550e8400-e29b-41d4-a716-446655440001")),
						VariantKey: datadog.PtrString("control"),
					},
				},
			},
			Type: datadogV2.ALLOCATIONDATATYPE_ALLOCATIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFeatureFlagsApi(apiClient)
	resp, r, err := api.CreateAllocationsForFeatureFlagInEnvironment(ctx, uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"), uuid.MustParse("550e8400-e29b-41d4-a716-446655440001"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.CreateAllocationsForFeatureFlagInEnvironment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.CreateAllocationsForFeatureFlagInEnvironment`:\n%s\n", responseContent)
}
