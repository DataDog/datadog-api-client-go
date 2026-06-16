// Update organization SAML preferences returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.OrgSAMLPreferencesUpdateRequest{
		Data: datadogV2.OrgSAMLPreferencesData{
			Attributes: datadogV2.OrgSAMLPreferencesAttributes{
				DefaultRoleUuids: []uuid.UUID{
					uuid.MustParse("8dd1cf3c-0c75-11ea-ad28-fb5701eabc7d"),
				},
				JitDomains: []string{
					"example.com",
				},
			},
			Id:   datadog.PtrString("00000000-0000-0000-0000-000000000000"),
			Type: datadogV2.ORGSAMLPREFERENCESTYPE_SAML_PREFERENCES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateOrgSamlConfigurations", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrganizationsApi(apiClient)
	r, err := api.UpdateOrgSamlConfigurations(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrgSamlConfigurations`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
