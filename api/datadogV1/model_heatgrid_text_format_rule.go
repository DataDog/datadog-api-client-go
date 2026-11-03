// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridTextFormatRule A rule for matching and optionally replacing a group label.
type HeatgridTextFormatRule struct {
	// Match a group label against a string.
	Match HeatgridTextFormatMatch `json:"match"`
	// Replace all or part of a matching group label.
	Replace *HeatgridTextFormatReplace `json:"replace,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridTextFormatRule instantiates a new HeatgridTextFormatRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridTextFormatRule(match HeatgridTextFormatMatch) *HeatgridTextFormatRule {
	this := HeatgridTextFormatRule{}
	this.Match = match
	return &this
}

// NewHeatgridTextFormatRuleWithDefaults instantiates a new HeatgridTextFormatRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridTextFormatRuleWithDefaults() *HeatgridTextFormatRule {
	this := HeatgridTextFormatRule{}
	return &this
}

// GetMatch returns the Match field value.
func (o *HeatgridTextFormatRule) GetMatch() HeatgridTextFormatMatch {
	if o == nil {
		var ret HeatgridTextFormatMatch
		return ret
	}
	return o.Match
}

// GetMatchOk returns a tuple with the Match field value
// and a boolean to check if the value has been set.
func (o *HeatgridTextFormatRule) GetMatchOk() (*HeatgridTextFormatMatch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Match, true
}

// SetMatch sets field value.
func (o *HeatgridTextFormatRule) SetMatch(v HeatgridTextFormatMatch) {
	o.Match = v
}

// GetReplace returns the Replace field value if set, zero value otherwise.
func (o *HeatgridTextFormatRule) GetReplace() HeatgridTextFormatReplace {
	if o == nil || o.Replace == nil {
		var ret HeatgridTextFormatReplace
		return ret
	}
	return *o.Replace
}

// GetReplaceOk returns a tuple with the Replace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridTextFormatRule) GetReplaceOk() (*HeatgridTextFormatReplace, bool) {
	if o == nil || o.Replace == nil {
		return nil, false
	}
	return o.Replace, true
}

// HasReplace returns a boolean if a field has been set.
func (o *HeatgridTextFormatRule) HasReplace() bool {
	return o != nil && o.Replace != nil
}

// SetReplace gets a reference to the given HeatgridTextFormatReplace and assigns it to the Replace field.
func (o *HeatgridTextFormatRule) SetReplace(v HeatgridTextFormatReplace) {
	o.Replace = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridTextFormatRule) MarshalJSON() ([]byte, error) {
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
func (o *HeatgridTextFormatRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Match   *HeatgridTextFormatMatch   `json:"match"`
		Replace *HeatgridTextFormatReplace `json:"replace,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Match == nil {
		return fmt.Errorf("required field match missing")
	}

	hasInvalidField := false
	if all.Match.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Match = *all.Match
	o.Replace = all.Replace

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
