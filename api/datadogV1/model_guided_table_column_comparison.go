// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableColumnComparison - Comparison to display in a guided table column.
type GuidedTableColumnComparison struct {
	GuidedTableColumnComparisonWithSelf        *GuidedTableColumnComparisonWithSelf
	GuidedTableColumnComparisonWithOtherColumn *GuidedTableColumnComparisonWithOtherColumn

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GuidedTableColumnComparisonWithSelfAsGuidedTableColumnComparison is a convenience function that returns GuidedTableColumnComparisonWithSelf wrapped in GuidedTableColumnComparison.
func GuidedTableColumnComparisonWithSelfAsGuidedTableColumnComparison(v *GuidedTableColumnComparisonWithSelf) GuidedTableColumnComparison {
	return GuidedTableColumnComparison{GuidedTableColumnComparisonWithSelf: v}
}

// GuidedTableColumnComparisonWithOtherColumnAsGuidedTableColumnComparison is a convenience function that returns GuidedTableColumnComparisonWithOtherColumn wrapped in GuidedTableColumnComparison.
func GuidedTableColumnComparisonWithOtherColumnAsGuidedTableColumnComparison(v *GuidedTableColumnComparisonWithOtherColumn) GuidedTableColumnComparison {
	return GuidedTableColumnComparison{GuidedTableColumnComparisonWithOtherColumn: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GuidedTableColumnComparison) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GuidedTableColumnComparisonWithSelf
	err = datadog.Unmarshal(data, &obj.GuidedTableColumnComparisonWithSelf)
	if err == nil {
		if obj.GuidedTableColumnComparisonWithSelf != nil && obj.GuidedTableColumnComparisonWithSelf.UnparsedObject == nil {
			jsonGuidedTableColumnComparisonWithSelf, _ := datadog.Marshal(obj.GuidedTableColumnComparisonWithSelf)
			if string(jsonGuidedTableColumnComparisonWithSelf) == "{}" { // empty struct
				obj.GuidedTableColumnComparisonWithSelf = nil
			} else {
				match++
			}
		} else {
			obj.GuidedTableColumnComparisonWithSelf = nil
		}
	} else {
		obj.GuidedTableColumnComparisonWithSelf = nil
	}

	// try to unmarshal data into GuidedTableColumnComparisonWithOtherColumn
	err = datadog.Unmarshal(data, &obj.GuidedTableColumnComparisonWithOtherColumn)
	if err == nil {
		if obj.GuidedTableColumnComparisonWithOtherColumn != nil && obj.GuidedTableColumnComparisonWithOtherColumn.UnparsedObject == nil {
			jsonGuidedTableColumnComparisonWithOtherColumn, _ := datadog.Marshal(obj.GuidedTableColumnComparisonWithOtherColumn)
			if string(jsonGuidedTableColumnComparisonWithOtherColumn) == "{}" { // empty struct
				obj.GuidedTableColumnComparisonWithOtherColumn = nil
			} else {
				match++
			}
		} else {
			obj.GuidedTableColumnComparisonWithOtherColumn = nil
		}
	} else {
		obj.GuidedTableColumnComparisonWithOtherColumn = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GuidedTableColumnComparisonWithSelf = nil
		obj.GuidedTableColumnComparisonWithOtherColumn = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GuidedTableColumnComparison) MarshalJSON() ([]byte, error) {
	if obj.GuidedTableColumnComparisonWithSelf != nil {
		return datadog.Marshal(&obj.GuidedTableColumnComparisonWithSelf)
	}

	if obj.GuidedTableColumnComparisonWithOtherColumn != nil {
		return datadog.Marshal(&obj.GuidedTableColumnComparisonWithOtherColumn)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GuidedTableColumnComparison) GetActualInstance() interface{} {
	if obj.GuidedTableColumnComparisonWithSelf != nil {
		return obj.GuidedTableColumnComparisonWithSelf
	}

	if obj.GuidedTableColumnComparisonWithOtherColumn != nil {
		return obj.GuidedTableColumnComparisonWithOtherColumn
	}

	// all schemas are nil
	return nil
}
