// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageLinkAttributes The attributes of a page link. A page can be the child of another Datadog resource
// (for example, when an incident triggers it), or the parent of one (for example, when the
// page is escalated to an incident).
type OnCallPageLinkAttributes struct {
	// The ID of the child resource.
	ChildEntityId string `json:"child_entity_id"`
	// The type of a resource linked to a page.
	ChildEntityType OnCallPageLinkEntityType `json:"child_entity_type"`
	// The ID of the parent resource.
	ParentEntityId string `json:"parent_entity_id"`
	// The type of a resource linked to a page.
	ParentEntityType OnCallPageLinkEntityType `json:"parent_entity_type"`
	// How the parent resource relates to the child resource.
	Relationship OnCallPageLinkRelationshipType `json:"relationship"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallPageLinkAttributes instantiates a new OnCallPageLinkAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallPageLinkAttributes(childEntityId string, childEntityType OnCallPageLinkEntityType, parentEntityId string, parentEntityType OnCallPageLinkEntityType, relationship OnCallPageLinkRelationshipType) *OnCallPageLinkAttributes {
	this := OnCallPageLinkAttributes{}
	this.ChildEntityId = childEntityId
	this.ChildEntityType = childEntityType
	this.ParentEntityId = parentEntityId
	this.ParentEntityType = parentEntityType
	this.Relationship = relationship
	return &this
}

// NewOnCallPageLinkAttributesWithDefaults instantiates a new OnCallPageLinkAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallPageLinkAttributesWithDefaults() *OnCallPageLinkAttributes {
	this := OnCallPageLinkAttributes{}
	return &this
}

// GetChildEntityId returns the ChildEntityId field value.
func (o *OnCallPageLinkAttributes) GetChildEntityId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ChildEntityId
}

// GetChildEntityIdOk returns a tuple with the ChildEntityId field value
// and a boolean to check if the value has been set.
func (o *OnCallPageLinkAttributes) GetChildEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChildEntityId, true
}

// SetChildEntityId sets field value.
func (o *OnCallPageLinkAttributes) SetChildEntityId(v string) {
	o.ChildEntityId = v
}

// GetChildEntityType returns the ChildEntityType field value.
func (o *OnCallPageLinkAttributes) GetChildEntityType() OnCallPageLinkEntityType {
	if o == nil {
		var ret OnCallPageLinkEntityType
		return ret
	}
	return o.ChildEntityType
}

// GetChildEntityTypeOk returns a tuple with the ChildEntityType field value
// and a boolean to check if the value has been set.
func (o *OnCallPageLinkAttributes) GetChildEntityTypeOk() (*OnCallPageLinkEntityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChildEntityType, true
}

// SetChildEntityType sets field value.
func (o *OnCallPageLinkAttributes) SetChildEntityType(v OnCallPageLinkEntityType) {
	o.ChildEntityType = v
}

// GetParentEntityId returns the ParentEntityId field value.
func (o *OnCallPageLinkAttributes) GetParentEntityId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ParentEntityId
}

// GetParentEntityIdOk returns a tuple with the ParentEntityId field value
// and a boolean to check if the value has been set.
func (o *OnCallPageLinkAttributes) GetParentEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentEntityId, true
}

// SetParentEntityId sets field value.
func (o *OnCallPageLinkAttributes) SetParentEntityId(v string) {
	o.ParentEntityId = v
}

// GetParentEntityType returns the ParentEntityType field value.
func (o *OnCallPageLinkAttributes) GetParentEntityType() OnCallPageLinkEntityType {
	if o == nil {
		var ret OnCallPageLinkEntityType
		return ret
	}
	return o.ParentEntityType
}

// GetParentEntityTypeOk returns a tuple with the ParentEntityType field value
// and a boolean to check if the value has been set.
func (o *OnCallPageLinkAttributes) GetParentEntityTypeOk() (*OnCallPageLinkEntityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentEntityType, true
}

// SetParentEntityType sets field value.
func (o *OnCallPageLinkAttributes) SetParentEntityType(v OnCallPageLinkEntityType) {
	o.ParentEntityType = v
}

// GetRelationship returns the Relationship field value.
func (o *OnCallPageLinkAttributes) GetRelationship() OnCallPageLinkRelationshipType {
	if o == nil {
		var ret OnCallPageLinkRelationshipType
		return ret
	}
	return o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value
// and a boolean to check if the value has been set.
func (o *OnCallPageLinkAttributes) GetRelationshipOk() (*OnCallPageLinkRelationshipType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationship, true
}

// SetRelationship sets field value.
func (o *OnCallPageLinkAttributes) SetRelationship(v OnCallPageLinkRelationshipType) {
	o.Relationship = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallPageLinkAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["child_entity_id"] = o.ChildEntityId
	toSerialize["child_entity_type"] = o.ChildEntityType
	toSerialize["parent_entity_id"] = o.ParentEntityId
	toSerialize["parent_entity_type"] = o.ParentEntityType
	toSerialize["relationship"] = o.Relationship

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnCallPageLinkAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChildEntityId    *string                         `json:"child_entity_id"`
		ChildEntityType  *OnCallPageLinkEntityType       `json:"child_entity_type"`
		ParentEntityId   *string                         `json:"parent_entity_id"`
		ParentEntityType *OnCallPageLinkEntityType       `json:"parent_entity_type"`
		Relationship     *OnCallPageLinkRelationshipType `json:"relationship"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ChildEntityId == nil {
		return fmt.Errorf("required field child_entity_id missing")
	}
	if all.ChildEntityType == nil {
		return fmt.Errorf("required field child_entity_type missing")
	}
	if all.ParentEntityId == nil {
		return fmt.Errorf("required field parent_entity_id missing")
	}
	if all.ParentEntityType == nil {
		return fmt.Errorf("required field parent_entity_type missing")
	}
	if all.Relationship == nil {
		return fmt.Errorf("required field relationship missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"child_entity_id", "child_entity_type", "parent_entity_id", "parent_entity_type", "relationship"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ChildEntityId = *all.ChildEntityId
	if !all.ChildEntityType.IsValid() {
		hasInvalidField = true
	} else {
		o.ChildEntityType = *all.ChildEntityType
	}
	o.ParentEntityId = *all.ParentEntityId
	if !all.ParentEntityType.IsValid() {
		hasInvalidField = true
	} else {
		o.ParentEntityType = *all.ParentEntityType
	}
	if !all.Relationship.IsValid() {
		hasInvalidField = true
	} else {
		o.Relationship = *all.Relationship
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
