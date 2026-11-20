// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776 Test-only generated schema for the dashboard versioning proof of concept.
type QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776 string

// List of QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776.
const (
	QUERYVALUEBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4450439776_DDSQL QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776 = "ddsql"
)

var allowedQueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776EnumValues = []QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776{
	QUERYVALUEBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4450439776_DDSQL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776) GetAllowedValues() []QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776 {
	return allowedQueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776(value)
	return nil
}

// NewQueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776FromValue returns a pointer to a valid QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776FromValue(v string) (*QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776, error) {
	ev := QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776: valid values are %v", v, allowedQueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776) IsValid() bool {
	for _, existing := range allowedQueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776 value.
func (v QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776) Ptr() *QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776 {
	return &v
}
