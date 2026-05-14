// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpecAttributesStatus The publication status of the spec.
type SpecAttributesStatus string

// List of SpecAttributesStatus.
const (
	SPECATTRIBUTESSTATUS_PUBLISHED  SpecAttributesStatus = "published"
	SPECATTRIBUTESSTATUS_DRAFT      SpecAttributesStatus = "draft"
	SPECATTRIBUTESSTATUS_DEPRECATED SpecAttributesStatus = "deprecated"
)

var allowedSpecAttributesStatusEnumValues = []SpecAttributesStatus{
	SPECATTRIBUTESSTATUS_PUBLISHED,
	SPECATTRIBUTESSTATUS_DRAFT,
	SPECATTRIBUTESSTATUS_DEPRECATED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SpecAttributesStatus) GetAllowedValues() []SpecAttributesStatus {
	return allowedSpecAttributesStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SpecAttributesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SpecAttributesStatus(value)
	return nil
}

// NewSpecAttributesStatusFromValue returns a pointer to a valid SpecAttributesStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSpecAttributesStatusFromValue(v string) (*SpecAttributesStatus, error) {
	ev := SpecAttributesStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SpecAttributesStatus: valid values are %v", v, allowedSpecAttributesStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SpecAttributesStatus) IsValid() bool {
	for _, existing := range allowedSpecAttributesStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SpecAttributesStatus value.
func (v SpecAttributesStatus) Ptr() *SpecAttributesStatus {
	return &v
}
