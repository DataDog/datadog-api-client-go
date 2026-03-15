// Import vulnerabilities returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.CycloneDXBOM{
		BomFormat: "CycloneDX",
		Components: []datadogV2.CycloneDXComponent{
			{
				BomRef:  "a3390fca-c315-41ae-ae05-af5e7859cdee",
				Name:    "lodash",
				Purl:    datadog.PtrString("pkg:npm/lodash@4.17.21"),
				Type:    datadogV2.CYCLONEDXCOMPONENTTYPE_LIBRARY,
				Version: "4.17.21",
			},
		},
		Metadata: datadogV2.CycloneDXMetadata{
			Component: datadogV2.CycloneDXAssetComponent{
				BomRef: datadog.PtrString("asset-ref-123"),
				Name:   "i-12345",
				Type:   datadog.PtrString("operating-system"),
			},
			Tools: datadogV2.CycloneDXTools{
				Components: []datadogV2.CycloneDXToolComponent{
					{
						Name: "my-scanner",
						Type: datadog.PtrString("application"),
					},
				},
			},
		},
		SpecVersion: "1.5",
		Version:     1,
		Vulnerabilities: []datadogV2.CycloneDXVulnerability{
			{
				Advisories: []datadogV2.CycloneDXAdvisory{
					{
						Url: datadog.PtrString("https://example.com/advisory/CVE-2021-1234"),
					},
				},
				Affects: []datadogV2.CycloneDXAffect{
					{
						Ref: "a3390fca-c315-41ae-ae05-af5e7859cdee",
					},
				},
				Cwes: []int32{
					123,
					345,
				},
				Description: datadog.PtrString("Sample vulnerability detected in the application."),
				Detail:      datadog.PtrString("Details about the vulnerability"),
				Id:          "CVE-2021-1234",
				Ratings: []datadogV2.CycloneDXRating{
					{
						Score:    datadog.PtrFloat64(9.0),
						Severity: datadog.PtrString("high"),
						Vector:   datadog.PtrString("CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:C/C:H/I:H/A:N"),
					},
				},
				References: []datadogV2.CycloneDXReference{
					{
						Id: datadog.PtrString("GHSA-35m5-8cvj-8783"),
						Source: &datadogV2.CycloneDXReferenceSource{
							Url: datadog.PtrString("https://example.com"),
						},
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ImportSecurityVulnerabilities", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	r, err := api.ImportSecurityVulnerabilities(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ImportSecurityVulnerabilities`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
