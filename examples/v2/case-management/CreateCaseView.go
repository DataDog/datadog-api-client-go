// Create a case view returns "Created" response

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
	body := datadogV2.CaseViewCreateRequest{
		Data: datadogV2.CaseViewCreate{
			Attributes: datadogV2.CaseViewCreateAttributes{
				Name:      "Open bugs",
				ProjectId: "e555e290-ed65-49bd-ae18-8acbfcf18db7",
				Query:     "status:open type:bug",
			},
			Type: datadogV2.CASEVIEWRESOURCETYPE_VIEW,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	resp, r, err := api.CreateCaseView(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.CreateCaseView`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementApi.CreateCaseView`:\n%s\n", responseContent)
}
