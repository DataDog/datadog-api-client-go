// Create a restriction query returns "OK" response

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
	body := datadogV2.RestrictionQueryCreatePayload{
		Data: &datadogV2.RestrictionQueryCreateData{
			Attributes: &datadogV2.RestrictionQueryCreateAttributes{
				RestrictionQuery: datadog.PtrString("env:sandbox"),
			},
			Type: datadogV2.LOGSRESTRICTIONQUERIESTYPE_LOGS_RESTRICTION_QUERIES.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateRestrictionQuery", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsRestrictionQueriesApi(apiClient)
	resp, r, err := api.CreateRestrictionQuery(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsRestrictionQueriesApi.CreateRestrictionQuery`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsRestrictionQueriesApi.CreateRestrictionQuery`:\n%s\n", responseContent)
}
