/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// TimeseriesWidgetExpressionAlias Define an expression alias.
type TimeseriesWidgetExpressionAlias struct {
	// Expression alias.
	AliasName *string `json:"alias_name,omitempty"`
	// Expression name.
	Expression string `json:"expression"`
}

// NewTimeseriesWidgetExpressionAlias instantiates a new TimeseriesWidgetExpressionAlias object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeseriesWidgetExpressionAlias(expression string) *TimeseriesWidgetExpressionAlias {
	this := TimeseriesWidgetExpressionAlias{}
	this.Expression = expression
	return &this
}

// NewTimeseriesWidgetExpressionAliasWithDefaults instantiates a new TimeseriesWidgetExpressionAlias object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeseriesWidgetExpressionAliasWithDefaults() *TimeseriesWidgetExpressionAlias {
	this := TimeseriesWidgetExpressionAlias{}
	return &this
}

// GetAliasName returns the AliasName field value if set, zero value otherwise.
func (o *TimeseriesWidgetExpressionAlias) GetAliasName() string {
	if o == nil || o.AliasName == nil {
		var ret string
		return ret
	}
	return *o.AliasName
}

// GetAliasNameOk returns a tuple with the AliasName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetExpressionAlias) GetAliasNameOk() (*string, bool) {
	if o == nil || o.AliasName == nil {
		return nil, false
	}
	return o.AliasName, true
}

// HasAliasName returns a boolean if a field has been set.
func (o *TimeseriesWidgetExpressionAlias) HasAliasName() bool {
	if o != nil && o.AliasName != nil {
		return true
	}

	return false
}

// SetAliasName gets a reference to the given string and assigns it to the AliasName field.
func (o *TimeseriesWidgetExpressionAlias) SetAliasName(v string) {
	o.AliasName = &v
}

// GetExpression returns the Expression field value
func (o *TimeseriesWidgetExpressionAlias) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetExpressionAlias) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *TimeseriesWidgetExpressionAlias) SetExpression(v string) {
	o.Expression = v
}

func (o TimeseriesWidgetExpressionAlias) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AliasName != nil {
		toSerialize["alias_name"] = o.AliasName
	}
	if true {
		toSerialize["expression"] = o.Expression
	}
	return json.Marshal(toSerialize)
}

func (o *TimeseriesWidgetExpressionAlias) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Expression *string `json:"expression"`
	}{}
	all := struct {
		AliasName  *string `json:"alias_name,omitempty"}`
		Expression string  `json:"expression"}`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Expression == nil {
		return fmt.Errorf("Required field expression missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.AliasName = all.AliasName
	o.Expression = all.Expression
	return nil
}

type NullableTimeseriesWidgetExpressionAlias struct {
	value *TimeseriesWidgetExpressionAlias
	isSet bool
}

func (v NullableTimeseriesWidgetExpressionAlias) Get() *TimeseriesWidgetExpressionAlias {
	return v.value
}

func (v *NullableTimeseriesWidgetExpressionAlias) Set(val *TimeseriesWidgetExpressionAlias) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeseriesWidgetExpressionAlias) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeseriesWidgetExpressionAlias) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeseriesWidgetExpressionAlias(val *TimeseriesWidgetExpressionAlias) *NullableTimeseriesWidgetExpressionAlias {
	return &NullableTimeseriesWidgetExpressionAlias{value: val, isSet: true}
}

func (v NullableTimeseriesWidgetExpressionAlias) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeseriesWidgetExpressionAlias) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
