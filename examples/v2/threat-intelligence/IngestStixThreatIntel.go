// Ingest STIX threat intelligence returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.STIXBundleRequest{
		Id: "bundle--44444444-4444-4444-8444-444444444444",
		Objects: []datadogV2.STIXIndicatorObject{
			{
				Created:     time.Date(2026, 7, 22, 12, 0, 0, 0, time.UTC),
				Id:          "indicator--55555555-5555-4555-8555-555555555555",
				Modified:    time.Date(2026, 7, 22, 12, 0, 0, 0, time.UTC),
				Pattern:     "[ipv4-addr:value = '198.51.100.42']",
				PatternType: datadogV2.STIXPATTERNTYPE_STIX,
				SpecVersion: datadogV2.STIXSPECVERSION_VERSION_2_1,
				Type:        datadogV2.STIXINDICATORTYPE_INDICATOR,
				ValidFrom:   time.Date(2026, 7, 22, 12, 0, 0, 0, time.UTC),
			},
		},
		SpecVersion: datadogV2.STIXSPECVERSION_VERSION_2_1.Ptr(),
		Type:        datadogV2.STIXBUNDLETYPE_BUNDLE,
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.IngestStixThreatIntel", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewThreatIntelligenceApi(apiClient)
	resp, r, err := api.IngestStixThreatIntel(ctx, "Acme-Inc", body, *datadogV2.NewIngestStixThreatIntelOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatIntelligenceApi.IngestStixThreatIntel`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ThreatIntelligenceApi.IngestStixThreatIntel`:\n%s\n", responseContent)
}
