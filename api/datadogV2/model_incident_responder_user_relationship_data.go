// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentResponderUserRelationshipData A user relationship data object for creating a responder.
type IncidentResponderUserRelationshipData struct {
	// The user identifier.
	Id uuid.UUID `json:"id"`
	// The user resource type.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentResponderUserRelationshipData instantiates a new IncidentResponderUserRelationshipData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentResponderUserRelationshipData(id uuid.UUID, typeVar string) *IncidentResponderUserRelationshipData {
	this := IncidentResponderUserRelationshipData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewIncidentResponderUserRelationshipDataWithDefaults instantiates a new IncidentResponderUserRelationshipData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentResponderUserRelationshipDataWithDefaults() *IncidentResponderUserRelationshipData {
	this := IncidentResponderUserRelationshipData{}
	return &this
}

// GetId returns the Id field value.
func (o *IncidentResponderUserRelationshipData) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncidentResponderUserRelationshipData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *IncidentResponderUserRelationshipData) SetId(v uuid.UUID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *IncidentResponderUserRelationshipData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentResponderUserRelationshipData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentResponderUserRelationshipData) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentResponderUserRelationshipData) MarshalJSON() ([]byte, error) {
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
func (o *IncidentResponderUserRelationshipData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *uuid.UUID `json:"id"`
		Type *string    `json:"type"`
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
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
