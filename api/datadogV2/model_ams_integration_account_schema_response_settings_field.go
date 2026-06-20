// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AmsIntegrationAccountSchemaResponseSettingsField JSON Schema definition for a single field within settings or secrets.
// The exact fields vary by integration.
type AmsIntegrationAccountSchemaResponseSettingsField struct {
	// Whether additional properties are allowed for this field.
	AdditionalProperties *bool `json:"additionalProperties,omitempty"`
	// Default value for the field if not provided.
	Default interface{} `json:"default,omitempty"`
	// Human-readable description of the field's purpose.
	Description *string `json:"description,omitempty"`
	// Schema for array items when type is "array".
	Items interface{} `json:"items,omitempty"`
	// Minimum length for string fields.
	MinLength *int64 `json:"minLength,omitempty"`
	// The data type of the field (string, boolean, integer, array, object).
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAmsIntegrationAccountSchemaResponseSettingsField instantiates a new AmsIntegrationAccountSchemaResponseSettingsField object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAmsIntegrationAccountSchemaResponseSettingsField() *AmsIntegrationAccountSchemaResponseSettingsField {
	this := AmsIntegrationAccountSchemaResponseSettingsField{}
	return &this
}

// NewAmsIntegrationAccountSchemaResponseSettingsFieldWithDefaults instantiates a new AmsIntegrationAccountSchemaResponseSettingsField object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAmsIntegrationAccountSchemaResponseSettingsFieldWithDefaults() *AmsIntegrationAccountSchemaResponseSettingsField {
	this := AmsIntegrationAccountSchemaResponseSettingsField{}
	return &this
}

// GetAdditionalProperties returns the AdditionalProperties field value if set, zero value otherwise.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) GetAdditionalProperties() bool {
	if o == nil || o.AdditionalProperties == nil {
		var ret bool
		return ret
	}
	return *o.AdditionalProperties
}

// GetAdditionalPropertiesOk returns a tuple with the AdditionalProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) GetAdditionalPropertiesOk() (*bool, bool) {
	if o == nil || o.AdditionalProperties == nil {
		return nil, false
	}
	return o.AdditionalProperties, true
}

// HasAdditionalProperties returns a boolean if a field has been set.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) HasAdditionalProperties() bool {
	return o != nil && o.AdditionalProperties != nil
}

// SetAdditionalProperties gets a reference to the given bool and assigns it to the AdditionalProperties field.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) SetAdditionalProperties(v bool) {
	o.AdditionalProperties = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) GetDefault() interface{} {
	if o == nil || o.Default == nil {
		var ret interface{}
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) GetDefaultOk() (*interface{}, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return &o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given interface{} and assigns it to the Default field.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) SetDefault(v interface{}) {
	o.Default = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) SetDescription(v string) {
	o.Description = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) GetItems() interface{} {
	if o == nil || o.Items == nil {
		var ret interface{}
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) GetItemsOk() (*interface{}, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given interface{} and assigns it to the Items field.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) SetItems(v interface{}) {
	o.Items = v
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) GetMinLength() int64 {
	if o == nil || o.MinLength == nil {
		var ret int64
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) GetMinLengthOk() (*int64, bool) {
	if o == nil || o.MinLength == nil {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) HasMinLength() bool {
	return o != nil && o.MinLength != nil
}

// SetMinLength gets a reference to the given int64 and assigns it to the MinLength field.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) SetMinLength(v int64) {
	o.MinLength = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AmsIntegrationAccountSchemaResponseSettingsField) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AmsIntegrationAccountSchemaResponseSettingsField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AdditionalProperties != nil {
		toSerialize["additionalProperties"] = o.AdditionalProperties
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.MinLength != nil {
		toSerialize["minLength"] = o.MinLength
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
func (o *AmsIntegrationAccountSchemaResponseSettingsField) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AdditionalProperties *bool       `json:"additionalProperties,omitempty"`
		Default              interface{} `json:"default,omitempty"`
		Description          *string     `json:"description,omitempty"`
		Items                interface{} `json:"items,omitempty"`
		MinLength            *int64      `json:"minLength,omitempty"`
		Type                 *string     `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"additionalProperties", "default", "description", "items", "minLength", "type"})
	} else {
		return err
	}
	o.AdditionalProperties = all.AdditionalProperties
	o.Default = all.Default
	o.Description = all.Description
	o.Items = all.Items
	o.MinLength = all.MinLength
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
