// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ManageStatusBackendDefinitionSummaryType Test-only generated schema for the dashboard versioning proof of concept.
type ManageStatusBackendDefinitionSummaryType string

// List of ManageStatusBackendDefinitionSummaryType.
const (
	MANAGESTATUSBACKENDDEFINITIONSUMMARYTYPE_GROUPS   ManageStatusBackendDefinitionSummaryType = "groups"
	MANAGESTATUSBACKENDDEFINITIONSUMMARYTYPE_MONITORS ManageStatusBackendDefinitionSummaryType = "monitors"
	MANAGESTATUSBACKENDDEFINITIONSUMMARYTYPE_COMBINED ManageStatusBackendDefinitionSummaryType = "combined"
)

var allowedManageStatusBackendDefinitionSummaryTypeEnumValues = []ManageStatusBackendDefinitionSummaryType{
	MANAGESTATUSBACKENDDEFINITIONSUMMARYTYPE_GROUPS,
	MANAGESTATUSBACKENDDEFINITIONSUMMARYTYPE_MONITORS,
	MANAGESTATUSBACKENDDEFINITIONSUMMARYTYPE_COMBINED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ManageStatusBackendDefinitionSummaryType) GetAllowedValues() []ManageStatusBackendDefinitionSummaryType {
	return allowedManageStatusBackendDefinitionSummaryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ManageStatusBackendDefinitionSummaryType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ManageStatusBackendDefinitionSummaryType(value)
	return nil
}

// NewManageStatusBackendDefinitionSummaryTypeFromValue returns a pointer to a valid ManageStatusBackendDefinitionSummaryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewManageStatusBackendDefinitionSummaryTypeFromValue(v string) (*ManageStatusBackendDefinitionSummaryType, error) {
	ev := ManageStatusBackendDefinitionSummaryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ManageStatusBackendDefinitionSummaryType: valid values are %v", v, allowedManageStatusBackendDefinitionSummaryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ManageStatusBackendDefinitionSummaryType) IsValid() bool {
	for _, existing := range allowedManageStatusBackendDefinitionSummaryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ManageStatusBackendDefinitionSummaryType value.
func (v ManageStatusBackendDefinitionSummaryType) Ptr() *ManageStatusBackendDefinitionSummaryType {
	return &v
}
