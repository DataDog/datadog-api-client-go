// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840 Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840 string

// List of QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840.
const (
	QUERYTABLEBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4450135840_ADVANCED_QUERY QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840 = "advanced_query"
)

var allowedQueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840EnumValues = []QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840{
	QUERYTABLEBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4450135840_ADVANCED_QUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840) GetAllowedValues() []QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840 {
	return allowedQueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840(value)
	return nil
}

// NewQueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840FromValue returns a pointer to a valid QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840FromValue(v string) (*QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840, error) {
	ev := QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840: valid values are %v", v, allowedQueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840) IsValid() bool {
	for _, existing := range allowedQueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840 value.
func (v QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840) Ptr() *QueryTableBackendRequestAnyOfAnyOfItemRequestType4450135840 {
	return &v
}
