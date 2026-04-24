// Bulk update org group memberships returns "OK" response

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
	body := datadogV2.OrgGroupMembershipBulkUpdateRequest{
		Data: datadogV2.OrgGroupMembershipBulkUpdateData{
			Attributes: datadogV2.OrgGroupMembershipBulkUpdateAttributes{
				Orgs: []datadogV2.GlobalOrgIdentifier{
					{
						OrgSite: "us1",
						OrgUuid: uuid.MustParse("c3d4e5f6-a7b8-9012-cdef-012345678901"),
					},
				},
			},
			Relationships: datadogV2.OrgGroupMembershipBulkUpdateRelationships{
				SourceOrgGroup: datadogV2.OrgGroupRelationshipToOne{
					Data: datadogV2.OrgGroupRelationshipToOneData{
						Id:   uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
						Type: datadogV2.ORGGROUPTYPE_ORG_GROUPS,
					},
				},
				TargetOrgGroup: datadogV2.OrgGroupRelationshipToOne{
					Data: datadogV2.OrgGroupRelationshipToOneData{
						Id:   uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
						Type: datadogV2.ORGGROUPTYPE_ORG_GROUPS,
					},
				},
			},
			Type: datadogV2.ORGGROUPMEMBERSHIPBULKUPDATETYPE_ORG_GROUP_MEMBERSHIP_BULK_UPDATES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.BulkUpdateOrgGroupMemberships", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgGroupsApi(apiClient)
	resp, r, err := api.BulkUpdateOrgGroupMemberships(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGroupsApi.BulkUpdateOrgGroupMemberships`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrgGroupsApi.BulkUpdateOrgGroupMemberships`:\n%s\n", responseContent)
}
