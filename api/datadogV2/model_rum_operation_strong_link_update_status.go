// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMOperationStrongLinkUpdateStatus The status of a RUM operation strong link. Can only be set to `CONFIRMED` or `REJECTED`.
type RUMOperationStrongLinkUpdateStatus string

// List of RUMOperationStrongLinkUpdateStatus.
const (
	RUMOPERATIONSTRONGLINKUPDATESTATUS_CONFIRMED RUMOperationStrongLinkUpdateStatus = "CONFIRMED"
	RUMOPERATIONSTRONGLINKUPDATESTATUS_REJECTED  RUMOperationStrongLinkUpdateStatus = "REJECTED"
)

var allowedRUMOperationStrongLinkUpdateStatusEnumValues = []RUMOperationStrongLinkUpdateStatus{
	RUMOPERATIONSTRONGLINKUPDATESTATUS_CONFIRMED,
	RUMOPERATIONSTRONGLINKUPDATESTATUS_REJECTED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMOperationStrongLinkUpdateStatus) GetAllowedValues() []RUMOperationStrongLinkUpdateStatus {
	return allowedRUMOperationStrongLinkUpdateStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMOperationStrongLinkUpdateStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMOperationStrongLinkUpdateStatus(value)
	return nil
}

// NewRUMOperationStrongLinkUpdateStatusFromValue returns a pointer to a valid RUMOperationStrongLinkUpdateStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMOperationStrongLinkUpdateStatusFromValue(v string) (*RUMOperationStrongLinkUpdateStatus, error) {
	ev := RUMOperationStrongLinkUpdateStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMOperationStrongLinkUpdateStatus: valid values are %v", v, allowedRUMOperationStrongLinkUpdateStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMOperationStrongLinkUpdateStatus) IsValid() bool {
	for _, existing := range allowedRUMOperationStrongLinkUpdateStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMOperationStrongLinkUpdateStatus value.
func (v RUMOperationStrongLinkUpdateStatus) Ptr() *RUMOperationStrongLinkUpdateStatus {
	return &v
}
