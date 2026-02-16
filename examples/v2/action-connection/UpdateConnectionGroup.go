// Update a connection group returns "OK" response

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
	body := datadogV2.UpdateConnectionGroupRequest{
		Data: datadogV2.ConnectionGroupDataRequest{
			Attributes: &datadogV2.ConnectionGroupDataAttributesRequest{
				Connections: []string{},
				Description: datadog.PtrString("An updated test connection group for AWS integrations"),
				Name:        datadog.PtrString("My Connection Group Updated"),
				TagKeys:     []string{},
			},
			Type: datadogV2.CONNECTIONGROUPTYPE_CONNECTION_GROUP,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateConnectionGroup", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewActionConnectionApi(apiClient)
	resp, r, err := api.UpdateConnectionGroup(ctx, "connection_group_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionConnectionApi.UpdateConnectionGroup`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ActionConnectionApi.UpdateConnectionGroup`:\n%s\n", responseContent)
}
