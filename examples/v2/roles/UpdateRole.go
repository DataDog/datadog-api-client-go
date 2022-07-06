// Update a role returns "OK" response

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

	body := datadog.RoleUpdateRequest{
		Data: datadog.RoleUpdateData{
			Id:   RoleDataID,
			Type: datadog.ROLESTYPE_ROLES,
			Attributes: datadog.RoleUpdateAttributes{
				Name: datadog.PtrString("developers-updated"),
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.RolesApi(apiClient)
	resp, r, err := api.UpdateRole(ctx, RoleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateRole`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RolesApi.UpdateRole`:\n%s\n", responseContent)
}
