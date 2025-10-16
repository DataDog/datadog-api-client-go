// Create Org Connection returns "OK" response

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
	body := datadogV2.OrgConnectionCreateRequest{
Data: datadogV2.OrgConnectionCreate{
Type: datadogV2.ORGCONNECTIONTYPE_ORG_CONNECTION,
Relationships: datadogV2.OrgConnectionCreateRelationships{
SinkOrg: datadogV2.OrgConnectionOrgRelationship{
Data: &datadogV2.OrgConnectionOrgRelationshipData{
Type: datadogV2.ORGCONNECTIONORGRELATIONSHIPDATATYPE_ORGS.Ptr(),
Id: datadog.PtrString("83999dcd-7f97-11f0-8de1-1ecf66f1aa85"),
},
},
},
Attributes: datadogV2.OrgConnectionCreateAttributes{
ConnectionTypes: []datadogV2.OrgConnectionTypeEnum{
datadogV2.ORGCONNECTIONTYPEENUM_LOGS,
},
},
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgConnectionsApi(apiClient)
	resp, r, err := api.CreateOrgConnections(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionsApi.CreateOrgConnections`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrgConnectionsApi.CreateOrgConnections`:\n%s\n", responseContent)
}