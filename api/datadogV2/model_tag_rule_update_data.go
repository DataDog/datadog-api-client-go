// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagRuleUpdateData Data object for updating a tag rule.
type TagRuleUpdateData struct {
	// Mutable attributes of a tag rule. Each field is optional; omitting a field leaves its
	// current value unchanged. The `source` of a rule cannot be changed.
	Attributes *TagRuleUpdateAttributes `json:"attributes,omitempty"`
	// The unique identifier of the tag rule being updated.
	Id string `json:"id"`
	// JSON:API resource type for a tag rule.
	Type TagRuleResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagRuleUpdateData instantiates a new TagRuleUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagRuleUpdateData(id string, typeVar TagRuleResourceType) *TagRuleUpdateData {
	this := TagRuleUpdateData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewTagRuleUpdateDataWithDefaults instantiates a new TagRuleUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagRuleUpdateDataWithDefaults() *TagRuleUpdateData {
	this := TagRuleUpdateData{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TagRuleUpdateData) GetAttributes() TagRuleUpdateAttributes {
	if o == nil || o.Attributes == nil {
		var ret TagRuleUpdateAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagRuleUpdateData) GetAttributesOk() (*TagRuleUpdateAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TagRuleUpdateData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given TagRuleUpdateAttributes and assigns it to the Attributes field.
func (o *TagRuleUpdateData) SetAttributes(v TagRuleUpdateAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *TagRuleUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TagRuleUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *TagRuleUpdateData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *TagRuleUpdateData) GetType() TagRuleResourceType {
	if o == nil {
		var ret TagRuleResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TagRuleUpdateData) GetTypeOk() (*TagRuleResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TagRuleUpdateData) SetType(v TagRuleResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagRuleUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagRuleUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *TagRuleUpdateAttributes `json:"attributes,omitempty"`
		Id         *string                  `json:"id"`
		Type       *TagRuleResourceType     `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
