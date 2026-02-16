// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataTransformationLanguage The programming language for the transformation.
type DataTransformationLanguage string

// List of DataTransformationLanguage.
const (
	DATATRANSFORMATIONLANGUAGE_JAVASCRIPT DataTransformationLanguage = "javascript"
	DATATRANSFORMATIONLANGUAGE_PYTHON     DataTransformationLanguage = "python"
)

var allowedDataTransformationLanguageEnumValues = []DataTransformationLanguage{
	DATATRANSFORMATIONLANGUAGE_JAVASCRIPT,
	DATATRANSFORMATIONLANGUAGE_PYTHON,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DataTransformationLanguage) GetAllowedValues() []DataTransformationLanguage {
	return allowedDataTransformationLanguageEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DataTransformationLanguage) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DataTransformationLanguage(value)
	return nil
}

// NewDataTransformationLanguageFromValue returns a pointer to a valid DataTransformationLanguage
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDataTransformationLanguageFromValue(v string) (*DataTransformationLanguage, error) {
	ev := DataTransformationLanguage(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DataTransformationLanguage: valid values are %v", v, allowedDataTransformationLanguageEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DataTransformationLanguage) IsValid() bool {
	for _, existing := range allowedDataTransformationLanguageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataTransformationLanguage value.
func (v DataTransformationLanguage) Ptr() *DataTransformationLanguage {
	return &v
}
