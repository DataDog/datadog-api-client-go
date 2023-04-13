// Create a service account returns "OK" response

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

	body := datadogV2.ServiceAccountCreateRequest{
		Data: datadogV2.ServiceAccountCreateData{
			Type: datadogV2.USERSTYPE_USERS,
			Attributes: datadogV2.ServiceAccountCreateAttributes{
				Name:           datadog.PtrString("Test API Client"),
				Email:          "Example-Service-Account@datadoghq.com",
				ServiceAccount: true,
			},
			Relationships: &datadogV2.UserRelationships{
				Roles: &datadogV2.RelationshipToRoles{
					Data: []datadogV2.RelationshipToRoleData{
						{
							Id:   datadog.PtrString(RoleDataID),
							Type: datadogV2.ROLESTYPE_ROLES.Ptr(),
						},
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceAccountsApi(apiClient)
	resp, r, err := api.CreateServiceAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.CreateServiceAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.CreateServiceAccount`:\n%s\n", responseContent)
}
