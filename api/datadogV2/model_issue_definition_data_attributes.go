// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueDefinitionDataAttributes Attributes of a single End User Device Monitoring issue definition.
type IssueDefinitionDataAttributes struct {
	// Category of the issue (for example, `battery`, `network`, or `performance`).
	Category string `json:"category"`
	// Human-readable label describing the issue, suitable for display in the Datadog UI.
	Label string `json:"label"`
	// Severity level of the issue (for example, `warning` or `critical`).
	Level string `json:"level"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueDefinitionDataAttributes instantiates a new IssueDefinitionDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueDefinitionDataAttributes(category string, label string, level string) *IssueDefinitionDataAttributes {
	this := IssueDefinitionDataAttributes{}
	this.Category = category
	this.Label = label
	this.Level = level
	return &this
}

// NewIssueDefinitionDataAttributesWithDefaults instantiates a new IssueDefinitionDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueDefinitionDataAttributesWithDefaults() *IssueDefinitionDataAttributes {
	this := IssueDefinitionDataAttributes{}
	return &this
}

// GetCategory returns the Category field value.
func (o *IssueDefinitionDataAttributes) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *IssueDefinitionDataAttributes) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *IssueDefinitionDataAttributes) SetCategory(v string) {
	o.Category = v
}

// GetLabel returns the Label field value.
func (o *IssueDefinitionDataAttributes) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *IssueDefinitionDataAttributes) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value.
func (o *IssueDefinitionDataAttributes) SetLabel(v string) {
	o.Label = v
}

// GetLevel returns the Level field value.
func (o *IssueDefinitionDataAttributes) GetLevel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *IssueDefinitionDataAttributes) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value.
func (o *IssueDefinitionDataAttributes) SetLevel(v string) {
	o.Level = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueDefinitionDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["category"] = o.Category
	toSerialize["label"] = o.Label
	toSerialize["level"] = o.Level

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueDefinitionDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category *string `json:"category"`
		Label    *string `json:"label"`
		Level    *string `json:"level"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.Label == nil {
		return fmt.Errorf("required field label missing")
	}
	if all.Level == nil {
		return fmt.Errorf("required field level missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "label", "level"})
	} else {
		return err
	}
	o.Category = *all.Category
	o.Label = *all.Label
	o.Level = *all.Level

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
