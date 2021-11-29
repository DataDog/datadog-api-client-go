// Grant permission to a role returns "OK" response

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
			Id:   datadog.PtrString("09166391-4377-3bf6-3ecc-817c2ef1b68e"),
			Type: datadog.PERMISSIONSTYPE_PERMISSIONS.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesApi.AddPermissionToRole(ctx, ROLE_DATA_ID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.AddPermissionToRole`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RolesApi.AddPermissionToRole`:\n%s\n", responseContent)
}
