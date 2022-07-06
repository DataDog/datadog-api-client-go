// Create role returns "OK" response

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
	body := datadog.RoleCreateRequest{
		Data: datadog.RoleCreateData{
			Type: datadog.ROLESTYPE_ROLES.Ptr(),
			Attributes: datadog.RoleCreateAttributes{
				Name: "Example-Create_role_returns_OK_response",
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.RolesApi(apiClient)
	resp, r, err := api.CreateRole(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.CreateRole`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RolesApi.CreateRole`:\n%s\n", responseContent)
}
