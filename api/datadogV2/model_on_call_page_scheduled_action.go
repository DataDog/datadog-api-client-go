// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageScheduledAction The next action scheduled on a page, such as an escalation or an automatic resolution.
type OnCallPageScheduledAction struct {
	// The attributes of a scheduled action.
	Attributes *OnCallPageScheduledActionAttributes `json:"attributes,omitempty"`
	// The ID of the scheduled action. Equal to the page ID.
	Id string `json:"id"`
	// The relationships of a scheduled action.
	Relationships *OnCallPageScheduledActionRelationships `json:"relationships,omitempty"`
	// Scheduled action resource type.
	Type OnCallPageScheduledActionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallPageScheduledAction instantiates a new OnCallPageScheduledAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallPageScheduledAction(id string, typeVar OnCallPageScheduledActionType) *OnCallPageScheduledAction {
	this := OnCallPageScheduledAction{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewOnCallPageScheduledActionWithDefaults instantiates a new OnCallPageScheduledAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallPageScheduledActionWithDefaults() *OnCallPageScheduledAction {
	this := OnCallPageScheduledAction{}
	var typeVar OnCallPageScheduledActionType = ONCALLPAGESCHEDULEDACTIONTYPE_SCHEDULED_ACTIONS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *OnCallPageScheduledAction) GetAttributes() OnCallPageScheduledActionAttributes {
	if o == nil || o.Attributes == nil {
		var ret OnCallPageScheduledActionAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageScheduledAction) GetAttributesOk() (*OnCallPageScheduledActionAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *OnCallPageScheduledAction) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given OnCallPageScheduledActionAttributes and assigns it to the Attributes field.
func (o *OnCallPageScheduledAction) SetAttributes(v OnCallPageScheduledActionAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *OnCallPageScheduledAction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OnCallPageScheduledAction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *OnCallPageScheduledAction) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *OnCallPageScheduledAction) GetRelationships() OnCallPageScheduledActionRelationships {
	if o == nil || o.Relationships == nil {
		var ret OnCallPageScheduledActionRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageScheduledAction) GetRelationshipsOk() (*OnCallPageScheduledActionRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *OnCallPageScheduledAction) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given OnCallPageScheduledActionRelationships and assigns it to the Relationships field.
func (o *OnCallPageScheduledAction) SetRelationships(v OnCallPageScheduledActionRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *OnCallPageScheduledAction) GetType() OnCallPageScheduledActionType {
	if o == nil {
		var ret OnCallPageScheduledActionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OnCallPageScheduledAction) GetTypeOk() (*OnCallPageScheduledActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OnCallPageScheduledAction) SetType(v OnCallPageScheduledActionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallPageScheduledAction) MarshalJSON() ([]byte, error) {
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
func (o *OnCallPageScheduledAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *OnCallPageScheduledActionAttributes    `json:"attributes,omitempty"`
		Id            *string                                 `json:"id"`
		Relationships *OnCallPageScheduledActionRelationships `json:"relationships,omitempty"`
		Type          *OnCallPageScheduledActionType          `json:"type"`
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
