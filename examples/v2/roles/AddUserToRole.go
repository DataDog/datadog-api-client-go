// Add a user to a role returns "OK" response

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
			Id:   "23d05e63-b1c2-4159-0b0d-533c9645335c",
			Type: datadog.USERSTYPE_USERS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesApi.AddUserToRole(ctx, ROLE_DATA_ID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.AddUserToRole`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RolesApi.AddUserToRole`:\n%s\n", responseContent)
}
