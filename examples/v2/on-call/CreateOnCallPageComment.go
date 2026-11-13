// Create On-Call Page Comment returns "No Content" response

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
	body := datadogV2.CreateOnCallPageCommentRequest{
		Data: datadogV2.CreateOnCallPageCommentRequestData{
			Attributes: datadogV2.CreateOnCallPageCommentRequestAttributes{
				Body: "Comment from the API client integration tests.",
			},
			Id:   datadog.PtrUUID(uuid.MustParse("3c5f5b5e-3a3e-4f8c-9b1d-2f6a7e8d9c0b")),
			Type: datadogV2.ONCALLPAGECOMMENTTYPE_COMMENTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnCallApi(apiClient)
	r, err := api.CreateOnCallPageComment(ctx, "f63eb677-07ce-46ab-b2e0-80ca1ffe91bf", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallApi.CreateOnCallPageComment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
