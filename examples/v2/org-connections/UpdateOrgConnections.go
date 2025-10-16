// Update Org Connection returns "OK" response

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
	// there is a valid "org_connection" in the system
	OrgConnectionDataID := uuid.MustParse(os.Getenv("ORG_CONNECTION_DATA_ID"))


	body := datadogV2.OrgConnectionUpdateRequest{
Data: datadogV2.OrgConnectionUpdate{
Type: datadogV2.ORGCONNECTIONTYPE_ORG_CONNECTION,
Id: OrgConnectionDataID,
Attributes: datadogV2.OrgConnectionUpdateAttributes{
ConnectionTypes: []datadogV2.OrgConnectionTypeEnum{
datadogV2.ORGCONNECTIONTYPEENUM_LOGS,
datadogV2.ORGCONNECTIONTYPEENUM_METRICS,
},
},
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrgConnectionsApi(apiClient)
	resp, r, err := api.UpdateOrgConnections(ctx, OrgConnectionDataID, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionsApi.UpdateOrgConnections`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OrgConnectionsApi.UpdateOrgConnections`:\n%s\n", responseContent)
}