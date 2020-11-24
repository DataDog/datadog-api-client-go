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

// EventAlertType If an alert event is enabled, set its type. For example, `error`, `warning`, `info`, `success`, `user_update`, `recommendation`, and `snapshot`.
type EventAlertType string

// List of EventAlertType
const (
	EVENTALERTTYPE_ERROR          EventAlertType = "error"
	EVENTALERTTYPE_WARNING        EventAlertType = "warning"
	EVENTALERTTYPE_INFO           EventAlertType = "info"
	EVENTALERTTYPE_SUCCESS        EventAlertType = "success"
	EVENTALERTTYPE_USER_UPDATE    EventAlertType = "user_update"
	EVENTALERTTYPE_RECOMMENDATION EventAlertType = "recommendation"
	EVENTALERTTYPE_SNAPSHOT       EventAlertType = "snapshot"
)

func (v *EventAlertType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventAlertType(value)
	for _, existing := range []EventAlertType{"error", "warning", "info", "success", "user_update", "recommendation", "snapshot"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventAlertType", value)
}

// Ptr returns reference to EventAlertType value
func (v EventAlertType) Ptr() *EventAlertType {
	return &v
}

type NullableEventAlertType struct {
	value *EventAlertType
	isSet bool
}

func (v NullableEventAlertType) Get() *EventAlertType {
	return v.value
}

func (v *NullableEventAlertType) Set(val *EventAlertType) {
	v.value = val
	v.isSet = true
}

func (v NullableEventAlertType) IsSet() bool {
	return v.isSet
}

func (v *NullableEventAlertType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventAlertType(val *EventAlertType) *NullableEventAlertType {
	return &NullableEventAlertType{value: val, isSet: true}
}

func (v NullableEventAlertType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventAlertType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
