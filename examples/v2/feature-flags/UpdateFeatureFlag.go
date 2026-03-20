// Update a feature flag returns "OK" response

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
	FeatureFlagDataID := uuid.MustParse(os.Getenv("FEATURE_FLAG_DATA_ID"))

	body := datadogV2.UpdateFeatureFlagRequest{
		Data: datadogV2.UpdateFeatureFlagData{
			Type: datadogV2.UPDATEFEATUREFLAGDATATYPE_FEATURE_FLAGS,
			Attributes: datadogV2.UpdateFeatureFlagAttributes{
				Description: datadog.PtrString("Updated description for the feature flag"),
				Name:        datadog.PtrString("Updated Test Feature Flag Example-Feature-Flag"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFeatureFlagsApi(apiClient)
	resp, r, err := api.UpdateFeatureFlag(ctx, FeatureFlagDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.UpdateFeatureFlag`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.UpdateFeatureFlag`:\n%s\n", responseContent)
}
