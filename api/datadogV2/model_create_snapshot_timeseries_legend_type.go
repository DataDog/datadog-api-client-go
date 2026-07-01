// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateSnapshotTimeseriesLegendType The legend display type for timeseries widgets. A value of `none` hides the legend entirely; omitting the field lets the frontend choose automatically.
type CreateSnapshotTimeseriesLegendType string

// List of CreateSnapshotTimeseriesLegendType.
const (
	CREATESNAPSHOTTIMESERIESLEGENDTYPE_COMPACT  CreateSnapshotTimeseriesLegendType = "compact"
	CREATESNAPSHOTTIMESERIESLEGENDTYPE_EXPANDED CreateSnapshotTimeseriesLegendType = "expanded"
	CREATESNAPSHOTTIMESERIESLEGENDTYPE_NONE     CreateSnapshotTimeseriesLegendType = "none"
)

var allowedCreateSnapshotTimeseriesLegendTypeEnumValues = []CreateSnapshotTimeseriesLegendType{
	CREATESNAPSHOTTIMESERIESLEGENDTYPE_COMPACT,
	CREATESNAPSHOTTIMESERIESLEGENDTYPE_EXPANDED,
	CREATESNAPSHOTTIMESERIESLEGENDTYPE_NONE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateSnapshotTimeseriesLegendType) GetAllowedValues() []CreateSnapshotTimeseriesLegendType {
	return allowedCreateSnapshotTimeseriesLegendTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateSnapshotTimeseriesLegendType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateSnapshotTimeseriesLegendType(value)
	return nil
}

// NewCreateSnapshotTimeseriesLegendTypeFromValue returns a pointer to a valid CreateSnapshotTimeseriesLegendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateSnapshotTimeseriesLegendTypeFromValue(v string) (*CreateSnapshotTimeseriesLegendType, error) {
	ev := CreateSnapshotTimeseriesLegendType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateSnapshotTimeseriesLegendType: valid values are %v", v, allowedCreateSnapshotTimeseriesLegendTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateSnapshotTimeseriesLegendType) IsValid() bool {
	for _, existing := range allowedCreateSnapshotTimeseriesLegendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateSnapshotTimeseriesLegendType value.
func (v CreateSnapshotTimeseriesLegendType) Ptr() *CreateSnapshotTimeseriesLegendType {
	return &v
}
