// Bulk delete datastore items returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "datastore" in the system
	DatastoreDataID := os.Getenv("DATASTORE_DATA_ID")


	body := datadogV2.BulkDeleteAppsDatastoreItemsRequest{
Data: &datadogV2.BulkDeleteAppsDatastoreItemsRequestData{
Attributes: &datadogV2.BulkDeleteAppsDatastoreItemsRequestDataAttributes{
ItemKeys: []string{
"test-key",
},
},
Type: datadogV2.BULKDELETEAPPSDATASTOREITEMSREQUESTDATATYPE_ITEMS,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewActionsDatastoresApi(apiClient)
	resp, r, err := api.BulkDeleteDatastoreItems(ctx, DatastoreDataID, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsDatastoresApi.BulkDeleteDatastoreItems`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ActionsDatastoresApi.BulkDeleteDatastoreItems`:\n%s\n", responseContent)
}