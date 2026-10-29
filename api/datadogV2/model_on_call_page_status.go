// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageStatus The current status of the On-Call Page.
type OnCallPageStatus string

// List of OnCallPageStatus.
const (
	ONCALLPAGESTATUS_TRIGGERED    OnCallPageStatus = "triggered"
	ONCALLPAGESTATUS_ACKNOWLEDGED OnCallPageStatus = "acknowledged"
	ONCALLPAGESTATUS_RESOLVED     OnCallPageStatus = "resolved"
)

var allowedOnCallPageStatusEnumValues = []OnCallPageStatus{
	ONCALLPAGESTATUS_TRIGGERED,
	ONCALLPAGESTATUS_ACKNOWLEDGED,
	ONCALLPAGESTATUS_RESOLVED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnCallPageStatus) GetAllowedValues() []OnCallPageStatus {
	return allowedOnCallPageStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnCallPageStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnCallPageStatus(value)
	return nil
}

// NewOnCallPageStatusFromValue returns a pointer to a valid OnCallPageStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnCallPageStatusFromValue(v string) (*OnCallPageStatus, error) {
	ev := OnCallPageStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnCallPageStatus: valid values are %v", v, allowedOnCallPageStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnCallPageStatus) IsValid() bool {
	for _, existing := range allowedOnCallPageStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnCallPageStatus value.
func (v OnCallPageStatus) Ptr() *OnCallPageStatus {
	return &v
}
