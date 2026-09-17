// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemCreateJourneyTestSuiteRequestType The resource type for a request to create a DEM journey test suite.
type DemCreateJourneyTestSuiteRequestType string

// List of DemCreateJourneyTestSuiteRequestType.
const (
	DEMCREATEJOURNEYTESTSUITEREQUESTTYPE_CREATE_TEST_SUITE_FOR_JOURNEY_REQUEST DemCreateJourneyTestSuiteRequestType = "create_test_suite_for_journey_request"
)

var allowedDemCreateJourneyTestSuiteRequestTypeEnumValues = []DemCreateJourneyTestSuiteRequestType{
	DEMCREATEJOURNEYTESTSUITEREQUESTTYPE_CREATE_TEST_SUITE_FOR_JOURNEY_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DemCreateJourneyTestSuiteRequestType) GetAllowedValues() []DemCreateJourneyTestSuiteRequestType {
	return allowedDemCreateJourneyTestSuiteRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DemCreateJourneyTestSuiteRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DemCreateJourneyTestSuiteRequestType(value)
	return nil
}

// NewDemCreateJourneyTestSuiteRequestTypeFromValue returns a pointer to a valid DemCreateJourneyTestSuiteRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDemCreateJourneyTestSuiteRequestTypeFromValue(v string) (*DemCreateJourneyTestSuiteRequestType, error) {
	ev := DemCreateJourneyTestSuiteRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DemCreateJourneyTestSuiteRequestType: valid values are %v", v, allowedDemCreateJourneyTestSuiteRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DemCreateJourneyTestSuiteRequestType) IsValid() bool {
	for _, existing := range allowedDemCreateJourneyTestSuiteRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DemCreateJourneyTestSuiteRequestType value.
func (v DemCreateJourneyTestSuiteRequestType) Ptr() *DemCreateJourneyTestSuiteRequestType {
	return &v
}
