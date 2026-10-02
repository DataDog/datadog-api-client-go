// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetIntegrationSchemaSpecPropertyV2 A property of an object-typed configuration value. A `oneOf` keyword (an array of exclusive alternative value specifications this property can match) can appear directly on this object when alternatives apply.
type FleetIntegrationSchemaSpecPropertyV2 struct {
	// Whether, or which, additional properties are allowed on the object. Can be a boolean or a nested schema. Present only when `type` is `object`.
	AdditionalProperties interface{} `json:"additionalProperties,omitempty"`
	// Alternative value specifications this property can match. Absent when none apply.
	AnyOf []FleetIntegrationSchemaSpecValueV2 `json:"anyOf,omitempty"`
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
	// The property name.
	Name string `json:"name"`
	// Nested properties. Present only when `type` is `object` and the object declares properties.
	Properties []FleetIntegrationSchemaSpecPropertyV2 `json:"properties,omitempty"`
	// The JSON Schema type of the property, such as `string` or `boolean`. Absent when not set.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetIntegrationSchemaSpecPropertyV2 instantiates a new FleetIntegrationSchemaSpecPropertyV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetIntegrationSchemaSpecPropertyV2(name string) *FleetIntegrationSchemaSpecPropertyV2 {
	this := FleetIntegrationSchemaSpecPropertyV2{}
	this.Name = name
	return &this
}

// NewFleetIntegrationSchemaSpecPropertyV2WithDefaults instantiates a new FleetIntegrationSchemaSpecPropertyV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetIntegrationSchemaSpecPropertyV2WithDefaults() *FleetIntegrationSchemaSpecPropertyV2 {
	this := FleetIntegrationSchemaSpecPropertyV2{}
	return &this
}

// GetAdditionalProperties returns the AdditionalProperties field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecPropertyV2) GetAdditionalProperties() interface{} {
	if o == nil || o.AdditionalProperties == nil {
		var ret interface{}
		return ret
	}
	return o.AdditionalProperties
}

// GetAdditionalPropertiesOk returns a tuple with the AdditionalProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecPropertyV2) GetAdditionalPropertiesOk() (*interface{}, bool) {
	if o == nil || o.AdditionalProperties == nil {
		return nil, false
	}
	return &o.AdditionalProperties, true
}

// HasAdditionalProperties returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecPropertyV2) HasAdditionalProperties() bool {
	return o != nil && o.AdditionalProperties != nil
}

// SetAdditionalProperties gets a reference to the given interface{} and assigns it to the AdditionalProperties field.
func (o *FleetIntegrationSchemaSpecPropertyV2) SetAdditionalProperties(v interface{}) {
	o.AdditionalProperties = v
}

// GetAnyOf returns the AnyOf field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecPropertyV2) GetAnyOf() []FleetIntegrationSchemaSpecValueV2 {
	if o == nil || o.AnyOf == nil {
		var ret []FleetIntegrationSchemaSpecValueV2
		return ret
	}
	return o.AnyOf
}

// GetAnyOfOk returns a tuple with the AnyOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecPropertyV2) GetAnyOfOk() (*[]FleetIntegrationSchemaSpecValueV2, bool) {
	if o == nil || o.AnyOf == nil {
		return nil, false
	}
	return &o.AnyOf, true
}

// HasAnyOf returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecPropertyV2) HasAnyOf() bool {
	return o != nil && o.AnyOf != nil
}

// SetAnyOf gets a reference to the given []FleetIntegrationSchemaSpecValueV2 and assigns it to the AnyOf field.
func (o *FleetIntegrationSchemaSpecPropertyV2) SetAnyOf(v []FleetIntegrationSchemaSpecValueV2) {
	o.AnyOf = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecPropertyV2) GetItems() FleetIntegrationSchemaSpecValueV2 {
	if o == nil || o.Items == nil {
		var ret FleetIntegrationSchemaSpecValueV2
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecPropertyV2) GetItemsOk() (*FleetIntegrationSchemaSpecValueV2, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecPropertyV2) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given FleetIntegrationSchemaSpecValueV2 and assigns it to the Items field.
func (o *FleetIntegrationSchemaSpecPropertyV2) SetItems(v FleetIntegrationSchemaSpecValueV2) {
	o.Items = &v
}

// GetName returns the Name field value.
func (o *FleetIntegrationSchemaSpecPropertyV2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecPropertyV2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FleetIntegrationSchemaSpecPropertyV2) SetName(v string) {
	o.Name = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecPropertyV2) GetProperties() []FleetIntegrationSchemaSpecPropertyV2 {
	if o == nil || o.Properties == nil {
		var ret []FleetIntegrationSchemaSpecPropertyV2
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecPropertyV2) GetPropertiesOk() (*[]FleetIntegrationSchemaSpecPropertyV2, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecPropertyV2) HasProperties() bool {
	return o != nil && o.Properties != nil
}

// SetProperties gets a reference to the given []FleetIntegrationSchemaSpecPropertyV2 and assigns it to the Properties field.
func (o *FleetIntegrationSchemaSpecPropertyV2) SetProperties(v []FleetIntegrationSchemaSpecPropertyV2) {
	o.Properties = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecPropertyV2) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecPropertyV2) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecPropertyV2) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FleetIntegrationSchemaSpecPropertyV2) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetIntegrationSchemaSpecPropertyV2) MarshalJSON() ([]byte, error) {
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
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	toSerialize["name"] = o.Name
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
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
func (o *FleetIntegrationSchemaSpecPropertyV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AdditionalProperties interface{}                            `json:"additionalProperties,omitempty"`
		AnyOf                []FleetIntegrationSchemaSpecValueV2    `json:"anyOf,omitempty"`
		Items                *FleetIntegrationSchemaSpecValueV2     `json:"items,omitempty"`
		Name                 *string                                `json:"name"`
		Properties           []FleetIntegrationSchemaSpecPropertyV2 `json:"properties,omitempty"`
		Type                 *string                                `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"additionalProperties", "anyOf", "items", "name", "properties", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AdditionalProperties = all.AdditionalProperties
	o.AnyOf = all.AnyOf
	if all.Items != nil && all.Items.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Items = all.Items
	o.Name = *all.Name
	o.Properties = all.Properties
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
