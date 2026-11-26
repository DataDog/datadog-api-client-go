// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HistogramMode Test-only generated schema for the dashboard versioning proof of concept.
type HistogramMode string

// List of HistogramMode.
const (
	HISTOGRAMMODE_DISTRIBUTION_METRIC HistogramMode = "distribution_metric"
	HISTOGRAMMODE_PRE_BINNED_METRIC   HistogramMode = "pre_binned_metric"
)

var allowedHistogramModeEnumValues = []HistogramMode{
	HISTOGRAMMODE_DISTRIBUTION_METRIC,
	HISTOGRAMMODE_PRE_BINNED_METRIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HistogramMode) GetAllowedValues() []HistogramMode {
	return allowedHistogramModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HistogramMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HistogramMode(value)
	return nil
}

// NewHistogramModeFromValue returns a pointer to a valid HistogramMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHistogramModeFromValue(v string) (*HistogramMode, error) {
	ev := HistogramMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HistogramMode: valid values are %v", v, allowedHistogramModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HistogramMode) IsValid() bool {
	for _, existing := range allowedHistogramModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HistogramMode value.
func (v HistogramMode) Ptr() *HistogramMode {
	return &v
}
