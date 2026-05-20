// Bulk update cases returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.CaseBulkUpdateRequest{
		Data: datadogV2.CaseBulkUpdateRequestData{
			Attributes: datadogV2.CaseBulkUpdateRequestAttributes{
				CaseIds: []string{
					"case-id-1",
					"case-id-2",
				},
				Payload: map[string]string{
					"priority": "P1",
				},
				Type: datadogV2.CASEBULKACTIONTYPE_PRIORITY,
			},
			Type: datadogV2.CASEBULKRESOURCETYPE_BULK,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.BulkUpdateCases", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	r, err := api.BulkUpdateCases(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.BulkUpdateCases`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
