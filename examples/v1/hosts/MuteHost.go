// Mute a host returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.HostMuteSettings{
		End:      datadog.PtrInt64(1579098130),
		Message:  datadog.PtrString("Muting this host for a test!"),
		Override: datadog.PtrBool(false),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsApi.MuteHost(ctx, "host_name", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.MuteHost`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `HostsApi.MuteHost`:\n%s\n", responseContent)
}
