// Update an org group membership returns "OK" response

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
	body := datadogV2.OrgGroupMembershipUpdateRequest{
		Data: datadogV2.OrgGroupMembershipUpdateData{
			Id: uuid.MustParse("f1e2d3c4-b5a6-7890-1234-567890abcdef"),
			Relationships: datadogV2.OrgGroupMembershipUpdateRelationships{
				OrgGroup: datadogV2.OrgGroupRelationshipToOne{
					Data: datadogV2.OrgGroupRelationshipToOneData{
						Id:   uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
						Type: datadogV2.ORGGROUPTYPE_ORG_GROUPS,
					},
				},
			},
			Type: datadogV2.ORGGROUPMEMBERSHIPTYPE_ORG_GROUP_MEMBERSHIPS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateOrgGroupMembership", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgGroupsApi(apiClient)
	resp, r, err := api.UpdateOrgGroupMembership(ctx, uuid.MustParse("f1e2d3c4-b5a6-7890-1234-567890abcdef"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGroupsApi.UpdateOrgGroupMembership`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrgGroupsApi.UpdateOrgGroupMembership`:\n%s\n", responseContent)
}
