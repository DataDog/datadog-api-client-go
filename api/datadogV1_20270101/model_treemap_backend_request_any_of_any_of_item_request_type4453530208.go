// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208 Test-only generated schema for the dashboard versioning proof of concept.
type TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208 string

// List of TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208.
const (
	TREEMAPBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4453530208_LOCAL_DATASET TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208 = "local_dataset"
)

var allowedTreemapBackendRequestAnyOfAnyOfItemRequestType4453530208EnumValues = []TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208{
	TREEMAPBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4453530208_LOCAL_DATASET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208) GetAllowedValues() []TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208 {
	return allowedTreemapBackendRequestAnyOfAnyOfItemRequestType4453530208EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208(value)
	return nil
}

// NewTreemapBackendRequestAnyOfAnyOfItemRequestType4453530208FromValue returns a pointer to a valid TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreemapBackendRequestAnyOfAnyOfItemRequestType4453530208FromValue(v string) (*TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208, error) {
	ev := TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208: valid values are %v", v, allowedTreemapBackendRequestAnyOfAnyOfItemRequestType4453530208EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208) IsValid() bool {
	for _, existing := range allowedTreemapBackendRequestAnyOfAnyOfItemRequestType4453530208EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208 value.
func (v TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208) Ptr() *TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208 {
	return &v
}
