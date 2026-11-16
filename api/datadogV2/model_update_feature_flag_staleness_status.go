// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateFeatureFlagStalenessStatus The lifecycle to set. `ACTIVE` makes the flag temporary; `PERMANENT` excludes it from
// staleness detection. Either value clears any snooze; omit the field to keep the lifecycle.
// Do not copy the top-level `staleness_status` from a response into this field: that legacy
// response field is always `ACTIVE`.
type UpdateFeatureFlagStalenessStatus string

// List of UpdateFeatureFlagStalenessStatus.
const (
	UPDATEFEATUREFLAGSTALENESSSTATUS_ACTIVE    UpdateFeatureFlagStalenessStatus = "ACTIVE"
	UPDATEFEATUREFLAGSTALENESSSTATUS_PERMANENT UpdateFeatureFlagStalenessStatus = "PERMANENT"
)

var allowedUpdateFeatureFlagStalenessStatusEnumValues = []UpdateFeatureFlagStalenessStatus{
	UPDATEFEATUREFLAGSTALENESSSTATUS_ACTIVE,
	UPDATEFEATUREFLAGSTALENESSSTATUS_PERMANENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UpdateFeatureFlagStalenessStatus) GetAllowedValues() []UpdateFeatureFlagStalenessStatus {
	return allowedUpdateFeatureFlagStalenessStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UpdateFeatureFlagStalenessStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UpdateFeatureFlagStalenessStatus(value)
	return nil
}

// NewUpdateFeatureFlagStalenessStatusFromValue returns a pointer to a valid UpdateFeatureFlagStalenessStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUpdateFeatureFlagStalenessStatusFromValue(v string) (*UpdateFeatureFlagStalenessStatus, error) {
	ev := UpdateFeatureFlagStalenessStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UpdateFeatureFlagStalenessStatus: valid values are %v", v, allowedUpdateFeatureFlagStalenessStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UpdateFeatureFlagStalenessStatus) IsValid() bool {
	for _, existing := range allowedUpdateFeatureFlagStalenessStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateFeatureFlagStalenessStatus value.
func (v UpdateFeatureFlagStalenessStatus) Ptr() *UpdateFeatureFlagStalenessStatus {
	return &v
}
