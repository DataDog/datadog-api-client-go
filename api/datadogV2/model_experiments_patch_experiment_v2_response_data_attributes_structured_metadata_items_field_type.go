// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType Type of value stored in the structured metadata field.
type ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType string

// List of ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType.
const (
	EXPERIMENTSPATCHEXPERIMENTV2RESPONSEDATAATTRIBUTESSTRUCTUREDMETADATAITEMSFIELDTYPE_FREETEXT ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType = "FREETEXT"
	EXPERIMENTSPATCHEXPERIMENTV2RESPONSEDATAATTRIBUTESSTRUCTUREDMETADATAITEMSFIELDTYPE_ENUM     ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType = "ENUM"
)

var allowedExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldTypeEnumValues = []ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType{
	EXPERIMENTSPATCHEXPERIMENTV2RESPONSEDATAATTRIBUTESSTRUCTUREDMETADATAITEMSFIELDTYPE_FREETEXT,
	EXPERIMENTSPATCHEXPERIMENTV2RESPONSEDATAATTRIBUTESSTRUCTUREDMETADATAITEMSFIELDTYPE_ENUM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType) GetAllowedValues() []ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType {
	return allowedExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType(value)
	return nil
}

// NewExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldTypeFromValue returns a pointer to a valid ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldTypeFromValue(v string) (*ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType, error) {
	ev := ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType: valid values are %v", v, allowedExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType) IsValid() bool {
	for _, existing := range allowedExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType value.
func (v ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType) Ptr() *ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType {
	return &v
}
