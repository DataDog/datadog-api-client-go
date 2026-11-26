// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionScaleType Test-only generated schema for the dashboard versioning proof of concept.
type DistributionScaleType string

// List of DistributionScaleType.
const (
	DISTRIBUTIONSCALETYPE_LINEAR DistributionScaleType = "linear"
	DISTRIBUTIONSCALETYPE_LOG    DistributionScaleType = "log"
)

var allowedDistributionScaleTypeEnumValues = []DistributionScaleType{
	DISTRIBUTIONSCALETYPE_LINEAR,
	DISTRIBUTIONSCALETYPE_LOG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DistributionScaleType) GetAllowedValues() []DistributionScaleType {
	return allowedDistributionScaleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DistributionScaleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DistributionScaleType(value)
	return nil
}

// NewDistributionScaleTypeFromValue returns a pointer to a valid DistributionScaleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDistributionScaleTypeFromValue(v string) (*DistributionScaleType, error) {
	ev := DistributionScaleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DistributionScaleType: valid values are %v", v, allowedDistributionScaleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DistributionScaleType) IsValid() bool {
	for _, existing := range allowedDistributionScaleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DistributionScaleType value.
func (v DistributionScaleType) Ptr() *DistributionScaleType {
	return &v
}
