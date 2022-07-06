// List roles returns "OK" response

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
	RoleDataAttributesName := os.Getenv("ROLE_DATA_ATTRIBUTES_NAME")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.RolesApi(apiClient)
	resp, r, err := api.ListRoles(ctx, *datadog.NewListRolesOptionalParameters().WithFilter(RoleDataAttributesName))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ListRoles`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RolesApi.ListRoles`:\n%s\n", responseContent)
}
