// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetRestrictionRestrictionMode Controls the default data visibility for the product type. `standard` makes data visible
// to all users with appropriate product access. `default_hide` hides data by default and
// requires explicit grants for each dataset.
type DatasetRestrictionRestrictionMode string

// List of DatasetRestrictionRestrictionMode.
const (
	DATASETRESTRICTIONRESTRICTIONMODE_STANDARD     DatasetRestrictionRestrictionMode = "standard"
	DATASETRESTRICTIONRESTRICTIONMODE_DEFAULT_HIDE DatasetRestrictionRestrictionMode = "default_hide"
)

var allowedDatasetRestrictionRestrictionModeEnumValues = []DatasetRestrictionRestrictionMode{
	DATASETRESTRICTIONRESTRICTIONMODE_STANDARD,
	DATASETRESTRICTIONRESTRICTIONMODE_DEFAULT_HIDE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatasetRestrictionRestrictionMode) GetAllowedValues() []DatasetRestrictionRestrictionMode {
	return allowedDatasetRestrictionRestrictionModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatasetRestrictionRestrictionMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatasetRestrictionRestrictionMode(value)
	return nil
}

// NewDatasetRestrictionRestrictionModeFromValue returns a pointer to a valid DatasetRestrictionRestrictionMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatasetRestrictionRestrictionModeFromValue(v string) (*DatasetRestrictionRestrictionMode, error) {
	ev := DatasetRestrictionRestrictionMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatasetRestrictionRestrictionMode: valid values are %v", v, allowedDatasetRestrictionRestrictionModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatasetRestrictionRestrictionMode) IsValid() bool {
	for _, existing := range allowedDatasetRestrictionRestrictionModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatasetRestrictionRestrictionMode value.
func (v DatasetRestrictionRestrictionMode) Ptr() *DatasetRestrictionRestrictionMode {
	return &v
}
