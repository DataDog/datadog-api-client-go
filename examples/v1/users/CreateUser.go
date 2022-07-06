// Create a user returns "User created" response

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
		Disabled:   datadog.PtrBool(false),
		Email:      datadog.PtrString("test@datadoghq.com"),
		Handle:     datadog.PtrString("test@datadoghq.com"),
		Name:       datadog.PtrString("test user"),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.UsersApi(apiClient)
	resp, r, err := api.CreateUser(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreateUser`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsersApi.CreateUser`:\n%s\n", responseContent)
}
