// Mute a host returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.HostMuteSettings{
		End:      common.PtrInt64(1579098130),
		Message:  common.PtrString("Muting this host for a test!"),
		Override: common.PtrBool(false),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewHostsApi(apiClient)
	resp, r, err := api.MuteHost(ctx, "host_name", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.MuteHost`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `HostsApi.MuteHost`:\n%s\n", responseContent)
}
