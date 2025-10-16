// Delete GCP Scan Options returns "No Content" response

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
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAgentlessScanningApi(apiClient)
	r, err := api.DeleteGcpScanOptions(ctx, "company-project-id", )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentlessScanningApi.DeleteGcpScanOptions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}