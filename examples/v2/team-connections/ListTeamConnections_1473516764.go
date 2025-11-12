// List team connections with filters returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ListTeamConnections", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamConnectionsApi(apiClient)
	resp, r, err := api.ListTeamConnections(ctx, *datadogV2.NewListTeamConnectionsOptionalParameters().WithFilterSources([]string{
		"github",
	}).WithPageSize(10))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamConnectionsApi.ListTeamConnections`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamConnectionsApi.ListTeamConnections`:\n%s\n", responseContent)
}
