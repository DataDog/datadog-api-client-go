// Create an org group policy override returns "Created" response

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
	body := datadogV2.OrgGroupPolicyOverrideCreateRequest{
		Data: datadogV2.OrgGroupPolicyOverrideCreateData{
			Attributes: datadogV2.OrgGroupPolicyOverrideCreateAttributes{
				OrgSite: "datadoghq.com",
				OrgUuid: uuid.MustParse("c3d4e5f6-a7b8-9012-cdef-012345678901"),
			},
			Relationships: datadogV2.OrgGroupPolicyOverrideCreateRelationships{
				OrgGroup: datadogV2.OrgGroupRelationshipToOne{
					Data: datadogV2.OrgGroupRelationshipToOneData{
						Id:   uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef0123456789"),
						Type: datadogV2.ORGGROUPTYPE_ORG_GROUPS,
					},
				},
				OrgGroupPolicy: datadogV2.OrgGroupPolicyRelationshipToOne{
					Data: datadogV2.OrgGroupPolicyRelationshipToOneData{
						Id:   uuid.MustParse("1a2b3c4d-5e6f-7890-abcd-ef0123456789"),
						Type: datadogV2.ORGGROUPPOLICYTYPE_ORG_GROUP_POLICIES,
					},
				},
			},
			Type: datadogV2.ORGGROUPPOLICYOVERRIDETYPE_ORG_GROUP_POLICY_OVERRIDES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateOrgGroupPolicyOverride", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgGroupsApi(apiClient)
	resp, r, err := api.CreateOrgGroupPolicyOverride(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGroupsApi.CreateOrgGroupPolicyOverride`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrgGroupsApi.CreateOrgGroupPolicyOverride`:\n%s\n", responseContent)
}
