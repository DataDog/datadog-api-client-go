// Delete datastore item returns "OK" response

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


	body := datadogV2.DeleteAppsDatastoreItemRequest{
Data: &datadogV2.DeleteAppsDatastoreItemRequestData{
Attributes: &datadogV2.DeleteAppsDatastoreItemRequestDataAttributes{
ItemKey: "test-key",
},
Type: datadogV2.DATASTOREITEMSDATATYPE_ITEMS,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewActionsDatastoresApi(apiClient)
	resp, r, err := api.DeleteDatastoreItem(ctx, DatastoreDataID, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsDatastoresApi.DeleteDatastoreItem`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ActionsDatastoresApi.DeleteDatastoreItem`:\n%s\n", responseContent)
}