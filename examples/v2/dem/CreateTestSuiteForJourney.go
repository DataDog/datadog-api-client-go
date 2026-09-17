// Create a test suite for a DEM journey returns "Created" response

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
	body := datadogV2.DemCreateJourneyTestSuiteRequest{
		Data: datadogV2.DemCreateJourneyTestSuiteData{
			Attributes: &datadogV2.DemCreateJourneyTestSuiteAttributes{
				IncludeTestsFromJourneyCoverage: *datadog.NewNullableBool(datadog.PtrBool(true)),
				TestSuiteName:                   *datadog.NewNullableString(datadog.PtrString("My Custom Suite")),
			},
			Type: datadogV2.DEMCREATEJOURNEYTESTSUITEREQUESTTYPE_CREATE_TEST_SUITE_FOR_JOURNEY_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDEMApi(apiClient)
	resp, r, err := api.CreateTestSuiteForJourney(ctx, "public_journey_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DEMApi.CreateTestSuiteForJourney`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DEMApi.CreateTestSuiteForJourney`:\n%s\n", responseContent)
}
