// Update your organization returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.Organization{
		Billing: &datadogV1.OrganizationBilling{
			Type: datadog.PtrString("parent_billing"),
		},
		Description: datadog.PtrString("some description"),
		Name:        datadog.PtrString("New child org"),
		PublicId:    datadog.PtrString("abcdef12345"),
		Settings: &datadogV1.OrganizationSettings{
			PrivateWidgetShare: datadog.PtrBool(false),
			Saml: &datadogV1.OrganizationSettingsSaml{
				Enabled: datadog.PtrBool(false),
			},
			SamlAutocreateAccessRole: *datadogV1.NewNullableAccessRole(datadogV1.ACCESSROLE_READ_ONLY.Ptr()),
			SamlAutocreateUsersDomains: &datadogV1.OrganizationSettingsSamlAutocreateUsersDomains{
				Domains: []string{
					"example.com",
				},
				Enabled: datadog.PtrBool(false),
			},
			SamlCanBeEnabled: datadog.PtrBool(false),
			SamlIdpEndpoint:  datadog.PtrString("https://my.saml.endpoint"),
			SamlIdpInitiatedLogin: &datadogV1.OrganizationSettingsSamlIdpInitiatedLogin{
				Enabled: datadog.PtrBool(false),
			},
			SamlIdpMetadataUploaded: datadog.PtrBool(false),
			SamlLoginUrl:            datadog.PtrString("https://my.saml.login.url"),
			SamlStrictMode: &datadogV1.OrganizationSettingsSamlStrictMode{
				Enabled: datadog.PtrBool(false),
			},
		},
		Subscription: &datadogV1.OrganizationSubscription{
			Type: datadog.PtrString("pro"),
		},
		Trial: datadog.PtrBool(false),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewOrganizationsApi(apiClient)
	resp, r, err := api.UpdateOrg(ctx, "abc123", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrg`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrg`:\n%s\n", responseContent)
}
