// Update reference table returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.PatchTableRequest{
		Data: &datadogV2.PatchTableRequestData{
			Attributes: &datadogV2.PatchTableRequestDataAttributes{
				Description: datadog.PtrString("this is a cloud table generated via a cloud bucket sync"),
				FileMetadata: &datadogV2.PatchTableRequestDataAttributesFileMetadata{
					PatchTableRequestDataAttributesFileMetadataCloudStorage: &datadogV2.PatchTableRequestDataAttributesFileMetadataCloudStorage{
						AccessDetails: &datadogV2.PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails{
							AwsDetail: &datadogV2.PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail{
								AwsAccountId:  datadog.PtrString("test-account-id"),
								AwsBucketName: datadog.PtrString("test-bucket"),
								FilePath:      datadog.PtrString("test_rt.csv"),
							},
						},
						SyncEnabled: datadog.PtrBool(true),
					}},
				Schema: &datadogV2.PatchTableRequestDataAttributesSchema{
					Fields: []datadogV2.PatchTableRequestDataAttributesSchemaFieldsItems{
						{
							Name: "id",
							Type: datadogV2.REFERENCETABLESCHEMAFIELDTYPE_INT32,
						},
						{
							Name: "name",
							Type: datadogV2.REFERENCETABLESCHEMAFIELDTYPE_STRING,
						},
					},
					PrimaryKeys: []string{
						"id",
					},
				},
				Tags: []string{
					"test_tag",
				},
			},
			Type: datadogV2.PATCHTABLEREQUESTDATATYPE_REFERENCE_TABLE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewReferenceTablesApi(apiClient)
	r, err := api.UpdateReferenceTable(ctx, "id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceTablesApi.UpdateReferenceTable`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
