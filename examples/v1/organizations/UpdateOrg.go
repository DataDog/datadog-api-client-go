// Update your organization returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v1/datadog"
)

func main() {
	body := datadog.Organization{
		Billing: &datadog.OrganizationBilling{
			Type: common.PtrString("parent_billing"),
		},
		Description: common.PtrString("some description"),
		Name:        common.PtrString("New child org"),
		PublicId:    common.PtrString("abcdef12345"),
		Settings: &datadog.OrganizationSettings{
			PrivateWidgetShare: common.PtrBool(false),
			Saml: &datadog.OrganizationSettingsSaml{
				Enabled: common.PtrBool(false),
			},
			SamlAutocreateAccessRole: datadog.ACCESSROLE_STANDARD.Ptr(),
			SamlAutocreateUsersDomains: &datadog.OrganizationSettingsSamlAutocreateUsersDomains{
				Domains: []string{
					"example.com",
				},
				Enabled: common.PtrBool(false),
			},
			SamlCanBeEnabled: common.PtrBool(false),
			SamlIdpEndpoint:  common.PtrString("https://my.saml.endpoint"),
			SamlIdpInitiatedLogin: &datadog.OrganizationSettingsSamlIdpInitiatedLogin{
				Enabled: common.PtrBool(false),
			},
			SamlIdpMetadataUploaded: common.PtrBool(false),
			SamlLoginUrl:            common.PtrString("https://my.saml.login.url"),
			SamlStrictMode: &datadog.OrganizationSettingsSamlStrictMode{
				Enabled: common.PtrBool(false),
			},
		},
		Subscription: &datadog.OrganizationSubscription{
			Type: common.PtrString("pro"),
		},
		Trial: common.PtrBool(false),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewOrganizationsApi(apiClient)
	resp, r, err := api.UpdateOrg(ctx, "abc123", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrg`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrg`:\n%s\n", responseContent)
}
