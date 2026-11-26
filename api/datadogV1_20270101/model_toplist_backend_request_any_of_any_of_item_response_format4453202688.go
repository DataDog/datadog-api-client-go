// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688 Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688 string

// List of ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688.
const (
	TOPLISTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4453202688_SCALAR ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688 = "scalar"
)

var allowedToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688EnumValues = []ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688{
	TOPLISTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4453202688_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688) GetAllowedValues() []ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688 {
	return allowedToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688(value)
	return nil
}

// NewToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688FromValue returns a pointer to a valid ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688FromValue(v string) (*ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688, error) {
	ev := ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688: valid values are %v", v, allowedToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688) IsValid() bool {
	for _, existing := range allowedToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688 value.
func (v ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688) Ptr() *ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453202688 {
	return &v
}
