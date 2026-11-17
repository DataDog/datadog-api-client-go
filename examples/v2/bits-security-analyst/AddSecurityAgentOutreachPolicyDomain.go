// Add an allowed outreach domain returns "OK" response

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
	body := datadogV2.SecurityAgentOutreachDomainAddRequest{
		Data: datadogV2.SecurityAgentOutreachDomainAddRequestData{
			Attributes: datadogV2.SecurityAgentOutreachDomainAddRequestDataAttributes{
				Domain: "acme.com",
			},
			Type: datadogV2.SECURITYAGENTOUTREACHDOMAINADDTYPE_OUTREACH_POLICY_DOMAIN_ADD,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.AddSecurityAgentOutreachPolicyDomain", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewBitsSecurityAnalystApi(apiClient)
	resp, r, err := api.AddSecurityAgentOutreachPolicyDomain(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BitsSecurityAnalystApi.AddSecurityAgentOutreachPolicyDomain`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `BitsSecurityAnalystApi.AddSecurityAgentOutreachPolicyDomain`:\n%s\n", responseContent)
}
