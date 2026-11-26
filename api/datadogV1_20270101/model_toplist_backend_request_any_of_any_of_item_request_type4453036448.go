// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendRequestAnyOfAnyOfItemRequestType4453036448 Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendRequestAnyOfAnyOfItemRequestType4453036448 string

// List of ToplistBackendRequestAnyOfAnyOfItemRequestType4453036448.
const (
	TOPLISTBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4453036448_LOCAL_DATASET ToplistBackendRequestAnyOfAnyOfItemRequestType4453036448 = "local_dataset"
)

var allowedToplistBackendRequestAnyOfAnyOfItemRequestType4453036448EnumValues = []ToplistBackendRequestAnyOfAnyOfItemRequestType4453036448{
	TOPLISTBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4453036448_LOCAL_DATASET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ToplistBackendRequestAnyOfAnyOfItemRequestType4453036448) GetAllowedValues() []ToplistBackendRequestAnyOfAnyOfItemRequestType4453036448 {
	return allowedToplistBackendRequestAnyOfAnyOfItemRequestType4453036448EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ToplistBackendRequestAnyOfAnyOfItemRequestType4453036448) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ToplistBackendRequestAnyOfAnyOfItemRequestType4453036448(value)
	return nil
}

// NewToplistBackendRequestAnyOfAnyOfItemRequestType4453036448FromValue returns a pointer to a valid ToplistBackendRequestAnyOfAnyOfItemRequestType4453036448
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewToplistBackendRequestAnyOfAnyOfItemRequestType4453036448FromValue(v string) (*ToplistBackendRequestAnyOfAnyOfItemRequestType4453036448, error) {
	ev := ToplistBackendRequestAnyOfAnyOfItemRequestType4453036448(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ToplistBackendRequestAnyOfAnyOfItemRequestType4453036448: valid values are %v", v, allowedToplistBackendRequestAnyOfAnyOfItemRequestType4453036448EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ToplistBackendRequestAnyOfAnyOfItemRequestType4453036448) IsValid() bool {
	for _, existing := range allowedToplistBackendRequestAnyOfAnyOfItemRequestType4453036448EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ToplistBackendRequestAnyOfAnyOfItemRequestType4453036448 value.
func (v ToplistBackendRequestAnyOfAnyOfItemRequestType4453036448) Ptr() *ToplistBackendRequestAnyOfAnyOfItemRequestType4453036448 {
	return &v
}
