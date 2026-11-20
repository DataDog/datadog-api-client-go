// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ImageBackendDefinitionVerticalAlign Test-only generated schema for the dashboard versioning proof of concept.
type ImageBackendDefinitionVerticalAlign string

// List of ImageBackendDefinitionVerticalAlign.
const (
	IMAGEBACKENDDEFINITIONVERTICALALIGN_CENTER ImageBackendDefinitionVerticalAlign = "center"
	IMAGEBACKENDDEFINITIONVERTICALALIGN_TOP    ImageBackendDefinitionVerticalAlign = "top"
	IMAGEBACKENDDEFINITIONVERTICALALIGN_BOTTOM ImageBackendDefinitionVerticalAlign = "bottom"
)

var allowedImageBackendDefinitionVerticalAlignEnumValues = []ImageBackendDefinitionVerticalAlign{
	IMAGEBACKENDDEFINITIONVERTICALALIGN_CENTER,
	IMAGEBACKENDDEFINITIONVERTICALALIGN_TOP,
	IMAGEBACKENDDEFINITIONVERTICALALIGN_BOTTOM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ImageBackendDefinitionVerticalAlign) GetAllowedValues() []ImageBackendDefinitionVerticalAlign {
	return allowedImageBackendDefinitionVerticalAlignEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ImageBackendDefinitionVerticalAlign) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ImageBackendDefinitionVerticalAlign(value)
	return nil
}

// NewImageBackendDefinitionVerticalAlignFromValue returns a pointer to a valid ImageBackendDefinitionVerticalAlign
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewImageBackendDefinitionVerticalAlignFromValue(v string) (*ImageBackendDefinitionVerticalAlign, error) {
	ev := ImageBackendDefinitionVerticalAlign(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ImageBackendDefinitionVerticalAlign: valid values are %v", v, allowedImageBackendDefinitionVerticalAlignEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ImageBackendDefinitionVerticalAlign) IsValid() bool {
	for _, existing := range allowedImageBackendDefinitionVerticalAlignEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageBackendDefinitionVerticalAlign value.
func (v ImageBackendDefinitionVerticalAlign) Ptr() *ImageBackendDefinitionVerticalAlign {
	return &v
}
