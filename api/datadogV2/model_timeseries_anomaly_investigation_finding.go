// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationFinding - Deterministic explanation for a detected anomaly.
type TimeseriesAnomalyInvestigationFinding struct {
	TimeseriesAnomalyInvestigationInfluentialTagFinding *TimeseriesAnomalyInvestigationInfluentialTagFinding
	TimeseriesAnomalyInvestigationAnomalyFinding        *TimeseriesAnomalyInvestigationAnomalyFinding

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TimeseriesAnomalyInvestigationInfluentialTagFindingAsTimeseriesAnomalyInvestigationFinding is a convenience function that returns TimeseriesAnomalyInvestigationInfluentialTagFinding wrapped in TimeseriesAnomalyInvestigationFinding.
func TimeseriesAnomalyInvestigationInfluentialTagFindingAsTimeseriesAnomalyInvestigationFinding(v *TimeseriesAnomalyInvestigationInfluentialTagFinding) TimeseriesAnomalyInvestigationFinding {
	return TimeseriesAnomalyInvestigationFinding{TimeseriesAnomalyInvestigationInfluentialTagFinding: v}
}

// TimeseriesAnomalyInvestigationAnomalyFindingAsTimeseriesAnomalyInvestigationFinding is a convenience function that returns TimeseriesAnomalyInvestigationAnomalyFinding wrapped in TimeseriesAnomalyInvestigationFinding.
func TimeseriesAnomalyInvestigationAnomalyFindingAsTimeseriesAnomalyInvestigationFinding(v *TimeseriesAnomalyInvestigationAnomalyFinding) TimeseriesAnomalyInvestigationFinding {
	return TimeseriesAnomalyInvestigationFinding{TimeseriesAnomalyInvestigationAnomalyFinding: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TimeseriesAnomalyInvestigationFinding) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TimeseriesAnomalyInvestigationInfluentialTagFinding
	err = datadog.Unmarshal(data, &obj.TimeseriesAnomalyInvestigationInfluentialTagFinding)
	if err == nil {
		if obj.TimeseriesAnomalyInvestigationInfluentialTagFinding != nil && obj.TimeseriesAnomalyInvestigationInfluentialTagFinding.UnparsedObject == nil {
			jsonTimeseriesAnomalyInvestigationInfluentialTagFinding, _ := datadog.Marshal(obj.TimeseriesAnomalyInvestigationInfluentialTagFinding)
			if string(jsonTimeseriesAnomalyInvestigationInfluentialTagFinding) == "{}" { // empty struct
				obj.TimeseriesAnomalyInvestigationInfluentialTagFinding = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesAnomalyInvestigationInfluentialTagFinding = nil
		}
	} else {
		obj.TimeseriesAnomalyInvestigationInfluentialTagFinding = nil
	}

	// try to unmarshal data into TimeseriesAnomalyInvestigationAnomalyFinding
	err = datadog.Unmarshal(data, &obj.TimeseriesAnomalyInvestigationAnomalyFinding)
	if err == nil {
		if obj.TimeseriesAnomalyInvestigationAnomalyFinding != nil && obj.TimeseriesAnomalyInvestigationAnomalyFinding.UnparsedObject == nil {
			jsonTimeseriesAnomalyInvestigationAnomalyFinding, _ := datadog.Marshal(obj.TimeseriesAnomalyInvestigationAnomalyFinding)
			if string(jsonTimeseriesAnomalyInvestigationAnomalyFinding) == "{}" { // empty struct
				obj.TimeseriesAnomalyInvestigationAnomalyFinding = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesAnomalyInvestigationAnomalyFinding = nil
		}
	} else {
		obj.TimeseriesAnomalyInvestigationAnomalyFinding = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TimeseriesAnomalyInvestigationInfluentialTagFinding = nil
		obj.TimeseriesAnomalyInvestigationAnomalyFinding = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TimeseriesAnomalyInvestigationFinding) MarshalJSON() ([]byte, error) {
	if obj.TimeseriesAnomalyInvestigationInfluentialTagFinding != nil {
		return datadog.Marshal(&obj.TimeseriesAnomalyInvestigationInfluentialTagFinding)
	}

	if obj.TimeseriesAnomalyInvestigationAnomalyFinding != nil {
		return datadog.Marshal(&obj.TimeseriesAnomalyInvestigationAnomalyFinding)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TimeseriesAnomalyInvestigationFinding) GetActualInstance() interface{} {
	if obj.TimeseriesAnomalyInvestigationInfluentialTagFinding != nil {
		return obj.TimeseriesAnomalyInvestigationInfluentialTagFinding
	}

	if obj.TimeseriesAnomalyInvestigationAnomalyFinding != nil {
		return obj.TimeseriesAnomalyInvestigationAnomalyFinding
	}

	// all schemas are nil
	return nil
}
