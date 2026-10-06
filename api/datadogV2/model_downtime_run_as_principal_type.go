// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeRunAsPrincipalType The type of principal allowed to act on behalf of the downtime.
type DowntimeRunAsPrincipalType string

// List of DowntimeRunAsPrincipalType.
const (
	DOWNTIMERUNASPRINCIPALTYPE_USER DowntimeRunAsPrincipalType = "user"
	DOWNTIMERUNASPRINCIPALTYPE_ROLE DowntimeRunAsPrincipalType = "role"
	DOWNTIMERUNASPRINCIPALTYPE_TEAM DowntimeRunAsPrincipalType = "team"
)

var allowedDowntimeRunAsPrincipalTypeEnumValues = []DowntimeRunAsPrincipalType{
	DOWNTIMERUNASPRINCIPALTYPE_USER,
	DOWNTIMERUNASPRINCIPALTYPE_ROLE,
	DOWNTIMERUNASPRINCIPALTYPE_TEAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DowntimeRunAsPrincipalType) GetAllowedValues() []DowntimeRunAsPrincipalType {
	return allowedDowntimeRunAsPrincipalTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DowntimeRunAsPrincipalType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DowntimeRunAsPrincipalType(value)
	return nil
}

// NewDowntimeRunAsPrincipalTypeFromValue returns a pointer to a valid DowntimeRunAsPrincipalType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDowntimeRunAsPrincipalTypeFromValue(v string) (*DowntimeRunAsPrincipalType, error) {
	ev := DowntimeRunAsPrincipalType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DowntimeRunAsPrincipalType: valid values are %v", v, allowedDowntimeRunAsPrincipalTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DowntimeRunAsPrincipalType) IsValid() bool {
	for _, existing := range allowedDowntimeRunAsPrincipalTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DowntimeRunAsPrincipalType value.
func (v DowntimeRunAsPrincipalType) Ptr() *DowntimeRunAsPrincipalType {
	return &v
}
