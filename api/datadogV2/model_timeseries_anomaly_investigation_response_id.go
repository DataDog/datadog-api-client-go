// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationResponseID Stable identifier for an anomaly investigation response resource.
type TimeseriesAnomalyInvestigationResponseID string

// List of TimeseriesAnomalyInvestigationResponseID.
const (
	TIMESERIESANOMALYINVESTIGATIONRESPONSEID_ZERO TimeseriesAnomalyInvestigationResponseID = "0"
)

var allowedTimeseriesAnomalyInvestigationResponseIDEnumValues = []TimeseriesAnomalyInvestigationResponseID{
	TIMESERIESANOMALYINVESTIGATIONRESPONSEID_ZERO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesAnomalyInvestigationResponseID) GetAllowedValues() []TimeseriesAnomalyInvestigationResponseID {
	return allowedTimeseriesAnomalyInvestigationResponseIDEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesAnomalyInvestigationResponseID) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesAnomalyInvestigationResponseID(value)
	return nil
}

// NewTimeseriesAnomalyInvestigationResponseIDFromValue returns a pointer to a valid TimeseriesAnomalyInvestigationResponseID
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesAnomalyInvestigationResponseIDFromValue(v string) (*TimeseriesAnomalyInvestigationResponseID, error) {
	ev := TimeseriesAnomalyInvestigationResponseID(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesAnomalyInvestigationResponseID: valid values are %v", v, allowedTimeseriesAnomalyInvestigationResponseIDEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesAnomalyInvestigationResponseID) IsValid() bool {
	for _, existing := range allowedTimeseriesAnomalyInvestigationResponseIDEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesAnomalyInvestigationResponseID value.
func (v TimeseriesAnomalyInvestigationResponseID) Ptr() *TimeseriesAnomalyInvestigationResponseID {
	return &v
}
