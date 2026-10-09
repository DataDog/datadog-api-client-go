// Bulk delete org group memberships returns "No Content" response

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
	body := datadogV2.OrgGroupMembershipBulkDeleteRequest{
		Data: []datadogV2.OrgGroupMembershipBulkDeleteRequestData{
			{
				Id:   uuid.MustParse("f1e2d3c4-b5a6-7890-1234-567890abcdef"),
				Type: datadogV2.ORGGROUPMEMBERSHIPTYPE_ORG_GROUP_MEMBERSHIPS,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.BulkDeleteOrgGroupMemberships", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgGroupsApi(apiClient)
	r, err := api.BulkDeleteOrgGroupMemberships(ctx, uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGroupsApi.BulkDeleteOrgGroupMemberships`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
