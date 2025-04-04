// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineKafkaSourceType The source type. The value should always be `kafka`.
type PipelineKafkaSourceType string

// List of PipelineKafkaSourceType.
const (
	PIPELINEKAFKASOURCETYPE_KAFKA PipelineKafkaSourceType = "kafka"
)

var allowedPipelineKafkaSourceTypeEnumValues = []PipelineKafkaSourceType{
	PIPELINEKAFKASOURCETYPE_KAFKA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PipelineKafkaSourceType) GetAllowedValues() []PipelineKafkaSourceType {
	return allowedPipelineKafkaSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PipelineKafkaSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PipelineKafkaSourceType(value)
	return nil
}

// NewPipelineKafkaSourceTypeFromValue returns a pointer to a valid PipelineKafkaSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPipelineKafkaSourceTypeFromValue(v string) (*PipelineKafkaSourceType, error) {
	ev := PipelineKafkaSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PipelineKafkaSourceType: valid values are %v", v, allowedPipelineKafkaSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PipelineKafkaSourceType) IsValid() bool {
	for _, existing := range allowedPipelineKafkaSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PipelineKafkaSourceType value.
func (v PipelineKafkaSourceType) Ptr() *PipelineKafkaSourceType {
	return &v
}
