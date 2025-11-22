// Create a WAF Policy returns "Created" response

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
	body := datadogV2.ApplicationSecurityPolicyCreateRequest{
		Data: datadogV2.ApplicationSecurityPolicyCreateData{
			Attributes: datadogV2.ApplicationSecurityPolicyCreateAttributes{
				BasedOn:     "recommended",
				Description: "Policy applied to internal web applications.",
				IsDefault:   datadog.PtrBool(false),
				Name:        "Internal Network Policy",
				ProtectionPresets: []string{
					"attack-tools",
				},
				Rules: []datadogV2.ApplicationSecurityPolicyRuleOverride{
					{
						Blocking: false,
						Enabled:  true,
						Id:       "rasp-001-002",
					},
				},
				Scope: []datadogV2.ApplicationSecurityPolicyScope{
					{
						Env:     "prod",
						Service: "billing-service",
					},
				},
				Version: datadog.PtrInt64(0),
			},
			Type: datadogV2.APPLICATIONSECURITYPOLICYTYPE_POLICY,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewApplicationSecurityApi(apiClient)
	resp, r, err := api.CreateApplicationSecurityWafPolicy(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSecurityApi.CreateApplicationSecurityWafPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSecurityApi.CreateApplicationSecurityWafPolicy`:\n%s\n", responseContent)
}
