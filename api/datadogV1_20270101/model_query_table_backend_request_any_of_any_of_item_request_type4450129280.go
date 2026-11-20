// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280 Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280 string

// List of QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280.
const (
	QUERYTABLEBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4450129280_TABLE QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280 = "table"
)

var allowedQueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280EnumValues = []QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280{
	QUERYTABLEBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4450129280_TABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280) GetAllowedValues() []QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280 {
	return allowedQueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280(value)
	return nil
}

// NewQueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280FromValue returns a pointer to a valid QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280FromValue(v string) (*QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280, error) {
	ev := QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280: valid values are %v", v, allowedQueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280) IsValid() bool {
	for _, existing := range allowedQueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280 value.
func (v QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280) Ptr() *QueryTableBackendRequestAnyOfAnyOfItemRequestType4450129280 {
	return &v
}
