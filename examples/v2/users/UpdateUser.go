// Update a user returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "user" in the system
	UserDataID := os.Getenv("USER_DATA_ID")

	body := datadog.UserUpdateRequest{
		Data: datadog.UserUpdateData{
			Id:   UserDataID,
			Type: datadog.USERSTYPE_USERS,
			Attributes: datadog.UserUpdateAttributes{
				Name:     datadog.PtrString("updated"),
				Disabled: datadog.PtrBool(true),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersApi.UpdateUser(ctx, UserDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateUser`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsersApi.UpdateUser`:\n%s\n", responseContent)
}
