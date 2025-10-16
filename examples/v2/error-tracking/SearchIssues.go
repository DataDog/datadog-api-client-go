// Search error tracking issues returns "OK" response

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
	body := datadogV2.IssuesSearchRequest{
Data: datadogV2.IssuesSearchRequestData{
Attributes: datadogV2.IssuesSearchRequestDataAttributes{
Query: "service:orders-* AND @language:go",
From: 1671612804000,
To: 1671620004000,
Track: datadogV2.ISSUESSEARCHREQUESTDATAATTRIBUTESTRACK_TRACE.Ptr(),
},
Type: datadogV2.ISSUESSEARCHREQUESTDATATYPE_SEARCH_REQUEST,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewErrorTrackingApi(apiClient)
	resp, r, err := api.SearchIssues(ctx, body, *datadogV2.NewSearchIssuesOptionalParameters(), )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ErrorTrackingApi.SearchIssues`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ErrorTrackingApi.SearchIssues`:\n%s\n", responseContent)
}