// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV2

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"

)


// CIAppPipelineEvent Object description of a pipeline event after being processed and stored by Datadog.
type CIAppPipelineEvent struct {
	// JSON object containing all event attributes and their associated values.
	Attributes *CIAppPipelineEventAttributes `json:"attributes,omitempty"`
	// Unique ID of the event.
	Id *string `json:"id,omitempty"`
	// Type of the event.
	Type *CIAppPipelineEventTypeName `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewCIAppPipelineEvent instantiates a new CIAppPipelineEvent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCIAppPipelineEvent() *CIAppPipelineEvent {
	this := CIAppPipelineEvent{}
	return &this
}

// NewCIAppPipelineEventWithDefaults instantiates a new CIAppPipelineEvent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCIAppPipelineEventWithDefaults() *CIAppPipelineEvent {
	this := CIAppPipelineEvent{}
	return &this
}
// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CIAppPipelineEvent) GetAttributes() CIAppPipelineEventAttributes {
	if o == nil || o.Attributes == nil {
		var ret CIAppPipelineEventAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEvent) GetAttributesOk() (*CIAppPipelineEventAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CIAppPipelineEvent) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given CIAppPipelineEventAttributes and assigns it to the Attributes field.
func (o *CIAppPipelineEvent) SetAttributes(v CIAppPipelineEventAttributes) {
	o.Attributes = &v
}


// GetId returns the Id field value if set, zero value otherwise.
func (o *CIAppPipelineEvent) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEvent) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CIAppPipelineEvent) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CIAppPipelineEvent) SetId(v string) {
	o.Id = &v
}


// GetType returns the Type field value if set, zero value otherwise.
func (o *CIAppPipelineEvent) GetType() CIAppPipelineEventTypeName {
	if o == nil || o.Type == nil {
		var ret CIAppPipelineEventTypeName
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEvent) GetTypeOk() (*CIAppPipelineEventTypeName, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CIAppPipelineEvent) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given CIAppPipelineEventTypeName and assigns it to the Type field.
func (o *CIAppPipelineEvent) SetType(v CIAppPipelineEventTypeName) {
	o.Type = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o CIAppPipelineEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CIAppPipelineEvent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CIAppPipelineEventAttributes `json:"attributes,omitempty"`
		Id *string `json:"id,omitempty"`
		Type *CIAppPipelineEventTypeName `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{ "attributes", "id", "type",  })
	} else {
		return err
	}

	hasInvalidField := false
	if  all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Type != nil &&!all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
