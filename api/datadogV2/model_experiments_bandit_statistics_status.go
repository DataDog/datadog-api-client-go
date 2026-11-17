// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditStatisticsStatus Whether the latest saved refresh calculated its statistics. Absent for older refreshes.
type ExperimentsBanditStatisticsStatus string

// List of ExperimentsBanditStatisticsStatus.
const (
	EXPERIMENTSBANDITSTATISTICSSTATUS_CALCULATED ExperimentsBanditStatisticsStatus = "calculated"
	EXPERIMENTSBANDITSTATISTICSSTATUS_FAILED     ExperimentsBanditStatisticsStatus = "failed"
)

var allowedExperimentsBanditStatisticsStatusEnumValues = []ExperimentsBanditStatisticsStatus{
	EXPERIMENTSBANDITSTATISTICSSTATUS_CALCULATED,
	EXPERIMENTSBANDITSTATISTICSSTATUS_FAILED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsBanditStatisticsStatus) GetAllowedValues() []ExperimentsBanditStatisticsStatus {
	return allowedExperimentsBanditStatisticsStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsBanditStatisticsStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsBanditStatisticsStatus(value)
	return nil
}

// NewExperimentsBanditStatisticsStatusFromValue returns a pointer to a valid ExperimentsBanditStatisticsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsBanditStatisticsStatusFromValue(v string) (*ExperimentsBanditStatisticsStatus, error) {
	ev := ExperimentsBanditStatisticsStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsBanditStatisticsStatus: valid values are %v", v, allowedExperimentsBanditStatisticsStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsBanditStatisticsStatus) IsValid() bool {
	for _, existing := range allowedExperimentsBanditStatisticsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsBanditStatisticsStatus value.
func (v ExperimentsBanditStatisticsStatus) Ptr() *ExperimentsBanditStatisticsStatus {
	return &v
}
