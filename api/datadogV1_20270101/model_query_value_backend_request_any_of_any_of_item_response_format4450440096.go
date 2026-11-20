// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096 Test-only generated schema for the dashboard versioning proof of concept.
type QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096 string

// List of QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096.
const (
	QUERYVALUEBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4450440096_SCALAR QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096 = "scalar"
)

var allowedQueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096EnumValues = []QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096{
	QUERYVALUEBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4450440096_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096) GetAllowedValues() []QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096 {
	return allowedQueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096(value)
	return nil
}

// NewQueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096FromValue returns a pointer to a valid QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096FromValue(v string) (*QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096, error) {
	ev := QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096: valid values are %v", v, allowedQueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096) IsValid() bool {
	for _, existing := range allowedQueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096 value.
func (v QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096) Ptr() *QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096 {
	return &v
}
