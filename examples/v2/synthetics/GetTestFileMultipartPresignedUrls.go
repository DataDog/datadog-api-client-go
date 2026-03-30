// Get presigned URLs for uploading a test file returns "OK" response

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
	body := datadogV2.SyntheticsTestFileMultipartPresignedUrlsRequest{
		BucketKeyPrefix: datadogV2.SYNTHETICSTESTFILEMULTIPARTPRESIGNEDURLSREQUESTBUCKETKEYPREFIX_API_UPLOAD_FILE,
		Parts: []datadogV2.SyntheticsTestFileMultipartPresignedUrlsPart{
			{
				Md5:        "1B2M2Y8AsgTpgAmY7PhCfg==",
				PartNumber: 1,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSyntheticsApi(apiClient)
	resp, r, err := api.GetTestFileMultipartPresignedUrls(ctx, "abc-def-123", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetTestFileMultipartPresignedUrls`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.GetTestFileMultipartPresignedUrls`:\n%s\n", responseContent)
}
