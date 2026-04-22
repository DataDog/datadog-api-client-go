// Get an org group policy returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetOrgGroupPolicy", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgGroupsApi(apiClient)
	resp, r, err := api.GetOrgGroupPolicy(ctx, uuid.MustParse("1a2b3c4d-5e6f-7890-abcd-ef0123456789"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGroupsApi.GetOrgGroupPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrgGroupsApi.GetOrgGroupPolicy`:\n%s\n", responseContent)
}
