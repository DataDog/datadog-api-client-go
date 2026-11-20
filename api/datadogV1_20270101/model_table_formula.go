// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableFormula - Test-only generated schema for the dashboard versioning proof of concept.
type TableFormula struct {
	TableFormula20270101OneOf           *TableFormula20270101OneOf
	TableFormula20270101OneOf4462669760 *TableFormula20270101OneOf4462669760

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TableFormula20270101OneOfAsTableFormula is a convenience function that returns TableFormula20270101OneOf wrapped in TableFormula.
func TableFormula20270101OneOfAsTableFormula(v *TableFormula20270101OneOf) TableFormula {
	return TableFormula{TableFormula20270101OneOf: v}
}

// TableFormula20270101OneOf4462669760AsTableFormula is a convenience function that returns TableFormula20270101OneOf4462669760 wrapped in TableFormula.
func TableFormula20270101OneOf4462669760AsTableFormula(v *TableFormula20270101OneOf4462669760) TableFormula {
	return TableFormula{TableFormula20270101OneOf4462669760: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TableFormula) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TableFormula20270101OneOf
	err = datadog.Unmarshal(data, &obj.TableFormula20270101OneOf)
	if err == nil {
		if obj.TableFormula20270101OneOf != nil && obj.TableFormula20270101OneOf.UnparsedObject == nil {
			jsonTableFormula20270101OneOf, _ := datadog.Marshal(obj.TableFormula20270101OneOf)
			if string(jsonTableFormula20270101OneOf) == "{}" { // empty struct
				obj.TableFormula20270101OneOf = nil
			} else {
				match++
			}
		} else {
			obj.TableFormula20270101OneOf = nil
		}
	} else {
		obj.TableFormula20270101OneOf = nil
	}

	// try to unmarshal data into TableFormula20270101OneOf4462669760
	err = datadog.Unmarshal(data, &obj.TableFormula20270101OneOf4462669760)
	if err == nil {
		if obj.TableFormula20270101OneOf4462669760 != nil && obj.TableFormula20270101OneOf4462669760.UnparsedObject == nil {
			jsonTableFormula20270101OneOf4462669760, _ := datadog.Marshal(obj.TableFormula20270101OneOf4462669760)
			if string(jsonTableFormula20270101OneOf4462669760) == "{}" { // empty struct
				obj.TableFormula20270101OneOf4462669760 = nil
			} else {
				match++
			}
		} else {
			obj.TableFormula20270101OneOf4462669760 = nil
		}
	} else {
		obj.TableFormula20270101OneOf4462669760 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TableFormula20270101OneOf = nil
		obj.TableFormula20270101OneOf4462669760 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TableFormula) MarshalJSON() ([]byte, error) {
	if obj.TableFormula20270101OneOf != nil {
		return datadog.Marshal(&obj.TableFormula20270101OneOf)
	}

	if obj.TableFormula20270101OneOf4462669760 != nil {
		return datadog.Marshal(&obj.TableFormula20270101OneOf4462669760)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TableFormula) GetActualInstance() interface{} {
	if obj.TableFormula20270101OneOf != nil {
		return obj.TableFormula20270101OneOf
	}

	if obj.TableFormula20270101OneOf4462669760 != nil {
		return obj.TableFormula20270101OneOf4462669760
	}

	// all schemas are nil
	return nil
}
