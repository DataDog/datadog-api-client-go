// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TextFormattingRule20270101Replace - Test-only generated schema for the dashboard versioning proof of concept.
type TextFormattingRule20270101Replace struct {
	TextFormattingRule20270101ReplaceOneOf           *TextFormattingRule20270101ReplaceOneOf
	TextFormattingRule20270101ReplaceOneOf4462863328 *TextFormattingRule20270101ReplaceOneOf4462863328

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TextFormattingRule20270101ReplaceOneOfAsTextFormattingRule20270101Replace is a convenience function that returns TextFormattingRule20270101ReplaceOneOf wrapped in TextFormattingRule20270101Replace.
func TextFormattingRule20270101ReplaceOneOfAsTextFormattingRule20270101Replace(v *TextFormattingRule20270101ReplaceOneOf) TextFormattingRule20270101Replace {
	return TextFormattingRule20270101Replace{TextFormattingRule20270101ReplaceOneOf: v}
}

// TextFormattingRule20270101ReplaceOneOf4462863328AsTextFormattingRule20270101Replace is a convenience function that returns TextFormattingRule20270101ReplaceOneOf4462863328 wrapped in TextFormattingRule20270101Replace.
func TextFormattingRule20270101ReplaceOneOf4462863328AsTextFormattingRule20270101Replace(v *TextFormattingRule20270101ReplaceOneOf4462863328) TextFormattingRule20270101Replace {
	return TextFormattingRule20270101Replace{TextFormattingRule20270101ReplaceOneOf4462863328: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TextFormattingRule20270101Replace) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TextFormattingRule20270101ReplaceOneOf
	err = datadog.Unmarshal(data, &obj.TextFormattingRule20270101ReplaceOneOf)
	if err == nil {
		if obj.TextFormattingRule20270101ReplaceOneOf != nil && obj.TextFormattingRule20270101ReplaceOneOf.UnparsedObject == nil {
			jsonTextFormattingRule20270101ReplaceOneOf, _ := datadog.Marshal(obj.TextFormattingRule20270101ReplaceOneOf)
			if string(jsonTextFormattingRule20270101ReplaceOneOf) == "{}" { // empty struct
				obj.TextFormattingRule20270101ReplaceOneOf = nil
			} else {
				match++
			}
		} else {
			obj.TextFormattingRule20270101ReplaceOneOf = nil
		}
	} else {
		obj.TextFormattingRule20270101ReplaceOneOf = nil
	}

	// try to unmarshal data into TextFormattingRule20270101ReplaceOneOf4462863328
	err = datadog.Unmarshal(data, &obj.TextFormattingRule20270101ReplaceOneOf4462863328)
	if err == nil {
		if obj.TextFormattingRule20270101ReplaceOneOf4462863328 != nil && obj.TextFormattingRule20270101ReplaceOneOf4462863328.UnparsedObject == nil {
			jsonTextFormattingRule20270101ReplaceOneOf4462863328, _ := datadog.Marshal(obj.TextFormattingRule20270101ReplaceOneOf4462863328)
			if string(jsonTextFormattingRule20270101ReplaceOneOf4462863328) == "{}" { // empty struct
				obj.TextFormattingRule20270101ReplaceOneOf4462863328 = nil
			} else {
				match++
			}
		} else {
			obj.TextFormattingRule20270101ReplaceOneOf4462863328 = nil
		}
	} else {
		obj.TextFormattingRule20270101ReplaceOneOf4462863328 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TextFormattingRule20270101ReplaceOneOf = nil
		obj.TextFormattingRule20270101ReplaceOneOf4462863328 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TextFormattingRule20270101Replace) MarshalJSON() ([]byte, error) {
	if obj.TextFormattingRule20270101ReplaceOneOf != nil {
		return datadog.Marshal(&obj.TextFormattingRule20270101ReplaceOneOf)
	}

	if obj.TextFormattingRule20270101ReplaceOneOf4462863328 != nil {
		return datadog.Marshal(&obj.TextFormattingRule20270101ReplaceOneOf4462863328)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TextFormattingRule20270101Replace) GetActualInstance() interface{} {
	if obj.TextFormattingRule20270101ReplaceOneOf != nil {
		return obj.TextFormattingRule20270101ReplaceOneOf
	}

	if obj.TextFormattingRule20270101ReplaceOneOf4462863328 != nil {
		return obj.TextFormattingRule20270101ReplaceOneOf4462863328
	}

	// all schemas are nil
	return nil
}
