// Create or update IDP configuration returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.IDPConfigRequest{
		Data: datadogV2.IDPConfigRequestData{
			Attributes: datadogV2.IDPConfigRequestAttributes{
				Value: []datadogV2.IDPConfigValueItem{
					{
						"displayName": "My Dashboard",
						"id":          "dashboard-1",
					},
				},
			},
			Type: datadogV2.IDPCONFIGTYPE_IDP_CONFIG,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpsertIDPConfigValue", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSoftwareCatalogApi(apiClient)
	r, err := api.UpsertIDPConfigValue(ctx, "idp_pinned_dashboards", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SoftwareCatalogApi.UpsertIDPConfigValue`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
