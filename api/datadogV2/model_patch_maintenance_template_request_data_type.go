// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchMaintenanceTemplateRequestDataType Maintenance templates resource type.
type PatchMaintenanceTemplateRequestDataType string

// List of PatchMaintenanceTemplateRequestDataType.
const (
	PATCHMAINTENANCETEMPLATEREQUESTDATATYPE_MAINTENANCE_TEMPLATES PatchMaintenanceTemplateRequestDataType = "maintenance_templates"
)

var allowedPatchMaintenanceTemplateRequestDataTypeEnumValues = []PatchMaintenanceTemplateRequestDataType{
	PATCHMAINTENANCETEMPLATEREQUESTDATATYPE_MAINTENANCE_TEMPLATES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PatchMaintenanceTemplateRequestDataType) GetAllowedValues() []PatchMaintenanceTemplateRequestDataType {
	return allowedPatchMaintenanceTemplateRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PatchMaintenanceTemplateRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PatchMaintenanceTemplateRequestDataType(value)
	return nil
}

// NewPatchMaintenanceTemplateRequestDataTypeFromValue returns a pointer to a valid PatchMaintenanceTemplateRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPatchMaintenanceTemplateRequestDataTypeFromValue(v string) (*PatchMaintenanceTemplateRequestDataType, error) {
	ev := PatchMaintenanceTemplateRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PatchMaintenanceTemplateRequestDataType: valid values are %v", v, allowedPatchMaintenanceTemplateRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PatchMaintenanceTemplateRequestDataType) IsValid() bool {
	for _, existing := range allowedPatchMaintenanceTemplateRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchMaintenanceTemplateRequestDataType value.
func (v PatchMaintenanceTemplateRequestDataType) Ptr() *PatchMaintenanceTemplateRequestDataType {
	return &v
}
