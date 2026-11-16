// Snooze an issue returns "OK" response

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
	// there is a valid "issue" in the system
	IssueID := os.Getenv("ISSUE_ID")

	body := datadogV2.IssueSnoozeRequest{
		Data: datadogV2.IssueSnoozeRequestData{
			Attributes: datadogV2.IssueSnoozeRequestDataAttributes{
				Condition: datadogV2.IssueWatchCondition{
					Kind:     datadogV2.ISSUEWATCHCONDITIONKIND_ELAPSED_TIME,
					Duration: datadog.PtrString("168h"),
				},
			},
			Id:   IssueID,
			Type: datadogV2.ISSUESNOOZEREQUESTDATATYPE_ERROR_TRACKING_ISSUE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.SnoozeIssue", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewErrorTrackingApi(apiClient)
	resp, r, err := api.SnoozeIssue(ctx, IssueID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ErrorTrackingApi.SnoozeIssue`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ErrorTrackingApi.SnoozeIssue`:\n%s\n", responseContent)
}
