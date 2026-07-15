// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedRolePatchDataRequest Data for updating an incident user-defined role.
type IncidentUserDefinedRolePatchDataRequest struct {
	// Attributes for updating an incident user-defined role.
	Attributes *IncidentUserDefinedRolePatchDataAttributesRequest `json:"attributes,omitempty"`
	// The ID of the user-defined role to update.
	Id uuid.UUID `json:"id"`
	// Incident user-defined role resource type.
	Type IncidentUserDefinedRoleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentUserDefinedRolePatchDataRequest instantiates a new IncidentUserDefinedRolePatchDataRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentUserDefinedRolePatchDataRequest(id uuid.UUID, typeVar IncidentUserDefinedRoleType) *IncidentUserDefinedRolePatchDataRequest {
	this := IncidentUserDefinedRolePatchDataRequest{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewIncidentUserDefinedRolePatchDataRequestWithDefaults instantiates a new IncidentUserDefinedRolePatchDataRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentUserDefinedRolePatchDataRequestWithDefaults() *IncidentUserDefinedRolePatchDataRequest {
	this := IncidentUserDefinedRolePatchDataRequest{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IncidentUserDefinedRolePatchDataRequest) GetAttributes() IncidentUserDefinedRolePatchDataAttributesRequest {
	if o == nil || o.Attributes == nil {
		var ret IncidentUserDefinedRolePatchDataAttributesRequest
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRolePatchDataRequest) GetAttributesOk() (*IncidentUserDefinedRolePatchDataAttributesRequest, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IncidentUserDefinedRolePatchDataRequest) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given IncidentUserDefinedRolePatchDataAttributesRequest and assigns it to the Attributes field.
func (o *IncidentUserDefinedRolePatchDataRequest) SetAttributes(v IncidentUserDefinedRolePatchDataAttributesRequest) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *IncidentUserDefinedRolePatchDataRequest) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRolePatchDataRequest) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *IncidentUserDefinedRolePatchDataRequest) SetId(v uuid.UUID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *IncidentUserDefinedRolePatchDataRequest) GetType() IncidentUserDefinedRoleType {
	if o == nil {
		var ret IncidentUserDefinedRoleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRolePatchDataRequest) GetTypeOk() (*IncidentUserDefinedRoleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentUserDefinedRolePatchDataRequest) SetType(v IncidentUserDefinedRoleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentUserDefinedRolePatchDataRequest) MarshalJSON() ([]byte, error) {
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
func (o *IncidentUserDefinedRolePatchDataRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *IncidentUserDefinedRolePatchDataAttributesRequest `json:"attributes,omitempty"`
		Id         *uuid.UUID                                         `json:"id"`
		Type       *IncidentUserDefinedRoleType                       `json:"type"`
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
