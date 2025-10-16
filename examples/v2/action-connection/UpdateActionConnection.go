// Update an existing Action Connection returns "Successfully updated Action Connection" response

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
	body := datadogV2.UpdateActionConnectionRequest{
Data: datadogV2.ActionConnectionDataUpdate{
Type: datadogV2.ACTIONCONNECTIONDATATYPE_ACTION_CONNECTION,
Attributes: datadogV2.ActionConnectionAttributesUpdate{
Name: datadog.PtrString("Cassette Connection"),
Integration: &datadogV2.ActionConnectionIntegrationUpdate{
AWSIntegrationUpdate: &datadogV2.AWSIntegrationUpdate{
Type: datadogV2.AWSINTEGRATIONTYPE_AWS,
Credentials: &datadogV2.AWSCredentialsUpdate{
AWSAssumeRoleUpdate: &datadogV2.AWSAssumeRoleUpdate{
Type: datadogV2.AWSASSUMEROLETYPE_AWSASSUMEROLE,
Role: datadog.PtrString("MyRoleUpdated"),
AccountId: datadog.PtrString("123456789123"),
}},
}},
},
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewActionConnectionApi(apiClient)
	resp, r, err := api.UpdateActionConnection(ctx, "cb460d51-3c88-4e87-adac-d47131d0423d", body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionConnectionApi.UpdateActionConnection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ActionConnectionApi.UpdateActionConnection`:\n%s\n", responseContent)
}