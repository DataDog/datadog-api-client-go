// Get a user invitation returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// the "user" has a "user_invitation"
	UserInvitationID := os.Getenv("USER_INVITATION_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersApi.GetInvitation(ctx, UserInvitationID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetInvitation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetInvitation`:\n%s\n", responseContent)
}
