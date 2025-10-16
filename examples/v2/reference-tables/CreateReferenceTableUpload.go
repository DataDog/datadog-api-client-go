// Create reference table upload returns "Created" response

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
	body := datadogV2.CreateUploadRequest{
Data: &datadogV2.CreateUploadRequestData{
Attributes: &datadogV2.CreateUploadRequestDataAttributes{
Headers: []string{
"id",
"name",
"value",
},
TableName: "test_upload_table_Example-Reference-Table",
PartCount: 1,
PartSize: 1024,
},
Type: datadogV2.CREATEUPLOADREQUESTDATATYPE_UPLOAD,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewReferenceTablesApi(apiClient)
	resp, r, err := api.CreateReferenceTableUpload(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceTablesApi.CreateReferenceTableUpload`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ReferenceTablesApi.CreateReferenceTableUpload`:\n%s\n", responseContent)
}