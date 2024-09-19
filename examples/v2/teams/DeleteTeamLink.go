// Remove a team link returns "No Content" response

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
	// there is a valid "dd_team" in the system
	DdTeamDataID := os.Getenv("DD_TEAM_DATA_ID")


	// there is a valid "team_link" in the system
	TeamLinkDataID := os.Getenv("TEAM_LINK_DATA_ID")


	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	r, err := api.DeleteTeamLink(ctx, DdTeamDataID, TeamLinkDataID, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.DeleteTeamLink`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}