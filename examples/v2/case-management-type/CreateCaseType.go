// Create a case type returns "CREATED" response

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
	body := datadogV2.CaseTypeCreateRequest{
		Data: datadogV2.CaseTypeCreate{
			Attributes: datadogV2.CaseTypeResourceAttributes{
				Description: datadog.PtrString("Investigations done in case management"),
				Emoji:       datadog.PtrString("👑"),
				Name:        "Investigation",
			},
			Type: datadogV2.CASETYPERESOURCETYPE_CASE_TYPE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementTypeApi(apiClient)
	resp, r, err := api.CreateCaseType(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementTypeApi.CreateCaseType`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementTypeApi.CreateCaseType`:\n%s\n", responseContent)
}
