// Update datastore returns "OK" response

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


	body := datadogV2.UpdateAppsDatastoreRequest{
Data: &datadogV2.UpdateAppsDatastoreRequestData{
Attributes: &datadogV2.UpdateAppsDatastoreRequestDataAttributes{
Name: datadog.PtrString("updated name"),
},
Type: datadogV2.DATASTOREDATATYPE_DATASTORES,
Id: datadog.PtrString(DatastoreDataID),
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewActionsDatastoresApi(apiClient)
	resp, r, err := api.UpdateDatastore(ctx, DatastoreDataID, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsDatastoresApi.UpdateDatastore`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ActionsDatastoresApi.UpdateDatastore`:\n%s\n", responseContent)
}