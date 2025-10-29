// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelSuggestionRequestData
type FunnelSuggestionRequestData struct {
	//
	Attributes *FunnelSuggestionRequestDataAttributes `json:"attributes,omitempty"`
	//
	Id *string `json:"id,omitempty"`
	//
	Type FunnelSuggestionRequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFunnelSuggestionRequestData instantiates a new FunnelSuggestionRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFunnelSuggestionRequestData(typeVar FunnelSuggestionRequestDataType) *FunnelSuggestionRequestData {
	this := FunnelSuggestionRequestData{}
	this.Type = typeVar
	return &this
}

// NewFunnelSuggestionRequestDataWithDefaults instantiates a new FunnelSuggestionRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFunnelSuggestionRequestDataWithDefaults() *FunnelSuggestionRequestData {
	this := FunnelSuggestionRequestData{}
	var typeVar FunnelSuggestionRequestDataType = FUNNELSUGGESTIONREQUESTDATATYPE_FUNNEL_SUGGESTION_REQUEST
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *FunnelSuggestionRequestData) GetAttributes() FunnelSuggestionRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret FunnelSuggestionRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelSuggestionRequestData) GetAttributesOk() (*FunnelSuggestionRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FunnelSuggestionRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given FunnelSuggestionRequestDataAttributes and assigns it to the Attributes field.
func (o *FunnelSuggestionRequestData) SetAttributes(v FunnelSuggestionRequestDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FunnelSuggestionRequestData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelSuggestionRequestData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FunnelSuggestionRequestData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FunnelSuggestionRequestData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *FunnelSuggestionRequestData) GetType() FunnelSuggestionRequestDataType {
	if o == nil {
		var ret FunnelSuggestionRequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FunnelSuggestionRequestData) GetTypeOk() (*FunnelSuggestionRequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *FunnelSuggestionRequestData) SetType(v FunnelSuggestionRequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FunnelSuggestionRequestData) MarshalJSON() ([]byte, error) {
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
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FunnelSuggestionRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *FunnelSuggestionRequestDataAttributes `json:"attributes,omitempty"`
		Id         *string                                `json:"id,omitempty"`
		Type       *FunnelSuggestionRequestDataType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
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
