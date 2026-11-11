// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAICodexIntegrationAccountUpdateData Data envelope for updating an OpenAI Codex integration account.
type OpenAICodexIntegrationAccountUpdateData struct {
	// Writable attributes used to update an OpenAI Codex integration account. Every field is optional; only the fields provided are changed. When `dataflows` is provided, only the dataflow ids included in the request are modified; dataflows omitted from the map keep their current configuration.
	Attributes OpenAICodexIntegrationAccountUpdateAttributes `json:"attributes"`
	// Unique identifier of the OpenAI Codex integration account to update.
	Id string `json:"id"`
	// The type of the integration account resource. Always `integration-account`.
	Type IntegrationAccountType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpenAICodexIntegrationAccountUpdateData instantiates a new OpenAICodexIntegrationAccountUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAICodexIntegrationAccountUpdateData(attributes OpenAICodexIntegrationAccountUpdateAttributes, id string, typeVar IntegrationAccountType) *OpenAICodexIntegrationAccountUpdateData {
	this := OpenAICodexIntegrationAccountUpdateData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewOpenAICodexIntegrationAccountUpdateDataWithDefaults instantiates a new OpenAICodexIntegrationAccountUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAICodexIntegrationAccountUpdateDataWithDefaults() *OpenAICodexIntegrationAccountUpdateData {
	this := OpenAICodexIntegrationAccountUpdateData{}
	var typeVar IntegrationAccountType = INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *OpenAICodexIntegrationAccountUpdateData) GetAttributes() OpenAICodexIntegrationAccountUpdateAttributes {
	if o == nil {
		var ret OpenAICodexIntegrationAccountUpdateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OpenAICodexIntegrationAccountUpdateData) GetAttributesOk() (*OpenAICodexIntegrationAccountUpdateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *OpenAICodexIntegrationAccountUpdateData) SetAttributes(v OpenAICodexIntegrationAccountUpdateAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *OpenAICodexIntegrationAccountUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OpenAICodexIntegrationAccountUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *OpenAICodexIntegrationAccountUpdateData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *OpenAICodexIntegrationAccountUpdateData) GetType() IntegrationAccountType {
	if o == nil {
		var ret IntegrationAccountType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OpenAICodexIntegrationAccountUpdateData) GetTypeOk() (*IntegrationAccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OpenAICodexIntegrationAccountUpdateData) SetType(v IntegrationAccountType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAICodexIntegrationAccountUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpenAICodexIntegrationAccountUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *OpenAICodexIntegrationAccountUpdateAttributes `json:"attributes"`
		Id         *string                                        `json:"id"`
		Type       *IntegrationAccountType                        `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
