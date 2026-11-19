// Update bandit returns "OK" response

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
	// there is a valid "bandit" in the system
	BanditDataID := uuid.MustParse(os.Getenv("BANDIT_DATA_ID"))

	// there is a valid "feature_flag" in the system
	FeatureFlagDataAttributesVariants0ID := uuid.MustParse(os.Getenv("FEATURE_FLAG_DATA_ATTRIBUTES_VARIANTS_0_ID"))
	FeatureFlagDataAttributesVariants1ID := uuid.MustParse(os.Getenv("FEATURE_FLAG_DATA_ATTRIBUTES_VARIANTS_1_ID"))

	body := datadogV2.ExperimentsPatchBanditV2Request{
		Data: datadogV2.ExperimentsPatchBanditV2RequestData{
			Type: datadogV2.EXPERIMENTSBANDITV2DTODATATYPE_BANDITS,
			Id:   datadog.PtrUUID(BanditDataID),
			Attributes: datadogV2.ExperimentsPatchBanditV2RequestDataAttributes{
				Name: datadog.PtrString("bandit-14bb9543f523edde-renamed"),
				Variants: []datadogV2.ExperimentsBanditVariantInputV2{
					{
						FeatureFlagVariantId: FeatureFlagDataAttributesVariants0ID,
					},
					{
						FeatureFlagVariantId: FeatureFlagDataAttributesVariants1ID,
						IsFallback:           datadog.PtrBool(true),
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateBandit", true)
	configuration.SetUnstableOperationEnabled("v2.PatchBandit", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	resp, r, err := api.PatchBandit(ctx, BanditDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.PatchBandit`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsApi.PatchBandit`:\n%s\n", responseContent)
}
