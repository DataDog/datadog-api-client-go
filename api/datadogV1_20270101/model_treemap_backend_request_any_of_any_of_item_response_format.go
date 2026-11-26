// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreemapBackendRequestAnyOfAnyOfItemResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type TreemapBackendRequestAnyOfAnyOfItemResponseFormat string

// List of TreemapBackendRequestAnyOfAnyOfItemResponseFormat.
const (
	TREEMAPBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT_TABULAR TreemapBackendRequestAnyOfAnyOfItemResponseFormat = "tabular"
)

var allowedTreemapBackendRequestAnyOfAnyOfItemResponseFormatEnumValues = []TreemapBackendRequestAnyOfAnyOfItemResponseFormat{
	TREEMAPBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT_TABULAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreemapBackendRequestAnyOfAnyOfItemResponseFormat) GetAllowedValues() []TreemapBackendRequestAnyOfAnyOfItemResponseFormat {
	return allowedTreemapBackendRequestAnyOfAnyOfItemResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreemapBackendRequestAnyOfAnyOfItemResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreemapBackendRequestAnyOfAnyOfItemResponseFormat(value)
	return nil
}

// NewTreemapBackendRequestAnyOfAnyOfItemResponseFormatFromValue returns a pointer to a valid TreemapBackendRequestAnyOfAnyOfItemResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreemapBackendRequestAnyOfAnyOfItemResponseFormatFromValue(v string) (*TreemapBackendRequestAnyOfAnyOfItemResponseFormat, error) {
	ev := TreemapBackendRequestAnyOfAnyOfItemResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreemapBackendRequestAnyOfAnyOfItemResponseFormat: valid values are %v", v, allowedTreemapBackendRequestAnyOfAnyOfItemResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreemapBackendRequestAnyOfAnyOfItemResponseFormat) IsValid() bool {
	for _, existing := range allowedTreemapBackendRequestAnyOfAnyOfItemResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreemapBackendRequestAnyOfAnyOfItemResponseFormat value.
func (v TreemapBackendRequestAnyOfAnyOfItemResponseFormat) Ptr() *TreemapBackendRequestAnyOfAnyOfItemResponseFormat {
	return &v
}
