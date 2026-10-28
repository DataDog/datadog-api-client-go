// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureDataExplorerDestinationIngestionMode The ingestion method used to send data to Azure Data Explorer.
//
// - `streaming`: Sends events directly to the table using streaming ingestion.
// - `queue`: Sends events through queued (batched) ingestion.
//
// Defaults to `streaming`.
type ObservabilityPipelineAzureDataExplorerDestinationIngestionMode string

// List of ObservabilityPipelineAzureDataExplorerDestinationIngestionMode.
const (
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONINGESTIONMODE_STREAMING ObservabilityPipelineAzureDataExplorerDestinationIngestionMode = "streaming"
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONINGESTIONMODE_QUEUE     ObservabilityPipelineAzureDataExplorerDestinationIngestionMode = "queue"
)

var allowedObservabilityPipelineAzureDataExplorerDestinationIngestionModeEnumValues = []ObservabilityPipelineAzureDataExplorerDestinationIngestionMode{
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONINGESTIONMODE_STREAMING,
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONINGESTIONMODE_QUEUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAzureDataExplorerDestinationIngestionMode) GetAllowedValues() []ObservabilityPipelineAzureDataExplorerDestinationIngestionMode {
	return allowedObservabilityPipelineAzureDataExplorerDestinationIngestionModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAzureDataExplorerDestinationIngestionMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAzureDataExplorerDestinationIngestionMode(value)
	return nil
}

// NewObservabilityPipelineAzureDataExplorerDestinationIngestionModeFromValue returns a pointer to a valid ObservabilityPipelineAzureDataExplorerDestinationIngestionMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAzureDataExplorerDestinationIngestionModeFromValue(v string) (*ObservabilityPipelineAzureDataExplorerDestinationIngestionMode, error) {
	ev := ObservabilityPipelineAzureDataExplorerDestinationIngestionMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAzureDataExplorerDestinationIngestionMode: valid values are %v", v, allowedObservabilityPipelineAzureDataExplorerDestinationIngestionModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAzureDataExplorerDestinationIngestionMode) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAzureDataExplorerDestinationIngestionModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAzureDataExplorerDestinationIngestionMode value.
func (v ObservabilityPipelineAzureDataExplorerDestinationIngestionMode) Ptr() *ObservabilityPipelineAzureDataExplorerDestinationIngestionMode {
	return &v
}
