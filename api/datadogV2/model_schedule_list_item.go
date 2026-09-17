// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleListItem Represents a summary of a schedule, linking its core attributes and team relationships.
type ScheduleListItem struct {
	// Provides core properties of a schedule object such as its name and time zone.
	Attributes *ScheduleDataAttributes `json:"attributes,omitempty"`
	// The schedule's unique identifier.
	Id string `json:"id"`
	// Groups the relationships for a schedule summary, referencing its teams.
	Relationships *ScheduleListItemRelationships `json:"relationships,omitempty"`
	// Schedules resource type.
	Type ScheduleDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleListItem instantiates a new ScheduleListItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleListItem(id string, typeVar ScheduleDataType) *ScheduleListItem {
	this := ScheduleListItem{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewScheduleListItemWithDefaults instantiates a new ScheduleListItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleListItemWithDefaults() *ScheduleListItem {
	this := ScheduleListItem{}
	var typeVar ScheduleDataType = SCHEDULEDATATYPE_SCHEDULES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ScheduleListItem) GetAttributes() ScheduleDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret ScheduleDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleListItem) GetAttributesOk() (*ScheduleDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ScheduleListItem) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given ScheduleDataAttributes and assigns it to the Attributes field.
func (o *ScheduleListItem) SetAttributes(v ScheduleDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *ScheduleListItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScheduleListItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ScheduleListItem) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ScheduleListItem) GetRelationships() ScheduleListItemRelationships {
	if o == nil || o.Relationships == nil {
		var ret ScheduleListItemRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleListItem) GetRelationshipsOk() (*ScheduleListItemRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ScheduleListItem) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given ScheduleListItemRelationships and assigns it to the Relationships field.
func (o *ScheduleListItem) SetRelationships(v ScheduleListItemRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *ScheduleListItem) GetType() ScheduleDataType {
	if o == nil {
		var ret ScheduleDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScheduleListItem) GetTypeOk() (*ScheduleDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ScheduleListItem) SetType(v ScheduleDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
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
func (o *ScheduleListItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *ScheduleDataAttributes        `json:"attributes,omitempty"`
		Id            *string                        `json:"id"`
		Relationships *ScheduleListItemRelationships `json:"relationships,omitempty"`
		Type          *ScheduleDataType              `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
