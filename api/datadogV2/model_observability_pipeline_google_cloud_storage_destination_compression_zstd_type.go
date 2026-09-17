// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType The compression type. Always `zstd`.
type ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType string

// List of ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType.
const (
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONCOMPRESSIONZSTDTYPE_ZSTD ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType = "zstd"
)

var allowedObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdTypeEnumValues = []ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType{
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONCOMPRESSIONZSTDTYPE_ZSTD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType) GetAllowedValues() []ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType {
	return allowedObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType(value)
	return nil
}

// NewObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdTypeFromValue returns a pointer to a valid ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdTypeFromValue(v string) (*ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType, error) {
	ev := ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType: valid values are %v", v, allowedObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType value.
func (v ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType) Ptr() *ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType {
	return &v
}
