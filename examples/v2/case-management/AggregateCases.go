// Aggregate cases returns "OK" response

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
	body := datadogV2.CaseAggregateRequest{
		Data: datadogV2.CaseAggregateRequestData{
			Attributes: datadogV2.CaseAggregateRequestAttributes{
				GroupBy: datadogV2.CaseAggregateGroupBy{
					Groups: []string{
						"status",
					},
					Limit: 14,
				},
				QueryFilter: "service:case-api",
			},
			Type: datadogV2.CASEAGGREGATERESOURCETYPE_AGGREGATE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.AggregateCases", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	resp, r, err := api.AggregateCases(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.AggregateCases`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementApi.AggregateCases`:\n%s\n", responseContent)
}
