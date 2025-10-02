// Delete case comment returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "case" in the system
	CaseID := os.Getenv("CASE_ID")

	// there is a valid "comment" in the system
	CommentID := os.Getenv("COMMENT_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	r, err := api.DeleteCaseComment(ctx, CaseID, CommentID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.DeleteCaseComment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
