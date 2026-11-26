// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256 Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256 string

// List of QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256.
const (
	QUERYTABLEBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4450272256_SCALAR QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256 = "scalar"
)

var allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256EnumValues = []QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256{
	QUERYTABLEBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4450272256_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256) GetAllowedValues() []QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256 {
	return allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256(value)
	return nil
}

// NewQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256FromValue returns a pointer to a valid QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256FromValue(v string) (*QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256, error) {
	ev := QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256: valid values are %v", v, allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256) IsValid() bool {
	for _, existing := range allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256 value.
func (v QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256) Ptr() *QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256 {
	return &v
}
