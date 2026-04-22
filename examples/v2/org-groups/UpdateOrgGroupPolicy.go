// Update an org group policy returns "OK" response

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
	body := datadogV2.OrgGroupPolicyUpdateRequest{
		Data: datadogV2.OrgGroupPolicyUpdateData{
			Attributes: datadogV2.OrgGroupPolicyUpdateAttributes{
				Content: map[string]interface{}{
					"value": "UTC",
				},
			},
			Id:   uuid.MustParse("1a2b3c4d-5e6f-7890-abcd-ef0123456789"),
			Type: datadogV2.ORGGROUPPOLICYTYPE_ORG_GROUP_POLICIES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateOrgGroupPolicy", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgGroupsApi(apiClient)
	resp, r, err := api.UpdateOrgGroupPolicy(ctx, uuid.MustParse("1a2b3c4d-5e6f-7890-abcd-ef0123456789"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGroupsApi.UpdateOrgGroupPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrgGroupsApi.UpdateOrgGroupPolicy`:\n%s\n", responseContent)
}
