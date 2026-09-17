// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScaRequestDataAttributesScanStartTimestamp The time when the SCA scan started.
type ScaRequestDataAttributesScanStartTimestamp struct {
	// Non-negative fractions of a second at nanosecond resolution.
	Nanos *int32 `json:"nanos,omitempty"`
	// Seconds of UTC time since Unix epoch.
	Seconds *int64 `json:"seconds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewScaRequestDataAttributesScanStartTimestamp instantiates a new ScaRequestDataAttributesScanStartTimestamp object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScaRequestDataAttributesScanStartTimestamp() *ScaRequestDataAttributesScanStartTimestamp {
	this := ScaRequestDataAttributesScanStartTimestamp{}
	return &this
}

// NewScaRequestDataAttributesScanStartTimestampWithDefaults instantiates a new ScaRequestDataAttributesScanStartTimestamp object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScaRequestDataAttributesScanStartTimestampWithDefaults() *ScaRequestDataAttributesScanStartTimestamp {
	this := ScaRequestDataAttributesScanStartTimestamp{}
	return &this
}

// GetNanos returns the Nanos field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesScanStartTimestamp) GetNanos() int32 {
	if o == nil || o.Nanos == nil {
		var ret int32
		return ret
	}
	return *o.Nanos
}

// GetNanosOk returns a tuple with the Nanos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesScanStartTimestamp) GetNanosOk() (*int32, bool) {
	if o == nil || o.Nanos == nil {
		return nil, false
	}
	return o.Nanos, true
}

// HasNanos returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesScanStartTimestamp) HasNanos() bool {
	return o != nil && o.Nanos != nil
}

// SetNanos gets a reference to the given int32 and assigns it to the Nanos field.
func (o *ScaRequestDataAttributesScanStartTimestamp) SetNanos(v int32) {
	o.Nanos = &v
}

// GetSeconds returns the Seconds field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesScanStartTimestamp) GetSeconds() int64 {
	if o == nil || o.Seconds == nil {
		var ret int64
		return ret
	}
	return *o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesScanStartTimestamp) GetSecondsOk() (*int64, bool) {
	if o == nil || o.Seconds == nil {
		return nil, false
	}
	return o.Seconds, true
}

// HasSeconds returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesScanStartTimestamp) HasSeconds() bool {
	return o != nil && o.Seconds != nil
}

// SetSeconds gets a reference to the given int64 and assigns it to the Seconds field.
func (o *ScaRequestDataAttributesScanStartTimestamp) SetSeconds(v int64) {
	o.Seconds = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScaRequestDataAttributesScanStartTimestamp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Nanos != nil {
		toSerialize["nanos"] = o.Nanos
	}
	if o.Seconds != nil {
		toSerialize["seconds"] = o.Seconds
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScaRequestDataAttributesScanStartTimestamp) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Nanos   *int32 `json:"nanos,omitempty"`
		Seconds *int64 `json:"seconds,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Nanos = all.Nanos
	o.Seconds = all.Seconds

	return nil
}

// NullableScaRequestDataAttributesScanStartTimestamp handles when a null is used for ScaRequestDataAttributesScanStartTimestamp.
type NullableScaRequestDataAttributesScanStartTimestamp struct {
	value *ScaRequestDataAttributesScanStartTimestamp
	isSet bool
}

// Get returns the associated value.
func (v NullableScaRequestDataAttributesScanStartTimestamp) Get() *ScaRequestDataAttributesScanStartTimestamp {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableScaRequestDataAttributesScanStartTimestamp) Set(val *ScaRequestDataAttributesScanStartTimestamp) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableScaRequestDataAttributesScanStartTimestamp) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableScaRequestDataAttributesScanStartTimestamp) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableScaRequestDataAttributesScanStartTimestamp initializes the struct as if Set has been called.
func NewNullableScaRequestDataAttributesScanStartTimestamp(val *ScaRequestDataAttributesScanStartTimestamp) *NullableScaRequestDataAttributesScanStartTimestamp {
	return &NullableScaRequestDataAttributesScanStartTimestamp{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableScaRequestDataAttributesScanStartTimestamp) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableScaRequestDataAttributesScanStartTimestamp) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
