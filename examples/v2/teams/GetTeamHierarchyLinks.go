// Get team hierarchy links returns "OK" response

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
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	resp, r, err := api.GetTeamHierarchyLinks(ctx, *datadogV2.NewGetTeamHierarchyLinksOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.GetTeamHierarchyLinks`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamsApi.GetTeamHierarchyLinks`:\n%s\n", responseContent)
}
