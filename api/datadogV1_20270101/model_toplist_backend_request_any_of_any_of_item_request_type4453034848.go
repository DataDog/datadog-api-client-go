// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848 Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848 string

// List of ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848.
const (
	TOPLISTBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4453034848_DDSQL ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848 = "ddsql"
)

var allowedToplistBackendRequestAnyOfAnyOfItemRequestType4453034848EnumValues = []ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848{
	TOPLISTBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4453034848_DDSQL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848) GetAllowedValues() []ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848 {
	return allowedToplistBackendRequestAnyOfAnyOfItemRequestType4453034848EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848(value)
	return nil
}

// NewToplistBackendRequestAnyOfAnyOfItemRequestType4453034848FromValue returns a pointer to a valid ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewToplistBackendRequestAnyOfAnyOfItemRequestType4453034848FromValue(v string) (*ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848, error) {
	ev := ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848: valid values are %v", v, allowedToplistBackendRequestAnyOfAnyOfItemRequestType4453034848EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848) IsValid() bool {
	for _, existing := range allowedToplistBackendRequestAnyOfAnyOfItemRequestType4453034848EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848 value.
func (v ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848) Ptr() *ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848 {
	return &v
}
