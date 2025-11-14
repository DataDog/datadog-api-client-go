// Create reference table returns "Created" response

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
				FileMetadata: &datadogV2.CreateTableRequestDataAttributesFileMetadata{
					CreateTableRequestDataAttributesFileMetadataCloudStorage: &datadogV2.CreateTableRequestDataAttributesFileMetadataCloudStorage{
						AccessDetails: datadogV2.CreateTableRequestDataAttributesFileMetadataOneOfAccessDetails{
							AwsDetail: &datadogV2.CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail{
								AwsAccountId:  "123456789000",
								AwsBucketName: "example-data-bucket",
								FilePath:      "reference-tables/users.csv",
							},
							AzureDetail: &datadogV2.CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail{
								AzureClientId:           "aaaaaaaa-1111-2222-3333-bbbbbbbbbbbb",
								AzureContainerName:      "reference-data",
								AzureStorageAccountName: "examplestorageaccount",
								AzureTenantId:           "cccccccc-4444-5555-6666-dddddddddddd",
								FilePath:                "tables/users.csv",
							},
							GcpDetail: &datadogV2.CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail{
								FilePath:               "data/reference_tables/users.csv",
								GcpBucketName:          "example-data-bucket",
								GcpProjectId:           "example-gcp-project-12345",
								GcpServiceAccountEmail: "example-service@example-gcp-project-12345.iam.gserviceaccount.com",
							},
						},
						SyncEnabled: false,
					}},
				Schema: datadogV2.CreateTableRequestDataAttributesSchema{
					Fields: []datadogV2.CreateTableRequestDataAttributesSchemaFieldsItems{
						{
							Name: "field_1",
							Type: datadogV2.REFERENCETABLESCHEMAFIELDTYPE_STRING,
						},
					},
					PrimaryKeys: []string{
						"field_1",
					},
				},
				Source:    datadogV2.REFERENCETABLECREATESOURCETYPE_LOCAL_FILE,
				TableName: "table_1",
				Tags: []string{
					"tag_1",
					"tag_2",
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
