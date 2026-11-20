// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BackendDistributionYAxisScale Test-only generated schema for the dashboard versioning proof of concept.
type BackendDistributionYAxisScale string

// List of BackendDistributionYAxisScale.
const (
	BACKENDDISTRIBUTIONYAXISSCALE_LINEAR BackendDistributionYAxisScale = "linear"
	BACKENDDISTRIBUTIONYAXISSCALE_LOG    BackendDistributionYAxisScale = "log"
)

var allowedBackendDistributionYAxisScaleEnumValues = []BackendDistributionYAxisScale{
	BACKENDDISTRIBUTIONYAXISSCALE_LINEAR,
	BACKENDDISTRIBUTIONYAXISSCALE_LOG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BackendDistributionYAxisScale) GetAllowedValues() []BackendDistributionYAxisScale {
	return allowedBackendDistributionYAxisScaleEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BackendDistributionYAxisScale) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BackendDistributionYAxisScale(value)
	return nil
}

// NewBackendDistributionYAxisScaleFromValue returns a pointer to a valid BackendDistributionYAxisScale
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBackendDistributionYAxisScaleFromValue(v string) (*BackendDistributionYAxisScale, error) {
	ev := BackendDistributionYAxisScale(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BackendDistributionYAxisScale: valid values are %v", v, allowedBackendDistributionYAxisScaleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BackendDistributionYAxisScale) IsValid() bool {
	for _, existing := range allowedBackendDistributionYAxisScaleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackendDistributionYAxisScale value.
func (v BackendDistributionYAxisScale) Ptr() *BackendDistributionYAxisScale {
	return &v
}
