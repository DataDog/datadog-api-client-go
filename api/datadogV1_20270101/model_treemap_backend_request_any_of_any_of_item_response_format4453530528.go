// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528 Test-only generated schema for the dashboard versioning proof of concept.
type TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528 string

// List of TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528.
const (
	TREEMAPBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4453530528_TABULAR TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528 = "tabular"
)

var allowedTreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528EnumValues = []TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528{
	TREEMAPBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4453530528_TABULAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528) GetAllowedValues() []TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528 {
	return allowedTreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528(value)
	return nil
}

// NewTreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528FromValue returns a pointer to a valid TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528FromValue(v string) (*TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528, error) {
	ev := TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528: valid values are %v", v, allowedTreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528) IsValid() bool {
	for _, existing := range allowedTreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528 value.
func (v TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528) Ptr() *TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528 {
	return &v
}
