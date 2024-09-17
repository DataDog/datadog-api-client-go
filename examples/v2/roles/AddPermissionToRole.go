// Grant permission to a role returns "OK" response

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
	// there is a valid "role" in the system
	RoleDataID := os.Getenv("ROLE_DATA_ID")

	// there is a valid "permission" in the system
	PermissionID := os.Getenv("PERMISSION_ID")

	body := datadogV2.RelationshipToPermission{
		Data: &datadogV2.RelationshipToPermissionData{
			Id:   datadog.PtrString(PermissionID),
			Type: datadogV2.PERMISSIONSTYPE_PERMISSIONS.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRolesApi(apiClient)
	resp, r, err := api.AddPermissionToRole(ctx, RoleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.AddPermissionToRole`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RolesApi.AddPermissionToRole`:\n%s\n", responseContent)
}
