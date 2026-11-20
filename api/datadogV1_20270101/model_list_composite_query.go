// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListCompositeQuery - Test-only generated schema for the dashboard versioning proof of concept.
type ListCompositeQuery struct {
	ListFilteredQuery_20270101  *ListFilteredQuery
	ListAugmentedQuery_20270101 *ListAugmentedQuery
	ListSegmentQuery_20270101   *ListSegmentQuery

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ListFilteredQuery_20270101AsListCompositeQuery is a convenience function that returns ListFilteredQuery wrapped in ListCompositeQuery.
func ListFilteredQuery_20270101AsListCompositeQuery(v *ListFilteredQuery) ListCompositeQuery {
	return ListCompositeQuery{ListFilteredQuery_20270101: v}
}

// ListAugmentedQuery_20270101AsListCompositeQuery is a convenience function that returns ListAugmentedQuery wrapped in ListCompositeQuery.
func ListAugmentedQuery_20270101AsListCompositeQuery(v *ListAugmentedQuery) ListCompositeQuery {
	return ListCompositeQuery{ListAugmentedQuery_20270101: v}
}

// ListSegmentQuery_20270101AsListCompositeQuery is a convenience function that returns ListSegmentQuery wrapped in ListCompositeQuery.
func ListSegmentQuery_20270101AsListCompositeQuery(v *ListSegmentQuery) ListCompositeQuery {
	return ListCompositeQuery{ListSegmentQuery_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ListCompositeQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListFilteredQuery_20270101
	err = datadog.Unmarshal(data, &obj.ListFilteredQuery_20270101)
	if err == nil {
		if obj.ListFilteredQuery_20270101 != nil && obj.ListFilteredQuery_20270101.UnparsedObject == nil {
			jsonListFilteredQuery_20270101, _ := datadog.Marshal(obj.ListFilteredQuery_20270101)
			if string(jsonListFilteredQuery_20270101) == "{}" { // empty struct
				obj.ListFilteredQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ListFilteredQuery_20270101 = nil
		}
	} else {
		obj.ListFilteredQuery_20270101 = nil
	}

	// try to unmarshal data into ListAugmentedQuery_20270101
	err = datadog.Unmarshal(data, &obj.ListAugmentedQuery_20270101)
	if err == nil {
		if obj.ListAugmentedQuery_20270101 != nil && obj.ListAugmentedQuery_20270101.UnparsedObject == nil {
			jsonListAugmentedQuery_20270101, _ := datadog.Marshal(obj.ListAugmentedQuery_20270101)
			if string(jsonListAugmentedQuery_20270101) == "{}" { // empty struct
				obj.ListAugmentedQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ListAugmentedQuery_20270101 = nil
		}
	} else {
		obj.ListAugmentedQuery_20270101 = nil
	}

	// try to unmarshal data into ListSegmentQuery_20270101
	err = datadog.Unmarshal(data, &obj.ListSegmentQuery_20270101)
	if err == nil {
		if obj.ListSegmentQuery_20270101 != nil && obj.ListSegmentQuery_20270101.UnparsedObject == nil {
			jsonListSegmentQuery_20270101, _ := datadog.Marshal(obj.ListSegmentQuery_20270101)
			if string(jsonListSegmentQuery_20270101) == "{}" { // empty struct
				obj.ListSegmentQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ListSegmentQuery_20270101 = nil
		}
	} else {
		obj.ListSegmentQuery_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ListFilteredQuery_20270101 = nil
		obj.ListAugmentedQuery_20270101 = nil
		obj.ListSegmentQuery_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ListCompositeQuery) MarshalJSON() ([]byte, error) {
	if obj.ListFilteredQuery_20270101 != nil {
		return datadog.Marshal(&obj.ListFilteredQuery_20270101)
	}

	if obj.ListAugmentedQuery_20270101 != nil {
		return datadog.Marshal(&obj.ListAugmentedQuery_20270101)
	}

	if obj.ListSegmentQuery_20270101 != nil {
		return datadog.Marshal(&obj.ListSegmentQuery_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ListCompositeQuery) GetActualInstance() interface{} {
	if obj.ListFilteredQuery_20270101 != nil {
		return obj.ListFilteredQuery_20270101
	}

	if obj.ListAugmentedQuery_20270101 != nil {
		return obj.ListAugmentedQuery_20270101
	}

	if obj.ListSegmentQuery_20270101 != nil {
		return obj.ListSegmentQuery_20270101
	}

	// all schemas are nil
	return nil
}
