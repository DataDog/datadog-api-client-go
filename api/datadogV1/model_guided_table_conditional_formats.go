// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableConditionalFormats - Conditional formatting rules for a guided table column. Either an array of threshold-based rules or a single range-based rule.
type GuidedTableConditionalFormats struct {
	GuidedTableThresholdConditionalFormats *[]GuidedTableThresholdFormattingRule
	GuidedTableRangeConditionalFormats     *[]GuidedTableRangeFormattingRule

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GuidedTableThresholdConditionalFormatsAsGuidedTableConditionalFormats is a convenience function that returns []GuidedTableThresholdFormattingRule wrapped in GuidedTableConditionalFormats.
func GuidedTableThresholdConditionalFormatsAsGuidedTableConditionalFormats(v *[]GuidedTableThresholdFormattingRule) GuidedTableConditionalFormats {
	return GuidedTableConditionalFormats{GuidedTableThresholdConditionalFormats: v}
}

// GuidedTableRangeConditionalFormatsAsGuidedTableConditionalFormats is a convenience function that returns []GuidedTableRangeFormattingRule wrapped in GuidedTableConditionalFormats.
func GuidedTableRangeConditionalFormatsAsGuidedTableConditionalFormats(v *[]GuidedTableRangeFormattingRule) GuidedTableConditionalFormats {
	return GuidedTableConditionalFormats{GuidedTableRangeConditionalFormats: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GuidedTableConditionalFormats) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GuidedTableThresholdConditionalFormats
	err = datadog.Unmarshal(data, &obj.GuidedTableThresholdConditionalFormats)
	if err == nil {
		if obj.GuidedTableThresholdConditionalFormats != nil {
			jsonGuidedTableThresholdConditionalFormats, _ := datadog.Marshal(obj.GuidedTableThresholdConditionalFormats)
			if string(jsonGuidedTableThresholdConditionalFormats) == "{}" && string(data) != "{}" { // empty struct
				obj.GuidedTableThresholdConditionalFormats = nil
			} else {
				match++
			}
		} else {
			obj.GuidedTableThresholdConditionalFormats = nil
		}
	} else {
		obj.GuidedTableThresholdConditionalFormats = nil
	}

	// try to unmarshal data into GuidedTableRangeConditionalFormats
	err = datadog.Unmarshal(data, &obj.GuidedTableRangeConditionalFormats)
	if err == nil {
		if obj.GuidedTableRangeConditionalFormats != nil {
			jsonGuidedTableRangeConditionalFormats, _ := datadog.Marshal(obj.GuidedTableRangeConditionalFormats)
			if string(jsonGuidedTableRangeConditionalFormats) == "{}" && string(data) != "{}" { // empty struct
				obj.GuidedTableRangeConditionalFormats = nil
			} else {
				match++
			}
		} else {
			obj.GuidedTableRangeConditionalFormats = nil
		}
	} else {
		obj.GuidedTableRangeConditionalFormats = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GuidedTableThresholdConditionalFormats = nil
		obj.GuidedTableRangeConditionalFormats = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GuidedTableConditionalFormats) MarshalJSON() ([]byte, error) {
	if obj.GuidedTableThresholdConditionalFormats != nil {
		return datadog.Marshal(&obj.GuidedTableThresholdConditionalFormats)
	}

	if obj.GuidedTableRangeConditionalFormats != nil {
		return datadog.Marshal(&obj.GuidedTableRangeConditionalFormats)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GuidedTableConditionalFormats) GetActualInstance() interface{} {
	if obj.GuidedTableThresholdConditionalFormats != nil {
		return obj.GuidedTableThresholdConditionalFormats
	}

	if obj.GuidedTableRangeConditionalFormats != nil {
		return obj.GuidedTableRangeConditionalFormats
	}

	// all schemas are nil
	return nil
}
