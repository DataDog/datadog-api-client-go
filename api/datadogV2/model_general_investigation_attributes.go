// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeneralInvestigationAttributes - Attributes for a general investigation, not tied to a specific monitor alert.
type GeneralInvestigationAttributes struct {
	GeneralInvestigationAttributesWithoutTimeBounds *GeneralInvestigationAttributesWithoutTimeBounds
	GeneralInvestigationAttributesWithTimeBounds    *GeneralInvestigationAttributesWithTimeBounds

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GeneralInvestigationAttributesWithoutTimeBoundsAsGeneralInvestigationAttributes is a convenience function that returns GeneralInvestigationAttributesWithoutTimeBounds wrapped in GeneralInvestigationAttributes.
func GeneralInvestigationAttributesWithoutTimeBoundsAsGeneralInvestigationAttributes(v *GeneralInvestigationAttributesWithoutTimeBounds) GeneralInvestigationAttributes {
	return GeneralInvestigationAttributes{GeneralInvestigationAttributesWithoutTimeBounds: v}
}

// GeneralInvestigationAttributesWithTimeBoundsAsGeneralInvestigationAttributes is a convenience function that returns GeneralInvestigationAttributesWithTimeBounds wrapped in GeneralInvestigationAttributes.
func GeneralInvestigationAttributesWithTimeBoundsAsGeneralInvestigationAttributes(v *GeneralInvestigationAttributesWithTimeBounds) GeneralInvestigationAttributes {
	return GeneralInvestigationAttributes{GeneralInvestigationAttributesWithTimeBounds: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GeneralInvestigationAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GeneralInvestigationAttributesWithoutTimeBounds
	err = datadog.Unmarshal(data, &obj.GeneralInvestigationAttributesWithoutTimeBounds)
	if err == nil {
		if obj.GeneralInvestigationAttributesWithoutTimeBounds != nil && obj.GeneralInvestigationAttributesWithoutTimeBounds.UnparsedObject == nil {
			jsonGeneralInvestigationAttributesWithoutTimeBounds, _ := datadog.Marshal(obj.GeneralInvestigationAttributesWithoutTimeBounds)
			if string(jsonGeneralInvestigationAttributesWithoutTimeBounds) == "{}" { // empty struct
				obj.GeneralInvestigationAttributesWithoutTimeBounds = nil
			} else {
				match++
			}
		} else {
			obj.GeneralInvestigationAttributesWithoutTimeBounds = nil
		}
	} else {
		obj.GeneralInvestigationAttributesWithoutTimeBounds = nil
	}

	// try to unmarshal data into GeneralInvestigationAttributesWithTimeBounds
	err = datadog.Unmarshal(data, &obj.GeneralInvestigationAttributesWithTimeBounds)
	if err == nil {
		if obj.GeneralInvestigationAttributesWithTimeBounds != nil && obj.GeneralInvestigationAttributesWithTimeBounds.UnparsedObject == nil {
			jsonGeneralInvestigationAttributesWithTimeBounds, _ := datadog.Marshal(obj.GeneralInvestigationAttributesWithTimeBounds)
			if string(jsonGeneralInvestigationAttributesWithTimeBounds) == "{}" { // empty struct
				obj.GeneralInvestigationAttributesWithTimeBounds = nil
			} else {
				match++
			}
		} else {
			obj.GeneralInvestigationAttributesWithTimeBounds = nil
		}
	} else {
		obj.GeneralInvestigationAttributesWithTimeBounds = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GeneralInvestigationAttributesWithoutTimeBounds = nil
		obj.GeneralInvestigationAttributesWithTimeBounds = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GeneralInvestigationAttributes) MarshalJSON() ([]byte, error) {
	if obj.GeneralInvestigationAttributesWithoutTimeBounds != nil {
		return datadog.Marshal(&obj.GeneralInvestigationAttributesWithoutTimeBounds)
	}

	if obj.GeneralInvestigationAttributesWithTimeBounds != nil {
		return datadog.Marshal(&obj.GeneralInvestigationAttributesWithTimeBounds)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GeneralInvestigationAttributes) GetActualInstance() interface{} {
	if obj.GeneralInvestigationAttributesWithoutTimeBounds != nil {
		return obj.GeneralInvestigationAttributesWithoutTimeBounds
	}

	if obj.GeneralInvestigationAttributesWithTimeBounds != nil {
		return obj.GeneralInvestigationAttributesWithTimeBounds
	}

	// all schemas are nil
	return nil
}
