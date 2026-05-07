// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagPipelinesRulesetStatusData Tag pipeline ruleset status data.
type TagPipelinesRulesetStatusData struct {
	// Attributes for a tag pipeline ruleset status.
	Attributes TagPipelinesRulesetStatusAttributes `json:"attributes"`
	// The unique identifier of the ruleset.
	Id string `json:"id"`
	// Tag pipeline ruleset status resource type.
	Type TagPipelinesRulesetStatusType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagPipelinesRulesetStatusData instantiates a new TagPipelinesRulesetStatusData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagPipelinesRulesetStatusData(attributes TagPipelinesRulesetStatusAttributes, id string, typeVar TagPipelinesRulesetStatusType) *TagPipelinesRulesetStatusData {
	this := TagPipelinesRulesetStatusData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewTagPipelinesRulesetStatusDataWithDefaults instantiates a new TagPipelinesRulesetStatusData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagPipelinesRulesetStatusDataWithDefaults() *TagPipelinesRulesetStatusData {
	this := TagPipelinesRulesetStatusData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *TagPipelinesRulesetStatusData) GetAttributes() TagPipelinesRulesetStatusAttributes {
	if o == nil {
		var ret TagPipelinesRulesetStatusAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TagPipelinesRulesetStatusData) GetAttributesOk() (*TagPipelinesRulesetStatusAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *TagPipelinesRulesetStatusData) SetAttributes(v TagPipelinesRulesetStatusAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *TagPipelinesRulesetStatusData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TagPipelinesRulesetStatusData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *TagPipelinesRulesetStatusData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *TagPipelinesRulesetStatusData) GetType() TagPipelinesRulesetStatusType {
	if o == nil {
		var ret TagPipelinesRulesetStatusType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TagPipelinesRulesetStatusData) GetTypeOk() (*TagPipelinesRulesetStatusType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TagPipelinesRulesetStatusData) SetType(v TagPipelinesRulesetStatusType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagPipelinesRulesetStatusData) MarshalJSON() ([]byte, error) {
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
func (o *TagPipelinesRulesetStatusData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *TagPipelinesRulesetStatusAttributes `json:"attributes"`
		Id         *string                              `json:"id"`
		Type       *TagPipelinesRulesetStatusType       `json:"type"`
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
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
