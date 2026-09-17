// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3DestinationCompressionGzipType The compression type. Always `gzip`.
type ObservabilityPipelineAmazonS3DestinationCompressionGzipType string

// List of ObservabilityPipelineAmazonS3DestinationCompressionGzipType.
const (
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONCOMPRESSIONGZIPTYPE_GZIP ObservabilityPipelineAmazonS3DestinationCompressionGzipType = "gzip"
)

var allowedObservabilityPipelineAmazonS3DestinationCompressionGzipTypeEnumValues = []ObservabilityPipelineAmazonS3DestinationCompressionGzipType{
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONCOMPRESSIONGZIPTYPE_GZIP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAmazonS3DestinationCompressionGzipType) GetAllowedValues() []ObservabilityPipelineAmazonS3DestinationCompressionGzipType {
	return allowedObservabilityPipelineAmazonS3DestinationCompressionGzipTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAmazonS3DestinationCompressionGzipType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAmazonS3DestinationCompressionGzipType(value)
	return nil
}

// NewObservabilityPipelineAmazonS3DestinationCompressionGzipTypeFromValue returns a pointer to a valid ObservabilityPipelineAmazonS3DestinationCompressionGzipType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAmazonS3DestinationCompressionGzipTypeFromValue(v string) (*ObservabilityPipelineAmazonS3DestinationCompressionGzipType, error) {
	ev := ObservabilityPipelineAmazonS3DestinationCompressionGzipType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAmazonS3DestinationCompressionGzipType: valid values are %v", v, allowedObservabilityPipelineAmazonS3DestinationCompressionGzipTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAmazonS3DestinationCompressionGzipType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAmazonS3DestinationCompressionGzipTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAmazonS3DestinationCompressionGzipType value.
func (v ObservabilityPipelineAmazonS3DestinationCompressionGzipType) Ptr() *ObservabilityPipelineAmazonS3DestinationCompressionGzipType {
	return &v
}
