// Creates a data deletion request returns "OK" response

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
	body := datadogV2.CreateDataDeletionRequestBody{
Data: datadogV2.CreateDataDeletionRequestBodyData{
Attributes: datadogV2.CreateDataDeletionRequestBodyAttributes{
From: 1672527600000,
Indexes: []string{
"test-index",
"test-index-2",
},
Query: map[string]string{
"host": "abc",
"service": "xyz",
},
To: 1704063600000,
},
Type: datadogV2.CREATEDATADELETIONREQUESTBODYDATATYPE_CREATE_DELETION_REQ,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateDataDeletionRequest", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDataDeletionApi(apiClient)
	resp, r, err := api.CreateDataDeletionRequest(ctx, "logs", body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataDeletionApi.CreateDataDeletionRequest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DataDeletionApi.CreateDataDeletionRequest`:\n%s\n", responseContent)
}