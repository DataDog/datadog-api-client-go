// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedRoleRelationshipsRequest Relationships for creating a user-defined role.
type IncidentUserDefinedRoleRelationshipsRequest struct {
	// Relationship to an incident type for a user-defined role.
	IncidentType IncidentUserDefinedRoleIncidentTypeRelationship `json:"incident_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentUserDefinedRoleRelationshipsRequest instantiates a new IncidentUserDefinedRoleRelationshipsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentUserDefinedRoleRelationshipsRequest(incidentType IncidentUserDefinedRoleIncidentTypeRelationship) *IncidentUserDefinedRoleRelationshipsRequest {
	this := IncidentUserDefinedRoleRelationshipsRequest{}
	this.IncidentType = incidentType
	return &this
}

// NewIncidentUserDefinedRoleRelationshipsRequestWithDefaults instantiates a new IncidentUserDefinedRoleRelationshipsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentUserDefinedRoleRelationshipsRequestWithDefaults() *IncidentUserDefinedRoleRelationshipsRequest {
	this := IncidentUserDefinedRoleRelationshipsRequest{}
	return &this
}

// GetIncidentType returns the IncidentType field value.
func (o *IncidentUserDefinedRoleRelationshipsRequest) GetIncidentType() IncidentUserDefinedRoleIncidentTypeRelationship {
	if o == nil {
		var ret IncidentUserDefinedRoleIncidentTypeRelationship
		return ret
	}
	return o.IncidentType
}

// GetIncidentTypeOk returns a tuple with the IncidentType field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRoleRelationshipsRequest) GetIncidentTypeOk() (*IncidentUserDefinedRoleIncidentTypeRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncidentType, true
}

// SetIncidentType sets field value.
func (o *IncidentUserDefinedRoleRelationshipsRequest) SetIncidentType(v IncidentUserDefinedRoleIncidentTypeRelationship) {
	o.IncidentType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentUserDefinedRoleRelationshipsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["incident_type"] = o.IncidentType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentUserDefinedRoleRelationshipsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncidentType *IncidentUserDefinedRoleIncidentTypeRelationship `json:"incident_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IncidentType == nil {
		return fmt.Errorf("required field incident_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"incident_type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.IncidentType.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IncidentType = *all.IncidentType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
