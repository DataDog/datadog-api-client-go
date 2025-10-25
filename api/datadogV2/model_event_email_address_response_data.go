// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventEmailAddressResponseData
type EventEmailAddressResponseData struct {
	//
	Attributes *EventEmailAddressResponseDataAttributes `json:"attributes,omitempty"`
	//
	Id string `json:"id"`
	//
	Relationships *EventEmailAddressResponseDataRelationships `json:"relationships,omitempty"`
	// Event emails resource type.
	Type EventEmailsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventEmailAddressResponseData instantiates a new EventEmailAddressResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventEmailAddressResponseData(id string, typeVar EventEmailsType) *EventEmailAddressResponseData {
	this := EventEmailAddressResponseData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewEventEmailAddressResponseDataWithDefaults instantiates a new EventEmailAddressResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventEmailAddressResponseDataWithDefaults() *EventEmailAddressResponseData {
	this := EventEmailAddressResponseData{}
	var typeVar EventEmailsType = EVENTEMAILSTYPE_EVENT_EMAILS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EventEmailAddressResponseData) GetAttributes() EventEmailAddressResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret EventEmailAddressResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseData) GetAttributesOk() (*EventEmailAddressResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EventEmailAddressResponseData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given EventEmailAddressResponseDataAttributes and assigns it to the Attributes field.
func (o *EventEmailAddressResponseData) SetAttributes(v EventEmailAddressResponseDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *EventEmailAddressResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *EventEmailAddressResponseData) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *EventEmailAddressResponseData) GetRelationships() EventEmailAddressResponseDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret EventEmailAddressResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseData) GetRelationshipsOk() (*EventEmailAddressResponseDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *EventEmailAddressResponseData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given EventEmailAddressResponseDataRelationships and assigns it to the Relationships field.
func (o *EventEmailAddressResponseData) SetRelationships(v EventEmailAddressResponseDataRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *EventEmailAddressResponseData) GetType() EventEmailsType {
	if o == nil {
		var ret EventEmailsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseData) GetTypeOk() (*EventEmailsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *EventEmailAddressResponseData) SetType(v EventEmailsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventEmailAddressResponseData) MarshalJSON() ([]byte, error) {
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
func (o *EventEmailAddressResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *EventEmailAddressResponseDataAttributes    `json:"attributes,omitempty"`
		Id            *string                                     `json:"id"`
		Relationships *EventEmailAddressResponseDataRelationships `json:"relationships,omitempty"`
		Type          *EventEmailsType                            `json:"type"`
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
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
