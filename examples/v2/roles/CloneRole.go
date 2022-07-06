// Create a new role by cloning an existing role returns "OK" response

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

	body := datadog.RoleCloneRequest{
		Data: datadog.RoleClone{
			Attributes: datadog.RoleCloneAttributes{
				Name: "Example-Create_a_new_role_by_cloning_an_existing_role_returns_OK_response clone",
			},
			Type: datadog.ROLESTYPE_ROLES,
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.RolesApi(apiClient)
	resp, r, err := api.CloneRole(ctx, RoleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.CloneRole`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RolesApi.CloneRole`:\n%s\n", responseContent)
}
