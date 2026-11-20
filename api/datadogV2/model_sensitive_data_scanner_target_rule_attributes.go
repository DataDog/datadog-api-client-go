// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerTargetRuleAttributes - Attributes of a rule of a scanning group of a target. They depend on the type of rule.
type SensitiveDataScannerTargetRuleAttributes struct {
	SensitiveDataScannerCustomRegexRuleAttributes     *SensitiveDataScannerCustomRegexRuleAttributes
	SensitiveDataScannerStandardRegexRuleAttributes   *SensitiveDataScannerStandardRegexRuleAttributes
	SensitiveDataScannerMachineLearningRuleAttributes *SensitiveDataScannerMachineLearningRuleAttributes

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SensitiveDataScannerCustomRegexRuleAttributesAsSensitiveDataScannerTargetRuleAttributes is a convenience function that returns SensitiveDataScannerCustomRegexRuleAttributes wrapped in SensitiveDataScannerTargetRuleAttributes.
func SensitiveDataScannerCustomRegexRuleAttributesAsSensitiveDataScannerTargetRuleAttributes(v *SensitiveDataScannerCustomRegexRuleAttributes) SensitiveDataScannerTargetRuleAttributes {
	return SensitiveDataScannerTargetRuleAttributes{SensitiveDataScannerCustomRegexRuleAttributes: v}
}

// SensitiveDataScannerStandardRegexRuleAttributesAsSensitiveDataScannerTargetRuleAttributes is a convenience function that returns SensitiveDataScannerStandardRegexRuleAttributes wrapped in SensitiveDataScannerTargetRuleAttributes.
func SensitiveDataScannerStandardRegexRuleAttributesAsSensitiveDataScannerTargetRuleAttributes(v *SensitiveDataScannerStandardRegexRuleAttributes) SensitiveDataScannerTargetRuleAttributes {
	return SensitiveDataScannerTargetRuleAttributes{SensitiveDataScannerStandardRegexRuleAttributes: v}
}

// SensitiveDataScannerMachineLearningRuleAttributesAsSensitiveDataScannerTargetRuleAttributes is a convenience function that returns SensitiveDataScannerMachineLearningRuleAttributes wrapped in SensitiveDataScannerTargetRuleAttributes.
func SensitiveDataScannerMachineLearningRuleAttributesAsSensitiveDataScannerTargetRuleAttributes(v *SensitiveDataScannerMachineLearningRuleAttributes) SensitiveDataScannerTargetRuleAttributes {
	return SensitiveDataScannerTargetRuleAttributes{SensitiveDataScannerMachineLearningRuleAttributes: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SensitiveDataScannerTargetRuleAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SensitiveDataScannerCustomRegexRuleAttributes
	err = datadog.Unmarshal(data, &obj.SensitiveDataScannerCustomRegexRuleAttributes)
	if err == nil {
		if obj.SensitiveDataScannerCustomRegexRuleAttributes != nil && obj.SensitiveDataScannerCustomRegexRuleAttributes.UnparsedObject == nil {
			jsonSensitiveDataScannerCustomRegexRuleAttributes, _ := datadog.Marshal(obj.SensitiveDataScannerCustomRegexRuleAttributes)
			if string(jsonSensitiveDataScannerCustomRegexRuleAttributes) == "{}" { // empty struct
				obj.SensitiveDataScannerCustomRegexRuleAttributes = nil
			} else {
				match++
			}
		} else {
			obj.SensitiveDataScannerCustomRegexRuleAttributes = nil
		}
	} else {
		obj.SensitiveDataScannerCustomRegexRuleAttributes = nil
	}

	// try to unmarshal data into SensitiveDataScannerStandardRegexRuleAttributes
	err = datadog.Unmarshal(data, &obj.SensitiveDataScannerStandardRegexRuleAttributes)
	if err == nil {
		if obj.SensitiveDataScannerStandardRegexRuleAttributes != nil && obj.SensitiveDataScannerStandardRegexRuleAttributes.UnparsedObject == nil {
			jsonSensitiveDataScannerStandardRegexRuleAttributes, _ := datadog.Marshal(obj.SensitiveDataScannerStandardRegexRuleAttributes)
			if string(jsonSensitiveDataScannerStandardRegexRuleAttributes) == "{}" { // empty struct
				obj.SensitiveDataScannerStandardRegexRuleAttributes = nil
			} else {
				match++
			}
		} else {
			obj.SensitiveDataScannerStandardRegexRuleAttributes = nil
		}
	} else {
		obj.SensitiveDataScannerStandardRegexRuleAttributes = nil
	}

	// try to unmarshal data into SensitiveDataScannerMachineLearningRuleAttributes
	err = datadog.Unmarshal(data, &obj.SensitiveDataScannerMachineLearningRuleAttributes)
	if err == nil {
		if obj.SensitiveDataScannerMachineLearningRuleAttributes != nil && obj.SensitiveDataScannerMachineLearningRuleAttributes.UnparsedObject == nil {
			jsonSensitiveDataScannerMachineLearningRuleAttributes, _ := datadog.Marshal(obj.SensitiveDataScannerMachineLearningRuleAttributes)
			if string(jsonSensitiveDataScannerMachineLearningRuleAttributes) == "{}" { // empty struct
				obj.SensitiveDataScannerMachineLearningRuleAttributes = nil
			} else {
				match++
			}
		} else {
			obj.SensitiveDataScannerMachineLearningRuleAttributes = nil
		}
	} else {
		obj.SensitiveDataScannerMachineLearningRuleAttributes = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SensitiveDataScannerCustomRegexRuleAttributes = nil
		obj.SensitiveDataScannerStandardRegexRuleAttributes = nil
		obj.SensitiveDataScannerMachineLearningRuleAttributes = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SensitiveDataScannerTargetRuleAttributes) MarshalJSON() ([]byte, error) {
	if obj.SensitiveDataScannerCustomRegexRuleAttributes != nil {
		return datadog.Marshal(&obj.SensitiveDataScannerCustomRegexRuleAttributes)
	}

	if obj.SensitiveDataScannerStandardRegexRuleAttributes != nil {
		return datadog.Marshal(&obj.SensitiveDataScannerStandardRegexRuleAttributes)
	}

	if obj.SensitiveDataScannerMachineLearningRuleAttributes != nil {
		return datadog.Marshal(&obj.SensitiveDataScannerMachineLearningRuleAttributes)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SensitiveDataScannerTargetRuleAttributes) GetActualInstance() interface{} {
	if obj.SensitiveDataScannerCustomRegexRuleAttributes != nil {
		return obj.SensitiveDataScannerCustomRegexRuleAttributes
	}

	if obj.SensitiveDataScannerStandardRegexRuleAttributes != nil {
		return obj.SensitiveDataScannerStandardRegexRuleAttributes
	}

	if obj.SensitiveDataScannerMachineLearningRuleAttributes != nil {
		return obj.SensitiveDataScannerMachineLearningRuleAttributes
	}

	// all schemas are nil
	return nil
}
