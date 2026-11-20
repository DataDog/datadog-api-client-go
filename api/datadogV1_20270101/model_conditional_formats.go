// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV1_20270101

import (
	_io "io"
	"time"
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"

)


// ConditionalFormats - Test-only generated schema for the dashboard versioning proof of concept.
type ConditionalFormats struct {
	[]ThresholdFormattingRule *[]ThresholdFormattingRule
	[]interface{} *[]interface{}

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// []ThresholdFormattingRuleAsConditionalFormats is a convenience function that returns []ThresholdFormattingRule wrapped in ConditionalFormats.
func []ThresholdFormattingRuleAsConditionalFormats(v *[]ThresholdFormattingRule) ConditionalFormats {
	return ConditionalFormats{[]ThresholdFormattingRule: v}
}

// []interface{}AsConditionalFormats is a convenience function that returns []interface{} wrapped in ConditionalFormats.
func []interface{}AsConditionalFormats(v *[]interface{}) ConditionalFormats {
	return ConditionalFormats{[]interface{}: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ConditionalFormats) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into []ThresholdFormattingRule
	err = datadog.Unmarshal(data, &obj.[]ThresholdFormattingRule)
	if err == nil {
		if obj.[]ThresholdFormattingRule != nil {
			json[]ThresholdFormattingRule, _ := datadog.Marshal(obj.[]ThresholdFormattingRule)
			if string(json[]ThresholdFormattingRule) == "{}" && string(data) != "{}"  { // empty struct
				obj.[]ThresholdFormattingRule = nil
			} else {
				match++
			}
		} else {
			obj.[]ThresholdFormattingRule = nil
		}
	} else {
		obj.[]ThresholdFormattingRule = nil
	}

	// try to unmarshal data into []interface{}
	err = datadog.Unmarshal(data, &obj.[]interface{})
	if err == nil {
		if obj.[]interface{} != nil {
			json[]interface{}, _ := datadog.Marshal(obj.[]interface{})
			if string(json[]interface{}) == "{}" && string(data) != "{}"  { // empty struct
				obj.[]interface{} = nil
			} else {
				match++
			}
		} else {
			obj.[]interface{} = nil
		}
	} else {
		obj.[]interface{} = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.[]ThresholdFormattingRule = nil
		obj.[]interface{} = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ConditionalFormats) MarshalJSON() ([]byte, error) {
	if obj.[]ThresholdFormattingRule != nil {
		return datadog.Marshal(&obj.[]ThresholdFormattingRule)
	}


	if obj.[]interface{} != nil {
		return datadog.Marshal(&obj.[]interface{})
	}


	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ConditionalFormats) GetActualInstance() (interface{}) {
	if obj.[]ThresholdFormattingRule != nil {
		return obj.[]ThresholdFormattingRule
	}


	if obj.[]interface{} != nil {
		return obj.[]interface{}
	}


	// all schemas are nil
	return nil
}
