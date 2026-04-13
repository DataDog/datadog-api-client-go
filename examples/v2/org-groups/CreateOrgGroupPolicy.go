// Create an org group policy returns "Created" response

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
	body := datadogV2.OrgGroupPolicyCreateRequest{
		Data: datadogV2.OrgGroupPolicyCreateData{
			Attributes: datadogV2.OrgGroupPolicyCreateAttributes{
				Content: map[string]interface{}{
					"value": "UTC",
				},
				PolicyName: "monitor_timezone",
			},
			Relationships: datadogV2.OrgGroupPolicyCreateRelationships{
				OrgGroup: datadogV2.OrgGroupRelationshipToOne{
					Data: datadogV2.OrgGroupRelationshipToOneData{
						Id:   uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
						Type: datadogV2.ORGGROUPTYPE_ORG_GROUPS,
					},
				},
			},
			Type: datadogV2.ORGGROUPPOLICYTYPE_ORG_GROUP_POLICIES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateOrgGroupPolicy", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgGroupsApi(apiClient)
	resp, r, err := api.CreateOrgGroupPolicy(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGroupsApi.CreateOrgGroupPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrgGroupsApi.CreateOrgGroupPolicy`:\n%s\n", responseContent)
}
