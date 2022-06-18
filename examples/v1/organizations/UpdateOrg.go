// Update your organization returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.Organization{
		Billing: &datadog.OrganizationBilling{
			Type: datadog.PtrString("parent_billing"),
		},
		Description: datadog.PtrString("some description"),
		Name:        datadog.PtrString("New child org"),
		PublicId:    datadog.PtrString("abcdef12345"),
		Settings: &datadog.OrganizationSettings{
			PrivateWidgetShare: datadog.PtrBool(false),
			Saml: &datadog.OrganizationSettingsSaml{
				Enabled: datadog.PtrBool(false),
			},
			SamlAutocreateAccessRole: datadog.ACCESSROLE_STANDARD.Ptr(),
			SamlAutocreateUsersDomains: &datadog.OrganizationSettingsSamlAutocreateUsersDomains{
				Domains: []string{
					"example.com",
				},
				Enabled: datadog.PtrBool(false),
			},
			SamlCanBeEnabled: datadog.PtrBool(false),
			SamlIdpEndpoint:  datadog.PtrString("https://my.saml.endpoint"),
			SamlIdpInitiatedLogin: &datadog.OrganizationSettingsSamlIdpInitiatedLogin{
				Enabled: datadog.PtrBool(false),
			},
			SamlIdpMetadataUploaded: datadog.PtrBool(false),
			SamlLoginUrl:            datadog.PtrString("https://my.saml.login.url"),
			SamlStrictMode: &datadog.OrganizationSettingsSamlStrictMode{
				Enabled: datadog.PtrBool(false),
			},
		},
		Subscription: &datadog.OrganizationSubscription{
			Type: datadog.PtrString("pro"),
		},
		Trial: datadog.PtrBool(false),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsApi.UpdateOrg(ctx, "abc123", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrg`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrg`:\n%s\n", responseContent)
}
