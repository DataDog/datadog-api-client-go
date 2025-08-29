// Create datastore from import returns "OK" response

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
	body := datadogV2.CreateAppsDatastoreFromImportRequest{
		Data: &datadogV2.CreateAppsDatastoreFromImportRequestData{
			Attributes: &datadogV2.CreateAppsDatastoreFromImportRequestDataAttributes{
				Name:              "datastore-name",
				PrimaryColumnName: "primaryKey",
				Values: []interface{}{
					{
						"primaryKey": "key1",
						"value":      "Newton",
					},
					{
						"primaryKey": "key2",
						"value":      "Leibniz",
					},
				},
			},
			Type: datadogV2.DATASTOREDATATYPE_DATASTORES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewActionsDatastoresApi(apiClient)
	resp, r, err := api.CreateDatastoreFromImport(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsDatastoresApi.CreateDatastoreFromImport`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ActionsDatastoresApi.CreateDatastoreFromImport`:\n%s\n", responseContent)
}
