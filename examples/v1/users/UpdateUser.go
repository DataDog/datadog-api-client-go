// Update a user returns "User updated" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.User{
		AccessRole: datadog.ACCESSROLE_STANDARD.Ptr(),
		Disabled:   common.PtrBool(false),
		Email:      common.PtrString("test@datadoghq.com"),
		Handle:     common.PtrString("test@datadoghq.com"),
		Name:       common.PtrString("test user"),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewUsersApi(apiClient)
	resp, r, err := api.UpdateUser(ctx, "test@datadoghq.com", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateUser`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsersApi.UpdateUser`:\n%s\n", responseContent)
}
