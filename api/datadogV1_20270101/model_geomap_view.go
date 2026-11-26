// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

// GeomapView - Test-only generated schema for the dashboard versioning proof of concept.
type GeomapView struct {
	GeomapView20270101OneOf           *datadogV1.GeomapView20270101OneOf
	GeomapView20270101OneOf4457650336 *datadogV1.GeomapView20270101OneOf4457650336

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GeomapView20270101OneOfAsGeomapView is a convenience function that returns datadogV1.GeomapView20270101OneOf wrapped in GeomapView.
func GeomapView20270101OneOfAsGeomapView(v *datadogV1.GeomapView20270101OneOf) GeomapView {
	return GeomapView{GeomapView20270101OneOf: v}
}

// GeomapView20270101OneOf4457650336AsGeomapView is a convenience function that returns datadogV1.GeomapView20270101OneOf4457650336 wrapped in GeomapView.
func GeomapView20270101OneOf4457650336AsGeomapView(v *datadogV1.GeomapView20270101OneOf4457650336) GeomapView {
	return GeomapView{GeomapView20270101OneOf4457650336: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GeomapView) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GeomapView20270101OneOf
	err = datadog.Unmarshal(data, &obj.GeomapView20270101OneOf)
	if err == nil {
		if obj.GeomapView20270101OneOf != nil && obj.GeomapView20270101OneOf.UnparsedObject == nil {
			jsonGeomapView20270101OneOf, _ := datadog.Marshal(obj.GeomapView20270101OneOf)
			if string(jsonGeomapView20270101OneOf) == "{}" { // empty struct
				obj.GeomapView20270101OneOf = nil
			} else {
				match++
			}
		} else {
			obj.GeomapView20270101OneOf = nil
		}
	} else {
		obj.GeomapView20270101OneOf = nil
	}

	// try to unmarshal data into GeomapView20270101OneOf4457650336
	err = datadog.Unmarshal(data, &obj.GeomapView20270101OneOf4457650336)
	if err == nil {
		if obj.GeomapView20270101OneOf4457650336 != nil && obj.GeomapView20270101OneOf4457650336.UnparsedObject == nil {
			jsonGeomapView20270101OneOf4457650336, _ := datadog.Marshal(obj.GeomapView20270101OneOf4457650336)
			if string(jsonGeomapView20270101OneOf4457650336) == "{}" { // empty struct
				obj.GeomapView20270101OneOf4457650336 = nil
			} else {
				match++
			}
		} else {
			obj.GeomapView20270101OneOf4457650336 = nil
		}
	} else {
		obj.GeomapView20270101OneOf4457650336 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GeomapView20270101OneOf = nil
		obj.GeomapView20270101OneOf4457650336 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GeomapView) MarshalJSON() ([]byte, error) {
	if obj.GeomapView20270101OneOf != nil {
		return datadog.Marshal(&obj.GeomapView20270101OneOf)
	}

	if obj.GeomapView20270101OneOf4457650336 != nil {
		return datadog.Marshal(&obj.GeomapView20270101OneOf4457650336)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GeomapView) GetActualInstance() interface{} {
	if obj.GeomapView20270101OneOf != nil {
		return obj.GeomapView20270101OneOf
	}

	if obj.GeomapView20270101OneOf4457650336 != nil {
		return obj.GeomapView20270101OneOf4457650336
	}

	// all schemas are nil
	return nil
}
