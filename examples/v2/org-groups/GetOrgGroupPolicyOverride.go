// Get an org group policy override returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetOrgGroupPolicyOverride", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgGroupsApi(apiClient)
	resp, r, err := api.GetOrgGroupPolicyOverride(ctx, uuid.MustParse("9f8e7d6c-5b4a-3210-fedc-ba0987654321"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGroupsApi.GetOrgGroupPolicyOverride`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrgGroupsApi.GetOrgGroupPolicyOverride`:\n%s\n", responseContent)
}
