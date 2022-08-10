// Send invitation emails returns "OK" response

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
	// there is a valid "user" in the system
	UserDataID := os.Getenv("USER_DATA_ID")

	body := datadogV2.UserInvitationsRequest{
		Data: []datadogV2.UserInvitationData{
			{
				Type: datadogV2.USERINVITATIONSTYPE_USER_INVITATIONS,
				Relationships: datadogV2.UserInvitationRelationships{
					User: datadogV2.RelationshipToUser{
						Data: datadogV2.RelationshipToUserData{
							Type: datadogV2.USERSTYPE_USERS,
							Id:   UserDataID,
						},
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewUsersApi(apiClient)
	resp, r, err := api.SendInvitations(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.SendInvitations`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsersApi.SendInvitations`:\n%s\n", responseContent)
}
