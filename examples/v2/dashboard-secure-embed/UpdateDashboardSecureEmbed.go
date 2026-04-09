// Update a secure embed for a dashboard returns "OK" response

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
	body := datadogV2.SecureEmbedUpdateRequest{
		Data: datadogV2.SecureEmbedUpdateRequestData{
			Attributes: datadogV2.SecureEmbedUpdateRequestAttributes{
				GlobalTime: &datadogV2.SecureEmbedGlobalTime{
					LiveSpan: datadogV2.SECUREEMBEDGLOBALTIMELIVESPAN_PAST_ONE_HOUR.Ptr(),
				},
				GlobalTimeSelectable: datadog.PtrBool(true),
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
				Status: datadogV2.SECUREEMBEDSTATUS_ACTIVE.Ptr(),
				Title:  datadog.PtrString("Q1 Metrics Dashboard (Updated)"),
				ViewingPreferences: &datadogV2.SecureEmbedViewingPreferences{
					HighDensity: datadog.PtrBool(false),
					Theme:       datadogV2.SECUREEMBEDVIEWINGPREFERENCESTHEME_SYSTEM.Ptr(),
				},
			},
			Type: datadogV2.SECUREEMBEDUPDATEREQUESTTYPE_SECURE_EMBED_UPDATE_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateDashboardSecureEmbed", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDashboardSecureEmbedApi(apiClient)
	resp, r, err := api.UpdateDashboardSecureEmbed(ctx, "dashboard_id", "token", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardSecureEmbedApi.UpdateDashboardSecureEmbed`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardSecureEmbedApi.UpdateDashboardSecureEmbed`:\n%s\n", responseContent)
}
