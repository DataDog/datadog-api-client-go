// Update case comment returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.CaseUpdateCommentRequest{
		Data: datadogV2.CaseUpdateComment{
			Attributes: datadogV2.CaseUpdateCommentAttributes{
				Comment: "Updated comment text",
			},
			Type: datadogV2.CASERESOURCETYPE_CASE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateCaseComment", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	r, err := api.UpdateCaseComment(ctx, "case_id", "cell_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.UpdateCaseComment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
