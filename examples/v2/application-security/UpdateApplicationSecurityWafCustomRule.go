// Update a WAF Custom Rule returns "OK" response

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
	// there is a valid "custom_rule" in the system
	CustomRuleDataID := os.Getenv("CUSTOM_RULE_DATA_ID")


	body := datadogV2.ApplicationSecurityWafCustomRuleUpdateRequest{
Data: datadogV2.ApplicationSecurityWafCustomRuleUpdateData{
Type: datadogV2.APPLICATIONSECURITYWAFCUSTOMRULETYPE_CUSTOM_RULE,
Attributes: datadogV2.ApplicationSecurityWafCustomRuleUpdateAttributes{
Blocking: false,
Conditions: []datadogV2.ApplicationSecurityWafCustomRuleCondition{
{
Operator: datadogV2.APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_MATCH_REGEX,
Parameters: datadogV2.ApplicationSecurityWafCustomRuleConditionParameters{
Inputs: []datadogV2.ApplicationSecurityWafCustomRuleConditionInput{
{
Address: datadogV2.APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_REQUEST_QUERY,
KeyPath: []string{
"id",
},
},
},
Regex: datadog.PtrString("badactor"),
},
},
},
Enabled: false,
Name: "test",
PathGlob: datadog.PtrString("/test"),
Scope: []datadogV2.ApplicationSecurityWafCustomRuleScope{
{
Env: "test",
Service: "test",
},
},
Tags: datadogV2.ApplicationSecurityWafCustomRuleTags{
Category: datadogV2.APPLICATIONSECURITYWAFCUSTOMRULETAGSCATEGORY_ATTACK_ATTEMPT,
Type: "test",
},
},
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewApplicationSecurityApi(apiClient)
	resp, r, err := api.UpdateApplicationSecurityWafCustomRule(ctx, CustomRuleDataID, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSecurityApi.UpdateApplicationSecurityWafCustomRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSecurityApi.UpdateApplicationSecurityWafCustomRule`:\n%s\n", responseContent)
}