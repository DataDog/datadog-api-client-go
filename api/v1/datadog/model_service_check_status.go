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

// ServiceCheckStatus the model 'ServiceCheckStatus'
type ServiceCheckStatus int32

// List of ServiceCheckStatus
const (
	SERVICECHECKSTATUS_OK       ServiceCheckStatus = 0
	SERVICECHECKSTATUS_WARNING  ServiceCheckStatus = 1
	SERVICECHECKSTATUS_CRITICAL ServiceCheckStatus = 2
	SERVICECHECKSTATUS_UNKNOWN  ServiceCheckStatus = 3
)

type NullableServiceCheckStatus struct {
	Value        ServiceCheckStatus
	ExplicitNull bool
}

func (v NullableServiceCheckStatus) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableServiceCheckStatus) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
