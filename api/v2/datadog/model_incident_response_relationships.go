// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadog

import (
	"encoding/json"
)

// IncidentResponseRelationships The incident's relationships from a response.
type IncidentResponseRelationships struct {
	// Relationship to user.
	CommanderUser *NullableRelationshipToUser `json:"commander_user,omitempty"`
	// Relationship to user.
	CreatedByUser *RelationshipToUser `json:"created_by_user,omitempty"`
	// A relationship reference for multiple integration metadata objects.
	Integrations *RelationshipToIncidentIntegrationMetadatas `json:"integrations,omitempty"`
	// Relationship to user.
	LastModifiedByUser *RelationshipToUser `json:"last_modified_by_user,omitempty"`
	// A relationship reference for postmortems.
	Postmortem *RelationshipToIncidentPostmortem `json:"postmortem,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewIncidentResponseRelationships instantiates a new IncidentResponseRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentResponseRelationships() *IncidentResponseRelationships {
	this := IncidentResponseRelationships{}
	return &this
}

// NewIncidentResponseRelationshipsWithDefaults instantiates a new IncidentResponseRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentResponseRelationshipsWithDefaults() *IncidentResponseRelationships {
	this := IncidentResponseRelationships{}
	return &this
}

// GetCommanderUser returns the CommanderUser field value if set, zero value otherwise.
func (o *IncidentResponseRelationships) GetCommanderUser() NullableRelationshipToUser {
	if o == nil || o.CommanderUser == nil {
		var ret NullableRelationshipToUser
		return ret
	}
	return *o.CommanderUser
}

// GetCommanderUserOk returns a tuple with the CommanderUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseRelationships) GetCommanderUserOk() (*NullableRelationshipToUser, bool) {
	if o == nil || o.CommanderUser == nil {
		return nil, false
	}
	return o.CommanderUser, true
}

// HasCommanderUser returns a boolean if a field has been set.
func (o *IncidentResponseRelationships) HasCommanderUser() bool {
	if o != nil && o.CommanderUser != nil {
		return true
	}

	return false
}

// SetCommanderUser gets a reference to the given NullableRelationshipToUser and assigns it to the CommanderUser field.
func (o *IncidentResponseRelationships) SetCommanderUser(v NullableRelationshipToUser) {
	o.CommanderUser = &v
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise.
func (o *IncidentResponseRelationships) GetCreatedByUser() RelationshipToUser {
	if o == nil || o.CreatedByUser == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseRelationships) GetCreatedByUserOk() (*RelationshipToUser, bool) {
	if o == nil || o.CreatedByUser == nil {
		return nil, false
	}
	return o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *IncidentResponseRelationships) HasCreatedByUser() bool {
	if o != nil && o.CreatedByUser != nil {
		return true
	}

	return false
}

// SetCreatedByUser gets a reference to the given RelationshipToUser and assigns it to the CreatedByUser field.
func (o *IncidentResponseRelationships) SetCreatedByUser(v RelationshipToUser) {
	o.CreatedByUser = &v
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise.
func (o *IncidentResponseRelationships) GetIntegrations() RelationshipToIncidentIntegrationMetadatas {
	if o == nil || o.Integrations == nil {
		var ret RelationshipToIncidentIntegrationMetadatas
		return ret
	}
	return *o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseRelationships) GetIntegrationsOk() (*RelationshipToIncidentIntegrationMetadatas, bool) {
	if o == nil || o.Integrations == nil {
		return nil, false
	}
	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *IncidentResponseRelationships) HasIntegrations() bool {
	if o != nil && o.Integrations != nil {
		return true
	}

	return false
}

// SetIntegrations gets a reference to the given RelationshipToIncidentIntegrationMetadatas and assigns it to the Integrations field.
func (o *IncidentResponseRelationships) SetIntegrations(v RelationshipToIncidentIntegrationMetadatas) {
	o.Integrations = &v
}

// GetLastModifiedByUser returns the LastModifiedByUser field value if set, zero value otherwise.
func (o *IncidentResponseRelationships) GetLastModifiedByUser() RelationshipToUser {
	if o == nil || o.LastModifiedByUser == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseRelationships) GetLastModifiedByUserOk() (*RelationshipToUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		return nil, false
	}
	return o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *IncidentResponseRelationships) HasLastModifiedByUser() bool {
	if o != nil && o.LastModifiedByUser != nil {
		return true
	}

	return false
}

// SetLastModifiedByUser gets a reference to the given RelationshipToUser and assigns it to the LastModifiedByUser field.
func (o *IncidentResponseRelationships) SetLastModifiedByUser(v RelationshipToUser) {
	o.LastModifiedByUser = &v
}

// GetPostmortem returns the Postmortem field value if set, zero value otherwise.
func (o *IncidentResponseRelationships) GetPostmortem() RelationshipToIncidentPostmortem {
	if o == nil || o.Postmortem == nil {
		var ret RelationshipToIncidentPostmortem
		return ret
	}
	return *o.Postmortem
}

// GetPostmortemOk returns a tuple with the Postmortem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseRelationships) GetPostmortemOk() (*RelationshipToIncidentPostmortem, bool) {
	if o == nil || o.Postmortem == nil {
		return nil, false
	}
	return o.Postmortem, true
}

// HasPostmortem returns a boolean if a field has been set.
func (o *IncidentResponseRelationships) HasPostmortem() bool {
	if o != nil && o.Postmortem != nil {
		return true
	}

	return false
}

// SetPostmortem gets a reference to the given RelationshipToIncidentPostmortem and assigns it to the Postmortem field.
func (o *IncidentResponseRelationships) SetPostmortem(v RelationshipToIncidentPostmortem) {
	o.Postmortem = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentResponseRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.CommanderUser != nil {
		toSerialize["commander_user"] = o.CommanderUser
	}
	if o.CreatedByUser != nil {
		toSerialize["created_by_user"] = o.CreatedByUser
	}
	if o.Integrations != nil {
		toSerialize["integrations"] = o.Integrations
	}
	if o.LastModifiedByUser != nil {
		toSerialize["last_modified_by_user"] = o.LastModifiedByUser
	}
	if o.Postmortem != nil {
		toSerialize["postmortem"] = o.Postmortem
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentResponseRelationships) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		CommanderUser      *NullableRelationshipToUser                 `json:"commander_user,omitempty"`
		CreatedByUser      *RelationshipToUser                         `json:"created_by_user,omitempty"`
		Integrations       *RelationshipToIncidentIntegrationMetadatas `json:"integrations,omitempty"`
		LastModifiedByUser *RelationshipToUser                         `json:"last_modified_by_user,omitempty"`
		Postmortem         *RelationshipToIncidentPostmortem           `json:"postmortem,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.CommanderUser != nil && all.CommanderUser.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.CommanderUser = all.CommanderUser
	if all.CreatedByUser != nil && all.CreatedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.CreatedByUser = all.CreatedByUser
	if all.Integrations != nil && all.Integrations.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Integrations = all.Integrations
	if all.LastModifiedByUser != nil && all.LastModifiedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.LastModifiedByUser = all.LastModifiedByUser
	if all.Postmortem != nil && all.Postmortem.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Postmortem = all.Postmortem
	return nil
}
