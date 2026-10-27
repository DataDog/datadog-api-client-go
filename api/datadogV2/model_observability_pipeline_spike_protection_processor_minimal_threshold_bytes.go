// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes Sets a specific byte threshold for the matching partition.
type ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes struct {
	// The minimal threshold, in bytes, for the matching partition.
	Bytes int64 `json:"bytes"`
	// The minimal threshold type. The value should always be `threshold`.
	Type ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes instantiates a new ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes(bytes int64, typeVar ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType) *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes {
	this := ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes{}
	this.Bytes = bytes
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesWithDefaults instantiates a new ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesWithDefaults() *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes {
	this := ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes{}
	var typeVar ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType = OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORMINIMALTHRESHOLDBYTESTYPE_THRESHOLD
	this.Type = typeVar
	return &this
}

// GetBytes returns the Bytes field value.
func (o *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes) GetBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Bytes
}

// GetBytesOk returns a tuple with the Bytes field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes) GetBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bytes, true
}

// SetBytes sets field value.
func (o *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes) SetBytes(v int64) {
	o.Bytes = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes) GetType() ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType {
	if o == nil {
		var ret ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes) GetTypeOk() (*ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes) SetType(v ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["bytes"] = o.Bytes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Bytes *int64                                                                  `json:"bytes"`
		Type  *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Bytes == nil {
		return fmt.Errorf("required field bytes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bytes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Bytes = *all.Bytes
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
