// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExternalTeamPermissionSettingDefaultsType Resource type for external team permission setting defaults.
type ExternalTeamPermissionSettingDefaultsType string

// List of ExternalTeamPermissionSettingDefaultsType.
const (
	EXTERNALTEAMPERMISSIONSETTINGDEFAULTSTYPE_EXTERNAL_TEAM_PERMISSION_SETTING_DEFAULTS ExternalTeamPermissionSettingDefaultsType = "external_team_permission_setting_defaults"
)

var allowedExternalTeamPermissionSettingDefaultsTypeEnumValues = []ExternalTeamPermissionSettingDefaultsType{
	EXTERNALTEAMPERMISSIONSETTINGDEFAULTSTYPE_EXTERNAL_TEAM_PERMISSION_SETTING_DEFAULTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExternalTeamPermissionSettingDefaultsType) GetAllowedValues() []ExternalTeamPermissionSettingDefaultsType {
	return allowedExternalTeamPermissionSettingDefaultsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExternalTeamPermissionSettingDefaultsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExternalTeamPermissionSettingDefaultsType(value)
	return nil
}

// NewExternalTeamPermissionSettingDefaultsTypeFromValue returns a pointer to a valid ExternalTeamPermissionSettingDefaultsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExternalTeamPermissionSettingDefaultsTypeFromValue(v string) (*ExternalTeamPermissionSettingDefaultsType, error) {
	ev := ExternalTeamPermissionSettingDefaultsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExternalTeamPermissionSettingDefaultsType: valid values are %v", v, allowedExternalTeamPermissionSettingDefaultsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExternalTeamPermissionSettingDefaultsType) IsValid() bool {
	for _, existing := range allowedExternalTeamPermissionSettingDefaultsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExternalTeamPermissionSettingDefaultsType value.
func (v ExternalTeamPermissionSettingDefaultsType) Ptr() *ExternalTeamPermissionSettingDefaultsType {
	return &v
}
