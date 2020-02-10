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

// EventPriority The priority of the event: normal or low.
type EventPriority string

// List of EventPriority
const (
	EVENTPRIORITY_NORMAL EventPriority = "normal"
	EVENTPRIORITY_LOW    EventPriority = "low"
)

// Ptr returns reference to EventPriority value
func (v EventPriority) Ptr() *EventPriority {
	return &v
}

type NullableEventPriority struct {
	Value        EventPriority
	ExplicitNull bool
}

func (v NullableEventPriority) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableEventPriority) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
