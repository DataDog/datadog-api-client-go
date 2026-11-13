// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditObservationCountsV2 How many observations the latest saved refresh saw, whatever the range. Null when its counts can't be read.
type ExperimentsBanditObservationCountsV2 struct {
	// Observations the refresh couldn't use.
	Invalid int64 `json:"invalid"`
	// Every observation. The sum of the other three.
	Total int64 `json:"total"`
	// Valid observations whose aggregation window is still open.
	ValidImmature int64 `json:"valid_immature"`
	// Valid observations whose aggregation window has closed. The estimates use these.
	ValidMature int64 `json:"valid_mature"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditObservationCountsV2 instantiates a new ExperimentsBanditObservationCountsV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditObservationCountsV2(invalid int64, total int64, validImmature int64, validMature int64) *ExperimentsBanditObservationCountsV2 {
	this := ExperimentsBanditObservationCountsV2{}
	this.Invalid = invalid
	this.Total = total
	this.ValidImmature = validImmature
	this.ValidMature = validMature
	return &this
}

// NewExperimentsBanditObservationCountsV2WithDefaults instantiates a new ExperimentsBanditObservationCountsV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditObservationCountsV2WithDefaults() *ExperimentsBanditObservationCountsV2 {
	this := ExperimentsBanditObservationCountsV2{}
	return &this
}

// GetInvalid returns the Invalid field value.
func (o *ExperimentsBanditObservationCountsV2) GetInvalid() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Invalid
}

// GetInvalidOk returns a tuple with the Invalid field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditObservationCountsV2) GetInvalidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Invalid, true
}

// SetInvalid sets field value.
func (o *ExperimentsBanditObservationCountsV2) SetInvalid(v int64) {
	o.Invalid = v
}

// GetTotal returns the Total field value.
func (o *ExperimentsBanditObservationCountsV2) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditObservationCountsV2) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *ExperimentsBanditObservationCountsV2) SetTotal(v int64) {
	o.Total = v
}

// GetValidImmature returns the ValidImmature field value.
func (o *ExperimentsBanditObservationCountsV2) GetValidImmature() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ValidImmature
}

// GetValidImmatureOk returns a tuple with the ValidImmature field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditObservationCountsV2) GetValidImmatureOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidImmature, true
}

// SetValidImmature sets field value.
func (o *ExperimentsBanditObservationCountsV2) SetValidImmature(v int64) {
	o.ValidImmature = v
}

// GetValidMature returns the ValidMature field value.
func (o *ExperimentsBanditObservationCountsV2) GetValidMature() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ValidMature
}

// GetValidMatureOk returns a tuple with the ValidMature field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditObservationCountsV2) GetValidMatureOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidMature, true
}

// SetValidMature sets field value.
func (o *ExperimentsBanditObservationCountsV2) SetValidMature(v int64) {
	o.ValidMature = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditObservationCountsV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["invalid"] = o.Invalid
	toSerialize["total"] = o.Total
	toSerialize["valid_immature"] = o.ValidImmature
	toSerialize["valid_mature"] = o.ValidMature

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditObservationCountsV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Invalid       *int64 `json:"invalid"`
		Total         *int64 `json:"total"`
		ValidImmature *int64 `json:"valid_immature"`
		ValidMature   *int64 `json:"valid_mature"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Invalid == nil {
		return fmt.Errorf("required field invalid missing")
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	if all.ValidImmature == nil {
		return fmt.Errorf("required field valid_immature missing")
	}
	if all.ValidMature == nil {
		return fmt.Errorf("required field valid_mature missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"invalid", "total", "valid_immature", "valid_mature"})
	} else {
		return err
	}
	o.Invalid = *all.Invalid
	o.Total = *all.Total
	o.ValidImmature = *all.ValidImmature
	o.ValidMature = *all.ValidMature

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableExperimentsBanditObservationCountsV2 handles when a null is used for ExperimentsBanditObservationCountsV2.
type NullableExperimentsBanditObservationCountsV2 struct {
	value *ExperimentsBanditObservationCountsV2
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsBanditObservationCountsV2) Get() *ExperimentsBanditObservationCountsV2 {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsBanditObservationCountsV2) Set(val *ExperimentsBanditObservationCountsV2) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsBanditObservationCountsV2) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableExperimentsBanditObservationCountsV2) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsBanditObservationCountsV2 initializes the struct as if Set has been called.
func NewNullableExperimentsBanditObservationCountsV2(val *ExperimentsBanditObservationCountsV2) *NullableExperimentsBanditObservationCountsV2 {
	return &NullableExperimentsBanditObservationCountsV2{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsBanditObservationCountsV2) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsBanditObservationCountsV2) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
