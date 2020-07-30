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

// LogStreamWidgetDefinitionType Type of the log stream widget.
type LogStreamWidgetDefinitionType string

// List of LogStreamWidgetDefinitionType
const (
	LOGSTREAMWIDGETDEFINITIONTYPE_LOG_STREAM LogStreamWidgetDefinitionType = "log_stream"
)

func (v *LogStreamWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogStreamWidgetDefinitionType(value)
	for _, existing := range []LogStreamWidgetDefinitionType{"log_stream"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogStreamWidgetDefinitionType", value)
}

// Ptr returns reference to LogStreamWidgetDefinitionType value
func (v LogStreamWidgetDefinitionType) Ptr() *LogStreamWidgetDefinitionType {
	return &v
}

type NullableLogStreamWidgetDefinitionType struct {
	value *LogStreamWidgetDefinitionType
	isSet bool
}

func (v NullableLogStreamWidgetDefinitionType) Get() *LogStreamWidgetDefinitionType {
	return v.value
}

func (v *NullableLogStreamWidgetDefinitionType) Set(val *LogStreamWidgetDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStreamWidgetDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStreamWidgetDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStreamWidgetDefinitionType(val *LogStreamWidgetDefinitionType) *NullableLogStreamWidgetDefinitionType {
	return &NullableLogStreamWidgetDefinitionType{value: val, isSet: true}
}

func (v NullableLogStreamWidgetDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStreamWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
