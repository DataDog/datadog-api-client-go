// Delete an incident type returns "OK" response

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
	// there is a valid "incident_type" in the system
	IncidentTypeDataID := os.Getenv("INCIDENT_TYPE_DATA_ID")


	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteIncidentType", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	r, err := api.DeleteIncidentType(ctx, IncidentTypeDataID, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.DeleteIncidentType`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}