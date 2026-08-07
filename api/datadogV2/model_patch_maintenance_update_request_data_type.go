// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchMaintenanceUpdateRequestDataType Maintenance updates resource type.
type PatchMaintenanceUpdateRequestDataType string

// List of PatchMaintenanceUpdateRequestDataType.
const (
	PATCHMAINTENANCEUPDATEREQUESTDATATYPE_MAINTENANCE_UPDATES PatchMaintenanceUpdateRequestDataType = "maintenance_updates"
)

var allowedPatchMaintenanceUpdateRequestDataTypeEnumValues = []PatchMaintenanceUpdateRequestDataType{
	PATCHMAINTENANCEUPDATEREQUESTDATATYPE_MAINTENANCE_UPDATES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PatchMaintenanceUpdateRequestDataType) GetAllowedValues() []PatchMaintenanceUpdateRequestDataType {
	return allowedPatchMaintenanceUpdateRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PatchMaintenanceUpdateRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PatchMaintenanceUpdateRequestDataType(value)
	return nil
}

// NewPatchMaintenanceUpdateRequestDataTypeFromValue returns a pointer to a valid PatchMaintenanceUpdateRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPatchMaintenanceUpdateRequestDataTypeFromValue(v string) (*PatchMaintenanceUpdateRequestDataType, error) {
	ev := PatchMaintenanceUpdateRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PatchMaintenanceUpdateRequestDataType: valid values are %v", v, allowedPatchMaintenanceUpdateRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PatchMaintenanceUpdateRequestDataType) IsValid() bool {
	for _, existing := range allowedPatchMaintenanceUpdateRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchMaintenanceUpdateRequestDataType value.
func (v PatchMaintenanceUpdateRequestDataType) Ptr() *PatchMaintenanceUpdateRequestDataType {
	return &v
}
