// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind The Azure credential kind. The value should always be `azure_cli`.
type ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind string

// List of ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind.
const (
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONAUTHAZURECLIKIND_AZURE_CLI ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind = "azure_cli"
)

var allowedObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKindEnumValues = []ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind{
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONAUTHAZURECLIKIND_AZURE_CLI,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind) GetAllowedValues() []ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind {
	return allowedObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind(value)
	return nil
}

// NewObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKindFromValue returns a pointer to a valid ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKindFromValue(v string) (*ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind, error) {
	ev := ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind: valid values are %v", v, allowedObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind value.
func (v ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind) Ptr() *ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind {
	return &v
}
