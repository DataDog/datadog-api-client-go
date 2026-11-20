// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableColumnComparison - Test-only generated schema for the dashboard versioning proof of concept.
type TableColumnComparison struct {
	TableColumnComparisonWithSelf_20270101                       *TableColumnComparisonWithSelf
	TableColumnComparisonWithOtherColumnDiffFrom_20270101        *TableColumnComparisonWithOtherColumnDiffFrom
	TableColumnComparisonWithOtherColumnPercentOf_20270101       *TableColumnComparisonWithOtherColumnPercentOf
	TableColumnComparisonWithOtherColumnPercentDiffFrom_20270101 *TableColumnComparisonWithOtherColumnPercentDiffFrom

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TableColumnComparisonWithSelf_20270101AsTableColumnComparison is a convenience function that returns TableColumnComparisonWithSelf wrapped in TableColumnComparison.
func TableColumnComparisonWithSelf_20270101AsTableColumnComparison(v *TableColumnComparisonWithSelf) TableColumnComparison {
	return TableColumnComparison{TableColumnComparisonWithSelf_20270101: v}
}

// TableColumnComparisonWithOtherColumnDiffFrom_20270101AsTableColumnComparison is a convenience function that returns TableColumnComparisonWithOtherColumnDiffFrom wrapped in TableColumnComparison.
func TableColumnComparisonWithOtherColumnDiffFrom_20270101AsTableColumnComparison(v *TableColumnComparisonWithOtherColumnDiffFrom) TableColumnComparison {
	return TableColumnComparison{TableColumnComparisonWithOtherColumnDiffFrom_20270101: v}
}

// TableColumnComparisonWithOtherColumnPercentOf_20270101AsTableColumnComparison is a convenience function that returns TableColumnComparisonWithOtherColumnPercentOf wrapped in TableColumnComparison.
func TableColumnComparisonWithOtherColumnPercentOf_20270101AsTableColumnComparison(v *TableColumnComparisonWithOtherColumnPercentOf) TableColumnComparison {
	return TableColumnComparison{TableColumnComparisonWithOtherColumnPercentOf_20270101: v}
}

// TableColumnComparisonWithOtherColumnPercentDiffFrom_20270101AsTableColumnComparison is a convenience function that returns TableColumnComparisonWithOtherColumnPercentDiffFrom wrapped in TableColumnComparison.
func TableColumnComparisonWithOtherColumnPercentDiffFrom_20270101AsTableColumnComparison(v *TableColumnComparisonWithOtherColumnPercentDiffFrom) TableColumnComparison {
	return TableColumnComparison{TableColumnComparisonWithOtherColumnPercentDiffFrom_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TableColumnComparison) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TableColumnComparisonWithSelf_20270101
	err = datadog.Unmarshal(data, &obj.TableColumnComparisonWithSelf_20270101)
	if err == nil {
		if obj.TableColumnComparisonWithSelf_20270101 != nil && obj.TableColumnComparisonWithSelf_20270101.UnparsedObject == nil {
			jsonTableColumnComparisonWithSelf_20270101, _ := datadog.Marshal(obj.TableColumnComparisonWithSelf_20270101)
			if string(jsonTableColumnComparisonWithSelf_20270101) == "{}" { // empty struct
				obj.TableColumnComparisonWithSelf_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TableColumnComparisonWithSelf_20270101 = nil
		}
	} else {
		obj.TableColumnComparisonWithSelf_20270101 = nil
	}

	// try to unmarshal data into TableColumnComparisonWithOtherColumnDiffFrom_20270101
	err = datadog.Unmarshal(data, &obj.TableColumnComparisonWithOtherColumnDiffFrom_20270101)
	if err == nil {
		if obj.TableColumnComparisonWithOtherColumnDiffFrom_20270101 != nil && obj.TableColumnComparisonWithOtherColumnDiffFrom_20270101.UnparsedObject == nil {
			jsonTableColumnComparisonWithOtherColumnDiffFrom_20270101, _ := datadog.Marshal(obj.TableColumnComparisonWithOtherColumnDiffFrom_20270101)
			if string(jsonTableColumnComparisonWithOtherColumnDiffFrom_20270101) == "{}" { // empty struct
				obj.TableColumnComparisonWithOtherColumnDiffFrom_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TableColumnComparisonWithOtherColumnDiffFrom_20270101 = nil
		}
	} else {
		obj.TableColumnComparisonWithOtherColumnDiffFrom_20270101 = nil
	}

	// try to unmarshal data into TableColumnComparisonWithOtherColumnPercentOf_20270101
	err = datadog.Unmarshal(data, &obj.TableColumnComparisonWithOtherColumnPercentOf_20270101)
	if err == nil {
		if obj.TableColumnComparisonWithOtherColumnPercentOf_20270101 != nil && obj.TableColumnComparisonWithOtherColumnPercentOf_20270101.UnparsedObject == nil {
			jsonTableColumnComparisonWithOtherColumnPercentOf_20270101, _ := datadog.Marshal(obj.TableColumnComparisonWithOtherColumnPercentOf_20270101)
			if string(jsonTableColumnComparisonWithOtherColumnPercentOf_20270101) == "{}" { // empty struct
				obj.TableColumnComparisonWithOtherColumnPercentOf_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TableColumnComparisonWithOtherColumnPercentOf_20270101 = nil
		}
	} else {
		obj.TableColumnComparisonWithOtherColumnPercentOf_20270101 = nil
	}

	// try to unmarshal data into TableColumnComparisonWithOtherColumnPercentDiffFrom_20270101
	err = datadog.Unmarshal(data, &obj.TableColumnComparisonWithOtherColumnPercentDiffFrom_20270101)
	if err == nil {
		if obj.TableColumnComparisonWithOtherColumnPercentDiffFrom_20270101 != nil && obj.TableColumnComparisonWithOtherColumnPercentDiffFrom_20270101.UnparsedObject == nil {
			jsonTableColumnComparisonWithOtherColumnPercentDiffFrom_20270101, _ := datadog.Marshal(obj.TableColumnComparisonWithOtherColumnPercentDiffFrom_20270101)
			if string(jsonTableColumnComparisonWithOtherColumnPercentDiffFrom_20270101) == "{}" { // empty struct
				obj.TableColumnComparisonWithOtherColumnPercentDiffFrom_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TableColumnComparisonWithOtherColumnPercentDiffFrom_20270101 = nil
		}
	} else {
		obj.TableColumnComparisonWithOtherColumnPercentDiffFrom_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TableColumnComparisonWithSelf_20270101 = nil
		obj.TableColumnComparisonWithOtherColumnDiffFrom_20270101 = nil
		obj.TableColumnComparisonWithOtherColumnPercentOf_20270101 = nil
		obj.TableColumnComparisonWithOtherColumnPercentDiffFrom_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TableColumnComparison) MarshalJSON() ([]byte, error) {
	if obj.TableColumnComparisonWithSelf_20270101 != nil {
		return datadog.Marshal(&obj.TableColumnComparisonWithSelf_20270101)
	}

	if obj.TableColumnComparisonWithOtherColumnDiffFrom_20270101 != nil {
		return datadog.Marshal(&obj.TableColumnComparisonWithOtherColumnDiffFrom_20270101)
	}

	if obj.TableColumnComparisonWithOtherColumnPercentOf_20270101 != nil {
		return datadog.Marshal(&obj.TableColumnComparisonWithOtherColumnPercentOf_20270101)
	}

	if obj.TableColumnComparisonWithOtherColumnPercentDiffFrom_20270101 != nil {
		return datadog.Marshal(&obj.TableColumnComparisonWithOtherColumnPercentDiffFrom_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TableColumnComparison) GetActualInstance() interface{} {
	if obj.TableColumnComparisonWithSelf_20270101 != nil {
		return obj.TableColumnComparisonWithSelf_20270101
	}

	if obj.TableColumnComparisonWithOtherColumnDiffFrom_20270101 != nil {
		return obj.TableColumnComparisonWithOtherColumnDiffFrom_20270101
	}

	if obj.TableColumnComparisonWithOtherColumnPercentOf_20270101 != nil {
		return obj.TableColumnComparisonWithOtherColumnPercentOf_20270101
	}

	if obj.TableColumnComparisonWithOtherColumnPercentDiffFrom_20270101 != nil {
		return obj.TableColumnComparisonWithOtherColumnPercentDiffFrom_20270101
	}

	// all schemas are nil
	return nil
}
