// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapLayoutDirection Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapLayoutDirection string

// List of CustomTopologyMapLayoutDirection.
const (
	CUSTOMTOPOLOGYMAPLAYOUTDIRECTION_TOP_TO_BOTTOM CustomTopologyMapLayoutDirection = "top-to-bottom"
	CUSTOMTOPOLOGYMAPLAYOUTDIRECTION_BOTTOM_TO_TOP CustomTopologyMapLayoutDirection = "bottom-to-top"
	CUSTOMTOPOLOGYMAPLAYOUTDIRECTION_LEFT_TO_RIGHT CustomTopologyMapLayoutDirection = "left-to-right"
	CUSTOMTOPOLOGYMAPLAYOUTDIRECTION_RIGHT_TO_LEFT CustomTopologyMapLayoutDirection = "right-to-left"
)

var allowedCustomTopologyMapLayoutDirectionEnumValues = []CustomTopologyMapLayoutDirection{
	CUSTOMTOPOLOGYMAPLAYOUTDIRECTION_TOP_TO_BOTTOM,
	CUSTOMTOPOLOGYMAPLAYOUTDIRECTION_BOTTOM_TO_TOP,
	CUSTOMTOPOLOGYMAPLAYOUTDIRECTION_LEFT_TO_RIGHT,
	CUSTOMTOPOLOGYMAPLAYOUTDIRECTION_RIGHT_TO_LEFT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomTopologyMapLayoutDirection) GetAllowedValues() []CustomTopologyMapLayoutDirection {
	return allowedCustomTopologyMapLayoutDirectionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomTopologyMapLayoutDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomTopologyMapLayoutDirection(value)
	return nil
}

// NewCustomTopologyMapLayoutDirectionFromValue returns a pointer to a valid CustomTopologyMapLayoutDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomTopologyMapLayoutDirectionFromValue(v string) (*CustomTopologyMapLayoutDirection, error) {
	ev := CustomTopologyMapLayoutDirection(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomTopologyMapLayoutDirection: valid values are %v", v, allowedCustomTopologyMapLayoutDirectionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomTopologyMapLayoutDirection) IsValid() bool {
	for _, existing := range allowedCustomTopologyMapLayoutDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomTopologyMapLayoutDirection value.
func (v CustomTopologyMapLayoutDirection) Ptr() *CustomTopologyMapLayoutDirection {
	return &v
}
