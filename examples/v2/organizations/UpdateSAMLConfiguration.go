// Update a SAML configuration returns "OK" response

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
	body := datadogV2.SAMLConfigurationUpdateRequest{
		Data: datadogV2.SAMLConfigurationUpdateData{
			Attributes: &datadogV2.SAMLConfigurationUpdateAttributes{
				IdpInitiated: datadog.PtrBool(true),
				JitDomains: []string{
					"example.com",
				},
			},
			Id: "3653d3c6-0c75-11ea-ad28-fb5701eabc7d",
			Relationships: &datadogV2.SAMLConfigurationRelationships{
				DefaultRoles: &datadogV2.RelationshipToRoles{
					Data: []datadogV2.RelationshipToRoleData{
						{
							Id:   datadog.PtrString("3653d3c6-0c75-11ea-ad28-fb5701eabc7d"),
							Type: datadogV2.ROLESTYPE_ROLES.Ptr(),
						},
					},
				},
			},
			Type: datadogV2.SAMLCONFIGURATIONSTYPE_SAML_CONFIGURATIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrganizationsApi(apiClient)
	resp, r, err := api.UpdateSAMLConfiguration(ctx, "3653d3c6-0c75-11ea-ad28-fb5701eabc7d", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateSAMLConfiguration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateSAMLConfiguration`:\n%s\n", responseContent)
}
