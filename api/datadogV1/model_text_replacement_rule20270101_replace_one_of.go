// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TextReplacementRule20270101ReplaceOneOf Test-only generated schema for the dashboard versioning proof of concept.
type TextReplacementRule20270101ReplaceOneOf struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Substring string `json:"substring"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type TextReplacementRuleReplaceAnyOfAnyOfItemType_20270101 `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	With string `json:"with"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTextReplacementRule20270101ReplaceOneOf instantiates a new TextReplacementRule20270101ReplaceOneOf object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTextReplacementRule20270101ReplaceOneOf(substring string, typeVar TextReplacementRuleReplaceAnyOfAnyOfItemType_20270101, with string) *TextReplacementRule20270101ReplaceOneOf {
	this := TextReplacementRule20270101ReplaceOneOf{}
	this.Substring = substring
	this.Type = typeVar
	this.With = with
	return &this
}

// NewTextReplacementRule20270101ReplaceOneOfWithDefaults instantiates a new TextReplacementRule20270101ReplaceOneOf object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTextReplacementRule20270101ReplaceOneOfWithDefaults() *TextReplacementRule20270101ReplaceOneOf {
	this := TextReplacementRule20270101ReplaceOneOf{}
	return &this
}

// GetSubstring returns the Substring field value.
func (o *TextReplacementRule20270101ReplaceOneOf) GetSubstring() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Substring
}

// GetSubstringOk returns a tuple with the Substring field value
// and a boolean to check if the value has been set.
func (o *TextReplacementRule20270101ReplaceOneOf) GetSubstringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Substring, true
}

// SetSubstring sets field value.
func (o *TextReplacementRule20270101ReplaceOneOf) SetSubstring(v string) {
	o.Substring = v
}

// GetType returns the Type field value.
func (o *TextReplacementRule20270101ReplaceOneOf) GetType() TextReplacementRuleReplaceAnyOfAnyOfItemType_20270101 {
	if o == nil {
		var ret TextReplacementRuleReplaceAnyOfAnyOfItemType_20270101
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TextReplacementRule20270101ReplaceOneOf) GetTypeOk() (*TextReplacementRuleReplaceAnyOfAnyOfItemType_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TextReplacementRule20270101ReplaceOneOf) SetType(v TextReplacementRuleReplaceAnyOfAnyOfItemType_20270101) {
	o.Type = v
}

// GetWith returns the With field value.
func (o *TextReplacementRule20270101ReplaceOneOf) GetWith() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.With
}

// GetWithOk returns a tuple with the With field value
// and a boolean to check if the value has been set.
func (o *TextReplacementRule20270101ReplaceOneOf) GetWithOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.With, true
}

// SetWith sets field value.
func (o *TextReplacementRule20270101ReplaceOneOf) SetWith(v string) {
	o.With = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TextReplacementRule20270101ReplaceOneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["substring"] = o.Substring
	toSerialize["type"] = o.Type
	toSerialize["with"] = o.With
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TextReplacementRule20270101ReplaceOneOf) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Substring *string                                                `json:"substring"`
		Type      *TextReplacementRuleReplaceAnyOfAnyOfItemType_20270101 `json:"type"`
		With      *string                                                `json:"with"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Substring == nil {
		return fmt.Errorf("required field substring missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.With == nil {
		return fmt.Errorf("required field with missing")
	}

	hasInvalidField := false
	o.Substring = *all.Substring
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.With = *all.With

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
