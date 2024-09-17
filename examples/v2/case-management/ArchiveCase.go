// Archive case returns "OK" response

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
	// there is a valid "case" in the system
	CaseID := os.Getenv("CASE_ID")


	body := datadogV2.CaseEmptyRequest{
Data: datadogV2.CaseEmpty{
Type: datadogV2.CASERESOURCETYPE_CASE,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	resp, r, err := api.ArchiveCase(ctx, CaseID, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.ArchiveCase`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementApi.ArchiveCase`:\n%s\n", responseContent)
}