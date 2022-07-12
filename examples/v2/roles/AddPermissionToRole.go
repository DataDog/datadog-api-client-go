// Grant permission to a role returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "role" in the system
	RoleDataID := os.Getenv("ROLE_DATA_ID")

	// there is a valid "permission" in the system
	PermissionID := os.Getenv("PERMISSION_ID")

	body := datadog.RelationshipToPermission{
		Data: &datadog.RelationshipToPermissionData{
			Id:   datadog.PtrString(PermissionID),
			Type: datadog.PERMISSIONSTYPE_PERMISSIONS.Ptr(),
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewRolesApi(apiClient)
	resp, r, err := api.AddPermissionToRole(ctx, RoleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.AddPermissionToRole`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RolesApi.AddPermissionToRole`:\n%s\n", responseContent)
}
