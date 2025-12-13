// Delete incident attachment returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteIncidentAttachment", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	r, err := api.DeleteIncidentAttachment(ctx, "incident_id", "00000000-0000-0000-0000-000000000002")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.DeleteIncidentAttachment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
