// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TextReplacementRule20270101Replace - Test-only generated schema for the dashboard versioning proof of concept.
type TextReplacementRule20270101Replace struct {
	TextReplacementRule20270101ReplaceOneOf           *TextReplacementRule20270101ReplaceOneOf
	TextReplacementRule20270101ReplaceOneOf4462870048 *TextReplacementRule20270101ReplaceOneOf4462870048

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TextReplacementRule20270101ReplaceOneOfAsTextReplacementRule20270101Replace is a convenience function that returns TextReplacementRule20270101ReplaceOneOf wrapped in TextReplacementRule20270101Replace.
func TextReplacementRule20270101ReplaceOneOfAsTextReplacementRule20270101Replace(v *TextReplacementRule20270101ReplaceOneOf) TextReplacementRule20270101Replace {
	return TextReplacementRule20270101Replace{TextReplacementRule20270101ReplaceOneOf: v}
}

// TextReplacementRule20270101ReplaceOneOf4462870048AsTextReplacementRule20270101Replace is a convenience function that returns TextReplacementRule20270101ReplaceOneOf4462870048 wrapped in TextReplacementRule20270101Replace.
func TextReplacementRule20270101ReplaceOneOf4462870048AsTextReplacementRule20270101Replace(v *TextReplacementRule20270101ReplaceOneOf4462870048) TextReplacementRule20270101Replace {
	return TextReplacementRule20270101Replace{TextReplacementRule20270101ReplaceOneOf4462870048: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TextReplacementRule20270101Replace) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TextReplacementRule20270101ReplaceOneOf
	err = datadog.Unmarshal(data, &obj.TextReplacementRule20270101ReplaceOneOf)
	if err == nil {
		if obj.TextReplacementRule20270101ReplaceOneOf != nil && obj.TextReplacementRule20270101ReplaceOneOf.UnparsedObject == nil {
			jsonTextReplacementRule20270101ReplaceOneOf, _ := datadog.Marshal(obj.TextReplacementRule20270101ReplaceOneOf)
			if string(jsonTextReplacementRule20270101ReplaceOneOf) == "{}" { // empty struct
				obj.TextReplacementRule20270101ReplaceOneOf = nil
			} else {
				match++
			}
		} else {
			obj.TextReplacementRule20270101ReplaceOneOf = nil
		}
	} else {
		obj.TextReplacementRule20270101ReplaceOneOf = nil
	}

	// try to unmarshal data into TextReplacementRule20270101ReplaceOneOf4462870048
	err = datadog.Unmarshal(data, &obj.TextReplacementRule20270101ReplaceOneOf4462870048)
	if err == nil {
		if obj.TextReplacementRule20270101ReplaceOneOf4462870048 != nil && obj.TextReplacementRule20270101ReplaceOneOf4462870048.UnparsedObject == nil {
			jsonTextReplacementRule20270101ReplaceOneOf4462870048, _ := datadog.Marshal(obj.TextReplacementRule20270101ReplaceOneOf4462870048)
			if string(jsonTextReplacementRule20270101ReplaceOneOf4462870048) == "{}" { // empty struct
				obj.TextReplacementRule20270101ReplaceOneOf4462870048 = nil
			} else {
				match++
			}
		} else {
			obj.TextReplacementRule20270101ReplaceOneOf4462870048 = nil
		}
	} else {
		obj.TextReplacementRule20270101ReplaceOneOf4462870048 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TextReplacementRule20270101ReplaceOneOf = nil
		obj.TextReplacementRule20270101ReplaceOneOf4462870048 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TextReplacementRule20270101Replace) MarshalJSON() ([]byte, error) {
	if obj.TextReplacementRule20270101ReplaceOneOf != nil {
		return datadog.Marshal(&obj.TextReplacementRule20270101ReplaceOneOf)
	}

	if obj.TextReplacementRule20270101ReplaceOneOf4462870048 != nil {
		return datadog.Marshal(&obj.TextReplacementRule20270101ReplaceOneOf4462870048)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TextReplacementRule20270101Replace) GetActualInstance() interface{} {
	if obj.TextReplacementRule20270101ReplaceOneOf != nil {
		return obj.TextReplacementRule20270101ReplaceOneOf
	}

	if obj.TextReplacementRule20270101ReplaceOneOf4462870048 != nil {
		return obj.TextReplacementRule20270101ReplaceOneOf4462870048
	}

	// all schemas are nil
	return nil
}
