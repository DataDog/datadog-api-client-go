// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemInferredJourneyItem - A single inferred journey item, either a candidate or an ignored journey.
type DemInferredJourneyItem struct {
	DemInferredJourneyCandidateData *DemInferredJourneyCandidateData
	DemInferredJourneyIgnoredData   *DemInferredJourneyIgnoredData

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DemInferredJourneyCandidateDataAsDemInferredJourneyItem is a convenience function that returns DemInferredJourneyCandidateData wrapped in DemInferredJourneyItem.
func DemInferredJourneyCandidateDataAsDemInferredJourneyItem(v *DemInferredJourneyCandidateData) DemInferredJourneyItem {
	return DemInferredJourneyItem{DemInferredJourneyCandidateData: v}
}

// DemInferredJourneyIgnoredDataAsDemInferredJourneyItem is a convenience function that returns DemInferredJourneyIgnoredData wrapped in DemInferredJourneyItem.
func DemInferredJourneyIgnoredDataAsDemInferredJourneyItem(v *DemInferredJourneyIgnoredData) DemInferredJourneyItem {
	return DemInferredJourneyItem{DemInferredJourneyIgnoredData: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DemInferredJourneyItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DemInferredJourneyCandidateData
	err = datadog.Unmarshal(data, &obj.DemInferredJourneyCandidateData)
	if err == nil {
		if obj.DemInferredJourneyCandidateData != nil && obj.DemInferredJourneyCandidateData.UnparsedObject == nil {
			jsonDemInferredJourneyCandidateData, _ := datadog.Marshal(obj.DemInferredJourneyCandidateData)
			if string(jsonDemInferredJourneyCandidateData) == "{}" { // empty struct
				obj.DemInferredJourneyCandidateData = nil
			} else {
				match++
			}
		} else {
			obj.DemInferredJourneyCandidateData = nil
		}
	} else {
		obj.DemInferredJourneyCandidateData = nil
	}

	// try to unmarshal data into DemInferredJourneyIgnoredData
	err = datadog.Unmarshal(data, &obj.DemInferredJourneyIgnoredData)
	if err == nil {
		if obj.DemInferredJourneyIgnoredData != nil && obj.DemInferredJourneyIgnoredData.UnparsedObject == nil {
			jsonDemInferredJourneyIgnoredData, _ := datadog.Marshal(obj.DemInferredJourneyIgnoredData)
			if string(jsonDemInferredJourneyIgnoredData) == "{}" { // empty struct
				obj.DemInferredJourneyIgnoredData = nil
			} else {
				match++
			}
		} else {
			obj.DemInferredJourneyIgnoredData = nil
		}
	} else {
		obj.DemInferredJourneyIgnoredData = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DemInferredJourneyCandidateData = nil
		obj.DemInferredJourneyIgnoredData = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DemInferredJourneyItem) MarshalJSON() ([]byte, error) {
	if obj.DemInferredJourneyCandidateData != nil {
		return datadog.Marshal(&obj.DemInferredJourneyCandidateData)
	}

	if obj.DemInferredJourneyIgnoredData != nil {
		return datadog.Marshal(&obj.DemInferredJourneyIgnoredData)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DemInferredJourneyItem) GetActualInstance() interface{} {
	if obj.DemInferredJourneyCandidateData != nil {
		return obj.DemInferredJourneyCandidateData
	}

	if obj.DemInferredJourneyIgnoredData != nil {
		return obj.DemInferredJourneyIgnoredData
	}

	// all schemas are nil
	return nil
}
