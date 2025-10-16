// Reorder tag pipeline rulesets returns "Successfully reordered rulesets" response

package main


import (
	"context"
	"encoding/json"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.ReorderRulesetResourceArray{
Data: []datadogV2.ReorderRulesetResourceData{
{
Id: datadog.PtrString("55ef2385-9ae1-4410-90c4-5ac1b60fec10"),
Type: datadogV2.REORDERRULESETRESOURCEDATATYPE_RULESET,
},
{
Id: datadog.PtrString("a7b8c9d0-1234-5678-9abc-def012345678"),
Type: datadogV2.REORDERRULESETRESOURCEDATATYPE_RULESET,
},
{
Id: datadog.PtrString("f1e2d3c4-b5a6-9780-1234-567890abcdef"),
Type: datadogV2.REORDERRULESETRESOURCEDATATYPE_RULESET,
},
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	r, err := api.ReorderTagPipelinesRulesets(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.ReorderTagPipelinesRulesets`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}