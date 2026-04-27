// Delete a pending user's invitations returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewUsersApi(apiClient)
	r, err := api.DeleteUserInvitations(ctx, uuid.MustParse("4dee724d-00cc-11ea-a77b-570c9d03c6c5"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.DeleteUserInvitations`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
