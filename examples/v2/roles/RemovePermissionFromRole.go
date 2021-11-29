// Revoke permission returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "role" in the system
	ROLE_DATA_ID := os.Getenv("ROLE_DATA_ID")

	// there is a valid "permission" in the system

	body := datadog.RelationshipToPermission{
		Data: &datadog.RelationshipToPermissionData{
			Id:   datadog.PtrString("241f89e1-cee0-d53b-8f91-daf1f1c64c82"),
			Type: datadog.PERMISSIONSTYPE_PERMISSIONS.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesApi.RemovePermissionFromRole(ctx, ROLE_DATA_ID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.RemovePermissionFromRole`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RolesApi.RemovePermissionFromRole`:\n%s\n", responseContent)
}
