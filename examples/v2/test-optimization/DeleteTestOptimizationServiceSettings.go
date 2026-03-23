// Delete Test Optimization service settings returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.TestOptimizationDeleteServiceSettingsRequest{
		Data: datadogV2.TestOptimizationDeleteServiceSettingsRequestData{
			Attributes: datadogV2.TestOptimizationDeleteServiceSettingsRequestAttributes{
				Env:          datadog.PtrString("prod"),
				RepositoryId: "github.com/datadog/shopist",
				ServiceName:  "shopist",
			},
			Type: datadogV2.TESTOPTIMIZATIONDELETESERVICESETTINGSREQUESTDATATYPE_TEST_OPTIMIZATION_DELETE_SERVICE_SETTINGS_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteTestOptimizationServiceSettings", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTestOptimizationApi(apiClient)
	r, err := api.DeleteTestOptimizationServiceSettings(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestOptimizationApi.DeleteTestOptimizationServiceSettings`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
