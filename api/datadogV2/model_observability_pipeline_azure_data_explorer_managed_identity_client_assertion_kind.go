// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind The Azure credential kind. The value should always be `managed_identity_client_assertion`.
type ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind string

// List of ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind.
const (
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERMANAGEDIDENTITYCLIENTASSERTIONKIND_MANAGED_IDENTITY_CLIENT_ASSERTION ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind = "managed_identity_client_assertion"
)

var allowedObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKindEnumValues = []ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind{
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERMANAGEDIDENTITYCLIENTASSERTIONKIND_MANAGED_IDENTITY_CLIENT_ASSERTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind) GetAllowedValues() []ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind {
	return allowedObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind(value)
	return nil
}

// NewObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKindFromValue returns a pointer to a valid ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKindFromValue(v string) (*ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind, error) {
	ev := ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind: valid values are %v", v, allowedObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind value.
func (v ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind) Ptr() *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind {
	return &v
}
