// Update flaky test states returns "OK" response

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
	body := datadogV2.UpdateFlakyTestsRequest{
		Data: datadogV2.UpdateFlakyTestsRequestData{
			Attributes: datadogV2.UpdateFlakyTestsRequestAttributes{
				Tests: []datadogV2.UpdateFlakyTestsRequestTest{
					{
						Id:       "4eb1887a8adb1847",
						NewState: datadogV2.UPDATEFLAKYTESTSREQUESTTESTNEWSTATE_ACTIVE,
					},
				},
			},
			Type: datadogV2.UPDATEFLAKYTESTSREQUESTDATATYPE_UPDATE_FLAKY_TEST_STATE_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateFlakyTests", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTestOptimizationApi(apiClient)
	resp, r, err := api.UpdateFlakyTests(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestOptimizationApi.UpdateFlakyTests`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TestOptimizationApi.UpdateFlakyTests`:\n%s\n", responseContent)
}
