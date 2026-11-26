// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CompositeSourceSubQuery - Test-only generated schema for the dashboard versioning proof of concept.
type CompositeSourceSubQuery struct {
	CompositeEventsPlatformSubQueryDefinition_20270101 *CompositeEventsPlatformSubQueryDefinition
	CompositeReferentialDataQuery_20270101             *CompositeReferentialDataQuery
	CompositeAudienceFiltersQuery_20270101             *CompositeAudienceFiltersQuery

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CompositeEventsPlatformSubQueryDefinition_20270101AsCompositeSourceSubQuery is a convenience function that returns CompositeEventsPlatformSubQueryDefinition wrapped in CompositeSourceSubQuery.
func CompositeEventsPlatformSubQueryDefinition_20270101AsCompositeSourceSubQuery(v *CompositeEventsPlatformSubQueryDefinition) CompositeSourceSubQuery {
	return CompositeSourceSubQuery{CompositeEventsPlatformSubQueryDefinition_20270101: v}
}

// CompositeReferentialDataQuery_20270101AsCompositeSourceSubQuery is a convenience function that returns CompositeReferentialDataQuery wrapped in CompositeSourceSubQuery.
func CompositeReferentialDataQuery_20270101AsCompositeSourceSubQuery(v *CompositeReferentialDataQuery) CompositeSourceSubQuery {
	return CompositeSourceSubQuery{CompositeReferentialDataQuery_20270101: v}
}

// CompositeAudienceFiltersQuery_20270101AsCompositeSourceSubQuery is a convenience function that returns CompositeAudienceFiltersQuery wrapped in CompositeSourceSubQuery.
func CompositeAudienceFiltersQuery_20270101AsCompositeSourceSubQuery(v *CompositeAudienceFiltersQuery) CompositeSourceSubQuery {
	return CompositeSourceSubQuery{CompositeAudienceFiltersQuery_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CompositeSourceSubQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CompositeEventsPlatformSubQueryDefinition_20270101
	err = datadog.Unmarshal(data, &obj.CompositeEventsPlatformSubQueryDefinition_20270101)
	if err == nil {
		if obj.CompositeEventsPlatformSubQueryDefinition_20270101 != nil && obj.CompositeEventsPlatformSubQueryDefinition_20270101.UnparsedObject == nil {
			jsonCompositeEventsPlatformSubQueryDefinition_20270101, _ := datadog.Marshal(obj.CompositeEventsPlatformSubQueryDefinition_20270101)
			if string(jsonCompositeEventsPlatformSubQueryDefinition_20270101) == "{}" { // empty struct
				obj.CompositeEventsPlatformSubQueryDefinition_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.CompositeEventsPlatformSubQueryDefinition_20270101 = nil
		}
	} else {
		obj.CompositeEventsPlatformSubQueryDefinition_20270101 = nil
	}

	// try to unmarshal data into CompositeReferentialDataQuery_20270101
	err = datadog.Unmarshal(data, &obj.CompositeReferentialDataQuery_20270101)
	if err == nil {
		if obj.CompositeReferentialDataQuery_20270101 != nil && obj.CompositeReferentialDataQuery_20270101.UnparsedObject == nil {
			jsonCompositeReferentialDataQuery_20270101, _ := datadog.Marshal(obj.CompositeReferentialDataQuery_20270101)
			if string(jsonCompositeReferentialDataQuery_20270101) == "{}" { // empty struct
				obj.CompositeReferentialDataQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.CompositeReferentialDataQuery_20270101 = nil
		}
	} else {
		obj.CompositeReferentialDataQuery_20270101 = nil
	}

	// try to unmarshal data into CompositeAudienceFiltersQuery_20270101
	err = datadog.Unmarshal(data, &obj.CompositeAudienceFiltersQuery_20270101)
	if err == nil {
		if obj.CompositeAudienceFiltersQuery_20270101 != nil && obj.CompositeAudienceFiltersQuery_20270101.UnparsedObject == nil {
			jsonCompositeAudienceFiltersQuery_20270101, _ := datadog.Marshal(obj.CompositeAudienceFiltersQuery_20270101)
			if string(jsonCompositeAudienceFiltersQuery_20270101) == "{}" { // empty struct
				obj.CompositeAudienceFiltersQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.CompositeAudienceFiltersQuery_20270101 = nil
		}
	} else {
		obj.CompositeAudienceFiltersQuery_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CompositeEventsPlatformSubQueryDefinition_20270101 = nil
		obj.CompositeReferentialDataQuery_20270101 = nil
		obj.CompositeAudienceFiltersQuery_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CompositeSourceSubQuery) MarshalJSON() ([]byte, error) {
	if obj.CompositeEventsPlatformSubQueryDefinition_20270101 != nil {
		return datadog.Marshal(&obj.CompositeEventsPlatformSubQueryDefinition_20270101)
	}

	if obj.CompositeReferentialDataQuery_20270101 != nil {
		return datadog.Marshal(&obj.CompositeReferentialDataQuery_20270101)
	}

	if obj.CompositeAudienceFiltersQuery_20270101 != nil {
		return datadog.Marshal(&obj.CompositeAudienceFiltersQuery_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CompositeSourceSubQuery) GetActualInstance() interface{} {
	if obj.CompositeEventsPlatformSubQueryDefinition_20270101 != nil {
		return obj.CompositeEventsPlatformSubQueryDefinition_20270101
	}

	if obj.CompositeReferentialDataQuery_20270101 != nil {
		return obj.CompositeReferentialDataQuery_20270101
	}

	if obj.CompositeAudienceFiltersQuery_20270101 != nil {
		return obj.CompositeAudienceFiltersQuery_20270101
	}

	// all schemas are nil
	return nil
}
