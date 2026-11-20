// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960 Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960 string

// List of QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960.
const (
	QUERYTABLEBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4450132960_SCALAR QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960 = "scalar"
)

var allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960EnumValues = []QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960{
	QUERYTABLEBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4450132960_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960) GetAllowedValues() []QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960 {
	return allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960(value)
	return nil
}

// NewQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960FromValue returns a pointer to a valid QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960FromValue(v string) (*QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960, error) {
	ev := QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960: valid values are %v", v, allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960) IsValid() bool {
	for _, existing := range allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960 value.
func (v QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960) Ptr() *QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450132960 {
	return &v
}
