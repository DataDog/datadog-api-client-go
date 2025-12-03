// Delete team connections returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.TeamConnectionDeleteRequest{
		Data: []datadogV2.TeamConnectionDeleteRequestDataItem{
			{
				Id:   "12345678-1234-5678-9abc-123456789012",
				Type: datadogV2.TEAMCONNECTIONTYPE_TEAM_CONNECTION,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteTeamConnections", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	r, err := api.DeleteTeamConnections(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.DeleteTeamConnections`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
