// Edit an AuthN Mapping returns "OK" response

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
	// there is a valid "authn_mapping" in the system
	AuthnMappingDataID := os.Getenv("AUTHN_MAPPING_DATA_ID")

	// there is a valid "role" in the system
	RoleDataID := os.Getenv("ROLE_DATA_ID")

	body := datadogV2.AuthNMappingUpdateRequest{
		Data: datadogV2.AuthNMappingUpdateData{
			Attributes: &datadogV2.AuthNMappingUpdateAttributes{
				AttributeKey:   datadog.PtrString("member-of"),
				AttributeValue: datadog.PtrString("Development"),
			},
			Id: AuthnMappingDataID,
			Relationships: &datadogV2.AuthNMappingUpdateRelationships{
				AuthNMappingRelationshipToRole: &datadogV2.AuthNMappingRelationshipToRole{
					Role: datadogV2.RelationshipToRole{
						Data: &datadogV2.RelationshipToRoleData{
							Id:   datadog.PtrString(RoleDataID),
							Type: datadogV2.ROLESTYPE_ROLES.Ptr(),
						},
					},
				}},
			Type: datadogV2.AUTHNMAPPINGSTYPE_AUTHN_MAPPINGS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAuthNMappingsApi(apiClient)
	resp, r, err := api.UpdateAuthNMapping(ctx, AuthnMappingDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthNMappingsApi.UpdateAuthNMapping`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AuthNMappingsApi.UpdateAuthNMapping`:\n%s\n", responseContent)
}
