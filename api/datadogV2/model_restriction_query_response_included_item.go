// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RestrictionQueryResponseIncludedItem - An object related to a restriction query.
type RestrictionQueryResponseIncludedItem struct {
	RestrictionQueryRole *RestrictionQueryRole

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// RestrictionQueryRoleAsRestrictionQueryResponseIncludedItem is a convenience function that returns RestrictionQueryRole wrapped in RestrictionQueryResponseIncludedItem.
func RestrictionQueryRoleAsRestrictionQueryResponseIncludedItem(v *RestrictionQueryRole) RestrictionQueryResponseIncludedItem {
	return RestrictionQueryResponseIncludedItem{RestrictionQueryRole: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *RestrictionQueryResponseIncludedItem) UnmarshalJSON(data []byte) error {
	var err error
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj RestrictionQueryResponseIncludedItem) MarshalJSON() ([]byte, error) {
	if obj.RestrictionQueryRole != nil {
		return datadog.Marshal(&obj.RestrictionQueryRole)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *RestrictionQueryResponseIncludedItem) GetActualInstance() interface{} {
	if obj.RestrictionQueryRole != nil {
		return obj.RestrictionQueryRole
	}

	// all schemas are nil
	return nil
}
