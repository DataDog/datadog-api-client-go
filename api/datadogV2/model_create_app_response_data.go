// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateAppResponseData The data object containing the app ID.
type CreateAppResponseData struct {
	// The ID of the created app.
	Id uuid.UUID `json:"id"`
	// The app definition type.
	Type AppDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateAppResponseData instantiates a new CreateAppResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateAppResponseData(id uuid.UUID, typeVar AppDefinitionType) *CreateAppResponseData {
	this := CreateAppResponseData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewCreateAppResponseDataWithDefaults instantiates a new CreateAppResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateAppResponseDataWithDefaults() *CreateAppResponseData {
	this := CreateAppResponseData{}
	var typeVar AppDefinitionType = APPDEFINITIONTYPE_APPDEFINITIONS
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *CreateAppResponseData) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateAppResponseData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CreateAppResponseData) SetId(v uuid.UUID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *CreateAppResponseData) GetType() AppDefinitionType {
	if o == nil {
		var ret AppDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateAppResponseData) GetTypeOk() (*AppDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateAppResponseData) SetType(v AppDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateAppResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateAppResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *uuid.UUID         `json:"id"`
		Type *AppDefinitionType `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
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
