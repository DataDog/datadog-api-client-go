// Create datastore returns "OK" response

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
	body := datadogV2.CreateAppsDatastoreRequest{
		Data: &datadogV2.CreateAppsDatastoreRequestData{
			Attributes: &datadogV2.CreateAppsDatastoreRequestDataAttributes{
				Name:              "datastore-name",
				PrimaryColumnName: "primaryKey",
			},
			Type: datadogV2.CREATEAPPSDATASTOREREQUESTDATATYPE_DATASTORES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewActionsDatastoresApi(apiClient)
	resp, r, err := api.CreateDatastore(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsDatastoresApi.CreateDatastore`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ActionsDatastoresApi.CreateDatastore`:\n%s\n", responseContent)
}
