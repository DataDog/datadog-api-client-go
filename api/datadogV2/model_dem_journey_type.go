// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemJourneyType The type identifier for DEM journeys.
type DemJourneyType string

// List of DemJourneyType.
const (
	DEMJOURNEYTYPE_JOURNEYS DemJourneyType = "journeys"
)

var allowedDemJourneyTypeEnumValues = []DemJourneyType{
	DEMJOURNEYTYPE_JOURNEYS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DemJourneyType) GetAllowedValues() []DemJourneyType {
	return allowedDemJourneyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DemJourneyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DemJourneyType(value)
	return nil
}

// NewDemJourneyTypeFromValue returns a pointer to a valid DemJourneyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDemJourneyTypeFromValue(v string) (*DemJourneyType, error) {
	ev := DemJourneyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DemJourneyType: valid values are %v", v, allowedDemJourneyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DemJourneyType) IsValid() bool {
	for _, existing := range allowedDemJourneyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DemJourneyType value.
func (v DemJourneyType) Ptr() *DemJourneyType {
	return &v
}
