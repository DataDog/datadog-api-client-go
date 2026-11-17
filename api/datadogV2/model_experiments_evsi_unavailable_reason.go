// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsEVSIUnavailableReason Recorded reason this analysis has no EVSI estimate. Omitted when unrecorded; absence does not establish availability for historical or imported results. Controls and non-Bayesian analyses omit this field.
type ExperimentsEVSIUnavailableReason string

// List of ExperimentsEVSIUnavailableReason.
const (
	EXPERIMENTSEVSIUNAVAILABLEREASON_TRAFFIC_UNAVAILABLE        ExperimentsEVSIUnavailableReason = "traffic_unavailable"
	EXPERIMENTSEVSIUNAVAILABLEREASON_INSUFFICIENT_COMPLETE_DAYS ExperimentsEVSIUnavailableReason = "insufficient_complete_days"
	EXPERIMENTSEVSIUNAVAILABLEREASON_NO_TOTAL_EXPOSURES         ExperimentsEVSIUnavailableReason = "no_total_exposures"
	EXPERIMENTSEVSIUNAVAILABLEREASON_NO_LOOKBACK_EXPOSURES      ExperimentsEVSIUnavailableReason = "no_lookback_exposures"
	EXPERIMENTSEVSIUNAVAILABLEREASON_INVALID_TRAFFIC            ExperimentsEVSIUnavailableReason = "invalid_traffic"
	EXPERIMENTSEVSIUNAVAILABLEREASON_INVALID_STANDARD_ERROR     ExperimentsEVSIUnavailableReason = "invalid_standard_error"
	EXPERIMENTSEVSIUNAVAILABLEREASON_COMPUTATION_FAILED         ExperimentsEVSIUnavailableReason = "computation_failed"
	EXPERIMENTSEVSIUNAVAILABLEREASON_ANALYSIS_UNAVAILABLE       ExperimentsEVSIUnavailableReason = "analysis_unavailable"
)

var allowedExperimentsEVSIUnavailableReasonEnumValues = []ExperimentsEVSIUnavailableReason{
	EXPERIMENTSEVSIUNAVAILABLEREASON_TRAFFIC_UNAVAILABLE,
	EXPERIMENTSEVSIUNAVAILABLEREASON_INSUFFICIENT_COMPLETE_DAYS,
	EXPERIMENTSEVSIUNAVAILABLEREASON_NO_TOTAL_EXPOSURES,
	EXPERIMENTSEVSIUNAVAILABLEREASON_NO_LOOKBACK_EXPOSURES,
	EXPERIMENTSEVSIUNAVAILABLEREASON_INVALID_TRAFFIC,
	EXPERIMENTSEVSIUNAVAILABLEREASON_INVALID_STANDARD_ERROR,
	EXPERIMENTSEVSIUNAVAILABLEREASON_COMPUTATION_FAILED,
	EXPERIMENTSEVSIUNAVAILABLEREASON_ANALYSIS_UNAVAILABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsEVSIUnavailableReason) GetAllowedValues() []ExperimentsEVSIUnavailableReason {
	return allowedExperimentsEVSIUnavailableReasonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsEVSIUnavailableReason) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsEVSIUnavailableReason(value)
	return nil
}

// NewExperimentsEVSIUnavailableReasonFromValue returns a pointer to a valid ExperimentsEVSIUnavailableReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsEVSIUnavailableReasonFromValue(v string) (*ExperimentsEVSIUnavailableReason, error) {
	ev := ExperimentsEVSIUnavailableReason(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsEVSIUnavailableReason: valid values are %v", v, allowedExperimentsEVSIUnavailableReasonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsEVSIUnavailableReason) IsValid() bool {
	for _, existing := range allowedExperimentsEVSIUnavailableReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsEVSIUnavailableReason value.
func (v ExperimentsEVSIUnavailableReason) Ptr() *ExperimentsEVSIUnavailableReason {
	return &v
}
