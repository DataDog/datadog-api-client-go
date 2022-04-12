// Post an event with a long title returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.EventCreateRequest{
		Title: "Example-Post_an_event_with_a_long_title_returns_OK_response very very very looooooooong looooooooooooong loooooooooooooooooooooong looooooooooooooooooooooooooong title with 100+ characters",
		Text:  "A text message.",
		Tags: []string{
			"test:ExamplePostaneventwithalongtitlereturnsOKresponse",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsApi.CreateEvent(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.CreateEvent`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `EventsApi.CreateEvent`:\n%s\n", responseContent)
}
