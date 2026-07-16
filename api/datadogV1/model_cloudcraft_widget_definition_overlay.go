// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudcraftWidgetDefinitionOverlay Overlay applied on top of the Cloudcraft topology.
type CloudcraftWidgetDefinitionOverlay string

// List of CloudcraftWidgetDefinitionOverlay.
const (
	CLOUDCRAFTWIDGETDEFINITIONOVERLAY_OBSERVABILITY    CloudcraftWidgetDefinitionOverlay = "Observability"
	CLOUDCRAFTWIDGETDEFINITIONOVERLAY_CLOUD_COST       CloudcraftWidgetDefinitionOverlay = "CloudCost"
	CLOUDCRAFTWIDGETDEFINITIONOVERLAY_SECURITY         CloudcraftWidgetDefinitionOverlay = "Security"
	CLOUDCRAFTWIDGETDEFINITIONOVERLAY_NDM_REACHABILITY CloudcraftWidgetDefinitionOverlay = "NDMReachability"
	CLOUDCRAFTWIDGETDEFINITIONOVERLAY_MONITORS         CloudcraftWidgetDefinitionOverlay = "Monitors"
	CLOUDCRAFTWIDGETDEFINITIONOVERLAY_APM              CloudcraftWidgetDefinitionOverlay = "APM"
	CLOUDCRAFTWIDGETDEFINITIONOVERLAY_DEFAULT          CloudcraftWidgetDefinitionOverlay = "Default"
)

var allowedCloudcraftWidgetDefinitionOverlayEnumValues = []CloudcraftWidgetDefinitionOverlay{
	CLOUDCRAFTWIDGETDEFINITIONOVERLAY_OBSERVABILITY,
	CLOUDCRAFTWIDGETDEFINITIONOVERLAY_CLOUD_COST,
	CLOUDCRAFTWIDGETDEFINITIONOVERLAY_SECURITY,
	CLOUDCRAFTWIDGETDEFINITIONOVERLAY_NDM_REACHABILITY,
	CLOUDCRAFTWIDGETDEFINITIONOVERLAY_MONITORS,
	CLOUDCRAFTWIDGETDEFINITIONOVERLAY_APM,
	CLOUDCRAFTWIDGETDEFINITIONOVERLAY_DEFAULT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CloudcraftWidgetDefinitionOverlay) GetAllowedValues() []CloudcraftWidgetDefinitionOverlay {
	return allowedCloudcraftWidgetDefinitionOverlayEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CloudcraftWidgetDefinitionOverlay) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CloudcraftWidgetDefinitionOverlay(value)
	return nil
}

// NewCloudcraftWidgetDefinitionOverlayFromValue returns a pointer to a valid CloudcraftWidgetDefinitionOverlay
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCloudcraftWidgetDefinitionOverlayFromValue(v string) (*CloudcraftWidgetDefinitionOverlay, error) {
	ev := CloudcraftWidgetDefinitionOverlay(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CloudcraftWidgetDefinitionOverlay: valid values are %v", v, allowedCloudcraftWidgetDefinitionOverlayEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CloudcraftWidgetDefinitionOverlay) IsValid() bool {
	for _, existing := range allowedCloudcraftWidgetDefinitionOverlayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudcraftWidgetDefinitionOverlay value.
func (v CloudcraftWidgetDefinitionOverlay) Ptr() *CloudcraftWidgetDefinitionOverlay {
	return &v
}
