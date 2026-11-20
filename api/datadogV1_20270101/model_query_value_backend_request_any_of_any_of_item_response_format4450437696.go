// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696 Test-only generated schema for the dashboard versioning proof of concept.
type QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696 string

// List of QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696.
const (
	QUERYVALUEBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4450437696_SCALAR QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696 = "scalar"
)

var allowedQueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696EnumValues = []QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696{
	QUERYVALUEBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4450437696_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696) GetAllowedValues() []QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696 {
	return allowedQueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696(value)
	return nil
}

// NewQueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696FromValue returns a pointer to a valid QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696FromValue(v string) (*QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696, error) {
	ev := QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696: valid values are %v", v, allowedQueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696) IsValid() bool {
	for _, existing := range allowedQueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696 value.
func (v QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696) Ptr() *QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450437696 {
	return &v
}
