// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3DestinationCompressionZstdType The compression type. Always `zstd`.
type ObservabilityPipelineAmazonS3DestinationCompressionZstdType string

// List of ObservabilityPipelineAmazonS3DestinationCompressionZstdType.
const (
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONCOMPRESSIONZSTDTYPE_ZSTD ObservabilityPipelineAmazonS3DestinationCompressionZstdType = "zstd"
)

var allowedObservabilityPipelineAmazonS3DestinationCompressionZstdTypeEnumValues = []ObservabilityPipelineAmazonS3DestinationCompressionZstdType{
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONCOMPRESSIONZSTDTYPE_ZSTD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAmazonS3DestinationCompressionZstdType) GetAllowedValues() []ObservabilityPipelineAmazonS3DestinationCompressionZstdType {
	return allowedObservabilityPipelineAmazonS3DestinationCompressionZstdTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAmazonS3DestinationCompressionZstdType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAmazonS3DestinationCompressionZstdType(value)
	return nil
}

// NewObservabilityPipelineAmazonS3DestinationCompressionZstdTypeFromValue returns a pointer to a valid ObservabilityPipelineAmazonS3DestinationCompressionZstdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAmazonS3DestinationCompressionZstdTypeFromValue(v string) (*ObservabilityPipelineAmazonS3DestinationCompressionZstdType, error) {
	ev := ObservabilityPipelineAmazonS3DestinationCompressionZstdType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAmazonS3DestinationCompressionZstdType: valid values are %v", v, allowedObservabilityPipelineAmazonS3DestinationCompressionZstdTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAmazonS3DestinationCompressionZstdType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAmazonS3DestinationCompressionZstdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAmazonS3DestinationCompressionZstdType value.
func (v ObservabilityPipelineAmazonS3DestinationCompressionZstdType) Ptr() *ObservabilityPipelineAmazonS3DestinationCompressionZstdType {
	return &v
}
