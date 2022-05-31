// Edit an AuthN Mapping returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "authn_mapping" in the system
	AuthnMappingDataID := os.Getenv("AUTHN_MAPPING_DATA_ID")

	// there is a valid "role" in the system
	RoleDataID := os.Getenv("ROLE_DATA_ID")

	body := datadog.AuthNMappingUpdateRequest{
		Data: datadog.AuthNMappingUpdateData{
			Attributes: &datadog.AuthNMappingUpdateAttributes{
				AttributeKey:   datadog.PtrString("member-of"),
				AttributeValue: datadog.PtrString("Development"),
			},
			Id: AuthnMappingDataID,
			Relationships: &datadog.AuthNMappingUpdateRelationships{
				Role: &datadog.RelationshipToRole{
					Data: &datadog.RelationshipToRoleData{
						Id:   datadog.PtrString(RoleDataID),
						Type: datadog.ROLESTYPE_ROLES.Ptr(),
					},
				},
			},
			Type: datadog.AUTHNMAPPINGSTYPE_AUTHN_MAPPINGS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthNMappingsApi.UpdateAuthNMapping(ctx, AuthnMappingDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthNMappingsApi.UpdateAuthNMapping`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AuthNMappingsApi.UpdateAuthNMapping`:\n%s\n", responseContent)
}
