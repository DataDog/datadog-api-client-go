// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168 Test-only generated schema for the dashboard versioning proof of concept.
type TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168 string

// List of TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168.
const (
	TREEMAPBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4453527168_SCALAR TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168 = "scalar"
)

var allowedTreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168EnumValues = []TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168{
	TREEMAPBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4453527168_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168) GetAllowedValues() []TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168 {
	return allowedTreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168(value)
	return nil
}

// NewTreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168FromValue returns a pointer to a valid TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168FromValue(v string) (*TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168, error) {
	ev := TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168: valid values are %v", v, allowedTreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168) IsValid() bool {
	for _, existing := range allowedTreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168 value.
func (v TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168) Ptr() *TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168 {
	return &v
}
