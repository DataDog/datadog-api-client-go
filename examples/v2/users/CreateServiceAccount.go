// Create a service account returns "OK" response

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

	body := datadog.ServiceAccountCreateRequest{
		Data: datadog.ServiceAccountCreateData{
			Type: datadog.USERSTYPE_USERS,
			Attributes: datadog.ServiceAccountCreateAttributes{
				Name:           datadog.PtrString("Test API Client"),
				Email:          "Example-Create_a_service_account_returns_OK_response@datadoghq.com",
				ServiceAccount: true,
			},
			Relationships: &datadog.UserRelationships{
				Roles: &datadog.RelationshipToRoles{
					Data: []datadog.RelationshipToRoleData{
						{
							Id:   datadog.PtrString(RoleDataID),
							Type: datadog.ROLESTYPE_ROLES.Ptr(),
						},
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersApi.CreateServiceAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreateServiceAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsersApi.CreateServiceAccount`:\n%s\n", responseContent)
}
