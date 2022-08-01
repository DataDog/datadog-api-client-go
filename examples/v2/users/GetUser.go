// Get user details returns "OK for get user" response

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
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewUsersApi(apiClient)
	resp, r, err := api.GetUser(ctx, "00000000-0000-9999-0000-000000000000")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUser`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUser`:\n%s\n", responseContent)
}
