// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetMultipleRulesetsResponseData The primary data object in the get-multiple-rulesets response, containing the response attributes and resource type.
type GetMultipleRulesetsResponseData struct {
	// The attributes of the get-multiple-rulesets response, containing the list of requested rulesets.
	Attributes GetMultipleRulesetsResponseDataAttributes `json:"attributes"`
	// The unique identifier of the get-multiple-rulesets response resource, echoed from the request.
	Id string `json:"id"`
	// Get multiple rulesets response resource type.
	Type GetMultipleRulesetsResponseDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetMultipleRulesetsResponseData instantiates a new GetMultipleRulesetsResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetMultipleRulesetsResponseData(attributes GetMultipleRulesetsResponseDataAttributes, id string, typeVar GetMultipleRulesetsResponseDataType) *GetMultipleRulesetsResponseData {
	this := GetMultipleRulesetsResponseData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewGetMultipleRulesetsResponseDataWithDefaults instantiates a new GetMultipleRulesetsResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetMultipleRulesetsResponseDataWithDefaults() *GetMultipleRulesetsResponseData {
	this := GetMultipleRulesetsResponseData{}
	var typeVar GetMultipleRulesetsResponseDataType = GETMULTIPLERULESETSRESPONSEDATATYPE_GET_MULTIPLE_RULESETS_RESPONSE
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *GetMultipleRulesetsResponseData) GetAttributes() GetMultipleRulesetsResponseDataAttributes {
	if o == nil {
		var ret GetMultipleRulesetsResponseDataAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseData) GetAttributesOk() (*GetMultipleRulesetsResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *GetMultipleRulesetsResponseData) SetAttributes(v GetMultipleRulesetsResponseDataAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *GetMultipleRulesetsResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *GetMultipleRulesetsResponseData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *GetMultipleRulesetsResponseData) GetType() GetMultipleRulesetsResponseDataType {
	if o == nil {
		var ret GetMultipleRulesetsResponseDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseData) GetTypeOk() (*GetMultipleRulesetsResponseDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GetMultipleRulesetsResponseData) SetType(v GetMultipleRulesetsResponseDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetMultipleRulesetsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetMultipleRulesetsResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *GetMultipleRulesetsResponseDataAttributes `json:"attributes"`
		Id         *string                                    `json:"id"`
		Type       *GetMultipleRulesetsResponseDataType       `json:"type"`
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
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
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
