// Enable a feature flag in an environment returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "feature_flag" in the system
	FeatureFlagDataID := uuid.MustParse(os.Getenv("FEATURE_FLAG_DATA_ID"))

	// there is a valid "environment" in the system
	EnvironmentDataID := uuid.MustParse(os.Getenv("ENVIRONMENT_DATA_ID"))

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFeatureFlagsApi(apiClient)
	r, err := api.EnableFeatureFlagEnvironment(ctx, FeatureFlagDataID, EnvironmentDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.EnableFeatureFlagEnvironment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
