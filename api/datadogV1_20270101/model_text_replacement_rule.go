// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TextReplacementRule Test-only generated schema for the dashboard versioning proof of concept.
type TextReplacementRule struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Match TextMatch `json:"match"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Replace *TextReplacementRule20270101Replace `json:"replace,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTextReplacementRule instantiates a new TextReplacementRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTextReplacementRule(match TextMatch) *TextReplacementRule {
	this := TextReplacementRule{}
	this.Match = match
	return &this
}

// NewTextReplacementRuleWithDefaults instantiates a new TextReplacementRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTextReplacementRuleWithDefaults() *TextReplacementRule {
	this := TextReplacementRule{}
	return &this
}

// GetMatch returns the Match field value.
func (o *TextReplacementRule) GetMatch() TextMatch {
	if o == nil {
		var ret TextMatch
		return ret
	}
	return o.Match
}

// GetMatchOk returns a tuple with the Match field value
// and a boolean to check if the value has been set.
func (o *TextReplacementRule) GetMatchOk() (*TextMatch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Match, true
}

// SetMatch sets field value.
func (o *TextReplacementRule) SetMatch(v TextMatch) {
	o.Match = v
}

// GetReplace returns the Replace field value if set, zero value otherwise.
func (o *TextReplacementRule) GetReplace() TextReplacementRule20270101Replace {
	if o == nil || o.Replace == nil {
		var ret TextReplacementRule20270101Replace
		return ret
	}
	return *o.Replace
}

// GetReplaceOk returns a tuple with the Replace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextReplacementRule) GetReplaceOk() (*TextReplacementRule20270101Replace, bool) {
	if o == nil || o.Replace == nil {
		return nil, false
	}
	return o.Replace, true
}

// HasReplace returns a boolean if a field has been set.
func (o *TextReplacementRule) HasReplace() bool {
	return o != nil && o.Replace != nil
}

// SetReplace gets a reference to the given TextReplacementRule20270101Replace and assigns it to the Replace field.
func (o *TextReplacementRule) SetReplace(v TextReplacementRule20270101Replace) {
	o.Replace = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TextReplacementRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["match"] = o.Match
	if o.Replace != nil {
		toSerialize["replace"] = o.Replace
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TextReplacementRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Match   *TextMatch                          `json:"match"`
		Replace *TextReplacementRule20270101Replace `json:"replace,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Match == nil {
		return fmt.Errorf("required field match missing")
	}
	o.Match = *all.Match
	o.Replace = all.Replace

	return nil
}
