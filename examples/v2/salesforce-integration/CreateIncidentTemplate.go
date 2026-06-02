// Create a Salesforce incident template returns "CREATED" response

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
	body := datadogV2.SalesforceIncidentsTemplateCreateRequest{
		Data: datadogV2.SalesforceIncidentsTemplateCreateData{
			Attributes: datadogV2.SalesforceIncidentsTemplateCreateAttributes{
				Description:     "An incident was detected by Datadog monitors.",
				Name:            "production-outage",
				OwnerId:         "005000000000000",
				Priority:        datadogV2.SALESFORCEINCIDENTSTEMPLATEPRIORITY_HIGH,
				SalesforceOrgId: uuid.MustParse("596da4af-0563-4097-90ff-07230c3f9db3"),
				Subject:         "Datadog Incident: Production Outage",
			},
			Type: datadogV2.SALESFORCEINCIDENTSTEMPLATETYPE_SALESFORCE_INCIDENTS_INCIDENT_TEMPLATE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSalesforceIntegrationApi(apiClient)
	resp, r, err := api.CreateIncidentTemplate(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesforceIntegrationApi.CreateIncidentTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SalesforceIntegrationApi.CreateIncidentTemplate`:\n%s\n", responseContent)
}
