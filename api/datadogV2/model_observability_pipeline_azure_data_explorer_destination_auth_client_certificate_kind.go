// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind The Azure credential kind. The value should always be `client_certificate_credential`.
type ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind string

// List of ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind.
const (
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONAUTHCLIENTCERTIFICATEKIND_CLIENT_CERTIFICATE_CREDENTIAL ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind = "client_certificate_credential"
)

var allowedObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKindEnumValues = []ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind{
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONAUTHCLIENTCERTIFICATEKIND_CLIENT_CERTIFICATE_CREDENTIAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind) GetAllowedValues() []ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind {
	return allowedObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind(value)
	return nil
}

// NewObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKindFromValue returns a pointer to a valid ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKindFromValue(v string) (*ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind, error) {
	ev := ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind: valid values are %v", v, allowedObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind value.
func (v ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind) Ptr() *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind {
	return &v
}
