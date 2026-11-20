// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ManageStatusBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type ManageStatusBackendDefinitionType string

// List of ManageStatusBackendDefinitionType.
const (
	MANAGESTATUSBACKENDDEFINITIONTYPE_MANAGE_STATUS ManageStatusBackendDefinitionType = "manage_status"
)

var allowedManageStatusBackendDefinitionTypeEnumValues = []ManageStatusBackendDefinitionType{
	MANAGESTATUSBACKENDDEFINITIONTYPE_MANAGE_STATUS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ManageStatusBackendDefinitionType) GetAllowedValues() []ManageStatusBackendDefinitionType {
	return allowedManageStatusBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ManageStatusBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ManageStatusBackendDefinitionType(value)
	return nil
}

// NewManageStatusBackendDefinitionTypeFromValue returns a pointer to a valid ManageStatusBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewManageStatusBackendDefinitionTypeFromValue(v string) (*ManageStatusBackendDefinitionType, error) {
	ev := ManageStatusBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ManageStatusBackendDefinitionType: valid values are %v", v, allowedManageStatusBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ManageStatusBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedManageStatusBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ManageStatusBackendDefinitionType value.
func (v ManageStatusBackendDefinitionType) Ptr() *ManageStatusBackendDefinitionType {
	return &v
}
