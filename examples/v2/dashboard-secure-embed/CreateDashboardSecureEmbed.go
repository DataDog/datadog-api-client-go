// Create a secure embed for a dashboard returns "OK" response

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
	body := datadogV2.SecureEmbedCreateRequest{
		Data: datadogV2.SecureEmbedCreateRequestData{
			Attributes: datadogV2.SecureEmbedCreateRequestAttributes{
				GlobalTime: datadogV2.SecureEmbedGlobalTime{
					LiveSpan: datadogV2.SECUREEMBEDGLOBALTIMELIVESPAN_PAST_ONE_HOUR.Ptr(),
				},
				GlobalTimeSelectable: true,
				SelectableTemplateVars: []datadogV2.SecureEmbedSelectableTemplateVariable{
					{
						DefaultValues: []string{
							"1",
						},
						Name:   datadog.PtrString("org_id"),
						Prefix: datadog.PtrString("org_id"),
						VisibleTags: []string{
							"1",
						},
					},
				},
				Status: datadogV2.SECUREEMBEDSTATUS_ACTIVE,
				Title:  "Q1 Metrics Dashboard",
				ViewingPreferences: datadogV2.SecureEmbedViewingPreferences{
					HighDensity: datadog.PtrBool(false),
					Theme:       datadogV2.SECUREEMBEDVIEWINGPREFERENCESTHEME_SYSTEM.Ptr(),
				},
			},
			Type: datadogV2.SECUREEMBEDREQUESTTYPE_SECURE_EMBED_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateDashboardSecureEmbed", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDashboardSecureEmbedApi(apiClient)
	resp, r, err := api.CreateDashboardSecureEmbed(ctx, "dashboard_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardSecureEmbedApi.CreateDashboardSecureEmbed`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardSecureEmbedApi.CreateDashboardSecureEmbed`:\n%s\n", responseContent)
}
