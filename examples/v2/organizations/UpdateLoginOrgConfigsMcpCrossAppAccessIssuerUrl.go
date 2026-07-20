// Update the MCP Cross-App Access issuer URL returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.McpCrossAppAccessIssuerUrlUpdateRequest{
		Data: datadogV2.McpCrossAppAccessIssuerUrlUpdateData{
			Attributes: datadogV2.McpCrossAppAccessIssuerUrlUpdateAttributes{
				IssuerUrl: "https://your-subdomain.okta.com",
			},
			Type: datadogV2.MCPCROSSAPPACCESSISSUERURLTYPE_ORG_CONFIG,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateLoginOrgConfigsMcpCrossAppAccessIssuerUrl", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrganizationsApi(apiClient)
	r, err := api.UpdateLoginOrgConfigsMcpCrossAppAccessIssuerUrl(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateLoginOrgConfigsMcpCrossAppAccessIssuerUrl`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
