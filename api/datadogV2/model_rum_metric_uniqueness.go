// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMMetricUniqueness The rule to count updatable events.
// Required when `event_type` is `session` or `view`, and must be omitted for other event types. It cannot be changed after creation.
type RUMMetricUniqueness struct {
	// When to count updatable events. `match` when the event is first seen, or `end` when the event is complete. Sessions and views are updated over time, and each one is counted once.
	When RUMMetricUniquenessWhen `json:"when"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMMetricUniqueness instantiates a new RUMMetricUniqueness object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMMetricUniqueness(when RUMMetricUniquenessWhen) *RUMMetricUniqueness {
	this := RUMMetricUniqueness{}
	this.When = when
	return &this
}

// NewRUMMetricUniquenessWithDefaults instantiates a new RUMMetricUniqueness object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMMetricUniquenessWithDefaults() *RUMMetricUniqueness {
	this := RUMMetricUniqueness{}
	return &this
}

// GetWhen returns the When field value.
func (o *RUMMetricUniqueness) GetWhen() RUMMetricUniquenessWhen {
	if o == nil {
		var ret RUMMetricUniquenessWhen
		return ret
	}
	return o.When
}

// GetWhenOk returns a tuple with the When field value
// and a boolean to check if the value has been set.
func (o *RUMMetricUniqueness) GetWhenOk() (*RUMMetricUniquenessWhen, bool) {
	if o == nil {
		return nil, false
	}
	return &o.When, true
}

// SetWhen sets field value.
func (o *RUMMetricUniqueness) SetWhen(v RUMMetricUniquenessWhen) {
	o.When = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMMetricUniqueness) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["when"] = o.When

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMMetricUniqueness) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		When *RUMMetricUniquenessWhen `json:"when"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.When == nil {
		return fmt.Errorf("required field when missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"when"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.When.IsValid() {
		hasInvalidField = true
	} else {
		o.When = *all.When
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
