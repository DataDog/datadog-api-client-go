// Update a GitHub personal access token Action Connection without replacing its secret

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
	// there is a valid "github_pat_action_connection" in the system
	GithubPatActionConnectionDataID := os.Getenv("GITHUB_PAT_ACTION_CONNECTION_DATA_ID")

	body := datadogV2.UpdateActionConnectionRequest{
		Data: datadogV2.ActionConnectionDataUpdate{
			Type: datadogV2.ACTIONCONNECTIONDATATYPE_ACTION_CONNECTION,
			Attributes: datadogV2.ActionConnectionAttributesUpdate{
				Name: datadog.PtrString("GitHub personal access token Updated exampleactionconnection"),
				Integration: &datadogV2.ActionConnectionIntegrationUpdate{
					GithubIntegrationUpdate: &datadogV2.GithubIntegrationUpdate{
						Type: datadogV2.GITHUBINTEGRATIONTYPE_GITHUB,
						Credentials: &datadogV2.GithubCredentialsUpdate{
							GithubPersonalAccessTokenUpdate: &datadogV2.GithubPersonalAccessTokenUpdate{
								Type: datadogV2.GITHUBPERSONALACCESSTOKENTYPE_GITHUBPERSONALACCESSTOKEN,
								Hostname: &datadogV2.GithubPersonalAccessTokenUpdateHostname{
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
	resp, r, err := api.UpdateActionConnection(ctx, GithubPatActionConnectionDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionConnectionApi.UpdateActionConnection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ActionConnectionApi.UpdateActionConnection`:\n%s\n", responseContent)
}
