// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType The type of the user-assigned managed identity ID.
type ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType string

// List of ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType.
const (
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONMANAGEDIDENTITYIDTYPE_CLIENT_ID   ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType = "client_id"
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONMANAGEDIDENTITYIDTYPE_OBJECT_ID   ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType = "object_id"
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONMANAGEDIDENTITYIDTYPE_RESOURCE_ID ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType = "resource_id"
)

var allowedObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdTypeEnumValues = []ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType{
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONMANAGEDIDENTITYIDTYPE_CLIENT_ID,
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONMANAGEDIDENTITYIDTYPE_OBJECT_ID,
	OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONMANAGEDIDENTITYIDTYPE_RESOURCE_ID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType) GetAllowedValues() []ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType {
	return allowedObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType(value)
	return nil
}

// NewObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdTypeFromValue returns a pointer to a valid ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdTypeFromValue(v string) (*ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType, error) {
	ev := ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType: valid values are %v", v, allowedObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType value.
func (v ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType) Ptr() *ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType {
	return &v
}
