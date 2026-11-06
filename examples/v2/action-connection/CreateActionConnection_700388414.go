// Create a GitHub personal access token Action Connection

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
	body := datadogV2.CreateActionConnectionRequest{
		Data: datadogV2.ActionConnectionData{
			Type: datadogV2.ACTIONCONNECTIONDATATYPE_ACTION_CONNECTION,
			Attributes: datadogV2.ActionConnectionAttributes{
				Name: "GitHub personal access token Connection exampleactionconnection",
				Integration: datadogV2.ActionConnectionIntegration{
					GithubIntegration: &datadogV2.GithubIntegration{
						Type: datadogV2.GITHUBINTEGRATIONTYPE_GITHUB,
						Credentials: datadogV2.GithubCredentials{
							GithubPersonalAccessToken: &datadogV2.GithubPersonalAccessToken{
								Type: datadogV2.GITHUBPERSONALACCESSTOKENTYPE_GITHUBPERSONALACCESSTOKEN,
								AccessToken: datadogV2.GithubPersonalAccessTokenAccessToken{
									ActionConnectionLiteralValue: datadog.PtrString("example-token")},
								Hostname: &datadogV2.GithubPersonalAccessTokenHostname{
									ActionConnectionLiteralValue: datadog.PtrString("github.example.com")},
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
	resp, r, err := api.CreateActionConnection(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionConnectionApi.CreateActionConnection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ActionConnectionApi.CreateActionConnection`:\n%s\n", responseContent)
}
