// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendPointsRequest - Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendPointsRequest struct {
	GeomapBackendPointsRequest20270101OneOf           *GeomapBackendPointsRequest20270101OneOf
	GeomapBackendPointsRequest20270101OneOf4457225472 *GeomapBackendPointsRequest20270101OneOf4457225472
	GeomapBackendPointsRequest20270101OneOf4457227232 *GeomapBackendPointsRequest20270101OneOf4457227232
	GeomapBackendPointsRequest20270101OneOf4457228992 *GeomapBackendPointsRequest20270101OneOf4457228992
	GeomapBackendPointsRequest20270101OneOf4457230272 *GeomapBackendPointsRequest20270101OneOf4457230272

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GeomapBackendPointsRequest20270101OneOfAsGeomapBackendPointsRequest is a convenience function that returns GeomapBackendPointsRequest20270101OneOf wrapped in GeomapBackendPointsRequest.
func GeomapBackendPointsRequest20270101OneOfAsGeomapBackendPointsRequest(v *GeomapBackendPointsRequest20270101OneOf) GeomapBackendPointsRequest {
	return GeomapBackendPointsRequest{GeomapBackendPointsRequest20270101OneOf: v}
}

// GeomapBackendPointsRequest20270101OneOf4457225472AsGeomapBackendPointsRequest is a convenience function that returns GeomapBackendPointsRequest20270101OneOf4457225472 wrapped in GeomapBackendPointsRequest.
func GeomapBackendPointsRequest20270101OneOf4457225472AsGeomapBackendPointsRequest(v *GeomapBackendPointsRequest20270101OneOf4457225472) GeomapBackendPointsRequest {
	return GeomapBackendPointsRequest{GeomapBackendPointsRequest20270101OneOf4457225472: v}
}

// GeomapBackendPointsRequest20270101OneOf4457227232AsGeomapBackendPointsRequest is a convenience function that returns GeomapBackendPointsRequest20270101OneOf4457227232 wrapped in GeomapBackendPointsRequest.
func GeomapBackendPointsRequest20270101OneOf4457227232AsGeomapBackendPointsRequest(v *GeomapBackendPointsRequest20270101OneOf4457227232) GeomapBackendPointsRequest {
	return GeomapBackendPointsRequest{GeomapBackendPointsRequest20270101OneOf4457227232: v}
}

// GeomapBackendPointsRequest20270101OneOf4457228992AsGeomapBackendPointsRequest is a convenience function that returns GeomapBackendPointsRequest20270101OneOf4457228992 wrapped in GeomapBackendPointsRequest.
func GeomapBackendPointsRequest20270101OneOf4457228992AsGeomapBackendPointsRequest(v *GeomapBackendPointsRequest20270101OneOf4457228992) GeomapBackendPointsRequest {
	return GeomapBackendPointsRequest{GeomapBackendPointsRequest20270101OneOf4457228992: v}
}

// GeomapBackendPointsRequest20270101OneOf4457230272AsGeomapBackendPointsRequest is a convenience function that returns GeomapBackendPointsRequest20270101OneOf4457230272 wrapped in GeomapBackendPointsRequest.
func GeomapBackendPointsRequest20270101OneOf4457230272AsGeomapBackendPointsRequest(v *GeomapBackendPointsRequest20270101OneOf4457230272) GeomapBackendPointsRequest {
	return GeomapBackendPointsRequest{GeomapBackendPointsRequest20270101OneOf4457230272: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GeomapBackendPointsRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GeomapBackendPointsRequest20270101OneOf
	err = datadog.Unmarshal(data, &obj.GeomapBackendPointsRequest20270101OneOf)
	if err == nil {
		if obj.GeomapBackendPointsRequest20270101OneOf != nil && obj.GeomapBackendPointsRequest20270101OneOf.UnparsedObject == nil {
			jsonGeomapBackendPointsRequest20270101OneOf, _ := datadog.Marshal(obj.GeomapBackendPointsRequest20270101OneOf)
			if string(jsonGeomapBackendPointsRequest20270101OneOf) == "{}" { // empty struct
				obj.GeomapBackendPointsRequest20270101OneOf = nil
			} else {
				match++
			}
		} else {
			obj.GeomapBackendPointsRequest20270101OneOf = nil
		}
	} else {
		obj.GeomapBackendPointsRequest20270101OneOf = nil
	}

	// try to unmarshal data into GeomapBackendPointsRequest20270101OneOf4457225472
	err = datadog.Unmarshal(data, &obj.GeomapBackendPointsRequest20270101OneOf4457225472)
	if err == nil {
		if obj.GeomapBackendPointsRequest20270101OneOf4457225472 != nil && obj.GeomapBackendPointsRequest20270101OneOf4457225472.UnparsedObject == nil {
			jsonGeomapBackendPointsRequest20270101OneOf4457225472, _ := datadog.Marshal(obj.GeomapBackendPointsRequest20270101OneOf4457225472)
			if string(jsonGeomapBackendPointsRequest20270101OneOf4457225472) == "{}" { // empty struct
				obj.GeomapBackendPointsRequest20270101OneOf4457225472 = nil
			} else {
				match++
			}
		} else {
			obj.GeomapBackendPointsRequest20270101OneOf4457225472 = nil
		}
	} else {
		obj.GeomapBackendPointsRequest20270101OneOf4457225472 = nil
	}

	// try to unmarshal data into GeomapBackendPointsRequest20270101OneOf4457227232
	err = datadog.Unmarshal(data, &obj.GeomapBackendPointsRequest20270101OneOf4457227232)
	if err == nil {
		if obj.GeomapBackendPointsRequest20270101OneOf4457227232 != nil && obj.GeomapBackendPointsRequest20270101OneOf4457227232.UnparsedObject == nil {
			jsonGeomapBackendPointsRequest20270101OneOf4457227232, _ := datadog.Marshal(obj.GeomapBackendPointsRequest20270101OneOf4457227232)
			if string(jsonGeomapBackendPointsRequest20270101OneOf4457227232) == "{}" { // empty struct
				obj.GeomapBackendPointsRequest20270101OneOf4457227232 = nil
			} else {
				match++
			}
		} else {
			obj.GeomapBackendPointsRequest20270101OneOf4457227232 = nil
		}
	} else {
		obj.GeomapBackendPointsRequest20270101OneOf4457227232 = nil
	}

	// try to unmarshal data into GeomapBackendPointsRequest20270101OneOf4457228992
	err = datadog.Unmarshal(data, &obj.GeomapBackendPointsRequest20270101OneOf4457228992)
	if err == nil {
		if obj.GeomapBackendPointsRequest20270101OneOf4457228992 != nil && obj.GeomapBackendPointsRequest20270101OneOf4457228992.UnparsedObject == nil {
			jsonGeomapBackendPointsRequest20270101OneOf4457228992, _ := datadog.Marshal(obj.GeomapBackendPointsRequest20270101OneOf4457228992)
			if string(jsonGeomapBackendPointsRequest20270101OneOf4457228992) == "{}" { // empty struct
				obj.GeomapBackendPointsRequest20270101OneOf4457228992 = nil
			} else {
				match++
			}
		} else {
			obj.GeomapBackendPointsRequest20270101OneOf4457228992 = nil
		}
	} else {
		obj.GeomapBackendPointsRequest20270101OneOf4457228992 = nil
	}

	// try to unmarshal data into GeomapBackendPointsRequest20270101OneOf4457230272
	err = datadog.Unmarshal(data, &obj.GeomapBackendPointsRequest20270101OneOf4457230272)
	if err == nil {
		if obj.GeomapBackendPointsRequest20270101OneOf4457230272 != nil && obj.GeomapBackendPointsRequest20270101OneOf4457230272.UnparsedObject == nil {
			jsonGeomapBackendPointsRequest20270101OneOf4457230272, _ := datadog.Marshal(obj.GeomapBackendPointsRequest20270101OneOf4457230272)
			if string(jsonGeomapBackendPointsRequest20270101OneOf4457230272) == "{}" { // empty struct
				obj.GeomapBackendPointsRequest20270101OneOf4457230272 = nil
			} else {
				match++
			}
		} else {
			obj.GeomapBackendPointsRequest20270101OneOf4457230272 = nil
		}
	} else {
		obj.GeomapBackendPointsRequest20270101OneOf4457230272 = nil
	}

	if match > 1 {
		// more than one variant matched, so tell them apart by the items of the arrays they hold
		if obj.GeomapBackendPointsRequest20270101OneOf != nil {
			allColumnsUnparsed := len(obj.GeomapBackendPointsRequest20270101OneOf.Columns) > 0
			for _, item := range obj.GeomapBackendPointsRequest20270101OneOf.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.GeomapBackendPointsRequest20270101OneOf = nil
				match--
			}
		}
		if obj.GeomapBackendPointsRequest20270101OneOf4457225472 != nil {
			allColumnsUnparsed := len(obj.GeomapBackendPointsRequest20270101OneOf4457225472.Columns) > 0
			for _, item := range obj.GeomapBackendPointsRequest20270101OneOf4457225472.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.GeomapBackendPointsRequest20270101OneOf4457225472 = nil
				match--
			}
		}
		if obj.GeomapBackendPointsRequest20270101OneOf4457227232 != nil {
			allColumnsUnparsed := len(obj.GeomapBackendPointsRequest20270101OneOf4457227232.Columns) > 0
			for _, item := range obj.GeomapBackendPointsRequest20270101OneOf4457227232.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.GeomapBackendPointsRequest20270101OneOf4457227232 = nil
				match--
			}
		}
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GeomapBackendPointsRequest20270101OneOf = nil
		obj.GeomapBackendPointsRequest20270101OneOf4457225472 = nil
		obj.GeomapBackendPointsRequest20270101OneOf4457227232 = nil
		obj.GeomapBackendPointsRequest20270101OneOf4457228992 = nil
		obj.GeomapBackendPointsRequest20270101OneOf4457230272 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GeomapBackendPointsRequest) MarshalJSON() ([]byte, error) {
	if obj.GeomapBackendPointsRequest20270101OneOf != nil {
		return datadog.Marshal(&obj.GeomapBackendPointsRequest20270101OneOf)
	}

	if obj.GeomapBackendPointsRequest20270101OneOf4457225472 != nil {
		return datadog.Marshal(&obj.GeomapBackendPointsRequest20270101OneOf4457225472)
	}

	if obj.GeomapBackendPointsRequest20270101OneOf4457227232 != nil {
		return datadog.Marshal(&obj.GeomapBackendPointsRequest20270101OneOf4457227232)
	}

	if obj.GeomapBackendPointsRequest20270101OneOf4457228992 != nil {
		return datadog.Marshal(&obj.GeomapBackendPointsRequest20270101OneOf4457228992)
	}

	if obj.GeomapBackendPointsRequest20270101OneOf4457230272 != nil {
		return datadog.Marshal(&obj.GeomapBackendPointsRequest20270101OneOf4457230272)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GeomapBackendPointsRequest) GetActualInstance() interface{} {
	if obj.GeomapBackendPointsRequest20270101OneOf != nil {
		return obj.GeomapBackendPointsRequest20270101OneOf
	}

	if obj.GeomapBackendPointsRequest20270101OneOf4457225472 != nil {
		return obj.GeomapBackendPointsRequest20270101OneOf4457225472
	}

	if obj.GeomapBackendPointsRequest20270101OneOf4457227232 != nil {
		return obj.GeomapBackendPointsRequest20270101OneOf4457227232
	}

	if obj.GeomapBackendPointsRequest20270101OneOf4457228992 != nil {
		return obj.GeomapBackendPointsRequest20270101OneOf4457228992
	}

	if obj.GeomapBackendPointsRequest20270101OneOf4457230272 != nil {
		return obj.GeomapBackendPointsRequest20270101OneOf4457230272
	}

	// all schemas are nil
	return nil
}
