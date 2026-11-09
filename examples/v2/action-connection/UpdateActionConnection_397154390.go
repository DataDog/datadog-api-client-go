// Update a GitHub App Action Connection without replacing its secret

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
	// there is a valid "github_app_action_connection" in the system
	GithubAppActionConnectionDataID := os.Getenv("GITHUB_APP_ACTION_CONNECTION_DATA_ID")

	body := datadogV2.UpdateActionConnectionRequest{
		Data: datadogV2.ActionConnectionDataUpdate{
			Type: datadogV2.ACTIONCONNECTIONDATATYPE_ACTION_CONNECTION,
			Attributes: datadogV2.ActionConnectionAttributesUpdate{
				Name: datadog.PtrString("GitHub App Updated exampleactionconnection"),
				Integration: &datadogV2.ActionConnectionIntegrationUpdate{
					GithubIntegrationUpdate: &datadogV2.GithubIntegrationUpdate{
						Type: datadogV2.GITHUBINTEGRATIONTYPE_GITHUB,
						Credentials: &datadogV2.GithubCredentialsUpdate{
							GithubAppUpdate: &datadogV2.GithubAppUpdate{
								Type: datadogV2.GITHUBAPPTYPE_GITHUBAPP,
								Hostname: &datadogV2.GithubAppUpdateHostname{
									ActionConnectionLiteralValue: datadog.PtrString("github.updated.example.com")},
							}},
					}},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewActionConnectionApi(apiClient)
	resp, r, err := api.UpdateActionConnection(ctx, GithubAppActionConnectionDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionConnectionApi.UpdateActionConnection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ActionConnectionApi.UpdateActionConnection`:\n%s\n", responseContent)
}
