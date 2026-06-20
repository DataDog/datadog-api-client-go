// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AmsIntegrationAccountSchemaResponse Response containing the JSON schema for an integration's account configuration.
// This schema defines the required and optional fields for both settings and secrets,
// including field types, validation rules, and descriptions.
//
// The response is a standard [JSON Schema (draft-07)](https://json-schema.org/draft-07/schema#) document describing the account
// configuration structure. Because this is a dynamic JSON Schema, the exact properties vary by integration.
type AmsIntegrationAccountSchemaResponse struct {
	// Whether additional properties are allowed at the root level (typically false).
	AdditionalProperties *bool `json:"additionalProperties,omitempty"`
	// The properties object containing settings and secrets schema definitions.
	// Both are always present in every integration schema, but the fields within each
	// vary depending on the specific integration.
	Properties AmsIntegrationAccountSchemaResponseProperties `json:"properties"`
	// List of required top-level properties.
	Required []string `json:"required"`
	// The root type of the schema (always "object").
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAmsIntegrationAccountSchemaResponse instantiates a new AmsIntegrationAccountSchemaResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAmsIntegrationAccountSchemaResponse(properties AmsIntegrationAccountSchemaResponseProperties, required []string, typeVar string) *AmsIntegrationAccountSchemaResponse {
	this := AmsIntegrationAccountSchemaResponse{}
	this.Properties = properties
	this.Required = required
	this.Type = typeVar
	return &this
}

// NewAmsIntegrationAccountSchemaResponseWithDefaults instantiates a new AmsIntegrationAccountSchemaResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAmsIntegrationAccountSchemaResponseWithDefaults() *AmsIntegrationAccountSchemaResponse {
	this := AmsIntegrationAccountSchemaResponse{}
	return &this
}

// GetAdditionalProperties returns the AdditionalProperties field value if set, zero value otherwise.
func (o *AmsIntegrationAccountSchemaResponse) GetAdditionalProperties() bool {
	if o == nil || o.AdditionalProperties == nil {
		var ret bool
		return ret
	}
	return *o.AdditionalProperties
}

// GetAdditionalPropertiesOk returns a tuple with the AdditionalProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountSchemaResponse) GetAdditionalPropertiesOk() (*bool, bool) {
	if o == nil || o.AdditionalProperties == nil {
		return nil, false
	}
	return o.AdditionalProperties, true
}

// HasAdditionalProperties returns a boolean if a field has been set.
func (o *AmsIntegrationAccountSchemaResponse) HasAdditionalProperties() bool {
	return o != nil && o.AdditionalProperties != nil
}

// SetAdditionalProperties gets a reference to the given bool and assigns it to the AdditionalProperties field.
func (o *AmsIntegrationAccountSchemaResponse) SetAdditionalProperties(v bool) {
	o.AdditionalProperties = &v
}

// GetProperties returns the Properties field value.
func (o *AmsIntegrationAccountSchemaResponse) GetProperties() AmsIntegrationAccountSchemaResponseProperties {
	if o == nil {
		var ret AmsIntegrationAccountSchemaResponseProperties
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountSchemaResponse) GetPropertiesOk() (*AmsIntegrationAccountSchemaResponseProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value.
func (o *AmsIntegrationAccountSchemaResponse) SetProperties(v AmsIntegrationAccountSchemaResponseProperties) {
	o.Properties = v
}

// GetRequired returns the Required field value.
func (o *AmsIntegrationAccountSchemaResponse) GetRequired() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountSchemaResponse) GetRequiredOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value.
func (o *AmsIntegrationAccountSchemaResponse) SetRequired(v []string) {
	o.Required = v
}

// GetType returns the Type field value.
func (o *AmsIntegrationAccountSchemaResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountSchemaResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AmsIntegrationAccountSchemaResponse) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AmsIntegrationAccountSchemaResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AdditionalProperties != nil {
		toSerialize["additionalProperties"] = o.AdditionalProperties
	}
	toSerialize["properties"] = o.Properties
	toSerialize["required"] = o.Required
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AmsIntegrationAccountSchemaResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AdditionalProperties *bool                                          `json:"additionalProperties,omitempty"`
		Properties           *AmsIntegrationAccountSchemaResponseProperties `json:"properties"`
		Required             *[]string                                      `json:"required"`
		Type                 *string                                        `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Properties == nil {
		return fmt.Errorf("required field properties missing")
	}
	if all.Required == nil {
		return fmt.Errorf("required field required missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"additionalProperties", "properties", "required", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AdditionalProperties = all.AdditionalProperties
	if all.Properties.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Properties = *all.Properties
	o.Required = *all.Required
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
