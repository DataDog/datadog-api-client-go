// Create a new Action Connection returns "Successfully created Action Connection" response

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
	body := datadogV2.CreateActionConnectionRequest{
Data: datadogV2.ActionConnectionData{
Type: datadogV2.ACTIONCONNECTIONDATATYPE_ACTION_CONNECTION,
Attributes: datadogV2.ActionConnectionAttributes{
Name: "Cassette Connection exampleactionconnection",
Integration: datadogV2.ActionConnectionIntegration{
AWSIntegration: &datadogV2.AWSIntegration{
Type: datadogV2.AWSINTEGRATIONTYPE_AWS,
Credentials: datadogV2.AWSCredentials{
AWSAssumeRole: &datadogV2.AWSAssumeRole{
Type: datadogV2.AWSASSUMEROLETYPE_AWSASSUMEROLE,
Role: "MyRoleUpdated",
AccountId: "123456789123",
}},
}},
},
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewActionConnectionApi(apiClient)
	resp, r, err := api.CreateActionConnection(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionConnectionApi.CreateActionConnection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ActionConnectionApi.CreateActionConnection`:\n%s\n", responseContent)
}