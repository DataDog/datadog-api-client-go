// Update the assignee of an issue returns "OK" response

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


	body := datadogV2.IssueUpdateAssigneeRequest{
Data: datadogV2.IssueUpdateAssigneeRequestData{
Id: "87cb11a0-278c-440a-99fe-701223c80296",
Type: datadogV2.ISSUEUPDATEASSIGNEEREQUESTDATATYPE_ASSIGNEE,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewErrorTrackingApi(apiClient)
	resp, r, err := api.UpdateIssueAssignee(ctx, IssueID, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ErrorTrackingApi.UpdateIssueAssignee`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ErrorTrackingApi.UpdateIssueAssignee`:\n%s\n", responseContent)
}