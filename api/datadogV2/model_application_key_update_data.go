// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationKeyUpdateData Object used to update an application key.
type ApplicationKeyUpdateData struct {
	// Attributes used to update an application Key.
	Attributes *ApplicationKeyUpdateAttributes `json:"attributes,omitempty"`
	// ID of the application key.
	Id string `json:"id"`
	// Application Keys resource type.
	Type ApplicationKeysType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewApplicationKeyUpdateData instantiates a new ApplicationKeyUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationKeyUpdateData(id string, typeVar ApplicationKeysType) *ApplicationKeyUpdateData {
	this := ApplicationKeyUpdateData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewApplicationKeyUpdateDataWithDefaults instantiates a new ApplicationKeyUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationKeyUpdateDataWithDefaults() *ApplicationKeyUpdateData {
	this := ApplicationKeyUpdateData{}
	var typeVar ApplicationKeysType = APPLICATIONKEYSTYPE_APPLICATION_KEYS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ApplicationKeyUpdateData) GetAttributes() ApplicationKeyUpdateAttributes {
	if o == nil || o.Attributes == nil {
		var ret ApplicationKeyUpdateAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationKeyUpdateData) GetAttributesOk() (*ApplicationKeyUpdateAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ApplicationKeyUpdateData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given ApplicationKeyUpdateAttributes and assigns it to the Attributes field.
func (o *ApplicationKeyUpdateData) SetAttributes(v ApplicationKeyUpdateAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *ApplicationKeyUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationKeyUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ApplicationKeyUpdateData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *ApplicationKeyUpdateData) GetType() ApplicationKeysType {
	if o == nil {
		var ret ApplicationKeysType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationKeyUpdateData) GetTypeOk() (*ApplicationKeysType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ApplicationKeyUpdateData) SetType(v ApplicationKeysType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationKeyUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationKeyUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ApplicationKeyUpdateAttributes `json:"attributes,omitempty"`
		Id         *string                         `json:"id"`
		Type       *ApplicationKeysType            `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
