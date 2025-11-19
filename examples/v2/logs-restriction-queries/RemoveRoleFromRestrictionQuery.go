// Revoke role from a restriction query returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "restriction_query" in the system
	RestrictionQueryDataID := os.Getenv("RESTRICTION_QUERY_DATA_ID")

	// there is a valid "role" in the system
	RoleDataID := os.Getenv("ROLE_DATA_ID")

	body := datadogV2.RelationshipToRole{
		Data: &datadogV2.RelationshipToRoleData{
			Id:   datadog.PtrString(RoleDataID),
			Type: datadogV2.ROLESTYPE_ROLES.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.RemoveRoleFromRestrictionQuery", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsRestrictionQueriesApi(apiClient)
	r, err := api.RemoveRoleFromRestrictionQuery(ctx, RestrictionQueryDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsRestrictionQueriesApi.RemoveRoleFromRestrictionQuery`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
