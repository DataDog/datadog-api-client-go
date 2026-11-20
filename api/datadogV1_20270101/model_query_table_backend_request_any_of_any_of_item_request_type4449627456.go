// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456 Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456 string

// List of QueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456.
const (
	QUERYTABLEBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4449627456_LOCAL_DATASET QueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456 = "local_dataset"
)

var allowedQueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456EnumValues = []QueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456{
	QUERYTABLEBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4449627456_LOCAL_DATASET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456) GetAllowedValues() []QueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456 {
	return allowedQueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456(value)
	return nil
}

// NewQueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456FromValue returns a pointer to a valid QueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456FromValue(v string) (*QueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456, error) {
	ev := QueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456: valid values are %v", v, allowedQueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456) IsValid() bool {
	for _, existing := range allowedQueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456 value.
func (v QueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456) Ptr() *QueryTableBackendRequestAnyOfAnyOfItemRequestType4449627456 {
	return &v
}
