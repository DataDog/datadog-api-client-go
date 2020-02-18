/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// LogsGrokParserRules Set of rules for the grok parser
type LogsGrokParserRules struct {
	// List of match rules for the grok parser, separated by a new line
	MatchRules string `json:"match_rules"`
	// List of support rules for the grok parser, separated by a new line
	SupportRules *string `json:"support_rules,omitempty"`
}

// NewLogsGrokParserRules instantiates a new LogsGrokParserRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsGrokParserRules(matchRules string) *LogsGrokParserRules {
	this := LogsGrokParserRules{}
	this.MatchRules = matchRules
	var supportRules string = ""
	this.SupportRules = &supportRules
	return &this
}

// NewLogsGrokParserRulesWithDefaults instantiates a new LogsGrokParserRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsGrokParserRulesWithDefaults() *LogsGrokParserRules {
	this := LogsGrokParserRules{}
	var supportRules string = ""
	this.SupportRules = &supportRules
	return &this
}

// GetMatchRules returns the MatchRules field value
func (o *LogsGrokParserRules) GetMatchRules() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MatchRules
}

// SetMatchRules sets field value
func (o *LogsGrokParserRules) SetMatchRules(v string) {
	o.MatchRules = v
}

// GetSupportRules returns the SupportRules field value if set, zero value otherwise.
func (o *LogsGrokParserRules) GetSupportRules() string {
	if o == nil || o.SupportRules == nil {
		var ret string
		return ret
	}
	return *o.SupportRules
}

// GetSupportRulesOk returns a tuple with the SupportRules field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LogsGrokParserRules) GetSupportRulesOk() (string, bool) {
	if o == nil || o.SupportRules == nil {
		var ret string
		return ret, false
	}
	return *o.SupportRules, true
}

// HasSupportRules returns a boolean if a field has been set.
func (o *LogsGrokParserRules) HasSupportRules() bool {
	if o != nil && o.SupportRules != nil {
		return true
	}

	return false
}

// SetSupportRules gets a reference to the given string and assigns it to the SupportRules field.
func (o *LogsGrokParserRules) SetSupportRules(v string) {
	o.SupportRules = &v
}

type NullableLogsGrokParserRules struct {
	Value        LogsGrokParserRules
	ExplicitNull bool
}

func (v NullableLogsGrokParserRules) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLogsGrokParserRules) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
