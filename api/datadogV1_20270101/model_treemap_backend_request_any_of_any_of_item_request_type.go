// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreemapBackendRequestAnyOfAnyOfItemRequestType Test-only generated schema for the dashboard versioning proof of concept.
type TreemapBackendRequestAnyOfAnyOfItemRequestType string

// List of TreemapBackendRequestAnyOfAnyOfItemRequestType.
const (
	TREEMAPBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE_ADVANCED_QUERY TreemapBackendRequestAnyOfAnyOfItemRequestType = "advanced_query"
)

var allowedTreemapBackendRequestAnyOfAnyOfItemRequestTypeEnumValues = []TreemapBackendRequestAnyOfAnyOfItemRequestType{
	TREEMAPBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE_ADVANCED_QUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreemapBackendRequestAnyOfAnyOfItemRequestType) GetAllowedValues() []TreemapBackendRequestAnyOfAnyOfItemRequestType {
	return allowedTreemapBackendRequestAnyOfAnyOfItemRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreemapBackendRequestAnyOfAnyOfItemRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreemapBackendRequestAnyOfAnyOfItemRequestType(value)
	return nil
}

// NewTreemapBackendRequestAnyOfAnyOfItemRequestTypeFromValue returns a pointer to a valid TreemapBackendRequestAnyOfAnyOfItemRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreemapBackendRequestAnyOfAnyOfItemRequestTypeFromValue(v string) (*TreemapBackendRequestAnyOfAnyOfItemRequestType, error) {
	ev := TreemapBackendRequestAnyOfAnyOfItemRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreemapBackendRequestAnyOfAnyOfItemRequestType: valid values are %v", v, allowedTreemapBackendRequestAnyOfAnyOfItemRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreemapBackendRequestAnyOfAnyOfItemRequestType) IsValid() bool {
	for _, existing := range allowedTreemapBackendRequestAnyOfAnyOfItemRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreemapBackendRequestAnyOfAnyOfItemRequestType value.
func (v TreemapBackendRequestAnyOfAnyOfItemRequestType) Ptr() *TreemapBackendRequestAnyOfAnyOfItemRequestType {
	return &v
}
