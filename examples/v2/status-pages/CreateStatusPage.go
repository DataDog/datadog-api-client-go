// Create status page returns "Created" response

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
	body := datadogV2.CreateStatusPageRequest{
		Data: &datadogV2.CreateStatusPageRequestData{
			Attributes: &datadogV2.CreateStatusPageRequestDataAttributes{
				Name:         "A Status Page",
				DomainPrefix: "status-page-5e2fd69be33e79aa",
				Components: []datadogV2.CreateStatusPageRequestDataAttributesComponentsItems{
					{
						Name:     datadog.PtrString("Login"),
						Type:     datadogV2.CREATECOMPONENTREQUESTDATAATTRIBUTESTYPE_COMPONENT.Ptr(),
						Position: datadog.PtrInt64(0),
					},
					{
						Name:     datadog.PtrString("Settings"),
						Type:     datadogV2.CREATECOMPONENTREQUESTDATAATTRIBUTESTYPE_COMPONENT.Ptr(),
						Position: datadog.PtrInt64(1),
					},
				},
				Enabled:           true,
				Type:              datadogV2.CREATESTATUSPAGEREQUESTDATAATTRIBUTESTYPE_INTERNAL,
				VisualizationType: datadogV2.CREATESTATUSPAGEREQUESTDATAATTRIBUTESVISUALIZATIONTYPE_BARS_AND_UPTIME_PERCENTAGE,
			},
			Type: datadogV2.STATUSPAGEDATATYPE_STATUS_PAGES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatusPagesApi(apiClient)
	resp, r, err := api.CreateStatusPage(ctx, body, *datadogV2.NewCreateStatusPageOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.CreateStatusPage`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StatusPagesApi.CreateStatusPage`:\n%s\n", responseContent)
}
