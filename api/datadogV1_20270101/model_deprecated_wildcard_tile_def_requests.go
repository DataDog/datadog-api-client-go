// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeprecatedWildcardTileDefRequests - Test-only generated schema for the dashboard versioning proof of concept.
type DeprecatedWildcardTileDefRequests struct {
	DDSQLTableRequest_20270101      *DDSQLTableRequest
	DDSQLTimeseriesRequest_20270101 *DDSQLTimeseriesRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DDSQLTableRequest_20270101AsDeprecatedWildcardTileDefRequests is a convenience function that returns DDSQLTableRequest wrapped in DeprecatedWildcardTileDefRequests.
func DDSQLTableRequest_20270101AsDeprecatedWildcardTileDefRequests(v *DDSQLTableRequest) DeprecatedWildcardTileDefRequests {
	return DeprecatedWildcardTileDefRequests{DDSQLTableRequest_20270101: v}
}

// DDSQLTimeseriesRequest_20270101AsDeprecatedWildcardTileDefRequests is a convenience function that returns DDSQLTimeseriesRequest wrapped in DeprecatedWildcardTileDefRequests.
func DDSQLTimeseriesRequest_20270101AsDeprecatedWildcardTileDefRequests(v *DDSQLTimeseriesRequest) DeprecatedWildcardTileDefRequests {
	return DeprecatedWildcardTileDefRequests{DDSQLTimeseriesRequest_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DeprecatedWildcardTileDefRequests) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DDSQLTableRequest_20270101
	err = datadog.Unmarshal(data, &obj.DDSQLTableRequest_20270101)
	if err == nil {
		if obj.DDSQLTableRequest_20270101 != nil && obj.DDSQLTableRequest_20270101.UnparsedObject == nil {
			jsonDDSQLTableRequest_20270101, _ := datadog.Marshal(obj.DDSQLTableRequest_20270101)
			if string(jsonDDSQLTableRequest_20270101) == "{}" { // empty struct
				obj.DDSQLTableRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.DDSQLTableRequest_20270101 = nil
		}
	} else {
		obj.DDSQLTableRequest_20270101 = nil
	}

	// try to unmarshal data into DDSQLTimeseriesRequest_20270101
	err = datadog.Unmarshal(data, &obj.DDSQLTimeseriesRequest_20270101)
	if err == nil {
		if obj.DDSQLTimeseriesRequest_20270101 != nil && obj.DDSQLTimeseriesRequest_20270101.UnparsedObject == nil {
			jsonDDSQLTimeseriesRequest_20270101, _ := datadog.Marshal(obj.DDSQLTimeseriesRequest_20270101)
			if string(jsonDDSQLTimeseriesRequest_20270101) == "{}" { // empty struct
				obj.DDSQLTimeseriesRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.DDSQLTimeseriesRequest_20270101 = nil
		}
	} else {
		obj.DDSQLTimeseriesRequest_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DDSQLTableRequest_20270101 = nil
		obj.DDSQLTimeseriesRequest_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DeprecatedWildcardTileDefRequests) MarshalJSON() ([]byte, error) {
	if obj.DDSQLTableRequest_20270101 != nil {
		return datadog.Marshal(&obj.DDSQLTableRequest_20270101)
	}

	if obj.DDSQLTimeseriesRequest_20270101 != nil {
		return datadog.Marshal(&obj.DDSQLTimeseriesRequest_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DeprecatedWildcardTileDefRequests) GetActualInstance() interface{} {
	if obj.DDSQLTableRequest_20270101 != nil {
		return obj.DDSQLTableRequest_20270101
	}

	if obj.DDSQLTimeseriesRequest_20270101 != nil {
		return obj.DDSQLTimeseriesRequest_20270101
	}

	// all schemas are nil
	return nil
}
