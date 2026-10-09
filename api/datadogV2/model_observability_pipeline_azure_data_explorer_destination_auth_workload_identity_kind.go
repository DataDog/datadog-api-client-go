// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind The Azure credential kind. The value should always be `workload_identity`.
type ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind string

// List of ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind.
const (
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONAUTHWORKLOADIDENTITYKIND_WORKLOAD_IDENTITY ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind = "workload_identity"
)

var allowedObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKindEnumValues = []ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind{
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONAUTHWORKLOADIDENTITYKIND_WORKLOAD_IDENTITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind) GetAllowedValues() []ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind {
	return allowedObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind(value)
	return nil
}

// NewObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKindFromValue returns a pointer to a valid ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKindFromValue(v string) (*ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind, error) {
	ev := ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind: valid values are %v", v, allowedObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind value.
func (v ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind) Ptr() *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind {
	return &v
}
