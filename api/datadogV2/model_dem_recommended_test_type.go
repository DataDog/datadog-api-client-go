// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemRecommendedTestType The type identifier for a recommended synthetic test.
type DemRecommendedTestType string

// List of DemRecommendedTestType.
const (
	DEMRECOMMENDEDTESTTYPE_RECOMMENDED_TESTS DemRecommendedTestType = "recommended_tests"
)

var allowedDemRecommendedTestTypeEnumValues = []DemRecommendedTestType{
	DEMRECOMMENDEDTESTTYPE_RECOMMENDED_TESTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DemRecommendedTestType) GetAllowedValues() []DemRecommendedTestType {
	return allowedDemRecommendedTestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DemRecommendedTestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DemRecommendedTestType(value)
	return nil
}

// NewDemRecommendedTestTypeFromValue returns a pointer to a valid DemRecommendedTestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDemRecommendedTestTypeFromValue(v string) (*DemRecommendedTestType, error) {
	ev := DemRecommendedTestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DemRecommendedTestType: valid values are %v", v, allowedDemRecommendedTestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DemRecommendedTestType) IsValid() bool {
	for _, existing := range allowedDemRecommendedTestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DemRecommendedTestType value.
func (v DemRecommendedTestType) Ptr() *DemRecommendedTestType {
	return &v
}
