// Get all CSM Agents returns "OK" response

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
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCSMAgentsApi(apiClient)
	resp, r, err := api.ListAllCSMAgents(ctx, *datadogV2.NewListAllCSMAgentsOptionalParameters(), )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CSMAgentsApi.ListAllCSMAgents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CSMAgentsApi.ListAllCSMAgents`:\n%s\n", responseContent)
}