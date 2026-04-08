// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventEmailAddressData A single event email address resource.
type EventEmailAddressData struct {
	// Attributes of an event email address resource.
	Attributes EventEmailAddressResponseAttributes `json:"attributes"`
	// The UUID of the event email address.
	Id string `json:"id"`
	// Relationships associated with an event email address resource.
	Relationships EventEmailAddressRelationships `json:"relationships"`
	// The type of the resource. Must be `event_emails`.
	Type EventEmailAddressResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventEmailAddressData instantiates a new EventEmailAddressData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventEmailAddressData(attributes EventEmailAddressResponseAttributes, id string, relationships EventEmailAddressRelationships, typeVar EventEmailAddressResourceType) *EventEmailAddressData {
	this := EventEmailAddressData{}
	this.Attributes = attributes
	this.Id = id
	this.Relationships = relationships
	this.Type = typeVar
	return &this
}

// NewEventEmailAddressDataWithDefaults instantiates a new EventEmailAddressData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventEmailAddressDataWithDefaults() *EventEmailAddressData {
	this := EventEmailAddressData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *EventEmailAddressData) GetAttributes() EventEmailAddressResponseAttributes {
	if o == nil {
		var ret EventEmailAddressResponseAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressData) GetAttributesOk() (*EventEmailAddressResponseAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *EventEmailAddressData) SetAttributes(v EventEmailAddressResponseAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *EventEmailAddressData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *EventEmailAddressData) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value.
func (o *EventEmailAddressData) GetRelationships() EventEmailAddressRelationships {
	if o == nil {
		var ret EventEmailAddressRelationships
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressData) GetRelationshipsOk() (*EventEmailAddressRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value.
func (o *EventEmailAddressData) SetRelationships(v EventEmailAddressRelationships) {
	o.Relationships = v
}

// GetType returns the Type field value.
func (o *EventEmailAddressData) GetType() EventEmailAddressResourceType {
	if o == nil {
		var ret EventEmailAddressResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressData) GetTypeOk() (*EventEmailAddressResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *EventEmailAddressData) SetType(v EventEmailAddressResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventEmailAddressData) MarshalJSON() ([]byte, error) {
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
func (o *EventEmailAddressData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *EventEmailAddressResponseAttributes `json:"attributes"`
		Id            *string                              `json:"id"`
		Relationships *EventEmailAddressRelationships      `json:"relationships"`
		Type          *EventEmailAddressResourceType       `json:"type"`
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
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
