// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768 Test-only generated schema for the dashboard versioning proof of concept.
type TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768 string

// List of TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768.
const (
	TREEMAPBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4453528768_DDSQL TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768 = "ddsql"
)

var allowedTreemapBackendRequestAnyOfAnyOfItemRequestType4453528768EnumValues = []TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768{
	TREEMAPBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4453528768_DDSQL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768) GetAllowedValues() []TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768 {
	return allowedTreemapBackendRequestAnyOfAnyOfItemRequestType4453528768EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768(value)
	return nil
}

// NewTreemapBackendRequestAnyOfAnyOfItemRequestType4453528768FromValue returns a pointer to a valid TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreemapBackendRequestAnyOfAnyOfItemRequestType4453528768FromValue(v string) (*TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768, error) {
	ev := TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768: valid values are %v", v, allowedTreemapBackendRequestAnyOfAnyOfItemRequestType4453528768EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768) IsValid() bool {
	for _, existing := range allowedTreemapBackendRequestAnyOfAnyOfItemRequestType4453528768EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768 value.
func (v TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768) Ptr() *TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768 {
	return &v
}
