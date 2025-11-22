// Update a restriction query returns "OK" response

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
	// there is a valid "restriction_query" in the system
	RestrictionQueryDataID := os.Getenv("RESTRICTION_QUERY_DATA_ID")

	body := datadogV2.RestrictionQueryUpdatePayload{
		Data: &datadogV2.RestrictionQueryUpdateData{
			Attributes: &datadogV2.RestrictionQueryUpdateAttributes{
				RestrictionQuery: "env:production",
			},
			Type: datadogV2.LOGSRESTRICTIONQUERIESTYPE_LOGS_RESTRICTION_QUERIES.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateRestrictionQuery", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsRestrictionQueriesApi(apiClient)
	resp, r, err := api.UpdateRestrictionQuery(ctx, RestrictionQueryDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsRestrictionQueriesApi.UpdateRestrictionQuery`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsRestrictionQueriesApi.UpdateRestrictionQuery`:\n%s\n", responseContent)
}
