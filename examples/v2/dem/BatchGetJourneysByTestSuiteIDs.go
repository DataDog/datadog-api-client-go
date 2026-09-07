// Batch get DEM journeys by test suite IDs returns "OK" response

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
	body := datadogV2.DemBatchGetJourneysRequest{
		Data: datadogV2.DemBatchGetJourneysData{
			Attributes: datadogV2.DemBatchGetJourneysAttributes{
				TestSuiteIds: []string{
					"suite-abc123",
					"suite-def456",
				},
			},
			Type: datadogV2.DEMBATCHGETJOURNEYSREQUESTTYPE_BATCH_GET_JOURNEYS_BY_TEST_SUITE_IDS_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDEMApi(apiClient)
	resp, r, err := api.BatchGetJourneysByTestSuiteIDs(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DEMApi.BatchGetJourneysByTestSuiteIDs`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DEMApi.BatchGetJourneysByTestSuiteIDs`:\n%s\n", responseContent)
}
