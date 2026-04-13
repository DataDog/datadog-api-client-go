// Update an org group returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.OrgGroupUpdateRequest{
		Data: datadogV2.OrgGroupUpdateData{
			Attributes: datadogV2.OrgGroupUpdateAttributes{
				Name: "Updated Org Group Name",
			},
			Id:   uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
			Type: datadogV2.ORGGROUPTYPE_ORG_GROUPS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateOrgGroup", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgGroupsApi(apiClient)
	resp, r, err := api.UpdateOrgGroup(ctx, uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGroupsApi.UpdateOrgGroup`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrgGroupsApi.UpdateOrgGroup`:\n%s\n", responseContent)
}
