// Update an existing Action Connection returns "Successfully updated an Action Connection." response

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
			Attributes: datadogV2.ActionConnectionAttributesUpdate{
				Integration: &datadogV2.ActionConnectionIntegrationUpdate{
					AWSIntegrationUpdate: &datadogV2.AWSIntegrationUpdate{
						Credentials: &datadogV2.AWSCredentialsUpdate{
							AWSAssumeRoleUpdate: &datadogV2.AWSAssumeRoleUpdate{
								AccountId: datadog.PtrString("111222333444"),
								Role:      datadog.PtrString("my-role"),
								Type:      datadogV2.AWSASSUMEROLETYPE_AWSASSUMEROLE,
							}},
						Type: datadogV2.AWSINTEGRATIONTYPE_AWS,
					}},
				Name: datadog.PtrString("My AWS Connection"),
			},
			Type: datadogV2.ACTIONCONNECTIONDATATYPE_ACTION_CONNECTION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewActionConnectionApi(apiClient)
	resp, r, err := api.UpdateActionConnection(ctx, "connection_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionConnectionApi.UpdateActionConnection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ActionConnectionApi.UpdateActionConnection`:\n%s\n", responseContent)
}
