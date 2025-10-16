// Update a WAF exclusion filter returns "OK" response

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
	// there is a valid "exclusion_filter" in the system
	ExclusionFilterDataID := os.Getenv("EXCLUSION_FILTER_DATA_ID")


	body := datadogV2.ApplicationSecurityWafExclusionFilterUpdateRequest{
Data: datadogV2.ApplicationSecurityWafExclusionFilterUpdateData{
Attributes: datadogV2.ApplicationSecurityWafExclusionFilterUpdateAttributes{
Description: "Exclude false positives on a path",
Enabled: false,
IpList: []string{
"198.51.100.72",
},
OnMatch: datadogV2.APPLICATIONSECURITYWAFEXCLUSIONFILTERONMATCH_MONITOR.Ptr(),
},
Type: datadogV2.APPLICATIONSECURITYWAFEXCLUSIONFILTERTYPE_EXCLUSION_FILTER,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewApplicationSecurityApi(apiClient)
	resp, r, err := api.UpdateApplicationSecurityWafExclusionFilter(ctx, ExclusionFilterDataID, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSecurityApi.UpdateApplicationSecurityWafExclusionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSecurityApi.UpdateApplicationSecurityWafExclusionFilter`:\n%s\n", responseContent)
}