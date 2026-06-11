// Update identity provider overrides for a user returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.UpdateUserIdentityProvidersRequest{
		Data: []datadogV2.UserRelationshipIdentityProviderData{
			{
				Id:   "00000000-0000-0000-0000-000000000001",
				Type: datadogV2.USERRELATIONSHIPIDENTITYPROVIDERDATATYPE_IDENTITY_PROVIDERS,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewUsersApi(apiClient)
	r, err := api.UpdateUserIdentityProviders(ctx, "00000000-0000-9999-0000-000000000000", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateUserIdentityProviders`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
