// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyMember A user relationship to the requested team.
type TeamHierarchyMember struct {
	// Membership attributes relative to the requested team.
	Attributes TeamHierarchyMemberAttributes `json:"attributes"`
	// Relationship identifier containing the team UUID and user UUID.
	Id string `json:"id"`
	// User and requested team linkage.
	Relationships TeamHierarchyMemberRelationships `json:"relationships"`
	// Team hierarchy membership resource type.
	Type TeamHierarchyMemberType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamHierarchyMember instantiates a new TeamHierarchyMember object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamHierarchyMember(attributes TeamHierarchyMemberAttributes, id string, relationships TeamHierarchyMemberRelationships, typeVar TeamHierarchyMemberType) *TeamHierarchyMember {
	this := TeamHierarchyMember{}
	this.Attributes = attributes
	this.Id = id
	this.Relationships = relationships
	this.Type = typeVar
	return &this
}

// NewTeamHierarchyMemberWithDefaults instantiates a new TeamHierarchyMember object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamHierarchyMemberWithDefaults() *TeamHierarchyMember {
	this := TeamHierarchyMember{}
	var typeVar TeamHierarchyMemberType = TEAMHIERARCHYMEMBERTYPE_USER_TEAM_RELATIONSHIP
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *TeamHierarchyMember) GetAttributes() TeamHierarchyMemberAttributes {
	if o == nil {
		var ret TeamHierarchyMemberAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMember) GetAttributesOk() (*TeamHierarchyMemberAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *TeamHierarchyMember) SetAttributes(v TeamHierarchyMemberAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *TeamHierarchyMember) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMember) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *TeamHierarchyMember) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value.
func (o *TeamHierarchyMember) GetRelationships() TeamHierarchyMemberRelationships {
	if o == nil {
		var ret TeamHierarchyMemberRelationships
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMember) GetRelationshipsOk() (*TeamHierarchyMemberRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value.
func (o *TeamHierarchyMember) SetRelationships(v TeamHierarchyMemberRelationships) {
	o.Relationships = v
}

// GetType returns the Type field value.
func (o *TeamHierarchyMember) GetType() TeamHierarchyMemberType {
	if o == nil {
		var ret TeamHierarchyMemberType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMember) GetTypeOk() (*TeamHierarchyMemberType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TeamHierarchyMember) SetType(v TeamHierarchyMemberType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamHierarchyMember) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["relationships"] = o.Relationships
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamHierarchyMember) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *TeamHierarchyMemberAttributes    `json:"attributes"`
		Id            *string                           `json:"id"`
		Relationships *TeamHierarchyMemberRelationships `json:"relationships"`
		Type          *TeamHierarchyMemberType          `json:"type"`
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
	if all.Relationships == nil {
		return fmt.Errorf("required field relationships missing")
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
