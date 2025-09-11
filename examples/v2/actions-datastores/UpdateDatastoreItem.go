// Update datastore item returns "OK" response

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

	body := datadogV2.UpdateAppsDatastoreItemRequest{
		Data: &datadogV2.UpdateAppsDatastoreItemRequestData{
			Attributes: &datadogV2.UpdateAppsDatastoreItemRequestDataAttributes{
				ItemChanges: datadogV2.UpdateAppsDatastoreItemRequestDataAttributesItemChanges{},
				ItemKey:     "test-key",
			},
			Type: datadogV2.UPDATEAPPSDATASTOREITEMREQUESTDATATYPE_ITEMS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewActionsDatastoresApi(apiClient)
	resp, r, err := api.UpdateDatastoreItem(ctx, DatastoreDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsDatastoresApi.UpdateDatastoreItem`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ActionsDatastoresApi.UpdateDatastoreItem`:\n%s\n", responseContent)
}
