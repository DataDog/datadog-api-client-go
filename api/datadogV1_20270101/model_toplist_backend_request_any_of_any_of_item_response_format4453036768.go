// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768 Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768 string

// List of ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768.
const (
	TOPLISTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4453036768_TABULAR ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768 = "tabular"
)

var allowedToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768EnumValues = []ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768{
	TOPLISTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4453036768_TABULAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768) GetAllowedValues() []ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768 {
	return allowedToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768(value)
	return nil
}

// NewToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768FromValue returns a pointer to a valid ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768FromValue(v string) (*ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768, error) {
	ev := ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768: valid values are %v", v, allowedToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768) IsValid() bool {
	for _, existing := range allowedToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768 value.
func (v ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768) Ptr() *ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453036768 {
	return &v
}
