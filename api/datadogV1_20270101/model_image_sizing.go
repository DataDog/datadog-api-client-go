// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ImageSizing Test-only generated schema for the dashboard versioning proof of concept.
type ImageSizing string

// List of ImageSizing.
const (
	IMAGESIZING_NONE       ImageSizing = "none"
	IMAGESIZING_CENTER     ImageSizing = "center"
	IMAGESIZING_FILL       ImageSizing = "fill"
	IMAGESIZING_CONTAIN    ImageSizing = "contain"
	IMAGESIZING_COVER      ImageSizing = "cover"
	IMAGESIZING_SCALE_DOWN ImageSizing = "scale-down"
	IMAGESIZING_ZOOM       ImageSizing = "zoom"
	IMAGESIZING_FIT        ImageSizing = "fit"
)

var allowedImageSizingEnumValues = []ImageSizing{
	IMAGESIZING_NONE,
	IMAGESIZING_CENTER,
	IMAGESIZING_FILL,
	IMAGESIZING_CONTAIN,
	IMAGESIZING_COVER,
	IMAGESIZING_SCALE_DOWN,
	IMAGESIZING_ZOOM,
	IMAGESIZING_FIT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ImageSizing) GetAllowedValues() []ImageSizing {
	return allowedImageSizingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ImageSizing) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ImageSizing(value)
	return nil
}

// NewImageSizingFromValue returns a pointer to a valid ImageSizing
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewImageSizingFromValue(v string) (*ImageSizing, error) {
	ev := ImageSizing(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ImageSizing: valid values are %v", v, allowedImageSizingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ImageSizing) IsValid() bool {
	for _, existing := range allowedImageSizingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageSizing value.
func (v ImageSizing) Ptr() *ImageSizing {
	return &v
}
