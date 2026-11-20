// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMMetricResponseUniqueness The rule to count updatable events. Only present if `event_type` is `session` or `view`, and absent on older metrics created before this setting existed.
type RUMMetricResponseUniqueness struct {
	// When to count updatable events. `match` when the event is first seen, or `end` when the event is complete. Sessions and views are updated over time, and each one is counted once.
	When *RUMMetricUniquenessWhen `json:"when,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMMetricResponseUniqueness instantiates a new RUMMetricResponseUniqueness object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMMetricResponseUniqueness() *RUMMetricResponseUniqueness {
	this := RUMMetricResponseUniqueness{}
	return &this
}

// NewRUMMetricResponseUniquenessWithDefaults instantiates a new RUMMetricResponseUniqueness object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMMetricResponseUniquenessWithDefaults() *RUMMetricResponseUniqueness {
	this := RUMMetricResponseUniqueness{}
	return &this
}

// GetWhen returns the When field value if set, zero value otherwise.
func (o *RUMMetricResponseUniqueness) GetWhen() RUMMetricUniquenessWhen {
	if o == nil || o.When == nil {
		var ret RUMMetricUniquenessWhen
		return ret
	}
	return *o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMMetricResponseUniqueness) GetWhenOk() (*RUMMetricUniquenessWhen, bool) {
	if o == nil || o.When == nil {
		return nil, false
	}
	return o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *RUMMetricResponseUniqueness) HasWhen() bool {
	return o != nil && o.When != nil
}

// SetWhen gets a reference to the given RUMMetricUniquenessWhen and assigns it to the When field.
func (o *RUMMetricResponseUniqueness) SetWhen(v RUMMetricUniquenessWhen) {
	o.When = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMMetricResponseUniqueness) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.When != nil {
		toSerialize["when"] = o.When
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMMetricResponseUniqueness) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		When *RUMMetricUniquenessWhen `json:"when,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"when"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.When != nil && !all.When.IsValid() {
		hasInvalidField = true
	} else {
		o.When = all.When
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
