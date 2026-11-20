// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ImageBackendDefinitionHorizontalAlign Test-only generated schema for the dashboard versioning proof of concept.
type ImageBackendDefinitionHorizontalAlign string

// List of ImageBackendDefinitionHorizontalAlign.
const (
	IMAGEBACKENDDEFINITIONHORIZONTALALIGN_LEFT   ImageBackendDefinitionHorizontalAlign = "left"
	IMAGEBACKENDDEFINITIONHORIZONTALALIGN_CENTER ImageBackendDefinitionHorizontalAlign = "center"
	IMAGEBACKENDDEFINITIONHORIZONTALALIGN_RIGHT  ImageBackendDefinitionHorizontalAlign = "right"
)

var allowedImageBackendDefinitionHorizontalAlignEnumValues = []ImageBackendDefinitionHorizontalAlign{
	IMAGEBACKENDDEFINITIONHORIZONTALALIGN_LEFT,
	IMAGEBACKENDDEFINITIONHORIZONTALALIGN_CENTER,
	IMAGEBACKENDDEFINITIONHORIZONTALALIGN_RIGHT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ImageBackendDefinitionHorizontalAlign) GetAllowedValues() []ImageBackendDefinitionHorizontalAlign {
	return allowedImageBackendDefinitionHorizontalAlignEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ImageBackendDefinitionHorizontalAlign) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ImageBackendDefinitionHorizontalAlign(value)
	return nil
}

// NewImageBackendDefinitionHorizontalAlignFromValue returns a pointer to a valid ImageBackendDefinitionHorizontalAlign
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewImageBackendDefinitionHorizontalAlignFromValue(v string) (*ImageBackendDefinitionHorizontalAlign, error) {
	ev := ImageBackendDefinitionHorizontalAlign(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ImageBackendDefinitionHorizontalAlign: valid values are %v", v, allowedImageBackendDefinitionHorizontalAlignEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ImageBackendDefinitionHorizontalAlign) IsValid() bool {
	for _, existing := range allowedImageBackendDefinitionHorizontalAlignEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageBackendDefinitionHorizontalAlign value.
func (v ImageBackendDefinitionHorizontalAlign) Ptr() *ImageBackendDefinitionHorizontalAlign {
	return &v
}
