// Create an AuthN Mapping returns "OK" response

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
	// there is a valid "role" in the system
	RoleDataID := os.Getenv("ROLE_DATA_ID")

	body := datadogV2.AuthNMappingCreateRequest{
		Data: datadogV2.AuthNMappingCreateData{
			Attributes: &datadogV2.AuthNMappingCreateAttributes{
				AttributeKey:   datadog.PtrString("exampleauthnmapping"),
				AttributeValue: datadog.PtrString("Example-AuthN-Mapping"),
			},
			Relationships: &datadogV2.AuthNMappingCreateRelationships{
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
	resp, r, err := api.CreateAuthNMapping(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthNMappingsApi.CreateAuthNMapping`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AuthNMappingsApi.CreateAuthNMapping`:\n%s\n", responseContent)
}
