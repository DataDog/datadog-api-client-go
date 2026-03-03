// Update maximum session duration returns "No Content - The maximum session duration was successfully updated." response

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
		Data: datadogV2.MaxSessionDurationUpdateRequestData{
			Attributes: datadogV2.MaxSessionDurationUpdateAttributes{
				MaxSessionDuration: 60,
			},
			Type: datadogV2.MAXSESSIONDURATIONUPDATEREQUESTDATATYPE_MAX_SESSION_DURATION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateLoginOrgConfigMaxSessionDuration", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrganizationsApi(apiClient)
	r, err := api.UpdateLoginOrgConfigMaxSessionDuration(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateLoginOrgConfigMaxSessionDuration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
