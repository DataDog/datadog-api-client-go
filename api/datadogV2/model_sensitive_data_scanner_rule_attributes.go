// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
)

// SensitiveDataScannerRuleAttributes Attributes of the Sensitive Data Scanner rule.
type SensitiveDataScannerRuleAttributes struct {
	// Description of the rule.
	Description *string `json:"description,omitempty"`
	// Attributes excluded from the scan.
	ExcludedAttributes []string `json:"excluded_attributes,omitempty"`
	// Whether or not the rule is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Name of the rule.
	Name *string `json:"name,omitempty"`
	// Not included if there is a relationship to a standard pattern.
	Pattern *string `json:"pattern,omitempty"`
	// List of tags.
	Tags []string `json:"tags,omitempty"`
	// Object describing how the scanned event will be replaced.
	TextReplacement *SensitiveDataScannerTextReplacement `json:"text_replacement,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSensitiveDataScannerRuleAttributes instantiates a new SensitiveDataScannerRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSensitiveDataScannerRuleAttributes() *SensitiveDataScannerRuleAttributes {
	this := SensitiveDataScannerRuleAttributes{}
	return &this
}

// NewSensitiveDataScannerRuleAttributesWithDefaults instantiates a new SensitiveDataScannerRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSensitiveDataScannerRuleAttributesWithDefaults() *SensitiveDataScannerRuleAttributes {
	this := SensitiveDataScannerRuleAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SensitiveDataScannerRuleAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerRuleAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SensitiveDataScannerRuleAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SensitiveDataScannerRuleAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetExcludedAttributes returns the ExcludedAttributes field value if set, zero value otherwise.
func (o *SensitiveDataScannerRuleAttributes) GetExcludedAttributes() []string {
	if o == nil || o.ExcludedAttributes == nil {
		var ret []string
		return ret
	}
	return o.ExcludedAttributes
}

// GetExcludedAttributesOk returns a tuple with the ExcludedAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerRuleAttributes) GetExcludedAttributesOk() (*[]string, bool) {
	if o == nil || o.ExcludedAttributes == nil {
		return nil, false
	}
	return &o.ExcludedAttributes, true
}

// HasExcludedAttributes returns a boolean if a field has been set.
func (o *SensitiveDataScannerRuleAttributes) HasExcludedAttributes() bool {
	return o != nil && o.ExcludedAttributes != nil
}

// SetExcludedAttributes gets a reference to the given []string and assigns it to the ExcludedAttributes field.
func (o *SensitiveDataScannerRuleAttributes) SetExcludedAttributes(v []string) {
	o.ExcludedAttributes = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *SensitiveDataScannerRuleAttributes) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerRuleAttributes) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *SensitiveDataScannerRuleAttributes) HasIsEnabled() bool {
	return o != nil && o.IsEnabled != nil
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *SensitiveDataScannerRuleAttributes) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SensitiveDataScannerRuleAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerRuleAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SensitiveDataScannerRuleAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SensitiveDataScannerRuleAttributes) SetName(v string) {
	o.Name = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *SensitiveDataScannerRuleAttributes) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerRuleAttributes) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *SensitiveDataScannerRuleAttributes) HasPattern() bool {
	return o != nil && o.Pattern != nil
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *SensitiveDataScannerRuleAttributes) SetPattern(v string) {
	o.Pattern = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SensitiveDataScannerRuleAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerRuleAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SensitiveDataScannerRuleAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SensitiveDataScannerRuleAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTextReplacement returns the TextReplacement field value if set, zero value otherwise.
func (o *SensitiveDataScannerRuleAttributes) GetTextReplacement() SensitiveDataScannerTextReplacement {
	if o == nil || o.TextReplacement == nil {
		var ret SensitiveDataScannerTextReplacement
		return ret
	}
	return *o.TextReplacement
}

// GetTextReplacementOk returns a tuple with the TextReplacement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerRuleAttributes) GetTextReplacementOk() (*SensitiveDataScannerTextReplacement, bool) {
	if o == nil || o.TextReplacement == nil {
		return nil, false
	}
	return o.TextReplacement, true
}

// HasTextReplacement returns a boolean if a field has been set.
func (o *SensitiveDataScannerRuleAttributes) HasTextReplacement() bool {
	return o != nil && o.TextReplacement != nil
}

// SetTextReplacement gets a reference to the given SensitiveDataScannerTextReplacement and assigns it to the TextReplacement field.
func (o *SensitiveDataScannerRuleAttributes) SetTextReplacement(v SensitiveDataScannerTextReplacement) {
	o.TextReplacement = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SensitiveDataScannerRuleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExcludedAttributes != nil {
		toSerialize["excluded_attributes"] = o.ExcludedAttributes
	}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TextReplacement != nil {
		toSerialize["text_replacement"] = o.TextReplacement
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SensitiveDataScannerRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Description        *string                              `json:"description,omitempty"`
		ExcludedAttributes []string                             `json:"excluded_attributes,omitempty"`
		IsEnabled          *bool                                `json:"is_enabled,omitempty"`
		Name               *string                              `json:"name,omitempty"`
		Pattern            *string                              `json:"pattern,omitempty"`
		Tags               []string                             `json:"tags,omitempty"`
		TextReplacement    *SensitiveDataScannerTextReplacement `json:"text_replacement,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Description = all.Description
	o.ExcludedAttributes = all.ExcludedAttributes
	o.IsEnabled = all.IsEnabled
	o.Name = all.Name
	o.Pattern = all.Pattern
	o.Tags = all.Tags
	if all.TextReplacement != nil && all.TextReplacement.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.TextReplacement = all.TextReplacement
	return nil
}
