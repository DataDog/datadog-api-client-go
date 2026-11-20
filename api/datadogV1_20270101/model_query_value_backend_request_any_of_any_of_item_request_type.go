// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryValueBackendRequestAnyOfAnyOfItemRequestType Test-only generated schema for the dashboard versioning proof of concept.
type QueryValueBackendRequestAnyOfAnyOfItemRequestType string

// List of QueryValueBackendRequestAnyOfAnyOfItemRequestType.
const (
	QUERYVALUEBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE_ADVANCED_QUERY QueryValueBackendRequestAnyOfAnyOfItemRequestType = "advanced_query"
)

var allowedQueryValueBackendRequestAnyOfAnyOfItemRequestTypeEnumValues = []QueryValueBackendRequestAnyOfAnyOfItemRequestType{
	QUERYVALUEBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE_ADVANCED_QUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryValueBackendRequestAnyOfAnyOfItemRequestType) GetAllowedValues() []QueryValueBackendRequestAnyOfAnyOfItemRequestType {
	return allowedQueryValueBackendRequestAnyOfAnyOfItemRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryValueBackendRequestAnyOfAnyOfItemRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryValueBackendRequestAnyOfAnyOfItemRequestType(value)
	return nil
}

// NewQueryValueBackendRequestAnyOfAnyOfItemRequestTypeFromValue returns a pointer to a valid QueryValueBackendRequestAnyOfAnyOfItemRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryValueBackendRequestAnyOfAnyOfItemRequestTypeFromValue(v string) (*QueryValueBackendRequestAnyOfAnyOfItemRequestType, error) {
	ev := QueryValueBackendRequestAnyOfAnyOfItemRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryValueBackendRequestAnyOfAnyOfItemRequestType: valid values are %v", v, allowedQueryValueBackendRequestAnyOfAnyOfItemRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryValueBackendRequestAnyOfAnyOfItemRequestType) IsValid() bool {
	for _, existing := range allowedQueryValueBackendRequestAnyOfAnyOfItemRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryValueBackendRequestAnyOfAnyOfItemRequestType value.
func (v QueryValueBackendRequestAnyOfAnyOfItemRequestType) Ptr() *QueryValueBackendRequestAnyOfAnyOfItemRequestType {
	return &v
}
