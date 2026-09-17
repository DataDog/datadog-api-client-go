// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemJourneyTestSuiteType The type identifier for DEM journey test suites.
type DemJourneyTestSuiteType string

// List of DemJourneyTestSuiteType.
const (
	DEMJOURNEYTESTSUITETYPE_JOURNEY_TEST_SUITE DemJourneyTestSuiteType = "journey_test_suite"
)

var allowedDemJourneyTestSuiteTypeEnumValues = []DemJourneyTestSuiteType{
	DEMJOURNEYTESTSUITETYPE_JOURNEY_TEST_SUITE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DemJourneyTestSuiteType) GetAllowedValues() []DemJourneyTestSuiteType {
	return allowedDemJourneyTestSuiteTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DemJourneyTestSuiteType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DemJourneyTestSuiteType(value)
	return nil
}

// NewDemJourneyTestSuiteTypeFromValue returns a pointer to a valid DemJourneyTestSuiteType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDemJourneyTestSuiteTypeFromValue(v string) (*DemJourneyTestSuiteType, error) {
	ev := DemJourneyTestSuiteType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DemJourneyTestSuiteType: valid values are %v", v, allowedDemJourneyTestSuiteTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DemJourneyTestSuiteType) IsValid() bool {
	for _, existing := range allowedDemJourneyTestSuiteTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DemJourneyTestSuiteType value.
func (v DemJourneyTestSuiteType) Ptr() *DemJourneyTestSuiteType {
	return &v
}
