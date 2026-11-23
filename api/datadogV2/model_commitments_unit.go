// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsUnit Unit metadata for a numeric metric. Unset fields are omitted, and the unit is `null` when it is unknown.
type CommitmentsUnit struct {
	// The unit family (for example, percentage or money).
	Family *string `json:"family,omitempty"`
	// The unit identifier.
	Id *int64 `json:"id,omitempty"`
	// The unit name (for example, percent or dollar).
	Name *string `json:"name,omitempty"`
	// The plural form of the unit name.
	Plural *string `json:"plural,omitempty"`
	// The scale factor for the unit.
	ScaleFactor *float64 `json:"scale_factor,omitempty"`
	// The abbreviated unit name (for example, % or $).
	ShortName *string `json:"short_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsUnit instantiates a new CommitmentsUnit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsUnit() *CommitmentsUnit {
	this := CommitmentsUnit{}
	return &this
}

// NewCommitmentsUnitWithDefaults instantiates a new CommitmentsUnit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsUnitWithDefaults() *CommitmentsUnit {
	this := CommitmentsUnit{}
	return &this
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *CommitmentsUnit) GetFamily() string {
	if o == nil || o.Family == nil {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsUnit) GetFamilyOk() (*string, bool) {
	if o == nil || o.Family == nil {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *CommitmentsUnit) HasFamily() bool {
	return o != nil && o.Family != nil
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *CommitmentsUnit) SetFamily(v string) {
	o.Family = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CommitmentsUnit) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsUnit) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CommitmentsUnit) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *CommitmentsUnit) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CommitmentsUnit) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsUnit) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CommitmentsUnit) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CommitmentsUnit) SetName(v string) {
	o.Name = &v
}

// GetPlural returns the Plural field value if set, zero value otherwise.
func (o *CommitmentsUnit) GetPlural() string {
	if o == nil || o.Plural == nil {
		var ret string
		return ret
	}
	return *o.Plural
}

// GetPluralOk returns a tuple with the Plural field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsUnit) GetPluralOk() (*string, bool) {
	if o == nil || o.Plural == nil {
		return nil, false
	}
	return o.Plural, true
}

// HasPlural returns a boolean if a field has been set.
func (o *CommitmentsUnit) HasPlural() bool {
	return o != nil && o.Plural != nil
}

// SetPlural gets a reference to the given string and assigns it to the Plural field.
func (o *CommitmentsUnit) SetPlural(v string) {
	o.Plural = &v
}

// GetScaleFactor returns the ScaleFactor field value if set, zero value otherwise.
func (o *CommitmentsUnit) GetScaleFactor() float64 {
	if o == nil || o.ScaleFactor == nil {
		var ret float64
		return ret
	}
	return *o.ScaleFactor
}

// GetScaleFactorOk returns a tuple with the ScaleFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsUnit) GetScaleFactorOk() (*float64, bool) {
	if o == nil || o.ScaleFactor == nil {
		return nil, false
	}
	return o.ScaleFactor, true
}

// HasScaleFactor returns a boolean if a field has been set.
func (o *CommitmentsUnit) HasScaleFactor() bool {
	return o != nil && o.ScaleFactor != nil
}

// SetScaleFactor gets a reference to the given float64 and assigns it to the ScaleFactor field.
func (o *CommitmentsUnit) SetScaleFactor(v float64) {
	o.ScaleFactor = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *CommitmentsUnit) GetShortName() string {
	if o == nil || o.ShortName == nil {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsUnit) GetShortNameOk() (*string, bool) {
	if o == nil || o.ShortName == nil {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *CommitmentsUnit) HasShortName() bool {
	return o != nil && o.ShortName != nil
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *CommitmentsUnit) SetShortName(v string) {
	o.ShortName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsUnit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Family != nil {
		toSerialize["family"] = o.Family
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Plural != nil {
		toSerialize["plural"] = o.Plural
	}
	if o.ScaleFactor != nil {
		toSerialize["scale_factor"] = o.ScaleFactor
	}
	if o.ShortName != nil {
		toSerialize["short_name"] = o.ShortName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsUnit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Family      *string  `json:"family,omitempty"`
		Id          *int64   `json:"id,omitempty"`
		Name        *string  `json:"name,omitempty"`
		Plural      *string  `json:"plural,omitempty"`
		ScaleFactor *float64 `json:"scale_factor,omitempty"`
		ShortName   *string  `json:"short_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"family", "id", "name", "plural", "scale_factor", "short_name"})
	} else {
		return err
	}
	o.Family = all.Family
	o.Id = all.Id
	o.Name = all.Name
	o.Plural = all.Plural
	o.ScaleFactor = all.ScaleFactor
	o.ShortName = all.ShortName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableCommitmentsUnit handles when a null is used for CommitmentsUnit.
type NullableCommitmentsUnit struct {
	value *CommitmentsUnit
	isSet bool
}

// Get returns the associated value.
func (v NullableCommitmentsUnit) Get() *CommitmentsUnit {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableCommitmentsUnit) Set(val *CommitmentsUnit) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableCommitmentsUnit) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableCommitmentsUnit) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCommitmentsUnit initializes the struct as if Set has been called.
func NewNullableCommitmentsUnit(val *CommitmentsUnit) *NullableCommitmentsUnit {
	return &NullableCommitmentsUnit{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableCommitmentsUnit) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableCommitmentsUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
