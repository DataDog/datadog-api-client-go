// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableTextFormattingRuleReplace - Optional replacement to apply to matched text.
type GuidedTableTextFormattingRuleReplace struct {
	TableWidgetTextFormatReplaceAll               *TableWidgetTextFormatReplaceAll
	GuidedTableTextFormattingRuleReplaceSubstring *GuidedTableTextFormattingRuleReplaceSubstring

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TableWidgetTextFormatReplaceAllAsGuidedTableTextFormattingRuleReplace is a convenience function that returns TableWidgetTextFormatReplaceAll wrapped in GuidedTableTextFormattingRuleReplace.
func TableWidgetTextFormatReplaceAllAsGuidedTableTextFormattingRuleReplace(v *TableWidgetTextFormatReplaceAll) GuidedTableTextFormattingRuleReplace {
	return GuidedTableTextFormattingRuleReplace{TableWidgetTextFormatReplaceAll: v}
}

// GuidedTableTextFormattingRuleReplaceSubstringAsGuidedTableTextFormattingRuleReplace is a convenience function that returns GuidedTableTextFormattingRuleReplaceSubstring wrapped in GuidedTableTextFormattingRuleReplace.
func GuidedTableTextFormattingRuleReplaceSubstringAsGuidedTableTextFormattingRuleReplace(v *GuidedTableTextFormattingRuleReplaceSubstring) GuidedTableTextFormattingRuleReplace {
	return GuidedTableTextFormattingRuleReplace{GuidedTableTextFormattingRuleReplaceSubstring: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GuidedTableTextFormattingRuleReplace) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TableWidgetTextFormatReplaceAll
	err = datadog.Unmarshal(data, &obj.TableWidgetTextFormatReplaceAll)
	if err == nil {
		if obj.TableWidgetTextFormatReplaceAll != nil && obj.TableWidgetTextFormatReplaceAll.UnparsedObject == nil {
			jsonTableWidgetTextFormatReplaceAll, _ := datadog.Marshal(obj.TableWidgetTextFormatReplaceAll)
			if string(jsonTableWidgetTextFormatReplaceAll) == "{}" { // empty struct
				obj.TableWidgetTextFormatReplaceAll = nil
			} else {
				match++
			}
		} else {
			obj.TableWidgetTextFormatReplaceAll = nil
		}
	} else {
		obj.TableWidgetTextFormatReplaceAll = nil
	}

	// try to unmarshal data into GuidedTableTextFormattingRuleReplaceSubstring
	err = datadog.Unmarshal(data, &obj.GuidedTableTextFormattingRuleReplaceSubstring)
	if err == nil {
		if obj.GuidedTableTextFormattingRuleReplaceSubstring != nil && obj.GuidedTableTextFormattingRuleReplaceSubstring.UnparsedObject == nil {
			jsonGuidedTableTextFormattingRuleReplaceSubstring, _ := datadog.Marshal(obj.GuidedTableTextFormattingRuleReplaceSubstring)
			if string(jsonGuidedTableTextFormattingRuleReplaceSubstring) == "{}" { // empty struct
				obj.GuidedTableTextFormattingRuleReplaceSubstring = nil
			} else {
				match++
			}
		} else {
			obj.GuidedTableTextFormattingRuleReplaceSubstring = nil
		}
	} else {
		obj.GuidedTableTextFormattingRuleReplaceSubstring = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TableWidgetTextFormatReplaceAll = nil
		obj.GuidedTableTextFormattingRuleReplaceSubstring = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GuidedTableTextFormattingRuleReplace) MarshalJSON() ([]byte, error) {
	if obj.TableWidgetTextFormatReplaceAll != nil {
		return datadog.Marshal(&obj.TableWidgetTextFormatReplaceAll)
	}

	if obj.GuidedTableTextFormattingRuleReplaceSubstring != nil {
		return datadog.Marshal(&obj.GuidedTableTextFormattingRuleReplaceSubstring)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GuidedTableTextFormattingRuleReplace) GetActualInstance() interface{} {
	if obj.TableWidgetTextFormatReplaceAll != nil {
		return obj.TableWidgetTextFormatReplaceAll
	}

	if obj.GuidedTableTextFormattingRuleReplaceSubstring != nil {
		return obj.GuidedTableTextFormattingRuleReplaceSubstring
	}

	// all schemas are nil
	return nil
}
