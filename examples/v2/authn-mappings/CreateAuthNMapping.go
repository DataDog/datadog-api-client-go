// Create an AuthN Mapping returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "role" in the system
	RoleDataID := os.Getenv("ROLE_DATA_ID")

	body := datadog.AuthNMappingCreateRequest{
		Data: datadog.AuthNMappingCreateData{
			Attributes: &datadog.AuthNMappingCreateAttributes{
				AttributeKey:   datadog.PtrString("examplecreateanauthnmappingreturnsokresponse"),
				AttributeValue: datadog.PtrString("Example-Create_an_AuthN_Mapping_returns_OK_response"),
			},
			Relationships: &datadog.AuthNMappingCreateRelationships{
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
	resp, r, err := apiClient.AuthNMappingsApi.CreateAuthNMapping(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthNMappingsApi.CreateAuthNMapping`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AuthNMappingsApi.CreateAuthNMapping`:\n%s\n", responseContent)
}
