// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DetectionRuleStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type DetectionRuleStreamQueryDataSource string

// List of DetectionRuleStreamQueryDataSource.
const (
	DETECTIONRULESTREAMQUERYDATASOURCE_DETECTION_RULES DetectionRuleStreamQueryDataSource = "detection_rules"
)

var allowedDetectionRuleStreamQueryDataSourceEnumValues = []DetectionRuleStreamQueryDataSource{
	DETECTIONRULESTREAMQUERYDATASOURCE_DETECTION_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DetectionRuleStreamQueryDataSource) GetAllowedValues() []DetectionRuleStreamQueryDataSource {
	return allowedDetectionRuleStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DetectionRuleStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DetectionRuleStreamQueryDataSource(value)
	return nil
}

// NewDetectionRuleStreamQueryDataSourceFromValue returns a pointer to a valid DetectionRuleStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDetectionRuleStreamQueryDataSourceFromValue(v string) (*DetectionRuleStreamQueryDataSource, error) {
	ev := DetectionRuleStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DetectionRuleStreamQueryDataSource: valid values are %v", v, allowedDetectionRuleStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DetectionRuleStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedDetectionRuleStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DetectionRuleStreamQueryDataSource value.
func (v DetectionRuleStreamQueryDataSource) Ptr() *DetectionRuleStreamQueryDataSource {
	return &v
}
