// Update a user returns "User updated" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.User{
		AccessRole: *datadogV1.NewNullableAccessRole(datadogV1.ACCESSROLE_READ_ONLY.Ptr()),
		Disabled:   datadog.PtrBool(false),
		Email:      datadog.PtrString("test@datadoghq.com"),
		Handle:     datadog.PtrString("test@datadoghq.com"),
		Name:       datadog.PtrString("test user"),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewUsersApi(apiClient)
	resp, r, err := api.UpdateUser(ctx, "test@datadoghq.com", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateUser`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsersApi.UpdateUser`:\n%s\n", responseContent)
}
