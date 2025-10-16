// Create a WAF exclusion filter returns "OK" response

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
	body := datadogV2.ApplicationSecurityWafExclusionFilterCreateRequest{
Data: datadogV2.ApplicationSecurityWafExclusionFilterCreateData{
Attributes: datadogV2.ApplicationSecurityWafExclusionFilterCreateAttributes{
Description: "Exclude false positives on a path",
Enabled: true,
Parameters: []string{
"list.search.query",
},
PathGlob: datadog.PtrString("/accounts/*"),
RulesTarget: []datadogV2.ApplicationSecurityWafExclusionFilterRulesTarget{
{
Tags: &datadogV2.ApplicationSecurityWafExclusionFilterRulesTargetTags{
Category: datadog.PtrString("attack_attempt"),
Type: datadog.PtrString("lfi"),
},
},
},
Scope: []datadogV2.ApplicationSecurityWafExclusionFilterScope{
{
Env: datadog.PtrString("www"),
Service: datadog.PtrString("prod"),
},
},
},
Type: datadogV2.APPLICATIONSECURITYWAFEXCLUSIONFILTERTYPE_EXCLUSION_FILTER,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewApplicationSecurityApi(apiClient)
	resp, r, err := api.CreateApplicationSecurityWafExclusionFilter(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSecurityApi.CreateApplicationSecurityWafExclusionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSecurityApi.CreateApplicationSecurityWafExclusionFilter`:\n%s\n", responseContent)
}