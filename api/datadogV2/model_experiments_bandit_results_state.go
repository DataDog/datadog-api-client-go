// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditResultsState WAITING_FOR_FIRST_REFRESH until a refresh saves results. STATS_UNAVAILABLE when the latest saved refresh has no complete probabilities of being optimal. ENDED once the bandit is stopped or completed; its last estimates stay. Otherwise WARMUP or OPTIMIZING says which policy the bandit serves. A paused bandit keeps its state.
type ExperimentsBanditResultsState string

// List of ExperimentsBanditResultsState.
const (
	EXPERIMENTSBANDITRESULTSSTATE_WAITING_FOR_FIRST_REFRESH ExperimentsBanditResultsState = "WAITING_FOR_FIRST_REFRESH"
	EXPERIMENTSBANDITRESULTSSTATE_WARMUP                    ExperimentsBanditResultsState = "WARMUP"
	EXPERIMENTSBANDITRESULTSSTATE_OPTIMIZING                ExperimentsBanditResultsState = "OPTIMIZING"
	EXPERIMENTSBANDITRESULTSSTATE_STATS_UNAVAILABLE         ExperimentsBanditResultsState = "STATS_UNAVAILABLE"
	EXPERIMENTSBANDITRESULTSSTATE_ENDED                     ExperimentsBanditResultsState = "ENDED"
)

var allowedExperimentsBanditResultsStateEnumValues = []ExperimentsBanditResultsState{
	EXPERIMENTSBANDITRESULTSSTATE_WAITING_FOR_FIRST_REFRESH,
	EXPERIMENTSBANDITRESULTSSTATE_WARMUP,
	EXPERIMENTSBANDITRESULTSSTATE_OPTIMIZING,
	EXPERIMENTSBANDITRESULTSSTATE_STATS_UNAVAILABLE,
	EXPERIMENTSBANDITRESULTSSTATE_ENDED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsBanditResultsState) GetAllowedValues() []ExperimentsBanditResultsState {
	return allowedExperimentsBanditResultsStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsBanditResultsState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsBanditResultsState(value)
	return nil
}

// NewExperimentsBanditResultsStateFromValue returns a pointer to a valid ExperimentsBanditResultsState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsBanditResultsStateFromValue(v string) (*ExperimentsBanditResultsState, error) {
	ev := ExperimentsBanditResultsState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsBanditResultsState: valid values are %v", v, allowedExperimentsBanditResultsStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsBanditResultsState) IsValid() bool {
	for _, existing := range allowedExperimentsBanditResultsStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsBanditResultsState value.
func (v ExperimentsBanditResultsState) Ptr() *ExperimentsBanditResultsState {
	return &v
}
