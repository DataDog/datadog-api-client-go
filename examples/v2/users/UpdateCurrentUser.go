// Update current user returns "OK" response

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
	body := datadogV2.UserUpdateRequest{
		Data: datadogV2.UserUpdateData{
			Attributes: datadogV2.UserUpdateAttributes{
				Title: *datadog.NewNullableString(nil),
			},
			Id:   "00000000-0000-feed-0000-000000000000",
			Type: datadogV2.USERSTYPE_USERS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewUsersApi(apiClient)
	resp, r, err := api.UpdateCurrentUser(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateCurrentUser`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsersApi.UpdateCurrentUser`:\n%s\n", responseContent)
}
