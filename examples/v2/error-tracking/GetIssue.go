// Get the details of an error tracking issue returns "OK" response

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
	// there is a valid "issue" in the system
	IssueID := os.Getenv("ISSUE_ID")


	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewErrorTrackingApi(apiClient)
	resp, r, err := api.GetIssue(ctx, IssueID, *datadogV2.NewGetIssueOptionalParameters(), )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ErrorTrackingApi.GetIssue`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ErrorTrackingApi.GetIssue`:\n%s\n", responseContent)
}