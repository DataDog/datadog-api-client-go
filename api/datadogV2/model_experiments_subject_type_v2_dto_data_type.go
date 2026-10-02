// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsSubjectTypeV2DTODataType Subject types resource type.
type ExperimentsSubjectTypeV2DTODataType string

// List of ExperimentsSubjectTypeV2DTODataType.
const (
	EXPERIMENTSSUBJECTTYPEV2DTODATATYPE_SUBJECT_TYPES ExperimentsSubjectTypeV2DTODataType = "subject-types"
)

var allowedExperimentsSubjectTypeV2DTODataTypeEnumValues = []ExperimentsSubjectTypeV2DTODataType{
	EXPERIMENTSSUBJECTTYPEV2DTODATATYPE_SUBJECT_TYPES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsSubjectTypeV2DTODataType) GetAllowedValues() []ExperimentsSubjectTypeV2DTODataType {
	return allowedExperimentsSubjectTypeV2DTODataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsSubjectTypeV2DTODataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsSubjectTypeV2DTODataType(value)
	return nil
}

// NewExperimentsSubjectTypeV2DTODataTypeFromValue returns a pointer to a valid ExperimentsSubjectTypeV2DTODataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsSubjectTypeV2DTODataTypeFromValue(v string) (*ExperimentsSubjectTypeV2DTODataType, error) {
	ev := ExperimentsSubjectTypeV2DTODataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsSubjectTypeV2DTODataType: valid values are %v", v, allowedExperimentsSubjectTypeV2DTODataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsSubjectTypeV2DTODataType) IsValid() bool {
	for _, existing := range allowedExperimentsSubjectTypeV2DTODataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsSubjectTypeV2DTODataType value.
func (v ExperimentsSubjectTypeV2DTODataType) Ptr() *ExperimentsSubjectTypeV2DTODataType {
	return &v
}
