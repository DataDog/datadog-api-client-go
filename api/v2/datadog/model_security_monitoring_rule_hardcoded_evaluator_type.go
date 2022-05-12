/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
	"fmt"
)

// SecurityMonitoringRuleHardcodedEvaluatorType Hardcoded evaluator type.
type SecurityMonitoringRuleHardcodedEvaluatorType string

// List of SecurityMonitoringRuleHardcodedEvaluatorType
const (
	SECURITYMONITORINGRULEHARDCODEDEVALUATORTYPE_LOG4SHELL SecurityMonitoringRuleHardcodedEvaluatorType = "log4shell"
)

var allowedSecurityMonitoringRuleHardcodedEvaluatorTypeEnumValues = []SecurityMonitoringRuleHardcodedEvaluatorType{
	SECURITYMONITORINGRULEHARDCODEDEVALUATORTYPE_LOG4SHELL,
}

func (w *SecurityMonitoringRuleHardcodedEvaluatorType) GetAllowedValues() []SecurityMonitoringRuleHardcodedEvaluatorType {
	return allowedSecurityMonitoringRuleHardcodedEvaluatorTypeEnumValues
}

func (v *SecurityMonitoringRuleHardcodedEvaluatorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleHardcodedEvaluatorType(value)
	return nil
}

// NewSecurityMonitoringRuleHardcodedEvaluatorTypeFromValue returns a pointer to a valid SecurityMonitoringRuleHardcodedEvaluatorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityMonitoringRuleHardcodedEvaluatorTypeFromValue(v string) (*SecurityMonitoringRuleHardcodedEvaluatorType, error) {
	ev := SecurityMonitoringRuleHardcodedEvaluatorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleHardcodedEvaluatorType: valid values are %v", v, allowedSecurityMonitoringRuleHardcodedEvaluatorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityMonitoringRuleHardcodedEvaluatorType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleHardcodedEvaluatorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleHardcodedEvaluatorType value
func (v SecurityMonitoringRuleHardcodedEvaluatorType) Ptr() *SecurityMonitoringRuleHardcodedEvaluatorType {
	return &v
}

type NullableSecurityMonitoringRuleHardcodedEvaluatorType struct {
	value *SecurityMonitoringRuleHardcodedEvaluatorType
	isSet bool
}

func (v NullableSecurityMonitoringRuleHardcodedEvaluatorType) Get() *SecurityMonitoringRuleHardcodedEvaluatorType {
	return v.value
}

func (v *NullableSecurityMonitoringRuleHardcodedEvaluatorType) Set(val *SecurityMonitoringRuleHardcodedEvaluatorType) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityMonitoringRuleHardcodedEvaluatorType) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityMonitoringRuleHardcodedEvaluatorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityMonitoringRuleHardcodedEvaluatorType(val *SecurityMonitoringRuleHardcodedEvaluatorType) *NullableSecurityMonitoringRuleHardcodedEvaluatorType {
	return &NullableSecurityMonitoringRuleHardcodedEvaluatorType{value: val, isSet: true}
}

func (v NullableSecurityMonitoringRuleHardcodedEvaluatorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityMonitoringRuleHardcodedEvaluatorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
