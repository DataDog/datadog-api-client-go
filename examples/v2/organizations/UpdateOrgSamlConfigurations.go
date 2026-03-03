// Update organization SAML preferences returns "No Content - The SAML preferences were successfully updated." response

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
	body := datadogV2.SamlConfigurationsUpdateRequest{
		Data: datadogV2.SamlConfigurationsUpdateRequestData{
			Attributes: datadogV2.SamlConfigurationsUpdateAttributes{
				DefaultRoleUuids: []uuid.UUID{
					uuid.MustParse("19fcc38b-b651-46a0-b463-1f8f56c6a862"),
				},
				JitDomains: []string{
					"example1.com",
					"example2.com",
				},
			},
			Type: datadogV2.SAMLCONFIGURATIONSUPDATEREQUESTDATATYPE_SAML_PREFERENCES,
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
