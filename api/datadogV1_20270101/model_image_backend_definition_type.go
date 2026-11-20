// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ImageBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type ImageBackendDefinitionType string

// List of ImageBackendDefinitionType.
const (
	IMAGEBACKENDDEFINITIONTYPE_IMAGE ImageBackendDefinitionType = "image"
)

var allowedImageBackendDefinitionTypeEnumValues = []ImageBackendDefinitionType{
	IMAGEBACKENDDEFINITIONTYPE_IMAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ImageBackendDefinitionType) GetAllowedValues() []ImageBackendDefinitionType {
	return allowedImageBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ImageBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ImageBackendDefinitionType(value)
	return nil
}

// NewImageBackendDefinitionTypeFromValue returns a pointer to a valid ImageBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewImageBackendDefinitionTypeFromValue(v string) (*ImageBackendDefinitionType, error) {
	ev := ImageBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ImageBackendDefinitionType: valid values are %v", v, allowedImageBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ImageBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedImageBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageBackendDefinitionType value.
func (v ImageBackendDefinitionType) Ptr() *ImageBackendDefinitionType {
	return &v
}
