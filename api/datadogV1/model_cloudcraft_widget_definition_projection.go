// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudcraftWidgetDefinitionProjection Projection used to render the Cloudcraft topology.
type CloudcraftWidgetDefinitionProjection string

// List of CloudcraftWidgetDefinitionProjection.
const (
	CLOUDCRAFTWIDGETDEFINITIONPROJECTION_ISOMETRIC CloudcraftWidgetDefinitionProjection = "isometric"
	CLOUDCRAFTWIDGETDEFINITIONPROJECTION_TWO_D     CloudcraftWidgetDefinitionProjection = "2d"
)

var allowedCloudcraftWidgetDefinitionProjectionEnumValues = []CloudcraftWidgetDefinitionProjection{
	CLOUDCRAFTWIDGETDEFINITIONPROJECTION_ISOMETRIC,
	CLOUDCRAFTWIDGETDEFINITIONPROJECTION_TWO_D,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CloudcraftWidgetDefinitionProjection) GetAllowedValues() []CloudcraftWidgetDefinitionProjection {
	return allowedCloudcraftWidgetDefinitionProjectionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CloudcraftWidgetDefinitionProjection) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CloudcraftWidgetDefinitionProjection(value)
	return nil
}

// NewCloudcraftWidgetDefinitionProjectionFromValue returns a pointer to a valid CloudcraftWidgetDefinitionProjection
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCloudcraftWidgetDefinitionProjectionFromValue(v string) (*CloudcraftWidgetDefinitionProjection, error) {
	ev := CloudcraftWidgetDefinitionProjection(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CloudcraftWidgetDefinitionProjection: valid values are %v", v, allowedCloudcraftWidgetDefinitionProjectionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CloudcraftWidgetDefinitionProjection) IsValid() bool {
	for _, existing := range allowedCloudcraftWidgetDefinitionProjectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudcraftWidgetDefinitionProjection value.
func (v CloudcraftWidgetDefinitionProjection) Ptr() *CloudcraftWidgetDefinitionProjection {
	return &v
}
