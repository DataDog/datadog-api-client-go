// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardExperienceType The experience type of the dashboard.
type DashboardExperienceType string

// List of DashboardExperienceType.
const (
	DASHBOARDEXPERIENCETYPE_DEFAULT           DashboardExperienceType = "default"
	DASHBOARDEXPERIENCETYPE_PRODUCT_ANALYTICS DashboardExperienceType = "product_analytics"
)

var allowedDashboardExperienceTypeEnumValues = []DashboardExperienceType{
	DASHBOARDEXPERIENCETYPE_DEFAULT,
	DASHBOARDEXPERIENCETYPE_PRODUCT_ANALYTICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardExperienceType) GetAllowedValues() []DashboardExperienceType {
	return allowedDashboardExperienceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardExperienceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardExperienceType(value)
	return nil
}

// NewDashboardExperienceTypeFromValue returns a pointer to a valid DashboardExperienceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardExperienceTypeFromValue(v string) (*DashboardExperienceType, error) {
	ev := DashboardExperienceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardExperienceType: valid values are %v", v, allowedDashboardExperienceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardExperienceType) IsValid() bool {
	for _, existing := range allowedDashboardExperienceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardExperienceType value.
func (v DashboardExperienceType) Ptr() *DashboardExperienceType {
	return &v
}
