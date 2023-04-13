// Create a new role by cloning an existing role returns "OK" response

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

	body := datadogV2.RoleCloneRequest{
		Data: datadogV2.RoleClone{
			Attributes: datadogV2.RoleCloneAttributes{
				Name: "Example-Role clone",
			},
			Type: datadogV2.ROLESTYPE_ROLES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRolesApi(apiClient)
	resp, r, err := api.CloneRole(ctx, RoleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.CloneRole`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RolesApi.CloneRole`:\n%s\n", responseContent)
}
