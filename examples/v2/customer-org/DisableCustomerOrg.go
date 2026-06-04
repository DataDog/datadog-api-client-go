// Disable the authenticated customer organization returns "OK" response

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
	body := datadogV2.CustomerOrgDisableRequest{
		Data: datadogV2.CustomerOrgDisableRequestData{
			Attributes: &datadogV2.CustomerOrgDisableRequestAttributes{
				OrgUuid: datadog.PtrString("abcdef01-2345-6789-abcd-ef0123456789"),
			},
			Id:   datadog.PtrString("1"),
			Type: datadogV2.CUSTOMERORGDISABLETYPE_CUSTOMER_ORG_DISABLE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DisableCustomerOrg", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCustomerOrgApi(apiClient)
	resp, r, err := api.DisableCustomerOrg(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerOrgApi.DisableCustomerOrg`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CustomerOrgApi.DisableCustomerOrg`:\n%s\n", responseContent)
}
