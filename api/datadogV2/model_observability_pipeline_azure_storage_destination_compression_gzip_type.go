// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureStorageDestinationCompressionGzipType The compression type. Always `gzip`.
type ObservabilityPipelineAzureStorageDestinationCompressionGzipType string

// List of ObservabilityPipelineAzureStorageDestinationCompressionGzipType.
const (
	OBSERVABILITYPIPELINEAZURESTORAGEDESTINATIONCOMPRESSIONGZIPTYPE_GZIP ObservabilityPipelineAzureStorageDestinationCompressionGzipType = "gzip"
)

var allowedObservabilityPipelineAzureStorageDestinationCompressionGzipTypeEnumValues = []ObservabilityPipelineAzureStorageDestinationCompressionGzipType{
	OBSERVABILITYPIPELINEAZURESTORAGEDESTINATIONCOMPRESSIONGZIPTYPE_GZIP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAzureStorageDestinationCompressionGzipType) GetAllowedValues() []ObservabilityPipelineAzureStorageDestinationCompressionGzipType {
	return allowedObservabilityPipelineAzureStorageDestinationCompressionGzipTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAzureStorageDestinationCompressionGzipType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAzureStorageDestinationCompressionGzipType(value)
	return nil
}

// NewObservabilityPipelineAzureStorageDestinationCompressionGzipTypeFromValue returns a pointer to a valid ObservabilityPipelineAzureStorageDestinationCompressionGzipType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAzureStorageDestinationCompressionGzipTypeFromValue(v string) (*ObservabilityPipelineAzureStorageDestinationCompressionGzipType, error) {
	ev := ObservabilityPipelineAzureStorageDestinationCompressionGzipType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAzureStorageDestinationCompressionGzipType: valid values are %v", v, allowedObservabilityPipelineAzureStorageDestinationCompressionGzipTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAzureStorageDestinationCompressionGzipType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAzureStorageDestinationCompressionGzipTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAzureStorageDestinationCompressionGzipType value.
func (v ObservabilityPipelineAzureStorageDestinationCompressionGzipType) Ptr() *ObservabilityPipelineAzureStorageDestinationCompressionGzipType {
	return &v
}
