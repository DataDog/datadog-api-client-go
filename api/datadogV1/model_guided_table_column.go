// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableColumn - Definition of a single column in a guided table widget. A column can be a computed value, a preset comparison, or a formula.
type GuidedTableColumn struct {
	GuidedTableComputeColumn *GuidedTableComputeColumn
	GuidedTablePresetColumn  *GuidedTablePresetColumn
	GuidedTableFormulaColumn *GuidedTableFormulaColumn

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GuidedTableComputeColumnAsGuidedTableColumn is a convenience function that returns GuidedTableComputeColumn wrapped in GuidedTableColumn.
func GuidedTableComputeColumnAsGuidedTableColumn(v *GuidedTableComputeColumn) GuidedTableColumn {
	return GuidedTableColumn{GuidedTableComputeColumn: v}
}

// GuidedTablePresetColumnAsGuidedTableColumn is a convenience function that returns GuidedTablePresetColumn wrapped in GuidedTableColumn.
func GuidedTablePresetColumnAsGuidedTableColumn(v *GuidedTablePresetColumn) GuidedTableColumn {
	return GuidedTableColumn{GuidedTablePresetColumn: v}
}

// GuidedTableFormulaColumnAsGuidedTableColumn is a convenience function that returns GuidedTableFormulaColumn wrapped in GuidedTableColumn.
func GuidedTableFormulaColumnAsGuidedTableColumn(v *GuidedTableFormulaColumn) GuidedTableColumn {
	return GuidedTableColumn{GuidedTableFormulaColumn: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GuidedTableColumn) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GuidedTableComputeColumn
	err = datadog.Unmarshal(data, &obj.GuidedTableComputeColumn)
	if err == nil {
		if obj.GuidedTableComputeColumn != nil && obj.GuidedTableComputeColumn.UnparsedObject == nil {
			jsonGuidedTableComputeColumn, _ := datadog.Marshal(obj.GuidedTableComputeColumn)
			if string(jsonGuidedTableComputeColumn) == "{}" { // empty struct
				obj.GuidedTableComputeColumn = nil
			} else {
				match++
			}
		} else {
			obj.GuidedTableComputeColumn = nil
		}
	} else {
		obj.GuidedTableComputeColumn = nil
	}

	// try to unmarshal data into GuidedTablePresetColumn
	err = datadog.Unmarshal(data, &obj.GuidedTablePresetColumn)
	if err == nil {
		if obj.GuidedTablePresetColumn != nil && obj.GuidedTablePresetColumn.UnparsedObject == nil {
			jsonGuidedTablePresetColumn, _ := datadog.Marshal(obj.GuidedTablePresetColumn)
			if string(jsonGuidedTablePresetColumn) == "{}" { // empty struct
				obj.GuidedTablePresetColumn = nil
			} else {
				match++
			}
		} else {
			obj.GuidedTablePresetColumn = nil
		}
	} else {
		obj.GuidedTablePresetColumn = nil
	}

	// try to unmarshal data into GuidedTableFormulaColumn
	err = datadog.Unmarshal(data, &obj.GuidedTableFormulaColumn)
	if err == nil {
		if obj.GuidedTableFormulaColumn != nil && obj.GuidedTableFormulaColumn.UnparsedObject == nil {
			jsonGuidedTableFormulaColumn, _ := datadog.Marshal(obj.GuidedTableFormulaColumn)
			if string(jsonGuidedTableFormulaColumn) == "{}" { // empty struct
				obj.GuidedTableFormulaColumn = nil
			} else {
				match++
			}
		} else {
			obj.GuidedTableFormulaColumn = nil
		}
	} else {
		obj.GuidedTableFormulaColumn = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GuidedTableComputeColumn = nil
		obj.GuidedTablePresetColumn = nil
		obj.GuidedTableFormulaColumn = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GuidedTableColumn) MarshalJSON() ([]byte, error) {
	if obj.GuidedTableComputeColumn != nil {
		return datadog.Marshal(&obj.GuidedTableComputeColumn)
	}

	if obj.GuidedTablePresetColumn != nil {
		return datadog.Marshal(&obj.GuidedTablePresetColumn)
	}

	if obj.GuidedTableFormulaColumn != nil {
		return datadog.Marshal(&obj.GuidedTableFormulaColumn)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GuidedTableColumn) GetActualInstance() interface{} {
	if obj.GuidedTableComputeColumn != nil {
		return obj.GuidedTableComputeColumn
	}

	if obj.GuidedTablePresetColumn != nil {
		return obj.GuidedTablePresetColumn
	}

	if obj.GuidedTableFormulaColumn != nil {
		return obj.GuidedTableFormulaColumn
	}

	// all schemas are nil
	return nil
}
