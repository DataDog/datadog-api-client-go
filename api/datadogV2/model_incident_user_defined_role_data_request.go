// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedRoleDataRequest Data for creating an incident user-defined role.
type IncidentUserDefinedRoleDataRequest struct {
	// Attributes for creating an incident user-defined role.
	Attributes IncidentUserDefinedRoleDataAttributesRequest `json:"attributes"`
	// Relationships for creating a user-defined role.
	Relationships IncidentUserDefinedRoleRelationshipsRequest `json:"relationships"`
	// Incident user-defined role resource type.
	Type IncidentUserDefinedRoleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentUserDefinedRoleDataRequest instantiates a new IncidentUserDefinedRoleDataRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentUserDefinedRoleDataRequest(attributes IncidentUserDefinedRoleDataAttributesRequest, relationships IncidentUserDefinedRoleRelationshipsRequest, typeVar IncidentUserDefinedRoleType) *IncidentUserDefinedRoleDataRequest {
	this := IncidentUserDefinedRoleDataRequest{}
	this.Attributes = attributes
	this.Relationships = relationships
	this.Type = typeVar
	return &this
}

// NewIncidentUserDefinedRoleDataRequestWithDefaults instantiates a new IncidentUserDefinedRoleDataRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentUserDefinedRoleDataRequestWithDefaults() *IncidentUserDefinedRoleDataRequest {
	this := IncidentUserDefinedRoleDataRequest{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *IncidentUserDefinedRoleDataRequest) GetAttributes() IncidentUserDefinedRoleDataAttributesRequest {
	if o == nil {
		var ret IncidentUserDefinedRoleDataAttributesRequest
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRoleDataRequest) GetAttributesOk() (*IncidentUserDefinedRoleDataAttributesRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *IncidentUserDefinedRoleDataRequest) SetAttributes(v IncidentUserDefinedRoleDataAttributesRequest) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value.
func (o *IncidentUserDefinedRoleDataRequest) GetRelationships() IncidentUserDefinedRoleRelationshipsRequest {
	if o == nil {
		var ret IncidentUserDefinedRoleRelationshipsRequest
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRoleDataRequest) GetRelationshipsOk() (*IncidentUserDefinedRoleRelationshipsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value.
func (o *IncidentUserDefinedRoleDataRequest) SetRelationships(v IncidentUserDefinedRoleRelationshipsRequest) {
	o.Relationships = v
}

// GetType returns the Type field value.
func (o *IncidentUserDefinedRoleDataRequest) GetType() IncidentUserDefinedRoleType {
	if o == nil {
		var ret IncidentUserDefinedRoleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRoleDataRequest) GetTypeOk() (*IncidentUserDefinedRoleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentUserDefinedRoleDataRequest) SetType(v IncidentUserDefinedRoleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentUserDefinedRoleDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentUserDefinedRoleDataRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *IncidentUserDefinedRoleDataAttributesRequest `json:"attributes"`
		Relationships *IncidentUserDefinedRoleRelationshipsRequest  `json:"relationships"`
		Type          *IncidentUserDefinedRoleType                  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Relationships == nil {
		return fmt.Errorf("required field relationships missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	if all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = *all.Relationships
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
