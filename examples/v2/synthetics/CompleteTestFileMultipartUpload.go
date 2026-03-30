// Complete a multipart upload of a test file returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.SyntheticsTestFileCompleteMultipartUploadRequest{
		Key: "org-123/api-upload-file/abc-def-123/2024-01-01T00:00:00_uuid.json",
		Parts: []datadogV2.SyntheticsTestFileCompleteMultipartUploadPart{
			{
				ETag:       `"d41d8cd98f00b204e9800998ecf8427e"`,
				PartNumber: 1,
			},
		},
		UploadId: "upload-id-abc123",
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSyntheticsApi(apiClient)
	r, err := api.CompleteTestFileMultipartUpload(ctx, "abc-def-123", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CompleteTestFileMultipartUpload`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
