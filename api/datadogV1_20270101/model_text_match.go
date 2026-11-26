// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TextMatch - Test-only generated schema for the dashboard versioning proof of concept.
type TextMatch struct {
	TextMatch20270101OneOf           *TextMatch20270101OneOf
	TextMatch20270101OneOf4462865408 *TextMatch20270101OneOf4462865408
	TextMatch20270101OneOf4462866048 *TextMatch20270101OneOf4462866048
	TextMatch20270101OneOf4462866688 *TextMatch20270101OneOf4462866688
	TextMatch20270101OneOf4462867328 *TextMatch20270101OneOf4462867328
	TextMatch20270101OneOf4462867968 *TextMatch20270101OneOf4462867968

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TextMatch20270101OneOfAsTextMatch is a convenience function that returns TextMatch20270101OneOf wrapped in TextMatch.
func TextMatch20270101OneOfAsTextMatch(v *TextMatch20270101OneOf) TextMatch {
	return TextMatch{TextMatch20270101OneOf: v}
}

// TextMatch20270101OneOf4462865408AsTextMatch is a convenience function that returns TextMatch20270101OneOf4462865408 wrapped in TextMatch.
func TextMatch20270101OneOf4462865408AsTextMatch(v *TextMatch20270101OneOf4462865408) TextMatch {
	return TextMatch{TextMatch20270101OneOf4462865408: v}
}

// TextMatch20270101OneOf4462866048AsTextMatch is a convenience function that returns TextMatch20270101OneOf4462866048 wrapped in TextMatch.
func TextMatch20270101OneOf4462866048AsTextMatch(v *TextMatch20270101OneOf4462866048) TextMatch {
	return TextMatch{TextMatch20270101OneOf4462866048: v}
}

// TextMatch20270101OneOf4462866688AsTextMatch is a convenience function that returns TextMatch20270101OneOf4462866688 wrapped in TextMatch.
func TextMatch20270101OneOf4462866688AsTextMatch(v *TextMatch20270101OneOf4462866688) TextMatch {
	return TextMatch{TextMatch20270101OneOf4462866688: v}
}

// TextMatch20270101OneOf4462867328AsTextMatch is a convenience function that returns TextMatch20270101OneOf4462867328 wrapped in TextMatch.
func TextMatch20270101OneOf4462867328AsTextMatch(v *TextMatch20270101OneOf4462867328) TextMatch {
	return TextMatch{TextMatch20270101OneOf4462867328: v}
}

// TextMatch20270101OneOf4462867968AsTextMatch is a convenience function that returns TextMatch20270101OneOf4462867968 wrapped in TextMatch.
func TextMatch20270101OneOf4462867968AsTextMatch(v *TextMatch20270101OneOf4462867968) TextMatch {
	return TextMatch{TextMatch20270101OneOf4462867968: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TextMatch) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TextMatch20270101OneOf
	err = datadog.Unmarshal(data, &obj.TextMatch20270101OneOf)
	if err == nil {
		if obj.TextMatch20270101OneOf != nil && obj.TextMatch20270101OneOf.UnparsedObject == nil {
			jsonTextMatch20270101OneOf, _ := datadog.Marshal(obj.TextMatch20270101OneOf)
			if string(jsonTextMatch20270101OneOf) == "{}" { // empty struct
				obj.TextMatch20270101OneOf = nil
			} else {
				match++
			}
		} else {
			obj.TextMatch20270101OneOf = nil
		}
	} else {
		obj.TextMatch20270101OneOf = nil
	}

	// try to unmarshal data into TextMatch20270101OneOf4462865408
	err = datadog.Unmarshal(data, &obj.TextMatch20270101OneOf4462865408)
	if err == nil {
		if obj.TextMatch20270101OneOf4462865408 != nil && obj.TextMatch20270101OneOf4462865408.UnparsedObject == nil {
			jsonTextMatch20270101OneOf4462865408, _ := datadog.Marshal(obj.TextMatch20270101OneOf4462865408)
			if string(jsonTextMatch20270101OneOf4462865408) == "{}" { // empty struct
				obj.TextMatch20270101OneOf4462865408 = nil
			} else {
				match++
			}
		} else {
			obj.TextMatch20270101OneOf4462865408 = nil
		}
	} else {
		obj.TextMatch20270101OneOf4462865408 = nil
	}

	// try to unmarshal data into TextMatch20270101OneOf4462866048
	err = datadog.Unmarshal(data, &obj.TextMatch20270101OneOf4462866048)
	if err == nil {
		if obj.TextMatch20270101OneOf4462866048 != nil && obj.TextMatch20270101OneOf4462866048.UnparsedObject == nil {
			jsonTextMatch20270101OneOf4462866048, _ := datadog.Marshal(obj.TextMatch20270101OneOf4462866048)
			if string(jsonTextMatch20270101OneOf4462866048) == "{}" { // empty struct
				obj.TextMatch20270101OneOf4462866048 = nil
			} else {
				match++
			}
		} else {
			obj.TextMatch20270101OneOf4462866048 = nil
		}
	} else {
		obj.TextMatch20270101OneOf4462866048 = nil
	}

	// try to unmarshal data into TextMatch20270101OneOf4462866688
	err = datadog.Unmarshal(data, &obj.TextMatch20270101OneOf4462866688)
	if err == nil {
		if obj.TextMatch20270101OneOf4462866688 != nil && obj.TextMatch20270101OneOf4462866688.UnparsedObject == nil {
			jsonTextMatch20270101OneOf4462866688, _ := datadog.Marshal(obj.TextMatch20270101OneOf4462866688)
			if string(jsonTextMatch20270101OneOf4462866688) == "{}" { // empty struct
				obj.TextMatch20270101OneOf4462866688 = nil
			} else {
				match++
			}
		} else {
			obj.TextMatch20270101OneOf4462866688 = nil
		}
	} else {
		obj.TextMatch20270101OneOf4462866688 = nil
	}

	// try to unmarshal data into TextMatch20270101OneOf4462867328
	err = datadog.Unmarshal(data, &obj.TextMatch20270101OneOf4462867328)
	if err == nil {
		if obj.TextMatch20270101OneOf4462867328 != nil && obj.TextMatch20270101OneOf4462867328.UnparsedObject == nil {
			jsonTextMatch20270101OneOf4462867328, _ := datadog.Marshal(obj.TextMatch20270101OneOf4462867328)
			if string(jsonTextMatch20270101OneOf4462867328) == "{}" { // empty struct
				obj.TextMatch20270101OneOf4462867328 = nil
			} else {
				match++
			}
		} else {
			obj.TextMatch20270101OneOf4462867328 = nil
		}
	} else {
		obj.TextMatch20270101OneOf4462867328 = nil
	}

	// try to unmarshal data into TextMatch20270101OneOf4462867968
	err = datadog.Unmarshal(data, &obj.TextMatch20270101OneOf4462867968)
	if err == nil {
		if obj.TextMatch20270101OneOf4462867968 != nil && obj.TextMatch20270101OneOf4462867968.UnparsedObject == nil {
			jsonTextMatch20270101OneOf4462867968, _ := datadog.Marshal(obj.TextMatch20270101OneOf4462867968)
			if string(jsonTextMatch20270101OneOf4462867968) == "{}" { // empty struct
				obj.TextMatch20270101OneOf4462867968 = nil
			} else {
				match++
			}
		} else {
			obj.TextMatch20270101OneOf4462867968 = nil
		}
	} else {
		obj.TextMatch20270101OneOf4462867968 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TextMatch20270101OneOf = nil
		obj.TextMatch20270101OneOf4462865408 = nil
		obj.TextMatch20270101OneOf4462866048 = nil
		obj.TextMatch20270101OneOf4462866688 = nil
		obj.TextMatch20270101OneOf4462867328 = nil
		obj.TextMatch20270101OneOf4462867968 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TextMatch) MarshalJSON() ([]byte, error) {
	if obj.TextMatch20270101OneOf != nil {
		return datadog.Marshal(&obj.TextMatch20270101OneOf)
	}

	if obj.TextMatch20270101OneOf4462865408 != nil {
		return datadog.Marshal(&obj.TextMatch20270101OneOf4462865408)
	}

	if obj.TextMatch20270101OneOf4462866048 != nil {
		return datadog.Marshal(&obj.TextMatch20270101OneOf4462866048)
	}

	if obj.TextMatch20270101OneOf4462866688 != nil {
		return datadog.Marshal(&obj.TextMatch20270101OneOf4462866688)
	}

	if obj.TextMatch20270101OneOf4462867328 != nil {
		return datadog.Marshal(&obj.TextMatch20270101OneOf4462867328)
	}

	if obj.TextMatch20270101OneOf4462867968 != nil {
		return datadog.Marshal(&obj.TextMatch20270101OneOf4462867968)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TextMatch) GetActualInstance() interface{} {
	if obj.TextMatch20270101OneOf != nil {
		return obj.TextMatch20270101OneOf
	}

	if obj.TextMatch20270101OneOf4462865408 != nil {
		return obj.TextMatch20270101OneOf4462865408
	}

	if obj.TextMatch20270101OneOf4462866048 != nil {
		return obj.TextMatch20270101OneOf4462866048
	}

	if obj.TextMatch20270101OneOf4462866688 != nil {
		return obj.TextMatch20270101OneOf4462866688
	}

	if obj.TextMatch20270101OneOf4462867328 != nil {
		return obj.TextMatch20270101OneOf4462867328
	}

	if obj.TextMatch20270101OneOf4462867968 != nil {
		return obj.TextMatch20270101OneOf4462867968
	}

	// all schemas are nil
	return nil
}
