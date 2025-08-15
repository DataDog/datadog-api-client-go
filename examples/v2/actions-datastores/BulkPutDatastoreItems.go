// Bulk put datastore items returns "OK" response

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
	// there is a valid "datastore" in the system
	DatastoreDataID := os.Getenv("DATASTORE_DATA_ID")

	body := datadogV2.BulkPutAppsDatastoreItemsRequest{
		Data: &datadogV2.BulkPutAppsDatastoreItemsRequestData{
			Attributes: &datadogV2.BulkPutAppsDatastoreItemsRequestDataAttributes{
				Values: []interface{}{
					{
						"28173b88-1a0e-001e-28c0-7664b6410518": "key1",
						"value":                                "{'data': 'example data 1', 'key': 'value'}",
					},
					{
						"28173b88-1a0e-001e-28c0-7664b6410518": "key2",
						"value":                                "{'data': 'example data 2', 'key': 'value'}",
					},
				},
			},
			Type: datadogV2.BULKPUTAPPSDATASTOREITEMSREQUESTDATATYPE_ITEMS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewActionsDatastoresApi(apiClient)
	resp, r, err := api.BulkPutDatastoreItems(ctx, DatastoreDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsDatastoresApi.BulkPutDatastoreItems`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ActionsDatastoresApi.BulkPutDatastoreItems`:\n%s\n", responseContent)
}
