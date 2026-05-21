// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentCaseLinkDataAttributes Attributes of a case link.
type IncidentCaseLinkDataAttributes struct {
	// The entity identifier.
	EntityId string `json:"entity_id"`
	// Whether this is a page link.
	IsPage bool `json:"is_page"`
	// The relationship type.
	Relationship string `json:"relationship"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentCaseLinkDataAttributes instantiates a new IncidentCaseLinkDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentCaseLinkDataAttributes(entityId string, isPage bool, relationship string) *IncidentCaseLinkDataAttributes {
	this := IncidentCaseLinkDataAttributes{}
	this.EntityId = entityId
	this.IsPage = isPage
	this.Relationship = relationship
	return &this
}

// NewIncidentCaseLinkDataAttributesWithDefaults instantiates a new IncidentCaseLinkDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentCaseLinkDataAttributesWithDefaults() *IncidentCaseLinkDataAttributes {
	this := IncidentCaseLinkDataAttributes{}
	return &this
}

// GetEntityId returns the EntityId field value.
func (o *IncidentCaseLinkDataAttributes) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *IncidentCaseLinkDataAttributes) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value.
func (o *IncidentCaseLinkDataAttributes) SetEntityId(v string) {
	o.EntityId = v
}

// GetIsPage returns the IsPage field value.
func (o *IncidentCaseLinkDataAttributes) GetIsPage() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsPage
}

// GetIsPageOk returns a tuple with the IsPage field value
// and a boolean to check if the value has been set.
func (o *IncidentCaseLinkDataAttributes) GetIsPageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPage, true
}

// SetIsPage sets field value.
func (o *IncidentCaseLinkDataAttributes) SetIsPage(v bool) {
	o.IsPage = v
}

// GetRelationship returns the Relationship field value.
func (o *IncidentCaseLinkDataAttributes) GetRelationship() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value
// and a boolean to check if the value has been set.
func (o *IncidentCaseLinkDataAttributes) GetRelationshipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationship, true
}

// SetRelationship sets field value.
func (o *IncidentCaseLinkDataAttributes) SetRelationship(v string) {
	o.Relationship = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentCaseLinkDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["entity_id"] = o.EntityId
	toSerialize["is_page"] = o.IsPage
	toSerialize["relationship"] = o.Relationship

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentCaseLinkDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EntityId     *string `json:"entity_id"`
		IsPage       *bool   `json:"is_page"`
		Relationship *string `json:"relationship"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EntityId == nil {
		return fmt.Errorf("required field entity_id missing")
	}
	if all.IsPage == nil {
		return fmt.Errorf("required field is_page missing")
	}
	if all.Relationship == nil {
		return fmt.Errorf("required field relationship missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"entity_id", "is_page", "relationship"})
	} else {
		return err
	}
	o.EntityId = *all.EntityId
	o.IsPage = *all.IsPage
	o.Relationship = *all.Relationship

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
