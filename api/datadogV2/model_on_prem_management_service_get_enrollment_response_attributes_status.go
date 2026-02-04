// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnPremManagementServiceGetEnrollmentResponseAttributesStatus The status of the enrollment.
type OnPremManagementServiceGetEnrollmentResponseAttributesStatus string

// List of OnPremManagementServiceGetEnrollmentResponseAttributesStatus.
const (
	ONPREMMANAGEMENTSERVICEGETENROLLMENTRESPONSEATTRIBUTESSTATUS_NEW     OnPremManagementServiceGetEnrollmentResponseAttributesStatus = "new"
	ONPREMMANAGEMENTSERVICEGETENROLLMENTRESPONSEATTRIBUTESSTATUS_SUCCESS OnPremManagementServiceGetEnrollmentResponseAttributesStatus = "success"
	ONPREMMANAGEMENTSERVICEGETENROLLMENTRESPONSEATTRIBUTESSTATUS_FAILED  OnPremManagementServiceGetEnrollmentResponseAttributesStatus = "failed"
)

var allowedOnPremManagementServiceGetEnrollmentResponseAttributesStatusEnumValues = []OnPremManagementServiceGetEnrollmentResponseAttributesStatus{
	ONPREMMANAGEMENTSERVICEGETENROLLMENTRESPONSEATTRIBUTESSTATUS_NEW,
	ONPREMMANAGEMENTSERVICEGETENROLLMENTRESPONSEATTRIBUTESSTATUS_SUCCESS,
	ONPREMMANAGEMENTSERVICEGETENROLLMENTRESPONSEATTRIBUTESSTATUS_FAILED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnPremManagementServiceGetEnrollmentResponseAttributesStatus) GetAllowedValues() []OnPremManagementServiceGetEnrollmentResponseAttributesStatus {
	return allowedOnPremManagementServiceGetEnrollmentResponseAttributesStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnPremManagementServiceGetEnrollmentResponseAttributesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnPremManagementServiceGetEnrollmentResponseAttributesStatus(value)
	return nil
}

// NewOnPremManagementServiceGetEnrollmentResponseAttributesStatusFromValue returns a pointer to a valid OnPremManagementServiceGetEnrollmentResponseAttributesStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnPremManagementServiceGetEnrollmentResponseAttributesStatusFromValue(v string) (*OnPremManagementServiceGetEnrollmentResponseAttributesStatus, error) {
	ev := OnPremManagementServiceGetEnrollmentResponseAttributesStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnPremManagementServiceGetEnrollmentResponseAttributesStatus: valid values are %v", v, allowedOnPremManagementServiceGetEnrollmentResponseAttributesStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnPremManagementServiceGetEnrollmentResponseAttributesStatus) IsValid() bool {
	for _, existing := range allowedOnPremManagementServiceGetEnrollmentResponseAttributesStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnPremManagementServiceGetEnrollmentResponseAttributesStatus value.
func (v OnPremManagementServiceGetEnrollmentResponseAttributesStatus) Ptr() *OnPremManagementServiceGetEnrollmentResponseAttributesStatus {
	return &v
}
