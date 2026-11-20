// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776 Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776 string

// List of QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776.
const (
	QUERYTABLEBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4449627776_TABULAR QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776 = "tabular"
)

var allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776EnumValues = []QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776{
	QUERYTABLEBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4449627776_TABULAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776) GetAllowedValues() []QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776 {
	return allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776(value)
	return nil
}

// NewQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776FromValue returns a pointer to a valid QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776FromValue(v string) (*QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776, error) {
	ev := QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776: valid values are %v", v, allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776) IsValid() bool {
	for _, existing := range allowedQueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776 value.
func (v QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776) Ptr() *QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4449627776 {
	return &v
}
