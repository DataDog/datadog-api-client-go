// Create a new Action Connection with HTTPTokenAuth returns "Successfully created Action Connection" response

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
				Name: "HTTP Token Connection exampleactionconnection",
				Integration: datadogV2.ActionConnectionIntegration{
					HTTPIntegration: &datadogV2.HTTPIntegration{
						Type:    datadogV2.HTTPINTEGRATIONTYPE_HTTP,
						BaseUrl: "https://api.example.com",
						Credentials: datadogV2.HTTPCredentials{
							HTTPTokenAuth: &datadogV2.HTTPTokenAuth{
								Type: datadogV2.HTTPTOKENAUTHTYPE_HTTPTOKENAUTH,
								Tokens: []datadogV2.HTTPToken{
									{
										Name:  "ApiKey",
										Type:  datadogV2.TOKENTYPE_SECRET,
										Value: "secret-token-value",
									},
								},
								Headers: []datadogV2.HTTPHeader{
									{
										Name:  "Authorization",
										Value: "Bearer token-value",
									},
								},
							}},
					}},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
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
