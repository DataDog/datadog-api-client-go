// Get a user invitation returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// the "user" has a "user_invitation"
	UserInvitationID := os.Getenv("USER_INVITATION_ID")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.UsersApi(apiClient)
	resp, r, err := api.GetInvitation(ctx, UserInvitationID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetInvitation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetInvitation`:\n%s\n", responseContent)
}
