// Update the maximum session duration returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.MaxSessionDurationUpdateRequest{
		Data: datadogV2.MaxSessionDurationUpdateData{
			Attributes: datadogV2.MaxSessionDurationUpdateAttributes{
				MaxSessionDuration: 604800,
			},
			Type: datadogV2.MAXSESSIONDURATIONTYPE_MAX_SESSION_DURATION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrganizationsApi(apiClient)
	r, err := api.UpdateLoginOrgConfigsMaxSessionDuration(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateLoginOrgConfigsMaxSessionDuration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
