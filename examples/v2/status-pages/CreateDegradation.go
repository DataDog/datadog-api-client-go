// Create degradation returns "Created" response

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
	// there is a valid "status_page" in the system
	StatusPageDataAttributesComponents0ID := os.Getenv("STATUS_PAGE_DATA_ATTRIBUTES_COMPONENTS_0_ID")
	StatusPageDataID := os.Getenv("STATUS_PAGE_DATA_ID")

	body := datadogV2.CreateDegradationRequest{
		Data: &datadogV2.CreateDegradationRequestData{
			Attributes: &datadogV2.CreateDegradationRequestDataAttributes{
				ComponentsAffected: []datadogV2.CreateDegradationRequestDataAttributesComponentsAffectedItems{
					{
						Id:     StatusPageDataAttributesComponents0ID,
						Status: datadogV2.STATUSPAGESCOMPONENTDATAATTRIBUTESSTATUS_MAJOR_OUTAGE,
					},
				},
				Description: datadog.PtrString("5e2fd69be33e79aa"),
				Status:      datadogV2.CREATEDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_INVESTIGATING,
				Title:       "5e2fd69be33e79aa",
			},
			Type: datadogV2.PATCHDEGRADATIONREQUESTDATATYPE_DEGRADATIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatusPagesApi(apiClient)
	resp, r, err := api.CreateDegradation(ctx, StatusPageDataID, body, *datadogV2.NewCreateDegradationOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.CreateDegradation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StatusPagesApi.CreateDegradation`:\n%s\n", responseContent)
}
