// Get AWS On Demand task by id returns "OK." response

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
	api := datadogV2.NewAgentlessScanningApi(apiClient)
	resp, r, err := api.GetAwsOnDemandTask(ctx, "63d6b4f5-e5d0-4d90-824a-9580f05f026a", )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentlessScanningApi.GetAwsOnDemandTask`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AgentlessScanningApi.GetAwsOnDemandTask`:\n%s\n", responseContent)
}