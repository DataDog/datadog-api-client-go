// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedRoleDataResponse Data for an incident user-defined role response.
type IncidentUserDefinedRoleDataResponse struct {
	// Attributes of an incident user-defined role.
	Attributes IncidentUserDefinedRoleDataAttributesResponse `json:"attributes"`
	// The ID of the user-defined role.
	Id uuid.UUID `json:"id"`
	// Relationships of a user-defined role response.
	Relationships *IncidentUserDefinedRoleRelationshipsResponse `json:"relationships,omitempty"`
	// Incident user-defined role resource type.
	Type IncidentUserDefinedRoleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentUserDefinedRoleDataResponse instantiates a new IncidentUserDefinedRoleDataResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentUserDefinedRoleDataResponse(attributes IncidentUserDefinedRoleDataAttributesResponse, id uuid.UUID, typeVar IncidentUserDefinedRoleType) *IncidentUserDefinedRoleDataResponse {
	this := IncidentUserDefinedRoleDataResponse{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewIncidentUserDefinedRoleDataResponseWithDefaults instantiates a new IncidentUserDefinedRoleDataResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentUserDefinedRoleDataResponseWithDefaults() *IncidentUserDefinedRoleDataResponse {
	this := IncidentUserDefinedRoleDataResponse{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *IncidentUserDefinedRoleDataResponse) GetAttributes() IncidentUserDefinedRoleDataAttributesResponse {
	if o == nil {
		var ret IncidentUserDefinedRoleDataAttributesResponse
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRoleDataResponse) GetAttributesOk() (*IncidentUserDefinedRoleDataAttributesResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *IncidentUserDefinedRoleDataResponse) SetAttributes(v IncidentUserDefinedRoleDataAttributesResponse) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *IncidentUserDefinedRoleDataResponse) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRoleDataResponse) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *IncidentUserDefinedRoleDataResponse) SetId(v uuid.UUID) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *IncidentUserDefinedRoleDataResponse) GetRelationships() IncidentUserDefinedRoleRelationshipsResponse {
	if o == nil || o.Relationships == nil {
		var ret IncidentUserDefinedRoleRelationshipsResponse
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRoleDataResponse) GetRelationshipsOk() (*IncidentUserDefinedRoleRelationshipsResponse, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *IncidentUserDefinedRoleDataResponse) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given IncidentUserDefinedRoleRelationshipsResponse and assigns it to the Relationships field.
func (o *IncidentUserDefinedRoleDataResponse) SetRelationships(v IncidentUserDefinedRoleRelationshipsResponse) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *IncidentUserDefinedRoleDataResponse) GetType() IncidentUserDefinedRoleType {
	if o == nil {
		var ret IncidentUserDefinedRoleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRoleDataResponse) GetTypeOk() (*IncidentUserDefinedRoleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentUserDefinedRoleDataResponse) SetType(v IncidentUserDefinedRoleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentUserDefinedRoleDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentUserDefinedRoleDataResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *IncidentUserDefinedRoleDataAttributesResponse `json:"attributes"`
		Id            *uuid.UUID                                     `json:"id"`
		Relationships *IncidentUserDefinedRoleRelationshipsResponse  `json:"relationships,omitempty"`
		Type          *IncidentUserDefinedRoleType                   `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
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
