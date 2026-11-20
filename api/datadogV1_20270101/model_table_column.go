// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableColumn - Test-only generated schema for the dashboard versioning proof of concept.
type TableColumn struct {
	ComputeTableColumn_20270101      *ComputeTableColumn
	PresetPeriodTableColumn_20270101 *PresetPeriodTableColumn
	FormulaTableColumn_20270101      *FormulaTableColumn
	ReferenceTableColumn_20270101    *ReferenceTableColumn

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ComputeTableColumn_20270101AsTableColumn is a convenience function that returns ComputeTableColumn wrapped in TableColumn.
func ComputeTableColumn_20270101AsTableColumn(v *ComputeTableColumn) TableColumn {
	return TableColumn{ComputeTableColumn_20270101: v}
}

// PresetPeriodTableColumn_20270101AsTableColumn is a convenience function that returns PresetPeriodTableColumn wrapped in TableColumn.
func PresetPeriodTableColumn_20270101AsTableColumn(v *PresetPeriodTableColumn) TableColumn {
	return TableColumn{PresetPeriodTableColumn_20270101: v}
}

// FormulaTableColumn_20270101AsTableColumn is a convenience function that returns FormulaTableColumn wrapped in TableColumn.
func FormulaTableColumn_20270101AsTableColumn(v *FormulaTableColumn) TableColumn {
	return TableColumn{FormulaTableColumn_20270101: v}
}

// ReferenceTableColumn_20270101AsTableColumn is a convenience function that returns ReferenceTableColumn wrapped in TableColumn.
func ReferenceTableColumn_20270101AsTableColumn(v *ReferenceTableColumn) TableColumn {
	return TableColumn{ReferenceTableColumn_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TableColumn) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ComputeTableColumn_20270101
	err = datadog.Unmarshal(data, &obj.ComputeTableColumn_20270101)
	if err == nil {
		if obj.ComputeTableColumn_20270101 != nil && obj.ComputeTableColumn_20270101.UnparsedObject == nil {
			jsonComputeTableColumn_20270101, _ := datadog.Marshal(obj.ComputeTableColumn_20270101)
			if string(jsonComputeTableColumn_20270101) == "{}" { // empty struct
				obj.ComputeTableColumn_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ComputeTableColumn_20270101 = nil
		}
	} else {
		obj.ComputeTableColumn_20270101 = nil
	}

	// try to unmarshal data into PresetPeriodTableColumn_20270101
	err = datadog.Unmarshal(data, &obj.PresetPeriodTableColumn_20270101)
	if err == nil {
		if obj.PresetPeriodTableColumn_20270101 != nil && obj.PresetPeriodTableColumn_20270101.UnparsedObject == nil {
			jsonPresetPeriodTableColumn_20270101, _ := datadog.Marshal(obj.PresetPeriodTableColumn_20270101)
			if string(jsonPresetPeriodTableColumn_20270101) == "{}" { // empty struct
				obj.PresetPeriodTableColumn_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.PresetPeriodTableColumn_20270101 = nil
		}
	} else {
		obj.PresetPeriodTableColumn_20270101 = nil
	}

	// try to unmarshal data into FormulaTableColumn_20270101
	err = datadog.Unmarshal(data, &obj.FormulaTableColumn_20270101)
	if err == nil {
		if obj.FormulaTableColumn_20270101 != nil && obj.FormulaTableColumn_20270101.UnparsedObject == nil {
			jsonFormulaTableColumn_20270101, _ := datadog.Marshal(obj.FormulaTableColumn_20270101)
			if string(jsonFormulaTableColumn_20270101) == "{}" { // empty struct
				obj.FormulaTableColumn_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaTableColumn_20270101 = nil
		}
	} else {
		obj.FormulaTableColumn_20270101 = nil
	}

	// try to unmarshal data into ReferenceTableColumn_20270101
	err = datadog.Unmarshal(data, &obj.ReferenceTableColumn_20270101)
	if err == nil {
		if obj.ReferenceTableColumn_20270101 != nil && obj.ReferenceTableColumn_20270101.UnparsedObject == nil {
			jsonReferenceTableColumn_20270101, _ := datadog.Marshal(obj.ReferenceTableColumn_20270101)
			if string(jsonReferenceTableColumn_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.ReferenceTableColumn_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ReferenceTableColumn_20270101 = nil
		}
	} else {
		obj.ReferenceTableColumn_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ComputeTableColumn_20270101 = nil
		obj.PresetPeriodTableColumn_20270101 = nil
		obj.FormulaTableColumn_20270101 = nil
		obj.ReferenceTableColumn_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TableColumn) MarshalJSON() ([]byte, error) {
	if obj.ComputeTableColumn_20270101 != nil {
		return datadog.Marshal(&obj.ComputeTableColumn_20270101)
	}

	if obj.PresetPeriodTableColumn_20270101 != nil {
		return datadog.Marshal(&obj.PresetPeriodTableColumn_20270101)
	}

	if obj.FormulaTableColumn_20270101 != nil {
		return datadog.Marshal(&obj.FormulaTableColumn_20270101)
	}

	if obj.ReferenceTableColumn_20270101 != nil {
		return datadog.Marshal(&obj.ReferenceTableColumn_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TableColumn) GetActualInstance() interface{} {
	if obj.ComputeTableColumn_20270101 != nil {
		return obj.ComputeTableColumn_20270101
	}

	if obj.PresetPeriodTableColumn_20270101 != nil {
		return obj.PresetPeriodTableColumn_20270101
	}

	if obj.FormulaTableColumn_20270101 != nil {
		return obj.FormulaTableColumn_20270101
	}

	if obj.ReferenceTableColumn_20270101 != nil {
		return obj.ReferenceTableColumn_20270101
	}

	// all schemas are nil
	return nil
}
