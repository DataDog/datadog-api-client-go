// Update a specific Org Config returns "OK" response

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
	body := datadogV2.OrgConfigWriteRequest{
		Data: datadogV2.OrgConfigWrite{
			Attributes: datadogV2.OrgConfigWriteAttributes{
				Value: "UTC",
			},
			Type: datadogV2.ORGCONFIGTYPE_ORG_CONFIGS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrganizationsApi(apiClient)
	resp, r, err := api.UpdateOrgConfig(ctx, "monitor_timezone", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrgConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrgConfig`:\n%s\n", responseContent)
}
