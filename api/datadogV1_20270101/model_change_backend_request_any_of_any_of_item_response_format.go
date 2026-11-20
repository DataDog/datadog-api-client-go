// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeBackendRequestAnyOfAnyOfItemResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type ChangeBackendRequestAnyOfAnyOfItemResponseFormat string

// List of ChangeBackendRequestAnyOfAnyOfItemResponseFormat.
const (
	CHANGEBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT_SCALAR ChangeBackendRequestAnyOfAnyOfItemResponseFormat = "scalar"
)

var allowedChangeBackendRequestAnyOfAnyOfItemResponseFormatEnumValues = []ChangeBackendRequestAnyOfAnyOfItemResponseFormat{
	CHANGEBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChangeBackendRequestAnyOfAnyOfItemResponseFormat) GetAllowedValues() []ChangeBackendRequestAnyOfAnyOfItemResponseFormat {
	return allowedChangeBackendRequestAnyOfAnyOfItemResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChangeBackendRequestAnyOfAnyOfItemResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeBackendRequestAnyOfAnyOfItemResponseFormat(value)
	return nil
}

// NewChangeBackendRequestAnyOfAnyOfItemResponseFormatFromValue returns a pointer to a valid ChangeBackendRequestAnyOfAnyOfItemResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChangeBackendRequestAnyOfAnyOfItemResponseFormatFromValue(v string) (*ChangeBackendRequestAnyOfAnyOfItemResponseFormat, error) {
	ev := ChangeBackendRequestAnyOfAnyOfItemResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChangeBackendRequestAnyOfAnyOfItemResponseFormat: valid values are %v", v, allowedChangeBackendRequestAnyOfAnyOfItemResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChangeBackendRequestAnyOfAnyOfItemResponseFormat) IsValid() bool {
	for _, existing := range allowedChangeBackendRequestAnyOfAnyOfItemResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeBackendRequestAnyOfAnyOfItemResponseFormat value.
func (v ChangeBackendRequestAnyOfAnyOfItemResponseFormat) Ptr() *ChangeBackendRequestAnyOfAnyOfItemResponseFormat {
	return &v
}
