// Create reference table with upload returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.CreateTableRequest{
		Data: &datadogV2.CreateTableRequestData{
			Attributes: &datadogV2.CreateTableRequestDataAttributes{
				Description: datadog.PtrString("Test reference table created via BDD test Example-Reference-Table"),
				Source:      datadogV2.REFERENCETABLECREATESOURCETYPE_LOCAL_FILE,
				FileMetadata: &datadogV2.CreateTableRequestDataAttributesFileMetadata{
					CreateTableRequestDataAttributesFileMetadataLocalFile: &datadogV2.CreateTableRequestDataAttributesFileMetadataLocalFile{
						UploadId: "test-upload-id-Example-Reference-Table",
					}},
				Schema: datadogV2.CreateTableRequestDataAttributesSchema{
					Fields: []datadogV2.CreateTableRequestDataAttributesSchemaFieldsItems{
						{
							Name: "id",
							Type: datadogV2.REFERENCETABLESCHEMAFIELDTYPE_STRING,
						},
						{
							Name: "name",
							Type: datadogV2.REFERENCETABLESCHEMAFIELDTYPE_STRING,
						},
						{
							Name: "value",
							Type: datadogV2.REFERENCETABLESCHEMAFIELDTYPE_INT32,
						},
					},
					PrimaryKeys: []string{
						"id",
					},
				},
				TableName: "test_reference_table_Example-Reference-Table",
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
	resp, r, err := api.CreateReferenceTable(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceTablesApi.CreateReferenceTable`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ReferenceTablesApi.CreateReferenceTable`:\n%s\n", responseContent)
}
