// Delete an org group policy override returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteOrgGroupPolicyOverride", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgGroupsApi(apiClient)
	r, err := api.DeleteOrgGroupPolicyOverride(ctx, uuid.MustParse("9f8e7d6c-5b4a-3210-fedc-ba0987654321"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGroupsApi.DeleteOrgGroupPolicyOverride`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
