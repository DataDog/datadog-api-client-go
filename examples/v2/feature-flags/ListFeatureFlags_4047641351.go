// List feature flags combines exact keys tags search types creators and IDs

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
	// there is a valid "feature_flag_list_a" in the system
	FeatureFlagListADataAttributesCreatedBy := uuid.MustParse(os.Getenv("FEATURE_FLAG_LIST_A_DATA_ATTRIBUTES_CREATED_BY"))
	FeatureFlagListADataID := uuid.MustParse(os.Getenv("FEATURE_FLAG_LIST_A_DATA_ID"))

	// there is a valid "feature_flag_list_b" in the system
	FeatureFlagListBDataID := uuid.MustParse(os.Getenv("FEATURE_FLAG_LIST_B_DATA_ID"))

	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFeatureFlagsApi(apiClient)
	resp, r, err := api.ListFeatureFlags(ctx, *datadogV2.NewListFeatureFlagsOptionalParameters().WithKey("list-a-example-feature-flag").WithTags([]string{
		"list-test:example-feature-flag",
		"color:blue",
	}).WithSearch("LIST-A-example-feature-flag").WithValueType([]datadogV2.ValueType{
		datadogV2.VALUETYPE_BOOLEAN,
		datadogV2.VALUETYPE_STRING,
	}).WithCreatedBy([]uuid.UUID{
		FeatureFlagListADataAttributesCreatedBy,
	}).WithFlagIds([]uuid.UUID{
		FeatureFlagListADataID,
		FeatureFlagListBDataID,
	}).WithIsArchived(false))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.ListFeatureFlags`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.ListFeatureFlags`:\n%s\n", responseContent)
}
