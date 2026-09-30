// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetIntegrationSchemaSpecValueV2 A JSON Schema-like specification for a configuration value.
//
// Object-typed values always include a `properties` array, even when empty.
// Non-object-typed values never include `properties`. Throughout this schema,
// an empty array is meaningfully different from an absent field.
//
// Three further JSON Schema keywords can appear directly on this object but are
// not listed among its properties below to avoid clashing with this document's
// own schema composition keywords: `enum` (an array of allowed values, present
// only when there are enum constraints), `required` (an array of required
// property names, present only when `type` is `object`), and `oneOf` (an array
// of exclusive alternative value specifications this value can match, present
// only when there are alternatives).
type FleetIntegrationSchemaSpecValueV2 struct {
	// Whether, or which, additional properties are allowed on the object. Can be a boolean or a nested schema. Present only when `type` is `object`.
	AdditionalProperties interface{} `json:"additionalProperties,omitempty"`
	// Alternative value specifications this value can match. Absent when none apply.
	AnyOf []FleetIntegrationSchemaSpecValueV2 `json:"anyOf,omitempty"`
	// The default value. Can be any JSON type. Absent when not set.
	Default interface{} `json:"default,omitempty"`
	// A human-readable description of the value. Absent when not set.
	Description *string `json:"description,omitempty"`
	// A legacy, display-formatted representation of the default value. Can be any JSON type. Absent when not set.
	DisplayDefault interface{} `json:"display_default,omitempty"`
	// An example value. Can be any JSON type. Absent when not set.
	Example interface{} `json:"example,omitempty"`
	// The maximum allowed numeric value, exclusive. Absent when not set.
	ExclusiveMaximum *float64 `json:"exclusiveMaximum,omitempty"`
	// The minimum allowed numeric value, exclusive. Absent when not set.
	ExclusiveMinimum *float64 `json:"exclusiveMinimum,omitempty"`
	// A JSON Schema-like specification for a configuration value.
	//
	// Object-typed values always include a `properties` array, even when empty.
	// Non-object-typed values never include `properties`. Throughout this schema,
	// an empty array is meaningfully different from an absent field.
	//
	// Three further JSON Schema keywords can appear directly on this object but are
	// not listed among its properties below to avoid clashing with this document's
	// own schema composition keywords: `enum` (an array of allowed values, present
	// only when there are enum constraints), `required` (an array of required
	// property names, present only when `type` is `object`), and `oneOf` (an array
	// of exclusive alternative value specifications this value can match, present
	// only when there are alternatives).
	Items *FleetIntegrationSchemaSpecValueV2 `json:"items,omitempty"`
	// The maximum allowed string length. Absent when not set.
	MaxLength *int64 `json:"maxLength,omitempty"`
	// The maximum allowed numeric value, inclusive. Absent when not set.
	Maximum *float64 `json:"maximum,omitempty"`
	// The minimum allowed string length. Absent when not set.
	MinLength *int64 `json:"minLength,omitempty"`
	// The minimum allowed numeric value, inclusive. Absent when not set.
	Minimum *float64 `json:"minimum,omitempty"`
	// A regular expression the string value must match. Absent when not set.
	Pattern *string `json:"pattern,omitempty"`
	// The object's declared properties. Present when `type` is `object`, including as an empty array when the object declares no properties. Absent for non-object types.
	Properties []FleetIntegrationSchemaSpecPropertyV2 `json:"properties,omitempty"`
	// Whether the value is a secret that should be masked. Absent when not set.
	Secret *bool `json:"secret,omitempty"`
	// The JSON Schema type of the value, such as `string` or `object`. Absent when not set.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetIntegrationSchemaSpecValueV2 instantiates a new FleetIntegrationSchemaSpecValueV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetIntegrationSchemaSpecValueV2() *FleetIntegrationSchemaSpecValueV2 {
	this := FleetIntegrationSchemaSpecValueV2{}
	return &this
}

// NewFleetIntegrationSchemaSpecValueV2WithDefaults instantiates a new FleetIntegrationSchemaSpecValueV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetIntegrationSchemaSpecValueV2WithDefaults() *FleetIntegrationSchemaSpecValueV2 {
	this := FleetIntegrationSchemaSpecValueV2{}
	return &this
}

// GetAdditionalProperties returns the AdditionalProperties field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecValueV2) GetAdditionalProperties() interface{} {
	if o == nil || o.AdditionalProperties == nil {
		var ret interface{}
		return ret
	}
	return o.AdditionalProperties
}

// GetAdditionalPropertiesOk returns a tuple with the AdditionalProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecValueV2) GetAdditionalPropertiesOk() (*interface{}, bool) {
	if o == nil || o.AdditionalProperties == nil {
		return nil, false
	}
	return &o.AdditionalProperties, true
}

// HasAdditionalProperties returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecValueV2) HasAdditionalProperties() bool {
	return o != nil && o.AdditionalProperties != nil
}

// SetAdditionalProperties gets a reference to the given interface{} and assigns it to the AdditionalProperties field.
func (o *FleetIntegrationSchemaSpecValueV2) SetAdditionalProperties(v interface{}) {
	o.AdditionalProperties = v
}

// GetAnyOf returns the AnyOf field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecValueV2) GetAnyOf() []FleetIntegrationSchemaSpecValueV2 {
	if o == nil || o.AnyOf == nil {
		var ret []FleetIntegrationSchemaSpecValueV2
		return ret
	}
	return o.AnyOf
}

// GetAnyOfOk returns a tuple with the AnyOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecValueV2) GetAnyOfOk() (*[]FleetIntegrationSchemaSpecValueV2, bool) {
	if o == nil || o.AnyOf == nil {
		return nil, false
	}
	return &o.AnyOf, true
}

// HasAnyOf returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecValueV2) HasAnyOf() bool {
	return o != nil && o.AnyOf != nil
}

// SetAnyOf gets a reference to the given []FleetIntegrationSchemaSpecValueV2 and assigns it to the AnyOf field.
func (o *FleetIntegrationSchemaSpecValueV2) SetAnyOf(v []FleetIntegrationSchemaSpecValueV2) {
	o.AnyOf = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecValueV2) GetDefault() interface{} {
	if o == nil || o.Default == nil {
		var ret interface{}
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecValueV2) GetDefaultOk() (*interface{}, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return &o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecValueV2) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given interface{} and assigns it to the Default field.
func (o *FleetIntegrationSchemaSpecValueV2) SetDefault(v interface{}) {
	o.Default = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecValueV2) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecValueV2) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecValueV2) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FleetIntegrationSchemaSpecValueV2) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayDefault returns the DisplayDefault field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecValueV2) GetDisplayDefault() interface{} {
	if o == nil || o.DisplayDefault == nil {
		var ret interface{}
		return ret
	}
	return o.DisplayDefault
}

// GetDisplayDefaultOk returns a tuple with the DisplayDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecValueV2) GetDisplayDefaultOk() (*interface{}, bool) {
	if o == nil || o.DisplayDefault == nil {
		return nil, false
	}
	return &o.DisplayDefault, true
}

// HasDisplayDefault returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecValueV2) HasDisplayDefault() bool {
	return o != nil && o.DisplayDefault != nil
}

// SetDisplayDefault gets a reference to the given interface{} and assigns it to the DisplayDefault field.
func (o *FleetIntegrationSchemaSpecValueV2) SetDisplayDefault(v interface{}) {
	o.DisplayDefault = v
}

// GetExample returns the Example field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecValueV2) GetExample() interface{} {
	if o == nil || o.Example == nil {
		var ret interface{}
		return ret
	}
	return o.Example
}

// GetExampleOk returns a tuple with the Example field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecValueV2) GetExampleOk() (*interface{}, bool) {
	if o == nil || o.Example == nil {
		return nil, false
	}
	return &o.Example, true
}

// HasExample returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecValueV2) HasExample() bool {
	return o != nil && o.Example != nil
}

// SetExample gets a reference to the given interface{} and assigns it to the Example field.
func (o *FleetIntegrationSchemaSpecValueV2) SetExample(v interface{}) {
	o.Example = v
}

// GetExclusiveMaximum returns the ExclusiveMaximum field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecValueV2) GetExclusiveMaximum() float64 {
	if o == nil || o.ExclusiveMaximum == nil {
		var ret float64
		return ret
	}
	return *o.ExclusiveMaximum
}

// GetExclusiveMaximumOk returns a tuple with the ExclusiveMaximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecValueV2) GetExclusiveMaximumOk() (*float64, bool) {
	if o == nil || o.ExclusiveMaximum == nil {
		return nil, false
	}
	return o.ExclusiveMaximum, true
}

// HasExclusiveMaximum returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecValueV2) HasExclusiveMaximum() bool {
	return o != nil && o.ExclusiveMaximum != nil
}

// SetExclusiveMaximum gets a reference to the given float64 and assigns it to the ExclusiveMaximum field.
func (o *FleetIntegrationSchemaSpecValueV2) SetExclusiveMaximum(v float64) {
	o.ExclusiveMaximum = &v
}

// GetExclusiveMinimum returns the ExclusiveMinimum field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecValueV2) GetExclusiveMinimum() float64 {
	if o == nil || o.ExclusiveMinimum == nil {
		var ret float64
		return ret
	}
	return *o.ExclusiveMinimum
}

// GetExclusiveMinimumOk returns a tuple with the ExclusiveMinimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecValueV2) GetExclusiveMinimumOk() (*float64, bool) {
	if o == nil || o.ExclusiveMinimum == nil {
		return nil, false
	}
	return o.ExclusiveMinimum, true
}

// HasExclusiveMinimum returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecValueV2) HasExclusiveMinimum() bool {
	return o != nil && o.ExclusiveMinimum != nil
}

// SetExclusiveMinimum gets a reference to the given float64 and assigns it to the ExclusiveMinimum field.
func (o *FleetIntegrationSchemaSpecValueV2) SetExclusiveMinimum(v float64) {
	o.ExclusiveMinimum = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecValueV2) GetItems() FleetIntegrationSchemaSpecValueV2 {
	if o == nil || o.Items == nil {
		var ret FleetIntegrationSchemaSpecValueV2
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecValueV2) GetItemsOk() (*FleetIntegrationSchemaSpecValueV2, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecValueV2) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given FleetIntegrationSchemaSpecValueV2 and assigns it to the Items field.
func (o *FleetIntegrationSchemaSpecValueV2) SetItems(v FleetIntegrationSchemaSpecValueV2) {
	o.Items = &v
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecValueV2) GetMaxLength() int64 {
	if o == nil || o.MaxLength == nil {
		var ret int64
		return ret
	}
	return *o.MaxLength
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecValueV2) GetMaxLengthOk() (*int64, bool) {
	if o == nil || o.MaxLength == nil {
		return nil, false
	}
	return o.MaxLength, true
}

// HasMaxLength returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecValueV2) HasMaxLength() bool {
	return o != nil && o.MaxLength != nil
}

// SetMaxLength gets a reference to the given int64 and assigns it to the MaxLength field.
func (o *FleetIntegrationSchemaSpecValueV2) SetMaxLength(v int64) {
	o.MaxLength = &v
}

// GetMaximum returns the Maximum field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecValueV2) GetMaximum() float64 {
	if o == nil || o.Maximum == nil {
		var ret float64
		return ret
	}
	return *o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecValueV2) GetMaximumOk() (*float64, bool) {
	if o == nil || o.Maximum == nil {
		return nil, false
	}
	return o.Maximum, true
}

// HasMaximum returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecValueV2) HasMaximum() bool {
	return o != nil && o.Maximum != nil
}

// SetMaximum gets a reference to the given float64 and assigns it to the Maximum field.
func (o *FleetIntegrationSchemaSpecValueV2) SetMaximum(v float64) {
	o.Maximum = &v
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecValueV2) GetMinLength() int64 {
	if o == nil || o.MinLength == nil {
		var ret int64
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecValueV2) GetMinLengthOk() (*int64, bool) {
	if o == nil || o.MinLength == nil {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecValueV2) HasMinLength() bool {
	return o != nil && o.MinLength != nil
}

// SetMinLength gets a reference to the given int64 and assigns it to the MinLength field.
func (o *FleetIntegrationSchemaSpecValueV2) SetMinLength(v int64) {
	o.MinLength = &v
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecValueV2) GetMinimum() float64 {
	if o == nil || o.Minimum == nil {
		var ret float64
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecValueV2) GetMinimumOk() (*float64, bool) {
	if o == nil || o.Minimum == nil {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecValueV2) HasMinimum() bool {
	return o != nil && o.Minimum != nil
}

// SetMinimum gets a reference to the given float64 and assigns it to the Minimum field.
func (o *FleetIntegrationSchemaSpecValueV2) SetMinimum(v float64) {
	o.Minimum = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecValueV2) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecValueV2) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecValueV2) HasPattern() bool {
	return o != nil && o.Pattern != nil
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *FleetIntegrationSchemaSpecValueV2) SetPattern(v string) {
	o.Pattern = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecValueV2) GetProperties() []FleetIntegrationSchemaSpecPropertyV2 {
	if o == nil || o.Properties == nil {
		var ret []FleetIntegrationSchemaSpecPropertyV2
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecValueV2) GetPropertiesOk() (*[]FleetIntegrationSchemaSpecPropertyV2, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecValueV2) HasProperties() bool {
	return o != nil && o.Properties != nil
}

// SetProperties gets a reference to the given []FleetIntegrationSchemaSpecPropertyV2 and assigns it to the Properties field.
func (o *FleetIntegrationSchemaSpecValueV2) SetProperties(v []FleetIntegrationSchemaSpecPropertyV2) {
	o.Properties = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecValueV2) GetSecret() bool {
	if o == nil || o.Secret == nil {
		var ret bool
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecValueV2) GetSecretOk() (*bool, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecValueV2) HasSecret() bool {
	return o != nil && o.Secret != nil
}

// SetSecret gets a reference to the given bool and assigns it to the Secret field.
func (o *FleetIntegrationSchemaSpecValueV2) SetSecret(v bool) {
	o.Secret = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecValueV2) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecValueV2) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecValueV2) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FleetIntegrationSchemaSpecValueV2) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetIntegrationSchemaSpecValueV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AdditionalProperties != nil {
		toSerialize["additionalProperties"] = o.AdditionalProperties
	}
	if o.AnyOf != nil {
		toSerialize["anyOf"] = o.AnyOf
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayDefault != nil {
		toSerialize["display_default"] = o.DisplayDefault
	}
	if o.Example != nil {
		toSerialize["example"] = o.Example
	}
	if o.ExclusiveMaximum != nil {
		toSerialize["exclusiveMaximum"] = o.ExclusiveMaximum
	}
	if o.ExclusiveMinimum != nil {
		toSerialize["exclusiveMinimum"] = o.ExclusiveMinimum
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.MaxLength != nil {
		toSerialize["maxLength"] = o.MaxLength
	}
	if o.Maximum != nil {
		toSerialize["maximum"] = o.Maximum
	}
	if o.MinLength != nil {
		toSerialize["minLength"] = o.MinLength
	}
	if o.Minimum != nil {
		toSerialize["minimum"] = o.Minimum
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetIntegrationSchemaSpecValueV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AdditionalProperties interface{}                            `json:"additionalProperties,omitempty"`
		AnyOf                []FleetIntegrationSchemaSpecValueV2    `json:"anyOf,omitempty"`
		Default              interface{}                            `json:"default,omitempty"`
		Description          *string                                `json:"description,omitempty"`
		DisplayDefault       interface{}                            `json:"display_default,omitempty"`
		Example              interface{}                            `json:"example,omitempty"`
		ExclusiveMaximum     *float64                               `json:"exclusiveMaximum,omitempty"`
		ExclusiveMinimum     *float64                               `json:"exclusiveMinimum,omitempty"`
		Items                *FleetIntegrationSchemaSpecValueV2     `json:"items,omitempty"`
		MaxLength            *int64                                 `json:"maxLength,omitempty"`
		Maximum              *float64                               `json:"maximum,omitempty"`
		MinLength            *int64                                 `json:"minLength,omitempty"`
		Minimum              *float64                               `json:"minimum,omitempty"`
		Pattern              *string                                `json:"pattern,omitempty"`
		Properties           []FleetIntegrationSchemaSpecPropertyV2 `json:"properties,omitempty"`
		Secret               *bool                                  `json:"secret,omitempty"`
		Type                 *string                                `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"additionalProperties", "anyOf", "default", "description", "display_default", "example", "exclusiveMaximum", "exclusiveMinimum", "items", "maxLength", "maximum", "minLength", "minimum", "pattern", "properties", "secret", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AdditionalProperties = all.AdditionalProperties
	o.AnyOf = all.AnyOf
	o.Default = all.Default
	o.Description = all.Description
	o.DisplayDefault = all.DisplayDefault
	o.Example = all.Example
	o.ExclusiveMaximum = all.ExclusiveMaximum
	o.ExclusiveMinimum = all.ExclusiveMinimum
	if all.Items != nil && all.Items.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Items = all.Items
	o.MaxLength = all.MaxLength
	o.Maximum = all.Maximum
	o.MinLength = all.MinLength
	o.Minimum = all.Minimum
	o.Pattern = all.Pattern
	o.Properties = all.Properties
	o.Secret = all.Secret
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
