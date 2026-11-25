// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPProjectFilterConfigType The GCP project field that the filters are matched against.
type GCPProjectFilterConfigType string

// List of GCPProjectFilterConfigType.
const (
	GCPPROJECTFILTERCONFIGTYPE_PROJECT_ID    GCPProjectFilterConfigType = "project_id"
	GCPPROJECTFILTERCONFIGTYPE_PROJECT_LABEL GCPProjectFilterConfigType = "project_label"
)

var allowedGCPProjectFilterConfigTypeEnumValues = []GCPProjectFilterConfigType{
	GCPPROJECTFILTERCONFIGTYPE_PROJECT_ID,
	GCPPROJECTFILTERCONFIGTYPE_PROJECT_LABEL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GCPProjectFilterConfigType) GetAllowedValues() []GCPProjectFilterConfigType {
	return allowedGCPProjectFilterConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GCPProjectFilterConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GCPProjectFilterConfigType(value)
	return nil
}

// NewGCPProjectFilterConfigTypeFromValue returns a pointer to a valid GCPProjectFilterConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGCPProjectFilterConfigTypeFromValue(v string) (*GCPProjectFilterConfigType, error) {
	ev := GCPProjectFilterConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GCPProjectFilterConfigType: valid values are %v", v, allowedGCPProjectFilterConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GCPProjectFilterConfigType) IsValid() bool {
	for _, existing := range allowedGCPProjectFilterConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GCPProjectFilterConfigType value.
func (v GCPProjectFilterConfigType) Ptr() *GCPProjectFilterConfigType {
	return &v
}
