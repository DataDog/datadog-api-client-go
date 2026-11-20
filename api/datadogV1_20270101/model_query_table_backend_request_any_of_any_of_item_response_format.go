// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableBackendRequestAnyOfAnyOfItemResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableBackendRequestAnyOfAnyOfItemResponseFormat string

// List of QueryTableBackendRequestAnyOfAnyOfItemResponseFormat.
const (
	QUERYTABLEBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT_SCALAR QueryTableBackendRequestAnyOfAnyOfItemResponseFormat = "scalar"
)

var allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormatEnumValues = []QueryTableBackendRequestAnyOfAnyOfItemResponseFormat{
	QUERYTABLEBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryTableBackendRequestAnyOfAnyOfItemResponseFormat) GetAllowedValues() []QueryTableBackendRequestAnyOfAnyOfItemResponseFormat {
	return allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryTableBackendRequestAnyOfAnyOfItemResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryTableBackendRequestAnyOfAnyOfItemResponseFormat(value)
	return nil
}

// NewQueryTableBackendRequestAnyOfAnyOfItemResponseFormatFromValue returns a pointer to a valid QueryTableBackendRequestAnyOfAnyOfItemResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryTableBackendRequestAnyOfAnyOfItemResponseFormatFromValue(v string) (*QueryTableBackendRequestAnyOfAnyOfItemResponseFormat, error) {
	ev := QueryTableBackendRequestAnyOfAnyOfItemResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryTableBackendRequestAnyOfAnyOfItemResponseFormat: valid values are %v", v, allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryTableBackendRequestAnyOfAnyOfItemResponseFormat) IsValid() bool {
	for _, existing := range allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryTableBackendRequestAnyOfAnyOfItemResponseFormat value.
func (v QueryTableBackendRequestAnyOfAnyOfItemResponseFormat) Ptr() *QueryTableBackendRequestAnyOfAnyOfItemResponseFormat {
	return &v
}
