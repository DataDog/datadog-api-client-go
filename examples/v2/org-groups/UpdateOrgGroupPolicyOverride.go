// Update an org group policy override returns "OK" response

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
	body := datadogV2.OrgGroupPolicyOverrideUpdateRequest{
		Data: datadogV2.OrgGroupPolicyOverrideUpdateData{
			Attributes: datadogV2.OrgGroupPolicyOverrideUpdateAttributes{
				OrgSite: "us1",
				OrgUuid: uuid.MustParse("c3d4e5f6-a7b8-9012-cdef-012345678901"),
			},
			Id:   uuid.MustParse("9f8e7d6c-5b4a-3210-fedc-ba0987654321"),
			Type: datadogV2.ORGGROUPPOLICYOVERRIDETYPE_ORG_GROUP_POLICY_OVERRIDES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateOrgGroupPolicyOverride", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgGroupsApi(apiClient)
	resp, r, err := api.UpdateOrgGroupPolicyOverride(ctx, uuid.MustParse("9f8e7d6c-5b4a-3210-fedc-ba0987654321"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGroupsApi.UpdateOrgGroupPolicyOverride`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrgGroupsApi.UpdateOrgGroupPolicyOverride`:\n%s\n", responseContent)
}
