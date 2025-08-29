// Write datastore item returns "OK" response

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

	body := datadogV2.PutAppsDatastoreItemRequest{
		Data: &datadogV2.PutAppsDatastoreItemRequestData{
			Attributes: &datadogV2.PutAppsDatastoreItemRequestDataAttributes{
				Value: map[string]interface{}{
					"id":   "new-item-key",
					"data": "example data",
					"key":  "value",
				},
			},
			Type: datadogV2.DATASTOREITEMSDATATYPE_ITEMS,
			Id:   datadog.PtrString("e7e64418-b60c-4789-9612-895ac8423207"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewActionsDatastoresApi(apiClient)
	resp, r, err := api.WriteDatastoreItem(ctx, DatastoreDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsDatastoresApi.WriteDatastoreItem`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ActionsDatastoresApi.WriteDatastoreItem`:\n%s\n", responseContent)
}
