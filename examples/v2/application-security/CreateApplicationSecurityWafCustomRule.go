// Create a WAF custom rule returns "Created" response

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
	body := datadogV2.ApplicationSecurityWafCustomRuleCreateRequest{
Data: datadogV2.ApplicationSecurityWafCustomRuleCreateData{
Attributes: datadogV2.ApplicationSecurityWafCustomRuleCreateAttributes{
Action: &datadogV2.ApplicationSecurityWafCustomRuleAction{
Action: datadogV2.APPLICATIONSECURITYWAFCUSTOMRULEACTIONACTION_BLOCK_REQUEST.Ptr(),
Parameters: &datadogV2.ApplicationSecurityWafCustomRuleActionParameters{
Location: datadog.PtrString("/blocking"),
StatusCode: datadog.PtrInt64(403),
},
},
Blocking: false,
Conditions: []datadogV2.ApplicationSecurityWafCustomRuleCondition{
{
Operator: datadogV2.APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_MATCH_REGEX,
Parameters: datadogV2.ApplicationSecurityWafCustomRuleConditionParameters{
Data: datadog.PtrString("blocked_users"),
Inputs: []datadogV2.ApplicationSecurityWafCustomRuleConditionInput{
{
Address: datadogV2.APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_DB_STATEMENT,
KeyPath: []string{
},
},
},
List: []string{
},
Options: &datadogV2.ApplicationSecurityWafCustomRuleConditionOptions{
CaseSensitive: datadog.PtrBool(false),
MinLength: datadog.PtrInt64(0),
},
Regex: datadog.PtrString("path.*"),
Value: datadog.PtrString("custom_tag"),
},
},
},
Enabled: false,
Name: "Block request from a bad useragent",
PathGlob: datadog.PtrString("/api/search/*"),
Scope: []datadogV2.ApplicationSecurityWafCustomRuleScope{
{
Env: "prod",
Service: "billing-service",
},
},
Tags: datadogV2.ApplicationSecurityWafCustomRuleTags{
Category: datadogV2.APPLICATIONSECURITYWAFCUSTOMRULETAGSCATEGORY_BUSINESS_LOGIC,
Type: "users.login.success",
},
},
Type: datadogV2.APPLICATIONSECURITYWAFCUSTOMRULETYPE_CUSTOM_RULE,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewApplicationSecurityApi(apiClient)
	resp, r, err := api.CreateApplicationSecurityWafCustomRule(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSecurityApi.CreateApplicationSecurityWafCustomRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSecurityApi.CreateApplicationSecurityWafCustomRule`:\n%s\n", responseContent)
}