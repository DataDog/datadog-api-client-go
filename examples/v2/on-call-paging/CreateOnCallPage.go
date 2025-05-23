// Create On-Call Page returns "OK." response

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
	body := datadogV2.CreatePageRequest{
		Data: &datadogV2.CreatePageRequestData{
			Attributes: &datadogV2.CreatePageRequestDataAttributes{
				Description: datadog.PtrString("Page details."),
				Tags: []string{
					"service:test",
				},
				Target: datadogV2.CreatePageRequestDataAttributesTarget{
					Identifier: datadog.PtrString("my-team"),
					Type:       datadogV2.ONCALLPAGETARGETTYPE_TEAM_HANDLE.Ptr(),
				},
				Title:   "Page title",
				Urgency: datadogV2.PAGEURGENCY_LOW,
			},
			Type: datadogV2.CREATEPAGEREQUESTDATATYPE_PAGES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnCallPagingApi(apiClient)
	resp, r, err := api.CreateOnCallPage(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallPagingApi.CreateOnCallPage`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OnCallPagingApi.CreateOnCallPage`:\n%s\n", responseContent)
}
