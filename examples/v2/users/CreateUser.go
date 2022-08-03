// Create a user returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v2/datadog"
)

func main() {
	body := datadog.UserCreateRequest{
		Data: datadog.UserCreateData{
			Type: datadog.USERSTYPE_USERS,
			Attributes: datadog.UserCreateAttributes{
				Name:  common.PtrString("Datadog API Client Python"),
				Email: "Example-Create_a_user_returns_OK_response@datadoghq.com",
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewUsersApi(apiClient)
	resp, r, err := api.CreateUser(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreateUser`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsersApi.CreateUser`:\n%s\n", responseContent)
}
