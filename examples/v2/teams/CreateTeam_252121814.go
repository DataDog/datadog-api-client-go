// Create a team with V2 fields returns "CREATED" response

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
	body := datadogV2.TeamCreateRequest{
		Data: datadogV2.TeamCreate{
			Attributes: datadogV2.TeamCreateAttributes{
				Handle: "test-handle-a0fc0297eb519635",
				Name:   "test-name-a0fc0297eb519635",
				Avatar: *datadog.NewNullableString(datadog.PtrString("🥑")),
				Banner: *datadog.NewNullableInt64(datadog.PtrInt64(7)),
				VisibleModules: []string{
					"m1",
					"m2",
				},
				HiddenModules: []string{
					"m3",
				},
			},
			Type: datadogV2.TEAMTYPE_TEAM,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	resp, r, err := api.CreateTeam(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.CreateTeam`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamsApi.CreateTeam`:\n%s\n", responseContent)
}
