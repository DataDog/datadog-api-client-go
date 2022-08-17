// Mute a host returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.HostMuteSettings{
		End:      datadog.PtrInt64(1579098130),
		Message:  datadog.PtrString("Muting this host for a test!"),
		Override: datadog.PtrBool(false),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewHostsApi(apiClient)
	resp, r, err := api.MuteHost(ctx, "host_name", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.MuteHost`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `HostsApi.MuteHost`:\n%s\n", responseContent)
}
