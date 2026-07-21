// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentResponderDataResponse Incident responder data in a response.
type IncidentResponderDataResponse struct {
	// Attributes of an incident responder in a response.
	Attributes IncidentResponderDataAttributesResponse `json:"attributes"`
	// The responder identifier.
	Id uuid.UUID `json:"id"`
	// Relationships for an incident responder.
	Relationships *IncidentResponderRelationships `json:"relationships,omitempty"`
	// Incident responder resource type.
	Type IncidentResponderType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentResponderDataResponse instantiates a new IncidentResponderDataResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentResponderDataResponse(attributes IncidentResponderDataAttributesResponse, id uuid.UUID, typeVar IncidentResponderType) *IncidentResponderDataResponse {
	this := IncidentResponderDataResponse{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewIncidentResponderDataResponseWithDefaults instantiates a new IncidentResponderDataResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentResponderDataResponseWithDefaults() *IncidentResponderDataResponse {
	this := IncidentResponderDataResponse{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *IncidentResponderDataResponse) GetAttributes() IncidentResponderDataAttributesResponse {
	if o == nil {
		var ret IncidentResponderDataAttributesResponse
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *IncidentResponderDataResponse) GetAttributesOk() (*IncidentResponderDataAttributesResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *IncidentResponderDataResponse) SetAttributes(v IncidentResponderDataAttributesResponse) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *IncidentResponderDataResponse) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncidentResponderDataResponse) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *IncidentResponderDataResponse) SetId(v uuid.UUID) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *IncidentResponderDataResponse) GetRelationships() IncidentResponderRelationships {
	if o == nil || o.Relationships == nil {
		var ret IncidentResponderRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponderDataResponse) GetRelationshipsOk() (*IncidentResponderRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *IncidentResponderDataResponse) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given IncidentResponderRelationships and assigns it to the Relationships field.
func (o *IncidentResponderDataResponse) SetRelationships(v IncidentResponderRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *IncidentResponderDataResponse) GetType() IncidentResponderType {
	if o == nil {
		var ret IncidentResponderType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentResponderDataResponse) GetTypeOk() (*IncidentResponderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentResponderDataResponse) SetType(v IncidentResponderType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentResponderDataResponse) MarshalJSON() ([]byte, error) {
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
func (o *IncidentResponderDataResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *IncidentResponderDataAttributesResponse `json:"attributes"`
		Id            *uuid.UUID                               `json:"id"`
		Relationships *IncidentResponderRelationships          `json:"relationships,omitempty"`
		Type          *IncidentResponderType                   `json:"type"`
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
