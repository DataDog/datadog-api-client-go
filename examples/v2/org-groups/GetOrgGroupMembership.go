// Get an org group membership returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetOrgGroupMembership", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgGroupsApi(apiClient)
	resp, r, err := api.GetOrgGroupMembership(ctx, uuid.MustParse("f1e2d3c4-b5a6-7890-1234-567890abcdef"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGroupsApi.GetOrgGroupMembership`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrgGroupsApi.GetOrgGroupMembership`:\n%s\n", responseContent)
}
