// Delete a tag configuration returns "No Content" response

package main


import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("DeleteTagConfiguration", true)
	apiClient := datadog.NewAPIClient(configuration)
	r, err := apiClient.MetricsApi.DeleteTagConfiguration(ctx, "ExampleDeleteatagconfigurationreturnsNoContentresponse", )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.DeleteTagConfiguration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}