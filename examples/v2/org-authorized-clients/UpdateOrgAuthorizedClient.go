// Update an org authorized client returns "OK" response

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
	body := datadogV2.OrgAuthorizedClientUpdateRequest{
		Data: datadogV2.OrgAuthorizedClientUpdateData{
			Attributes: &datadogV2.OrgAuthorizedClientUpdateAttributes{
				Disabled: datadog.PtrBool(true),
			},
			Id:   "00000000-0000-0000-0000-000000000001",
			Type: datadogV2.ORGAUTHORIZEDCLIENTTYPE_ORG_AUTHORIZED_CLIENTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgAuthorizedClientsApi(apiClient)
	resp, r, err := api.UpdateOrgAuthorizedClient(ctx, "00000000-0000-0000-0000-000000000001", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgAuthorizedClientsApi.UpdateOrgAuthorizedClient`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrgAuthorizedClientsApi.UpdateOrgAuthorizedClient`:\n%s\n", responseContent)
}
