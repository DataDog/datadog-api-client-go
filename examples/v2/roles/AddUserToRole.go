// Add a user to a role returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "role" in the system
	RoleDataID := os.Getenv("ROLE_DATA_ID")

	// there is a valid "user" in the system
	UserDataID := os.Getenv("USER_DATA_ID")

	body := datadog.RelationshipToUser{
		Data: datadog.RelationshipToUserData{
			Id:   UserDataID,
			Type: datadog.USERSTYPE_USERS,
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.RolesApi(apiClient)
	resp, r, err := api.AddUserToRole(ctx, RoleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.AddUserToRole`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RolesApi.AddUserToRole`:\n%s\n", responseContent)
}
