// Create or update HAMR organization connection returns "OK" response

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
	body := datadogV2.HamrOrgConnectionRequest{
		Data: datadogV2.HamrOrgConnectionDataRequest{
			Attributes: datadogV2.HamrOrgConnectionAttributesRequest{
				HamrStatus:          datadogV2.HAMRORGCONNECTIONSTATUS_ACTIVE,
				IsPrimary:           true,
				ModifiedBy:          "admin@example.com",
				TargetOrgDatacenter: "us1",
				TargetOrgName:       "Production Backup Org",
				TargetOrgUuid:       "660f9511-f3ac-52e5-b827-557766551111",
			},
			Id:   "550e8400-e29b-41d4-a716-446655440000",
			Type: datadogV2.HAMRORGCONNECTIONTYPE_HAMR_ORG_CONNECTIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateHamrOrgConnection", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewHighAvailabilityMultiRegionApi(apiClient)
	resp, r, err := api.CreateHamrOrgConnection(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityMultiRegionApi.CreateHamrOrgConnection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `HighAvailabilityMultiRegionApi.CreateHamrOrgConnection`:\n%s\n", responseContent)
}
