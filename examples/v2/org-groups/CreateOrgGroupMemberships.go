// Create org group memberships returns "Created" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.OrgGroupMembershipCreateRequest{
		Data: datadogV2.OrgGroupMembershipCreateData{
			Attributes: datadogV2.OrgGroupMembershipCreateAttributes{
				Orgs: []datadogV2.GlobalOrgIdentifier{
					{
						OrgSite: "us1",
						OrgUuid: uuid.MustParse("c3d4e5f6-a7b8-9012-cdef-012345678901"),
					},
				},
			},
			Relationships: datadogV2.OrgGroupMembershipCreateRelationships{
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
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateOrgGroupMemberships", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgGroupsApi(apiClient)
	r, err := api.CreateOrgGroupMemberships(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGroupsApi.CreateOrgGroupMemberships`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
