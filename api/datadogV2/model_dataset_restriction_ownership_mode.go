// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetRestrictionOwnershipMode Controls how dataset ownership is determined. `disabled` turns off ownership-based access
// entirely. `team_tag_based` assigns dataset ownership based on the team tags applied to the
// data, allowing team members to see their own team's datasets.
type DatasetRestrictionOwnershipMode string

// List of DatasetRestrictionOwnershipMode.
const (
	DATASETRESTRICTIONOWNERSHIPMODE_DISABLED       DatasetRestrictionOwnershipMode = "disabled"
	DATASETRESTRICTIONOWNERSHIPMODE_TEAM_TAG_BASED DatasetRestrictionOwnershipMode = "team_tag_based"
)

var allowedDatasetRestrictionOwnershipModeEnumValues = []DatasetRestrictionOwnershipMode{
	DATASETRESTRICTIONOWNERSHIPMODE_DISABLED,
	DATASETRESTRICTIONOWNERSHIPMODE_TEAM_TAG_BASED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatasetRestrictionOwnershipMode) GetAllowedValues() []DatasetRestrictionOwnershipMode {
	return allowedDatasetRestrictionOwnershipModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatasetRestrictionOwnershipMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatasetRestrictionOwnershipMode(value)
	return nil
}

// NewDatasetRestrictionOwnershipModeFromValue returns a pointer to a valid DatasetRestrictionOwnershipMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatasetRestrictionOwnershipModeFromValue(v string) (*DatasetRestrictionOwnershipMode, error) {
	ev := DatasetRestrictionOwnershipMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatasetRestrictionOwnershipMode: valid values are %v", v, allowedDatasetRestrictionOwnershipModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatasetRestrictionOwnershipMode) IsValid() bool {
	for _, existing := range allowedDatasetRestrictionOwnershipModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatasetRestrictionOwnershipMode value.
func (v DatasetRestrictionOwnershipMode) Ptr() *DatasetRestrictionOwnershipMode {
	return &v
}
