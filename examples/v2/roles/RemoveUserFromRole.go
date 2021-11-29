// Remove a user from a role returns "OK" response

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
	ROLE_DATA_ID := os.Getenv("ROLE_DATA_ID")

	// there is a valid "user" in the system

	body := datadog.RelationshipToUser{
		Data: datadog.RelationshipToUserData{
			Id:   "a3b035bc-8978-d113-1b9c-27d99bfe4b97",
			Type: datadog.USERSTYPE_USERS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesApi.RemoveUserFromRole(ctx, ROLE_DATA_ID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.RemoveUserFromRole`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RolesApi.RemoveUserFromRole`:\n%s\n", responseContent)
}
