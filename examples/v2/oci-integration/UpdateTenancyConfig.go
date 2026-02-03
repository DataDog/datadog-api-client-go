// Update tenancy config returns "OK" response

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
	body := datadogV2.UpdateTenancyConfigRequest{
		Data: datadogV2.UpdateTenancyConfigData{
			Attributes: &datadogV2.UpdateTenancyConfigDataAttributes{
				AuthCredentials: &datadogV2.UpdateTenancyConfigDataAttributesAuthCredentials{
					Fingerprint: datadog.PtrString(""),
					PrivateKey: `----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCdvSMmlfLyeD4M
QsA3WlrWBqKdWa5eVV3/uODyqT3wWMEMIJHcG3/quNs8nh9xrK1/JkQT2qoKEHqR
C5k59jN6Vp8em8ARJthMgam9K37ELt+IQ/G8ySTSuqZG8T4cHp/cs3fAclNqttOl
YnGr4RbVAgMBAAECggEAGZNLGbyCUbIRTW6Kh4d8ZVC+eZtJMqGmGJ3KfVaW8Pjn
QGWfSuJCEe2o2Y8G3phlidFauICnZ44enXA17Rhi+I/whnr7FIyQk2bR7rv+1Uhc
mOJygWX5eFFMsledgVAdIAl9Luk2nykx7Un3g6rtbl/Vs+5k4m7ITLFMpCHzsJLU
nm8kBzDOqY2JUkMd08nL88KL6QywWtal05UESzQpNFXd0e5kxYfexeMCsLsWP0mc
quMLRbn7NuBjCbe9VU2kmIvcfDDaWjurT7d5m1OXx1cc8p6P4PFZTVyCjdhiWOr3
LQXZ4/vdZNR3zgEHypRoM6D9Yq99LWUOUEMrdiSLQQKBgQDQkh7C1OtAXnpy7F6R
W+/I3zBHici2p7A57UT7VECQ1IVGg37/uus83DkuOtdZ33JmHLAVrwLFJvUlbyjx
l6dc/1ms40L5HFdLgaVtd4k0rSPFeOSDr6evz0lX4yBuzlP0fEh+o3XHW7mwe2G+
rWCULF/Uqza66fjbCSKMNgLIXQKBgQDBm9nZg/s4S0THWCFNWcB1tXBG0p/sH5eY
PC1H/VmTEINIixStrS4ufczf31X8rcoSjSbO7+vZDTTATdk7OLn1I2uGFVYl8M59
86BYT2Hi7cwp7YVzOc/cJigVeBAqSRW/iYYyWBEUTiW1gbkV0sRWwhPp67m+c0sP
XpY/iEZA2QKBgB1w8tynt4l/jKNaUEMOijt9ndALWATIiOy0XG9pxi9rgGCiwTOS
DBCsOXoYHjv2eayGUijNaoOv6xzcoxfvQ1WySdNIxTRq1ru20kYwgHKqGgmO9hrM
mcwMY5r/WZ2qjFlPjeAqbL62aPDLidGjoaVo2iIoBPK/gjxQ/5f0MS4N/YQ0zWoYBueSQ0DGs
-----END PRIVATE KEY-----`,
				},
				CostCollectionEnabled: datadog.PtrBool(true),
				HomeRegion:            datadog.PtrString("us-ashburn-1"),
				LogsConfig: &datadogV2.UpdateTenancyConfigDataAttributesLogsConfig{
					CompartmentTagFilters: []string{
						"datadog:true",
						"env:prod",
					},
					Enabled: datadog.PtrBool(true),
					EnabledServices: []string{
						"service_1",
						"service_1",
					},
				},
				MetricsConfig: &datadogV2.UpdateTenancyConfigDataAttributesMetricsConfig{
					CompartmentTagFilters: []string{
						"datadog:true",
						"env:prod",
					},
					Enabled: datadog.PtrBool(true),
					ExcludedServices: []string{
						"service_1",
						"service_1",
					},
				},
				RegionsConfig: &datadogV2.UpdateTenancyConfigDataAttributesRegionsConfig{
					Available: []string{
						"us-ashburn-1",
						"us-phoenix-1",
					},
					Disabled: []string{
						"us-phoenix-1",
					},
					Enabled: []string{
						"us-ashburn-1",
					},
				},
				ResourceCollectionEnabled: datadog.PtrBool(true),
				UserOcid:                  datadog.PtrString("ocid.user.test"),
			},
			Id:   "ocid.tenancy.test",
			Type: datadogV2.UPDATETENANCYCONFIGDATATYPE_OCI_TENANCY,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOCIIntegrationApi(apiClient)
	resp, r, err := api.UpdateTenancyConfig(ctx, "tenancy_ocid", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OCIIntegrationApi.UpdateTenancyConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OCIIntegrationApi.UpdateTenancyConfig`:\n%s\n", responseContent)
}
