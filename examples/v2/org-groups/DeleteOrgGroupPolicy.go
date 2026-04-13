// Delete an org group policy returns "No Content" response

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
	configuration.SetUnstableOperationEnabled("v2.DeleteOrgGroupPolicy", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgGroupsApi(apiClient)
	r, err := api.DeleteOrgGroupPolicy(ctx, uuid.MustParse("1a2b3c4d-5e6f-7890-abcd-ef0123456789"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGroupsApi.DeleteOrgGroupPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
