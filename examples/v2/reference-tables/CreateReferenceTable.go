// Create reference table returns "Created" response

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
	body := datadogV2.CreateTableRequest{
Data: &datadogV2.CreateTableRequestData{
Attributes: &datadogV2.CreateTableRequestDataAttributes{
Description: datadog.PtrString("this is a cloud table generated via a cloud bucket sync"),
FileMetadata: &datadogV2.CreateTableRequestDataAttributesFileMetadata{
CreateTableRequestDataAttributesFileMetadataCloudStorage: &datadogV2.CreateTableRequestDataAttributesFileMetadataCloudStorage{
AccessDetails: datadogV2.CreateTableRequestDataAttributesFileMetadataOneOfAccessDetails{
AwsDetail: &datadogV2.CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail{
AwsAccountId: "test-account-id",
AwsBucketName: "test-bucket",
FilePath: "test_rt.csv",
},
},
SyncEnabled: true,
}},
Schema: datadogV2.CreateTableRequestDataAttributesSchema{
Fields: []datadogV2.CreateTableRequestDataAttributesSchemaFieldsItems{
{
Name: "name",
Type: datadogV2.REFERENCETABLESCHEMAFIELDTYPE_STRING,
},
{
Name: "account_id",
Type: datadogV2.REFERENCETABLESCHEMAFIELDTYPE_STRING,
},
},
PrimaryKeys: []string{
"account_id",
},
},
Source: datadogV2.REFERENCETABLECREATESOURCETYPE_S3,
TableName: "test_reference_table",
Tags: []string{
"test_tag",
},
},
Type: datadogV2.CREATETABLEREQUESTDATATYPE_REFERENCE_TABLE,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewReferenceTablesApi(apiClient)
	resp, r, err := api.CreateReferenceTable(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceTablesApi.CreateReferenceTable`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ReferenceTablesApi.CreateReferenceTable`:\n%s\n", responseContent)
}