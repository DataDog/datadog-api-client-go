// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKind The Azure credential kind. The value should always be `managed_identity_client_assertion`.
type ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKind string

// List of ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKind.
const (
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONAUTHMANAGEDIDENTITYCLIENTASSERTIONKIND_MANAGED_IDENTITY_CLIENT_ASSERTION ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKind = "managed_identity_client_assertion"
)

var allowedObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKindEnumValues = []ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKind{
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONAUTHMANAGEDIDENTITYCLIENTASSERTIONKIND_MANAGED_IDENTITY_CLIENT_ASSERTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKind) GetAllowedValues() []ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKind {
	return allowedObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKind(value)
	return nil
}

// NewObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKindFromValue returns a pointer to a valid ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKindFromValue(v string) (*ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKind, error) {
	ev := ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKind: valid values are %v", v, allowedObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKind) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKind value.
func (v ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKind) Ptr() *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionKind {
	return &v
}
