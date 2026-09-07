// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemInferredJourneyStatus The status of an inferred DEM journey.
type DemInferredJourneyStatus string

// List of DemInferredJourneyStatus.
const (
	DEMINFERREDJOURNEYSTATUS_CANDIDATE DemInferredJourneyStatus = "candidate"
	DEMINFERREDJOURNEYSTATUS_IGNORED   DemInferredJourneyStatus = "ignored"
)

var allowedDemInferredJourneyStatusEnumValues = []DemInferredJourneyStatus{
	DEMINFERREDJOURNEYSTATUS_CANDIDATE,
	DEMINFERREDJOURNEYSTATUS_IGNORED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DemInferredJourneyStatus) GetAllowedValues() []DemInferredJourneyStatus {
	return allowedDemInferredJourneyStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DemInferredJourneyStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DemInferredJourneyStatus(value)
	return nil
}

// NewDemInferredJourneyStatusFromValue returns a pointer to a valid DemInferredJourneyStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDemInferredJourneyStatusFromValue(v string) (*DemInferredJourneyStatus, error) {
	ev := DemInferredJourneyStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DemInferredJourneyStatus: valid values are %v", v, allowedDemInferredJourneyStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DemInferredJourneyStatus) IsValid() bool {
	for _, existing := range allowedDemInferredJourneyStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DemInferredJourneyStatus value.
func (v DemInferredJourneyStatus) Ptr() *DemInferredJourneyStatus {
	return &v
}
