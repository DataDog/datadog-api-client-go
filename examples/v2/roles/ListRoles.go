// List roles returns "OK" response

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
	RoleDataAttributesName := os.Getenv("ROLE_DATA_ATTRIBUTES_NAME")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesApi.ListRoles(ctx, *datadog.NewListRolesOptionalParameters().WithFilter(RoleDataAttributesName))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ListRoles`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RolesApi.ListRoles`:\n%s\n", responseContent)
}
