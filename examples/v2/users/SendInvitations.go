// Send invitation emails returns "OK" response

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
	// there is a valid "user" in the system
	UserDataID := os.Getenv("USER_DATA_ID")

	body := datadog.UserInvitationsRequest{
		Data: []datadog.UserInvitationData{
			{
				Type: datadog.USERINVITATIONSTYPE_USER_INVITATIONS,
				Relationships: datadog.UserInvitationRelationships{
					User: datadog.RelationshipToUser{
						Data: datadog.RelationshipToUserData{
							Type: datadog.USERSTYPE_USERS,
							Id:   UserDataID,
						},
					},
				},
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewUsersApi(apiClient)
	resp, r, err := api.SendInvitations(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.SendInvitations`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsersApi.SendInvitations`:\n%s\n", responseContent)
}
