// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityExclusionFilterListRulesTargetTags Tags identifying the category and type of the targeted rule.
type ApplicationSecurityExclusionFilterListRulesTargetTags struct {
	// The category of the rule.
	Category *string `json:"category,omitempty"`
	// The type of the rule.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityExclusionFilterListRulesTargetTags instantiates a new ApplicationSecurityExclusionFilterListRulesTargetTags object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityExclusionFilterListRulesTargetTags() *ApplicationSecurityExclusionFilterListRulesTargetTags {
	this := ApplicationSecurityExclusionFilterListRulesTargetTags{}
	return &this
}

// NewApplicationSecurityExclusionFilterListRulesTargetTagsWithDefaults instantiates a new ApplicationSecurityExclusionFilterListRulesTargetTags object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityExclusionFilterListRulesTargetTagsWithDefaults() *ApplicationSecurityExclusionFilterListRulesTargetTags {
	this := ApplicationSecurityExclusionFilterListRulesTargetTags{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterListRulesTargetTags) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterListRulesTargetTags) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterListRulesTargetTags) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ApplicationSecurityExclusionFilterListRulesTargetTags) SetCategory(v string) {
	o.Category = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterListRulesTargetTags) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterListRulesTargetTags) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterListRulesTargetTags) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApplicationSecurityExclusionFilterListRulesTargetTags) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityExclusionFilterListRulesTargetTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
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
func (o *ApplicationSecurityExclusionFilterListRulesTargetTags) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category *string `json:"category,omitempty"`
		Type     *string `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "type"})
	} else {
		return err
	}
	o.Category = all.Category
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
