// Create bandit returns "Created" response

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
	FeatureFlagDataAttributesVariants1ID := uuid.MustParse(os.Getenv("FEATURE_FLAG_DATA_ATTRIBUTES_VARIANTS_1_ID"))
	FeatureFlagDataID := uuid.MustParse(os.Getenv("FEATURE_FLAG_DATA_ID"))

	// there is a valid "environment" in the system
	EnvironmentDataID := uuid.MustParse(os.Getenv("ENVIRONMENT_DATA_ID"))

	// there is a valid "experiment_subject_type" in the system
	ExperimentSubjectTypeDataID := uuid.MustParse(os.Getenv("EXPERIMENT_SUBJECT_TYPE_DATA_ID"))

	body := datadogV2.ExperimentsCreateBanditV2Request{
		Data: datadogV2.ExperimentsCreateBanditV2RequestData{
			Type: datadogV2.EXPERIMENTSBANDITV2DTODATATYPE_BANDITS,
			Attributes: datadogV2.ExperimentsCreateBanditV2RequestDataAttributes{
				Name: "bandit-14bb9543f523edde",
				DatadogFlagConfiguration: *datadogV2.NewNullableExperimentsBanditFlagConfigurationInputV2(&datadogV2.ExperimentsBanditFlagConfigurationInputV2{
					FeatureFlagId:  datadog.PtrUUID(FeatureFlagDataID),
					EnvironmentId:  datadog.PtrUUID(EnvironmentDataID),
					TargetingRules: []datadogV2.ExperimentsTargetingRule{},
				}),
				Variants: []datadogV2.ExperimentsBanditVariantInputV2{
					{
						FeatureFlagVariantId: FeatureFlagDataAttributesVariants0ID,
						IsFallback:           datadog.PtrBool(true),
					},
					{
						FeatureFlagVariantId: FeatureFlagDataAttributesVariants1ID,
					},
				},
				SubjectTypeId: datadog.PtrUUID(ExperimentSubjectTypeDataID),
				Tags: []string{
					"bdd",
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateBandit", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	resp, r, err := api.CreateBandit(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.CreateBandit`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsApi.CreateBandit`:\n%s\n", responseContent)
}
