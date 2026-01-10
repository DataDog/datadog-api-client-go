// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineKafkaSasl Specifies the SASL mechanism for authenticating with a Kafka cluster.
type ObservabilityPipelineKafkaSasl struct {
	// SASL mechanism used for Kafka authentication.
	Mechanism *ObservabilityPipelineKafkaSaslMechanism `json:"mechanism,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineKafkaSasl instantiates a new ObservabilityPipelineKafkaSasl object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineKafkaSasl() *ObservabilityPipelineKafkaSasl {
	this := ObservabilityPipelineKafkaSasl{}
	return &this
}

// NewObservabilityPipelineKafkaSaslWithDefaults instantiates a new ObservabilityPipelineKafkaSasl object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineKafkaSaslWithDefaults() *ObservabilityPipelineKafkaSasl {
	this := ObservabilityPipelineKafkaSasl{}
	return &this
}

// GetMechanism returns the Mechanism field value if set, zero value otherwise.
func (o *ObservabilityPipelineKafkaSasl) GetMechanism() ObservabilityPipelineKafkaSaslMechanism {
	if o == nil || o.Mechanism == nil {
		var ret ObservabilityPipelineKafkaSaslMechanism
		return ret
	}
	return *o.Mechanism
}

// GetMechanismOk returns a tuple with the Mechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaSasl) GetMechanismOk() (*ObservabilityPipelineKafkaSaslMechanism, bool) {
	if o == nil || o.Mechanism == nil {
		return nil, false
	}
	return o.Mechanism, true
}

// HasMechanism returns a boolean if a field has been set.
func (o *ObservabilityPipelineKafkaSasl) HasMechanism() bool {
	return o != nil && o.Mechanism != nil
}

// SetMechanism gets a reference to the given ObservabilityPipelineKafkaSaslMechanism and assigns it to the Mechanism field.
func (o *ObservabilityPipelineKafkaSasl) SetMechanism(v ObservabilityPipelineKafkaSaslMechanism) {
	o.Mechanism = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineKafkaSasl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Mechanism != nil {
		toSerialize["mechanism"] = o.Mechanism
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineKafkaSasl) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mechanism *ObservabilityPipelineKafkaSaslMechanism `json:"mechanism,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"mechanism"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Mechanism != nil && !all.Mechanism.IsValid() {
		hasInvalidField = true
	} else {
		o.Mechanism = all.Mechanism
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
