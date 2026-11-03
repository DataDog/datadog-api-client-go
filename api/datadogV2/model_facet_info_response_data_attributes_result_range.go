// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FacetInfoResponseDataAttributesResultRange The numeric range of a facet attribute, representing the minimum and maximum observed values.
type FacetInfoResponseDataAttributesResultRange struct {
	// The maximum observed value for the numeric facet attribute. Null when no records match.
	Max datadog.NullableFloat64 `json:"max,omitempty"`
	// The minimum observed value for the numeric facet attribute. Null when no records match.
	Min datadog.NullableFloat64 `json:"min,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFacetInfoResponseDataAttributesResultRange instantiates a new FacetInfoResponseDataAttributesResultRange object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFacetInfoResponseDataAttributesResultRange() *FacetInfoResponseDataAttributesResultRange {
	this := FacetInfoResponseDataAttributesResultRange{}
	return &this
}

// NewFacetInfoResponseDataAttributesResultRangeWithDefaults instantiates a new FacetInfoResponseDataAttributesResultRange object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFacetInfoResponseDataAttributesResultRangeWithDefaults() *FacetInfoResponseDataAttributesResultRange {
	this := FacetInfoResponseDataAttributesResultRange{}
	return &this
}

// GetMax returns the Max field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FacetInfoResponseDataAttributesResultRange) GetMax() float64 {
	if o == nil || o.Max.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Max.Get()
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FacetInfoResponseDataAttributesResultRange) GetMaxOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Max.Get(), o.Max.IsSet()
}

// HasMax returns a boolean if a field has been set.
func (o *FacetInfoResponseDataAttributesResultRange) HasMax() bool {
	return o != nil && o.Max.IsSet()
}

// SetMax gets a reference to the given datadog.NullableFloat64 and assigns it to the Max field.
func (o *FacetInfoResponseDataAttributesResultRange) SetMax(v float64) {
	o.Max.Set(&v)
}

// SetMaxNil sets the value for Max to be an explicit nil.
func (o *FacetInfoResponseDataAttributesResultRange) SetMaxNil() {
	o.Max.Set(nil)
}

// UnsetMax ensures that no value is present for Max, not even an explicit nil.
func (o *FacetInfoResponseDataAttributesResultRange) UnsetMax() {
	o.Max.Unset()
}

// GetMin returns the Min field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FacetInfoResponseDataAttributesResultRange) GetMin() float64 {
	if o == nil || o.Min.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Min.Get()
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FacetInfoResponseDataAttributesResultRange) GetMinOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Min.Get(), o.Min.IsSet()
}

// HasMin returns a boolean if a field has been set.
func (o *FacetInfoResponseDataAttributesResultRange) HasMin() bool {
	return o != nil && o.Min.IsSet()
}

// SetMin gets a reference to the given datadog.NullableFloat64 and assigns it to the Min field.
func (o *FacetInfoResponseDataAttributesResultRange) SetMin(v float64) {
	o.Min.Set(&v)
}

// SetMinNil sets the value for Min to be an explicit nil.
func (o *FacetInfoResponseDataAttributesResultRange) SetMinNil() {
	o.Min.Set(nil)
}

// UnsetMin ensures that no value is present for Min, not even an explicit nil.
func (o *FacetInfoResponseDataAttributesResultRange) UnsetMin() {
	o.Min.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o FacetInfoResponseDataAttributesResultRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Max.IsSet() {
		toSerialize["max"] = o.Max.Get()
	}
	if o.Min.IsSet() {
		toSerialize["min"] = o.Min.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FacetInfoResponseDataAttributesResultRange) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Max datadog.NullableFloat64 `json:"max,omitempty"`
		Min datadog.NullableFloat64 `json:"min,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"max", "min"})
	} else {
		return err
	}
	o.Max = all.Max
	o.Min = all.Min

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
