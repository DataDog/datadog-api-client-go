// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168 Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168 string

// List of ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168.
const (
	TOPLISTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4453035168_SCALAR ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168 = "scalar"
)

var allowedToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168EnumValues = []ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168{
	TOPLISTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4453035168_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168) GetAllowedValues() []ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168 {
	return allowedToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168(value)
	return nil
}

// NewToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168FromValue returns a pointer to a valid ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168FromValue(v string) (*ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168, error) {
	ev := ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168: valid values are %v", v, allowedToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168) IsValid() bool {
	for _, existing := range allowedToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168 value.
func (v ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168) Ptr() *ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168 {
	return &v
}
