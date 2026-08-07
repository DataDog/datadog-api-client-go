// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MaintenanceUpdateDataAttributesStatus The status of the maintenance update.
type MaintenanceUpdateDataAttributesStatus string

// List of MaintenanceUpdateDataAttributesStatus.
const (
	MAINTENANCEUPDATEDATAATTRIBUTESSTATUS_SCHEDULED   MaintenanceUpdateDataAttributesStatus = "scheduled"
	MAINTENANCEUPDATEDATAATTRIBUTESSTATUS_IN_PROGRESS MaintenanceUpdateDataAttributesStatus = "in_progress"
	MAINTENANCEUPDATEDATAATTRIBUTESSTATUS_COMPLETED   MaintenanceUpdateDataAttributesStatus = "completed"
	MAINTENANCEUPDATEDATAATTRIBUTESSTATUS_CANCELED    MaintenanceUpdateDataAttributesStatus = "canceled"
)

var allowedMaintenanceUpdateDataAttributesStatusEnumValues = []MaintenanceUpdateDataAttributesStatus{
	MAINTENANCEUPDATEDATAATTRIBUTESSTATUS_SCHEDULED,
	MAINTENANCEUPDATEDATAATTRIBUTESSTATUS_IN_PROGRESS,
	MAINTENANCEUPDATEDATAATTRIBUTESSTATUS_COMPLETED,
	MAINTENANCEUPDATEDATAATTRIBUTESSTATUS_CANCELED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MaintenanceUpdateDataAttributesStatus) GetAllowedValues() []MaintenanceUpdateDataAttributesStatus {
	return allowedMaintenanceUpdateDataAttributesStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MaintenanceUpdateDataAttributesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MaintenanceUpdateDataAttributesStatus(value)
	return nil
}

// NewMaintenanceUpdateDataAttributesStatusFromValue returns a pointer to a valid MaintenanceUpdateDataAttributesStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMaintenanceUpdateDataAttributesStatusFromValue(v string) (*MaintenanceUpdateDataAttributesStatus, error) {
	ev := MaintenanceUpdateDataAttributesStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MaintenanceUpdateDataAttributesStatus: valid values are %v", v, allowedMaintenanceUpdateDataAttributesStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MaintenanceUpdateDataAttributesStatus) IsValid() bool {
	for _, existing := range allowedMaintenanceUpdateDataAttributesStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MaintenanceUpdateDataAttributesStatus value.
func (v MaintenanceUpdateDataAttributesStatus) Ptr() *MaintenanceUpdateDataAttributesStatus {
	return &v
}
