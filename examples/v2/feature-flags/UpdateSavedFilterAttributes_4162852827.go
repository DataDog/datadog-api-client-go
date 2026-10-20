// Patch a saved filter clears the description with an empty string

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
				Description: *datadog.NewNullableString(datadog.PtrString("")),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFeatureFlagsApi(apiClient)
	resp, r, err := api.UpdateSavedFilterAttributes(ctx, uuid.MustParse("2cfc7949-7eb4-4d76-b1b8-d1aa9b565535"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.UpdateSavedFilterAttributes`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.UpdateSavedFilterAttributes`:\n%s\n", responseContent)
}
