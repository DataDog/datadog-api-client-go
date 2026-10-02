// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsMetricV2DTODataAttributesDesiredChange Direction of metric change considered desirable.
type ExperimentsMetricV2DTODataAttributesDesiredChange string

// List of ExperimentsMetricV2DTODataAttributesDesiredChange.
const (
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESDESIREDCHANGE_METRIC_INCREASES ExperimentsMetricV2DTODataAttributesDesiredChange = "METRIC_INCREASES"
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESDESIREDCHANGE_METRIC_DECREASES ExperimentsMetricV2DTODataAttributesDesiredChange = "METRIC_DECREASES"
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESDESIREDCHANGE_UNKNOWN          ExperimentsMetricV2DTODataAttributesDesiredChange = "UNKNOWN"
)

var allowedExperimentsMetricV2DTODataAttributesDesiredChangeEnumValues = []ExperimentsMetricV2DTODataAttributesDesiredChange{
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESDESIREDCHANGE_METRIC_INCREASES,
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESDESIREDCHANGE_METRIC_DECREASES,
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESDESIREDCHANGE_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsMetricV2DTODataAttributesDesiredChange) GetAllowedValues() []ExperimentsMetricV2DTODataAttributesDesiredChange {
	return allowedExperimentsMetricV2DTODataAttributesDesiredChangeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsMetricV2DTODataAttributesDesiredChange) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsMetricV2DTODataAttributesDesiredChange(value)
	return nil
}

// NewExperimentsMetricV2DTODataAttributesDesiredChangeFromValue returns a pointer to a valid ExperimentsMetricV2DTODataAttributesDesiredChange
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsMetricV2DTODataAttributesDesiredChangeFromValue(v string) (*ExperimentsMetricV2DTODataAttributesDesiredChange, error) {
	ev := ExperimentsMetricV2DTODataAttributesDesiredChange(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsMetricV2DTODataAttributesDesiredChange: valid values are %v", v, allowedExperimentsMetricV2DTODataAttributesDesiredChangeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsMetricV2DTODataAttributesDesiredChange) IsValid() bool {
	for _, existing := range allowedExperimentsMetricV2DTODataAttributesDesiredChangeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsMetricV2DTODataAttributesDesiredChange value.
func (v ExperimentsMetricV2DTODataAttributesDesiredChange) Ptr() *ExperimentsMetricV2DTODataAttributesDesiredChange {
	return &v
}
