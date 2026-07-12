// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AmsIntegrationAccountSchemaResponseSecretsObject JSON Schema definition for the secrets object.
// Contains sensitive credentials required for the integration such as API keys,
// tokens, and passwords. These values are write-only and never returned in responses.
type AmsIntegrationAccountSchemaResponseSecretsObject struct {
	// Whether additional properties are allowed (typically false).
	AdditionalProperties *bool `json:"additionalProperties,omitempty"`
	// The individual secret fields for this integration.
	// Field names and types vary by integration.
	Properties map[string]AmsIntegrationAccountSchemaResponseSettingsField `json:"properties,omitempty"`
	// List of required secret field names.
	Required []string `json:"required,omitempty"`
	// Always "object" for the secrets container.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAmsIntegrationAccountSchemaResponseSecretsObject instantiates a new AmsIntegrationAccountSchemaResponseSecretsObject object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAmsIntegrationAccountSchemaResponseSecretsObject() *AmsIntegrationAccountSchemaResponseSecretsObject {
	this := AmsIntegrationAccountSchemaResponseSecretsObject{}
	return &this
}

// NewAmsIntegrationAccountSchemaResponseSecretsObjectWithDefaults instantiates a new AmsIntegrationAccountSchemaResponseSecretsObject object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAmsIntegrationAccountSchemaResponseSecretsObjectWithDefaults() *AmsIntegrationAccountSchemaResponseSecretsObject {
	this := AmsIntegrationAccountSchemaResponseSecretsObject{}
	return &this
}

// GetAdditionalProperties returns the AdditionalProperties field value if set, zero value otherwise.
func (o *AmsIntegrationAccountSchemaResponseSecretsObject) GetAdditionalProperties() bool {
	if o == nil || o.AdditionalProperties == nil {
		var ret bool
		return ret
	}
	return *o.AdditionalProperties
}

// GetAdditionalPropertiesOk returns a tuple with the AdditionalProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountSchemaResponseSecretsObject) GetAdditionalPropertiesOk() (*bool, bool) {
	if o == nil || o.AdditionalProperties == nil {
		return nil, false
	}
	return o.AdditionalProperties, true
}

// HasAdditionalProperties returns a boolean if a field has been set.
func (o *AmsIntegrationAccountSchemaResponseSecretsObject) HasAdditionalProperties() bool {
	return o != nil && o.AdditionalProperties != nil
}

// SetAdditionalProperties gets a reference to the given bool and assigns it to the AdditionalProperties field.
func (o *AmsIntegrationAccountSchemaResponseSecretsObject) SetAdditionalProperties(v bool) {
	o.AdditionalProperties = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *AmsIntegrationAccountSchemaResponseSecretsObject) GetProperties() map[string]AmsIntegrationAccountSchemaResponseSettingsField {
	if o == nil || o.Properties == nil {
		var ret map[string]AmsIntegrationAccountSchemaResponseSettingsField
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountSchemaResponseSecretsObject) GetPropertiesOk() (*map[string]AmsIntegrationAccountSchemaResponseSettingsField, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *AmsIntegrationAccountSchemaResponseSecretsObject) HasProperties() bool {
	return o != nil && o.Properties != nil
}

// SetProperties gets a reference to the given map[string]AmsIntegrationAccountSchemaResponseSettingsField and assigns it to the Properties field.
func (o *AmsIntegrationAccountSchemaResponseSecretsObject) SetProperties(v map[string]AmsIntegrationAccountSchemaResponseSettingsField) {
	o.Properties = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *AmsIntegrationAccountSchemaResponseSecretsObject) GetRequired() []string {
	if o == nil || o.Required == nil {
		var ret []string
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountSchemaResponseSecretsObject) GetRequiredOk() (*[]string, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return &o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *AmsIntegrationAccountSchemaResponseSecretsObject) HasRequired() bool {
	return o != nil && o.Required != nil
}

// SetRequired gets a reference to the given []string and assigns it to the Required field.
func (o *AmsIntegrationAccountSchemaResponseSecretsObject) SetRequired(v []string) {
	o.Required = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AmsIntegrationAccountSchemaResponseSecretsObject) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountSchemaResponseSecretsObject) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AmsIntegrationAccountSchemaResponseSecretsObject) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AmsIntegrationAccountSchemaResponseSecretsObject) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AmsIntegrationAccountSchemaResponseSecretsObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AdditionalProperties != nil {
		toSerialize["additionalProperties"] = o.AdditionalProperties
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
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
func (o *AmsIntegrationAccountSchemaResponseSecretsObject) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AdditionalProperties *bool                                                       `json:"additionalProperties,omitempty"`
		Properties           map[string]AmsIntegrationAccountSchemaResponseSettingsField `json:"properties,omitempty"`
		Required             []string                                                    `json:"required,omitempty"`
		Type                 *string                                                     `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"additionalProperties", "properties", "required", "type"})
	} else {
		return err
	}
	o.AdditionalProperties = all.AdditionalProperties
	o.Properties = all.Properties
	o.Required = all.Required
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
