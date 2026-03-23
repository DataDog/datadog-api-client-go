// Get Test Optimization service settings returns "OK" response

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
	body := datadogV2.TestOptimizationGetServiceSettingsRequest{
		Data: datadogV2.TestOptimizationGetServiceSettingsRequestData{
			Attributes: datadogV2.TestOptimizationGetServiceSettingsRequestAttributes{
				Env:          datadog.PtrString("prod"),
				RepositoryId: "github.com/datadog/shopist",
				ServiceName:  "shopist",
			},
			Type: datadogV2.TESTOPTIMIZATIONGETSERVICESETTINGSREQUESTDATATYPE_TEST_OPTIMIZATION_GET_SERVICE_SETTINGS_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetTestOptimizationServiceSettings", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTestOptimizationApi(apiClient)
	resp, r, err := api.GetTestOptimizationServiceSettings(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestOptimizationApi.GetTestOptimizationServiceSettings`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TestOptimizationApi.GetTestOptimizationServiceSettings`:\n%s\n", responseContent)
}
