// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ImageMargin Test-only generated schema for the dashboard versioning proof of concept.
type ImageMargin string

// List of ImageMargin.
const (
	IMAGEMARGIN_      ImageMargin = ""
	IMAGEMARGIN_SM    ImageMargin = "sm"
	IMAGEMARGIN_MD    ImageMargin = "md"
	IMAGEMARGIN_LG    ImageMargin = "lg"
	IMAGEMARGIN_SMALL ImageMargin = "small"
	IMAGEMARGIN_LARGE ImageMargin = "large"
)

var allowedImageMarginEnumValues = []ImageMargin{
	IMAGEMARGIN_,
	IMAGEMARGIN_SM,
	IMAGEMARGIN_MD,
	IMAGEMARGIN_LG,
	IMAGEMARGIN_SMALL,
	IMAGEMARGIN_LARGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ImageMargin) GetAllowedValues() []ImageMargin {
	return allowedImageMarginEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ImageMargin) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ImageMargin(value)
	return nil
}

// NewImageMarginFromValue returns a pointer to a valid ImageMargin
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewImageMarginFromValue(v string) (*ImageMargin, error) {
	ev := ImageMargin(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ImageMargin: valid values are %v", v, allowedImageMarginEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ImageMargin) IsValid() bool {
	for _, existing := range allowedImageMarginEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageMargin value.
func (v ImageMargin) Ptr() *ImageMargin {
	return &v
}
