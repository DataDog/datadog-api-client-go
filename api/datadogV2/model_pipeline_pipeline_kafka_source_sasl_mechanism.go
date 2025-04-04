// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelinePipelineKafkaSourceSaslMechanism SASL mechanism used for Kafka authentication.
type PipelinePipelineKafkaSourceSaslMechanism string

// List of PipelinePipelineKafkaSourceSaslMechanism.
const (
	PIPELINEPIPELINEKAFKASOURCESASLMECHANISM_PLAIN               PipelinePipelineKafkaSourceSaslMechanism = "PLAIN"
	PIPELINEPIPELINEKAFKASOURCESASLMECHANISM_SCRAMNOT_SHANOT_256 PipelinePipelineKafkaSourceSaslMechanism = "SCRAM-SHA-256"
	PIPELINEPIPELINEKAFKASOURCESASLMECHANISM_SCRAMNOT_SHANOT_512 PipelinePipelineKafkaSourceSaslMechanism = "SCRAM-SHA-512"
)

var allowedPipelinePipelineKafkaSourceSaslMechanismEnumValues = []PipelinePipelineKafkaSourceSaslMechanism{
	PIPELINEPIPELINEKAFKASOURCESASLMECHANISM_PLAIN,
	PIPELINEPIPELINEKAFKASOURCESASLMECHANISM_SCRAMNOT_SHANOT_256,
	PIPELINEPIPELINEKAFKASOURCESASLMECHANISM_SCRAMNOT_SHANOT_512,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PipelinePipelineKafkaSourceSaslMechanism) GetAllowedValues() []PipelinePipelineKafkaSourceSaslMechanism {
	return allowedPipelinePipelineKafkaSourceSaslMechanismEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PipelinePipelineKafkaSourceSaslMechanism) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PipelinePipelineKafkaSourceSaslMechanism(value)
	return nil
}

// NewPipelinePipelineKafkaSourceSaslMechanismFromValue returns a pointer to a valid PipelinePipelineKafkaSourceSaslMechanism
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPipelinePipelineKafkaSourceSaslMechanismFromValue(v string) (*PipelinePipelineKafkaSourceSaslMechanism, error) {
	ev := PipelinePipelineKafkaSourceSaslMechanism(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PipelinePipelineKafkaSourceSaslMechanism: valid values are %v", v, allowedPipelinePipelineKafkaSourceSaslMechanismEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PipelinePipelineKafkaSourceSaslMechanism) IsValid() bool {
	for _, existing := range allowedPipelinePipelineKafkaSourceSaslMechanismEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PipelinePipelineKafkaSourceSaslMechanism value.
func (v PipelinePipelineKafkaSourceSaslMechanism) Ptr() *PipelinePipelineKafkaSourceSaslMechanism {
	return &v
}
