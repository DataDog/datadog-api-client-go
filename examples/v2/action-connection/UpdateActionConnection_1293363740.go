// Update an existing Action Connection with HTTPBasicAuth returns "Successfully updated Action Connection" response

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
	body := datadogV2.UpdateActionConnectionRequest{
		Data: datadogV2.ActionConnectionDataUpdate{
			Type: datadogV2.ACTIONCONNECTIONDATATYPE_ACTION_CONNECTION,
			Attributes: datadogV2.ActionConnectionAttributesUpdate{
				Name: datadog.PtrString("HTTP Basic Auth Updated"),
				Integration: &datadogV2.ActionConnectionIntegrationUpdate{
					HTTPIntegrationUpdate: &datadogV2.HTTPIntegrationUpdate{
						Type:    datadogV2.HTTPINTEGRATIONTYPE_HTTP,
						BaseUrl: datadog.PtrString("https://api.updated.com"),
						Credentials: &datadogV2.HTTPCredentialsUpdate{
							HTTPBasicAuthUpdate: &datadogV2.HTTPBasicAuthUpdate{
								Type:     datadogV2.HTTPBASICAUTHTYPE_HTTPBASICAUTH,
								Username: datadog.PtrString("updated-user"),
								Password: datadog.PtrString("updated-password"),
							}},
					}},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewActionConnectionApi(apiClient)
	resp, r, err := api.UpdateActionConnection(ctx, "cb460d51-3c88-4e87-adac-d47131d0423d", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionConnectionApi.UpdateActionConnection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ActionConnectionApi.UpdateActionConnection`:\n%s\n", responseContent)
}
