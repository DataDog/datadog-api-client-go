// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind The Azure credential kind. The value should always be `client_secret_credential`.
type ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind string

// List of ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind.
const (
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONAUTHCLIENTSECRETKIND_CLIENT_SECRET_CREDENTIAL ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind = "client_secret_credential"
)

var allowedObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKindEnumValues = []ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind{
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONAUTHCLIENTSECRETKIND_CLIENT_SECRET_CREDENTIAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind) GetAllowedValues() []ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind {
	return allowedObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind(value)
	return nil
}

// NewObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKindFromValue returns a pointer to a valid ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKindFromValue(v string) (*ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind, error) {
	ev := ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind: valid values are %v", v, allowedObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind value.
func (v ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind) Ptr() *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind {
	return &v
}
