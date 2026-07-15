// Update account filters returns "OK" response

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
	body := datadogV2.AccountFiltersPatchRequest{
		Data: datadogV2.AccountFiltersPatchData{
			Attributes: datadogV2.AccountFiltersPatchRequestAttributes{
				AccountFilters: datadogV2.AccountFilteringConfig{
					ExcludedAccounts: []string{
						"123456789123",
						"123456789143",
					},
					IncludeNewAccounts: *datadog.NewNullableBool(datadog.PtrBool(true)),
					IncludedAccounts: []string{
						"123456789123",
						"123456789143",
					},
				},
			},
			Type: datadogV2.ACCOUNTFILTERSPATCHREQUESTTYPE_ACCOUNT_FILTERS_PATCH_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.UpdateCostAccountFilters(ctx, 9223372036854775807, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.UpdateCostAccountFilters`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.UpdateCostAccountFilters`:\n%s\n", responseContent)
}
