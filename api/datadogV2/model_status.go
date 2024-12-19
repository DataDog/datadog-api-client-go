// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Status The vulnerability status.
type Status string

// List of Status.
const (
	STATUS_OPEN       Status = "Open"
	STATUS_MUTED      Status = "Muted"
	STATUS_REMEDIATED Status = "Remediated"
	STATUS_INPROGRESS Status = "InProgress"
	STATUS_AUTOCLOSED Status = "AutoClosed"
)

var allowedStatusEnumValues = []Status{
	STATUS_OPEN,
	STATUS_MUTED,
	STATUS_REMEDIATED,
	STATUS_INPROGRESS,
	STATUS_AUTOCLOSED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *Status) GetAllowedValues() []Status {
	return allowedStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *Status) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = Status(value)
	return nil
}

// NewStatusFromValue returns a pointer to a valid Status
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStatusFromValue(v string) (*Status, error) {
	ev := Status(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for Status: valid values are %v", v, allowedStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v Status) IsValid() bool {
	for _, existing := range allowedStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Status value.
func (v Status) Ptr() *Status {
	return &v
}
