// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640 Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640 string

// List of QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640.
const (
	QUERYTABLEBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4450136640_TABULAR QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640 = "tabular"
)

var allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640EnumValues = []QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640{
	QUERYTABLEBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4450136640_TABULAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640) GetAllowedValues() []QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640 {
	return allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640(value)
	return nil
}

// NewQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640FromValue returns a pointer to a valid QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640FromValue(v string) (*QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640, error) {
	ev := QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640: valid values are %v", v, allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640) IsValid() bool {
	for _, existing := range allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640 value.
func (v QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640) Ptr() *QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450136640 {
	return &v
}
