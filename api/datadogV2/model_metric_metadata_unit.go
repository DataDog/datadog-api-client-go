// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricMetadataUnit A unit of measure from the Datadog unit catalog, or `null` when no unit is set.
type MetricMetadataUnit struct {
	// Unit family, such as `bytes`, `time`, or `percentage`.
	Family string `json:"family"`
	// Full name of the unit, such as `byte`.
	Name string `json:"name"`
	// Plural form of the unit name.
	Plural string `json:"plural"`
	// Factor that converts a value in this unit to the base unit of its family.
	ScaleFactor float64 `json:"scale_factor"`
	// Abbreviated unit name, or `null` if the unit has none.
	ShortName datadog.NullableString `json:"short_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricMetadataUnit instantiates a new MetricMetadataUnit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricMetadataUnit(family string, name string, plural string, scaleFactor float64, shortName datadog.NullableString) *MetricMetadataUnit {
	this := MetricMetadataUnit{}
	this.Family = family
	this.Name = name
	this.Plural = plural
	this.ScaleFactor = scaleFactor
	this.ShortName = shortName
	return &this
}

// NewMetricMetadataUnitWithDefaults instantiates a new MetricMetadataUnit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricMetadataUnitWithDefaults() *MetricMetadataUnit {
	this := MetricMetadataUnit{}
	return &this
}

// GetFamily returns the Family field value.
func (o *MetricMetadataUnit) GetFamily() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *MetricMetadataUnit) GetFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value.
func (o *MetricMetadataUnit) SetFamily(v string) {
	o.Family = v
}

// GetName returns the Name field value.
func (o *MetricMetadataUnit) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetricMetadataUnit) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MetricMetadataUnit) SetName(v string) {
	o.Name = v
}

// GetPlural returns the Plural field value.
func (o *MetricMetadataUnit) GetPlural() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Plural
}

// GetPluralOk returns a tuple with the Plural field value
// and a boolean to check if the value has been set.
func (o *MetricMetadataUnit) GetPluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plural, true
}

// SetPlural sets field value.
func (o *MetricMetadataUnit) SetPlural(v string) {
	o.Plural = v
}

// GetScaleFactor returns the ScaleFactor field value.
func (o *MetricMetadataUnit) GetScaleFactor() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.ScaleFactor
}

// GetScaleFactorOk returns a tuple with the ScaleFactor field value
// and a boolean to check if the value has been set.
func (o *MetricMetadataUnit) GetScaleFactorOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScaleFactor, true
}

// SetScaleFactor sets field value.
func (o *MetricMetadataUnit) SetScaleFactor(v float64) {
	o.ScaleFactor = v
}

// GetShortName returns the ShortName field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *MetricMetadataUnit) GetShortName() string {
	if o == nil || o.ShortName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ShortName.Get()
}

// GetShortNameOk returns a tuple with the ShortName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadataUnit) GetShortNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortName.Get(), o.ShortName.IsSet()
}

// SetShortName sets field value.
func (o *MetricMetadataUnit) SetShortName(v string) {
	o.ShortName.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricMetadataUnit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["family"] = o.Family
	toSerialize["name"] = o.Name
	toSerialize["plural"] = o.Plural
	toSerialize["scale_factor"] = o.ScaleFactor
	toSerialize["short_name"] = o.ShortName.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricMetadataUnit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Family      *string                `json:"family"`
		Name        *string                `json:"name"`
		Plural      *string                `json:"plural"`
		ScaleFactor *float64               `json:"scale_factor"`
		ShortName   datadog.NullableString `json:"short_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Family == nil {
		return fmt.Errorf("required field family missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Plural == nil {
		return fmt.Errorf("required field plural missing")
	}
	if all.ScaleFactor == nil {
		return fmt.Errorf("required field scale_factor missing")
	}
	if !all.ShortName.IsSet() {
		return fmt.Errorf("required field short_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"family", "name", "plural", "scale_factor", "short_name"})
	} else {
		return err
	}
	o.Family = *all.Family
	o.Name = *all.Name
	o.Plural = *all.Plural
	o.ScaleFactor = *all.ScaleFactor
	o.ShortName = all.ShortName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableMetricMetadataUnit handles when a null is used for MetricMetadataUnit.
type NullableMetricMetadataUnit struct {
	value *MetricMetadataUnit
	isSet bool
}

// Get returns the associated value.
func (v NullableMetricMetadataUnit) Get() *MetricMetadataUnit {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableMetricMetadataUnit) Set(val *MetricMetadataUnit) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableMetricMetadataUnit) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableMetricMetadataUnit) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableMetricMetadataUnit initializes the struct as if Set has been called.
func NewNullableMetricMetadataUnit(val *MetricMetadataUnit) *NullableMetricMetadataUnit {
	return &NullableMetricMetadataUnit{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableMetricMetadataUnit) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableMetricMetadataUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
