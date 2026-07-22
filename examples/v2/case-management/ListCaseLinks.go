// List case links returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	resp, r, err := api.ListCaseLinks(ctx, "CASE", "bf0cbac6-4c16-4cfb-b6bf-ca5e0ec37a4f", *datadogV2.NewListCaseLinksOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.ListCaseLinks`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementApi.ListCaseLinks`:\n%s\n", responseContent)
}
