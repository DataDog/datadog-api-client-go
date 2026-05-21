// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRoleAssignmentDataRequest Role assignment data for a request.
type IncidentRoleAssignmentDataRequest struct {
	// Attributes for creating a role assignment.
	Attributes *IncidentRoleAssignmentDataAttributesRequest `json:"attributes,omitempty"`
	// Relationships for creating a role assignment.
	Relationships IncidentRoleAssignmentRelationshipsRequest `json:"relationships"`
	// Incident role assignment resource type.
	Type IncidentRoleAssignmentType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentRoleAssignmentDataRequest instantiates a new IncidentRoleAssignmentDataRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentRoleAssignmentDataRequest(relationships IncidentRoleAssignmentRelationshipsRequest, typeVar IncidentRoleAssignmentType) *IncidentRoleAssignmentDataRequest {
	this := IncidentRoleAssignmentDataRequest{}
	this.Relationships = relationships
	this.Type = typeVar
	return &this
}

// NewIncidentRoleAssignmentDataRequestWithDefaults instantiates a new IncidentRoleAssignmentDataRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentRoleAssignmentDataRequestWithDefaults() *IncidentRoleAssignmentDataRequest {
	this := IncidentRoleAssignmentDataRequest{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IncidentRoleAssignmentDataRequest) GetAttributes() IncidentRoleAssignmentDataAttributesRequest {
	if o == nil || o.Attributes == nil {
		var ret IncidentRoleAssignmentDataAttributesRequest
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRoleAssignmentDataRequest) GetAttributesOk() (*IncidentRoleAssignmentDataAttributesRequest, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IncidentRoleAssignmentDataRequest) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given IncidentRoleAssignmentDataAttributesRequest and assigns it to the Attributes field.
func (o *IncidentRoleAssignmentDataRequest) SetAttributes(v IncidentRoleAssignmentDataAttributesRequest) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value.
func (o *IncidentRoleAssignmentDataRequest) GetRelationships() IncidentRoleAssignmentRelationshipsRequest {
	if o == nil {
		var ret IncidentRoleAssignmentRelationshipsRequest
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *IncidentRoleAssignmentDataRequest) GetRelationshipsOk() (*IncidentRoleAssignmentRelationshipsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value.
func (o *IncidentRoleAssignmentDataRequest) SetRelationships(v IncidentRoleAssignmentRelationshipsRequest) {
	o.Relationships = v
}

// GetType returns the Type field value.
func (o *IncidentRoleAssignmentDataRequest) GetType() IncidentRoleAssignmentType {
	if o == nil {
		var ret IncidentRoleAssignmentType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentRoleAssignmentDataRequest) GetTypeOk() (*IncidentRoleAssignmentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentRoleAssignmentDataRequest) SetType(v IncidentRoleAssignmentType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentRoleAssignmentDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["relationships"] = o.Relationships
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentRoleAssignmentDataRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *IncidentRoleAssignmentDataAttributesRequest `json:"attributes,omitempty"`
		Relationships *IncidentRoleAssignmentRelationshipsRequest  `json:"relationships"`
		Type          *IncidentRoleAssignmentType                  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Relationships == nil {
		return fmt.Errorf("required field relationships missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
