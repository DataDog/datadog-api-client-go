// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMOperationStrongLinkStatus The status of a RUM operation strong link.
type RUMOperationStrongLinkStatus string

// List of RUMOperationStrongLinkStatus.
const (
	RUMOPERATIONSTRONGLINKSTATUS_DRAFT     RUMOperationStrongLinkStatus = "DRAFT"
	RUMOPERATIONSTRONGLINKSTATUS_CONFIRMED RUMOperationStrongLinkStatus = "CONFIRMED"
	RUMOPERATIONSTRONGLINKSTATUS_REJECTED  RUMOperationStrongLinkStatus = "REJECTED"
)

var allowedRUMOperationStrongLinkStatusEnumValues = []RUMOperationStrongLinkStatus{
	RUMOPERATIONSTRONGLINKSTATUS_DRAFT,
	RUMOPERATIONSTRONGLINKSTATUS_CONFIRMED,
	RUMOPERATIONSTRONGLINKSTATUS_REJECTED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMOperationStrongLinkStatus) GetAllowedValues() []RUMOperationStrongLinkStatus {
	return allowedRUMOperationStrongLinkStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMOperationStrongLinkStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMOperationStrongLinkStatus(value)
	return nil
}

// NewRUMOperationStrongLinkStatusFromValue returns a pointer to a valid RUMOperationStrongLinkStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMOperationStrongLinkStatusFromValue(v string) (*RUMOperationStrongLinkStatus, error) {
	ev := RUMOperationStrongLinkStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMOperationStrongLinkStatus: valid values are %v", v, allowedRUMOperationStrongLinkStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMOperationStrongLinkStatus) IsValid() bool {
	for _, existing := range allowedRUMOperationStrongLinkStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMOperationStrongLinkStatus value.
func (v RUMOperationStrongLinkStatus) Ptr() *RUMOperationStrongLinkStatus {
	return &v
}
