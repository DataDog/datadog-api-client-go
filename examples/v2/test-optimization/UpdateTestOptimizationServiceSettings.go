// Update Test Optimization service settings returns "OK" response

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
	body := datadogV2.TestOptimizationUpdateServiceSettingsRequest{
		Data: datadogV2.TestOptimizationUpdateServiceSettingsRequestData{
			Attributes: datadogV2.TestOptimizationUpdateServiceSettingsRequestAttributes{
				AutoTestRetriesEnabled:     datadog.PtrBool(false),
				CodeCoverageEnabled:        datadog.PtrBool(false),
				EarlyFlakeDetectionEnabled: datadog.PtrBool(false),
				Env:                        datadog.PtrString("prod"),
				FailedTestReplayEnabled:    datadog.PtrBool(false),
				PrCommentsEnabled:          datadog.PtrBool(true),
				RepositoryId:               "github.com/datadog/shopist",
				ServiceName:                "shopist",
				TestImpactAnalysisEnabled:  datadog.PtrBool(false),
			},
			Type: datadogV2.TESTOPTIMIZATIONUPDATESERVICESETTINGSREQUESTDATATYPE_TEST_OPTIMIZATION_UPDATE_SERVICE_SETTINGS_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTestOptimizationApi(apiClient)
	resp, r, err := api.UpdateTestOptimizationServiceSettings(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestOptimizationApi.UpdateTestOptimizationServiceSettings`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TestOptimizationApi.UpdateTestOptimizationServiceSettings`:\n%s\n", responseContent)
}
