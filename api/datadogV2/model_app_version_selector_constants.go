// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppVersionSelectorConstants Constants that always select a particular version of an app.
type AppVersionSelectorConstants string

// List of AppVersionSelectorConstants.
const (
	APPVERSIONSELECTORCONSTANTS_LATEST   AppVersionSelectorConstants = "latest"
	APPVERSIONSELECTORCONSTANTS_DEPLOYED AppVersionSelectorConstants = "deployed"
)

var allowedAppVersionSelectorConstantsEnumValues = []AppVersionSelectorConstants{
	APPVERSIONSELECTORCONSTANTS_LATEST,
	APPVERSIONSELECTORCONSTANTS_DEPLOYED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AppVersionSelectorConstants) GetAllowedValues() []AppVersionSelectorConstants {
	return allowedAppVersionSelectorConstantsEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AppVersionSelectorConstants) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AppVersionSelectorConstants(value)
	return nil
}

// NewAppVersionSelectorConstantsFromValue returns a pointer to a valid AppVersionSelectorConstants
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAppVersionSelectorConstantsFromValue(v string) (*AppVersionSelectorConstants, error) {
	ev := AppVersionSelectorConstants(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AppVersionSelectorConstants: valid values are %v", v, allowedAppVersionSelectorConstantsEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AppVersionSelectorConstants) IsValid() bool {
	for _, existing := range allowedAppVersionSelectorConstantsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppVersionSelectorConstants value.
func (v AppVersionSelectorConstants) Ptr() *AppVersionSelectorConstants {
	return &v
}
