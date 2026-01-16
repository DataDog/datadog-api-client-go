// Revert Custom Rule Revision returns "Successfully reverted" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.RevertCustomRuleRevisionRequest{
		Data: &datadogV2.RevertCustomRuleRevisionRequestData{
			Attributes: &datadogV2.RevertCustomRuleRevisionRequestDataAttributes{},
			Type:       datadogV2.REVERTCUSTOMRULEREVISIONDATATYPE_REVERT_CUSTOM_RULE_REVISION_REQUEST.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.RevertCustomRuleRevision", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStaticAnalysisApi(apiClient)
	r, err := api.RevertCustomRuleRevision(ctx, "ruleset_name", "rule_name", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticAnalysisApi.RevertCustomRuleRevision`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
