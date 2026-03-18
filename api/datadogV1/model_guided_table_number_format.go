// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableNumberFormat Number display format for a guided table column value.
type GuidedTableNumberFormat struct {
	// Number of digits after the decimal point. Use `*` for full precision, `integer` to show full integer values, or omit for automatic precision.
	Precision *GuidedTableNumberFormatPrecision `json:"precision,omitempty"`
	// Number format unit.
	Unit *NumberFormatUnit `json:"unit,omitempty"`
	// The definition of `NumberFormatUnitScale` object.
	UnitScale NullableNumberFormatUnitScale `json:"unit_scale,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGuidedTableNumberFormat instantiates a new GuidedTableNumberFormat object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGuidedTableNumberFormat() *GuidedTableNumberFormat {
	this := GuidedTableNumberFormat{}
	return &this
}

// NewGuidedTableNumberFormatWithDefaults instantiates a new GuidedTableNumberFormat object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGuidedTableNumberFormatWithDefaults() *GuidedTableNumberFormat {
	this := GuidedTableNumberFormat{}
	return &this
}

// GetPrecision returns the Precision field value if set, zero value otherwise.
func (o *GuidedTableNumberFormat) GetPrecision() GuidedTableNumberFormatPrecision {
	if o == nil || o.Precision == nil {
		var ret GuidedTableNumberFormatPrecision
		return ret
	}
	return *o.Precision
}

// GetPrecisionOk returns a tuple with the Precision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableNumberFormat) GetPrecisionOk() (*GuidedTableNumberFormatPrecision, bool) {
	if o == nil || o.Precision == nil {
		return nil, false
	}
	return o.Precision, true
}

// HasPrecision returns a boolean if a field has been set.
func (o *GuidedTableNumberFormat) HasPrecision() bool {
	return o != nil && o.Precision != nil
}

// SetPrecision gets a reference to the given GuidedTableNumberFormatPrecision and assigns it to the Precision field.
func (o *GuidedTableNumberFormat) SetPrecision(v GuidedTableNumberFormatPrecision) {
	o.Precision = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *GuidedTableNumberFormat) GetUnit() NumberFormatUnit {
	if o == nil || o.Unit == nil {
		var ret NumberFormatUnit
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableNumberFormat) GetUnitOk() (*NumberFormatUnit, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *GuidedTableNumberFormat) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given NumberFormatUnit and assigns it to the Unit field.
func (o *GuidedTableNumberFormat) SetUnit(v NumberFormatUnit) {
	o.Unit = &v
}

// GetUnitScale returns the UnitScale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidedTableNumberFormat) GetUnitScale() NumberFormatUnitScale {
	if o == nil || o.UnitScale.Get() == nil {
		var ret NumberFormatUnitScale
		return ret
	}
	return *o.UnitScale.Get()
}

// GetUnitScaleOk returns a tuple with the UnitScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GuidedTableNumberFormat) GetUnitScaleOk() (*NumberFormatUnitScale, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitScale.Get(), o.UnitScale.IsSet()
}

// HasUnitScale returns a boolean if a field has been set.
func (o *GuidedTableNumberFormat) HasUnitScale() bool {
	return o != nil && o.UnitScale.IsSet()
}

// SetUnitScale gets a reference to the given NullableNumberFormatUnitScale and assigns it to the UnitScale field.
func (o *GuidedTableNumberFormat) SetUnitScale(v NumberFormatUnitScale) {
	o.UnitScale.Set(&v)
}

// SetUnitScaleNil sets the value for UnitScale to be an explicit nil.
func (o *GuidedTableNumberFormat) SetUnitScaleNil() {
	o.UnitScale.Set(nil)
}

// UnsetUnitScale ensures that no value is present for UnitScale, not even an explicit nil.
func (o *GuidedTableNumberFormat) UnsetUnitScale() {
	o.UnitScale.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o GuidedTableNumberFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Precision != nil {
		toSerialize["precision"] = o.Precision
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.UnitScale.IsSet() {
		toSerialize["unit_scale"] = o.UnitScale.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GuidedTableNumberFormat) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Precision *GuidedTableNumberFormatPrecision `json:"precision,omitempty"`
		Unit      *NumberFormatUnit                 `json:"unit,omitempty"`
		UnitScale NullableNumberFormatUnitScale     `json:"unit_scale,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"precision", "unit", "unit_scale"})
	} else {
		return err
	}
	o.Precision = all.Precision
	o.Unit = all.Unit
	o.UnitScale = all.UnitScale

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
