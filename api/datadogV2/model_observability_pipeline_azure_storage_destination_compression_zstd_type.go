// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureStorageDestinationCompressionZstdType The compression type. Always `zstd`.
type ObservabilityPipelineAzureStorageDestinationCompressionZstdType string

// List of ObservabilityPipelineAzureStorageDestinationCompressionZstdType.
const (
	OBSERVABILITYPIPELINEAZURESTORAGEDESTINATIONCOMPRESSIONZSTDTYPE_ZSTD ObservabilityPipelineAzureStorageDestinationCompressionZstdType = "zstd"
)

var allowedObservabilityPipelineAzureStorageDestinationCompressionZstdTypeEnumValues = []ObservabilityPipelineAzureStorageDestinationCompressionZstdType{
	OBSERVABILITYPIPELINEAZURESTORAGEDESTINATIONCOMPRESSIONZSTDTYPE_ZSTD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAzureStorageDestinationCompressionZstdType) GetAllowedValues() []ObservabilityPipelineAzureStorageDestinationCompressionZstdType {
	return allowedObservabilityPipelineAzureStorageDestinationCompressionZstdTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAzureStorageDestinationCompressionZstdType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAzureStorageDestinationCompressionZstdType(value)
	return nil
}

// NewObservabilityPipelineAzureStorageDestinationCompressionZstdTypeFromValue returns a pointer to a valid ObservabilityPipelineAzureStorageDestinationCompressionZstdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAzureStorageDestinationCompressionZstdTypeFromValue(v string) (*ObservabilityPipelineAzureStorageDestinationCompressionZstdType, error) {
	ev := ObservabilityPipelineAzureStorageDestinationCompressionZstdType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAzureStorageDestinationCompressionZstdType: valid values are %v", v, allowedObservabilityPipelineAzureStorageDestinationCompressionZstdTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAzureStorageDestinationCompressionZstdType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAzureStorageDestinationCompressionZstdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAzureStorageDestinationCompressionZstdType value.
func (v ObservabilityPipelineAzureStorageDestinationCompressionZstdType) Ptr() *ObservabilityPipelineAzureStorageDestinationCompressionZstdType {
	return &v
}
