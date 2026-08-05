// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchDegradationTemplateRequestDataType Degradation templates resource type.
type PatchDegradationTemplateRequestDataType string

// List of PatchDegradationTemplateRequestDataType.
const (
	PATCHDEGRADATIONTEMPLATEREQUESTDATATYPE_DEGRADATION_TEMPLATES PatchDegradationTemplateRequestDataType = "degradation_templates"
)

var allowedPatchDegradationTemplateRequestDataTypeEnumValues = []PatchDegradationTemplateRequestDataType{
	PATCHDEGRADATIONTEMPLATEREQUESTDATATYPE_DEGRADATION_TEMPLATES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PatchDegradationTemplateRequestDataType) GetAllowedValues() []PatchDegradationTemplateRequestDataType {
	return allowedPatchDegradationTemplateRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PatchDegradationTemplateRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PatchDegradationTemplateRequestDataType(value)
	return nil
}

// NewPatchDegradationTemplateRequestDataTypeFromValue returns a pointer to a valid PatchDegradationTemplateRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPatchDegradationTemplateRequestDataTypeFromValue(v string) (*PatchDegradationTemplateRequestDataType, error) {
	ev := PatchDegradationTemplateRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PatchDegradationTemplateRequestDataType: valid values are %v", v, allowedPatchDegradationTemplateRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PatchDegradationTemplateRequestDataType) IsValid() bool {
	for _, existing := range allowedPatchDegradationTemplateRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchDegradationTemplateRequestDataType value.
func (v PatchDegradationTemplateRequestDataType) Ptr() *PatchDegradationTemplateRequestDataType {
	return &v
}
