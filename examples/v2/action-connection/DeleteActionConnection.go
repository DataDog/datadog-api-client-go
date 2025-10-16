// Delete an existing Action Connection returns "The resource was deleted successfully." response

package main


import (
	"context"
	"encoding/json"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "action_connection" in the system
	ActionConnectionDataID := os.Getenv("ACTION_CONNECTION_DATA_ID")


	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewActionConnectionApi(apiClient)
	r, err := api.DeleteActionConnection(ctx, ActionConnectionDataID, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionConnectionApi.DeleteActionConnection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}