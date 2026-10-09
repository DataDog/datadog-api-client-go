// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureDataExplorerDestinationType The destination type. The value should always be `azure_data_explorer`.
type ObservabilityPipelineAzureDataExplorerDestinationType string

// List of ObservabilityPipelineAzureDataExplorerDestinationType.
const (
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONTYPE_AZURE_DATA_EXPLORER ObservabilityPipelineAzureDataExplorerDestinationType = "azure_data_explorer"
)

var allowedObservabilityPipelineAzureDataExplorerDestinationTypeEnumValues = []ObservabilityPipelineAzureDataExplorerDestinationType{
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONTYPE_AZURE_DATA_EXPLORER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAzureDataExplorerDestinationType) GetAllowedValues() []ObservabilityPipelineAzureDataExplorerDestinationType {
	return allowedObservabilityPipelineAzureDataExplorerDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAzureDataExplorerDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAzureDataExplorerDestinationType(value)
	return nil
}

// NewObservabilityPipelineAzureDataExplorerDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineAzureDataExplorerDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAzureDataExplorerDestinationTypeFromValue(v string) (*ObservabilityPipelineAzureDataExplorerDestinationType, error) {
	ev := ObservabilityPipelineAzureDataExplorerDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAzureDataExplorerDestinationType: valid values are %v", v, allowedObservabilityPipelineAzureDataExplorerDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAzureDataExplorerDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAzureDataExplorerDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAzureDataExplorerDestinationType value.
func (v ObservabilityPipelineAzureDataExplorerDestinationType) Ptr() *ObservabilityPipelineAzureDataExplorerDestinationType {
	return &v
}
