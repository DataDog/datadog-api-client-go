// Create a feature flag returns "Created" response

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
	body := datadogV2.CreateFeatureFlagRequest{
		Data: datadogV2.CreateFeatureFlagData{
			Type: datadogV2.CREATEFEATUREFLAGDATATYPE_FEATURE_FLAGS,
			Attributes: datadogV2.CreateFeatureFlagAttributes{
				DefaultVariantKey: *datadog.NewNullableString(datadog.PtrString("variant-Example-Feature-Flag-1")),
				Description:       "Test feature flag for BDD scenarios",
				Key:               "test-feature-flag-Example-Feature-Flag",
				Name:              "Test Feature Flag Example-Feature-Flag",
				ValueType:         datadogV2.VALUETYPE_BOOLEAN,
				Variants: []datadogV2.CreateVariant{
					{
						Key:   "variant-Example-Feature-Flag-1",
						Name:  "Variant Example-Feature-Flag A",
						Value: "true",
					},
					{
						Key:   "variant-Example-Feature-Flag-2",
						Name:  "Variant Example-Feature-Flag B",
						Value: "false",
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFeatureFlagsApi(apiClient)
	resp, r, err := api.CreateFeatureFlag(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.CreateFeatureFlag`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.CreateFeatureFlag`:\n%s\n", responseContent)
}
