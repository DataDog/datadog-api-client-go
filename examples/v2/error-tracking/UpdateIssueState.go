// Update the state of an issue returns "OK" response

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


	body := datadogV2.IssueUpdateStateRequest{
Data: datadogV2.IssueUpdateStateRequestData{
Attributes: datadogV2.IssueUpdateStateRequestDataAttributes{
State: datadogV2.ISSUESTATE_RESOLVED,
},
Id: IssueID,
Type: datadogV2.ISSUEUPDATESTATEREQUESTDATATYPE_ERROR_TRACKING_ISSUE,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewErrorTrackingApi(apiClient)
	resp, r, err := api.UpdateIssueState(ctx, IssueID, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ErrorTrackingApi.UpdateIssueState`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ErrorTrackingApi.UpdateIssueState`:\n%s\n", responseContent)
}