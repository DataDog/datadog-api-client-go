// Create backfilled degradation returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "status_page" in the system
	StatusPageDataAttributesComponents0Components0ID := uuid.MustParse(os.Getenv("STATUS_PAGE_DATA_ATTRIBUTES_COMPONENTS_0_COMPONENTS_0_ID"))
	StatusPageDataID := uuid.MustParse(os.Getenv("STATUS_PAGE_DATA_ID"))

	body := datadogV2.CreateBackfilledDegradationRequest{
		Data: &datadogV2.CreateBackfilledDegradationRequestData{
			Attributes: &datadogV2.CreateBackfilledDegradationRequestDataAttributes{
				Title: "Past API Outage",
				Updates: []datadogV2.CreateBackfilledDegradationRequestDataAttributesUpdatesItems{
					{
						ComponentsAffected: []datadogV2.CreateDegradationRequestDataAttributesComponentsAffectedItems{
							{
								Id:     StatusPageDataAttributesComponents0Components0ID,
								Status: datadogV2.STATUSPAGESCOMPONENTDATAATTRIBUTESSTATUS_DEGRADED,
							},
						},
						Description: datadog.PtrString("We detected elevated error rates in the API."),
						StartedAt:   time.Now().Add(time.Hour * -1),
						Status:      datadogV2.CREATEDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_INVESTIGATING,
					},
					{
						ComponentsAffected: []datadogV2.CreateDegradationRequestDataAttributesComponentsAffectedItems{
							{
								Id:     StatusPageDataAttributesComponents0Components0ID,
								Status: datadogV2.STATUSPAGESCOMPONENTDATAATTRIBUTESSTATUS_DEGRADED,
							},
						},
						Description: datadog.PtrString("Root cause identified as a misconfigured deployment."),
						StartedAt:   time.Now().Add(time.Minute * -30),
						Status:      datadogV2.CREATEDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_IDENTIFIED,
					},
					{
						ComponentsAffected: []datadogV2.CreateDegradationRequestDataAttributesComponentsAffectedItems{
							{
								Id:     StatusPageDataAttributesComponents0Components0ID,
								Status: datadogV2.STATUSPAGESCOMPONENTDATAATTRIBUTESSTATUS_OPERATIONAL,
							},
						},
						Description: datadog.PtrString("The issue has been resolved and API is operating normally."),
						StartedAt:   time.Now(),
						Status:      datadogV2.CREATEDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_RESOLVED,
					},
				},
			},
			Type: datadogV2.PATCHDEGRADATIONREQUESTDATATYPE_DEGRADATIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatusPagesApi(apiClient)
	resp, r, err := api.CreateBackfilledDegradation(ctx, StatusPageDataID, body, *datadogV2.NewCreateBackfilledDegradationOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.CreateBackfilledDegradation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StatusPagesApi.CreateBackfilledDegradation`:\n%s\n", responseContent)
}
