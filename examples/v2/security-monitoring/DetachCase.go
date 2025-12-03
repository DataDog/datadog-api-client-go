// Detach security findings from their case returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.DetachCaseRequest{
		Data: &datadogV2.DetachCaseRequestData{
			Relationships: &datadogV2.DetachCaseRequestDataRelationships{
				Findings: datadogV2.Findings{
					Data: []datadogV2.FindingData{
						{
							Id:   "YzM2MTFjYzcyNmY0Zjg4MTAxZmRlNjQ1MWU1ZGQwYzR-YzI5NzE5Y2Y4MzU4ZjliNzhkNjYxNTY0ODIzZDQ2YTM=",
							Type: datadogV2.FINDINGDATATYPE_FINDINGS,
						},
					},
				},
			},
			Type: datadogV2.CASEDATATYPE_CASES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	r, err := api.DetachCase(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.DetachCase`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
