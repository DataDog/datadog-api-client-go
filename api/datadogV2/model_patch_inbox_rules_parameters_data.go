// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchInboxRulesParametersData Data of the inbox rule patch request: the rule UUID, the rule type, and the rule attributes. All fields are required.
type PatchInboxRulesParametersData struct {
	// Attributes of the inbox rule patch request: the rule name, the rule details, the associated action, and the enabled field.
	Attributes PatchInboxRulesParametersDataAttributes `json:"attributes"`
	// The ID of a pipeline rule
	Id uuid.UUID `json:"id"`
	// The pipeline rule type associated to inbox rules
	Type InboxRulesType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPatchInboxRulesParametersData instantiates a new PatchInboxRulesParametersData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPatchInboxRulesParametersData(attributes PatchInboxRulesParametersDataAttributes, id uuid.UUID, typeVar InboxRulesType) *PatchInboxRulesParametersData {
	this := PatchInboxRulesParametersData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewPatchInboxRulesParametersDataWithDefaults instantiates a new PatchInboxRulesParametersData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPatchInboxRulesParametersDataWithDefaults() *PatchInboxRulesParametersData {
	this := PatchInboxRulesParametersData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *PatchInboxRulesParametersData) GetAttributes() PatchInboxRulesParametersDataAttributes {
	if o == nil {
		var ret PatchInboxRulesParametersDataAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PatchInboxRulesParametersData) GetAttributesOk() (*PatchInboxRulesParametersDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *PatchInboxRulesParametersData) SetAttributes(v PatchInboxRulesParametersDataAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *PatchInboxRulesParametersData) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchInboxRulesParametersData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *PatchInboxRulesParametersData) SetId(v uuid.UUID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *PatchInboxRulesParametersData) GetType() InboxRulesType {
	if o == nil {
		var ret InboxRulesType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PatchInboxRulesParametersData) GetTypeOk() (*InboxRulesType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *PatchInboxRulesParametersData) SetType(v InboxRulesType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PatchInboxRulesParametersData) MarshalJSON() ([]byte, error) {
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
func (o *PatchInboxRulesParametersData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *PatchInboxRulesParametersDataAttributes `json:"attributes"`
		Id         *uuid.UUID                               `json:"id"`
		Type       *InboxRulesType                          `json:"type"`
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
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