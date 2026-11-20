// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartBackendDefinitionStyle20270101Display - Test-only generated schema for the dashboard versioning proof of concept.
type BarChartBackendDefinitionStyle20270101Display struct {
	BarChartBackendDefinitionStyle20270101DisplayOneOf           *BarChartBackendDefinitionStyle20270101DisplayOneOf
	BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856 *BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// BarChartBackendDefinitionStyle20270101DisplayOneOfAsBarChartBackendDefinitionStyle20270101Display is a convenience function that returns BarChartBackendDefinitionStyle20270101DisplayOneOf wrapped in BarChartBackendDefinitionStyle20270101Display.
func BarChartBackendDefinitionStyle20270101DisplayOneOfAsBarChartBackendDefinitionStyle20270101Display(v *BarChartBackendDefinitionStyle20270101DisplayOneOf) BarChartBackendDefinitionStyle20270101Display {
	return BarChartBackendDefinitionStyle20270101Display{BarChartBackendDefinitionStyle20270101DisplayOneOf: v}
}

// BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856AsBarChartBackendDefinitionStyle20270101Display is a convenience function that returns BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856 wrapped in BarChartBackendDefinitionStyle20270101Display.
func BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856AsBarChartBackendDefinitionStyle20270101Display(v *BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856) BarChartBackendDefinitionStyle20270101Display {
	return BarChartBackendDefinitionStyle20270101Display{BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *BarChartBackendDefinitionStyle20270101Display) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BarChartBackendDefinitionStyle20270101DisplayOneOf
	err = datadog.Unmarshal(data, &obj.BarChartBackendDefinitionStyle20270101DisplayOneOf)
	if err == nil {
		if obj.BarChartBackendDefinitionStyle20270101DisplayOneOf != nil && obj.BarChartBackendDefinitionStyle20270101DisplayOneOf.UnparsedObject == nil {
			jsonBarChartBackendDefinitionStyle20270101DisplayOneOf, _ := datadog.Marshal(obj.BarChartBackendDefinitionStyle20270101DisplayOneOf)
			if string(jsonBarChartBackendDefinitionStyle20270101DisplayOneOf) == "{}" { // empty struct
				obj.BarChartBackendDefinitionStyle20270101DisplayOneOf = nil
			} else {
				match++
			}
		} else {
			obj.BarChartBackendDefinitionStyle20270101DisplayOneOf = nil
		}
	} else {
		obj.BarChartBackendDefinitionStyle20270101DisplayOneOf = nil
	}

	// try to unmarshal data into BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856
	err = datadog.Unmarshal(data, &obj.BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856)
	if err == nil {
		if obj.BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856 != nil && obj.BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856.UnparsedObject == nil {
			jsonBarChartBackendDefinitionStyle20270101DisplayOneOf4453561856, _ := datadog.Marshal(obj.BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856)
			if string(jsonBarChartBackendDefinitionStyle20270101DisplayOneOf4453561856) == "{}" { // empty struct
				obj.BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856 = nil
			} else {
				match++
			}
		} else {
			obj.BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856 = nil
		}
	} else {
		obj.BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.BarChartBackendDefinitionStyle20270101DisplayOneOf = nil
		obj.BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj BarChartBackendDefinitionStyle20270101Display) MarshalJSON() ([]byte, error) {
	if obj.BarChartBackendDefinitionStyle20270101DisplayOneOf != nil {
		return datadog.Marshal(&obj.BarChartBackendDefinitionStyle20270101DisplayOneOf)
	}

	if obj.BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856 != nil {
		return datadog.Marshal(&obj.BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *BarChartBackendDefinitionStyle20270101Display) GetActualInstance() interface{} {
	if obj.BarChartBackendDefinitionStyle20270101DisplayOneOf != nil {
		return obj.BarChartBackendDefinitionStyle20270101DisplayOneOf
	}

	if obj.BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856 != nil {
		return obj.BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856
	}

	// all schemas are nil
	return nil
}
