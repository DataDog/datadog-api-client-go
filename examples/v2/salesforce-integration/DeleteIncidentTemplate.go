// Delete a Salesforce incident template returns "OK" response

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
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSalesforceIntegrationApi(apiClient)
	r, err := api.DeleteIncidentTemplate(ctx, "incident_template_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesforceIntegrationApi.DeleteIncidentTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
