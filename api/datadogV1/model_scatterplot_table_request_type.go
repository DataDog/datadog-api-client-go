// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotTableRequestType The type of the scatterplot table request.
type ScatterplotTableRequestType string

// List of ScatterplotTableRequestType.
const (
	SCATTERPLOTTABLEREQUESTTYPE_TABLE           ScatterplotTableRequestType = "table"
	SCATTERPLOTTABLEREQUESTTYPE_DATA_PROJECTION ScatterplotTableRequestType = "data_projection"
)

var allowedScatterplotTableRequestTypeEnumValues = []ScatterplotTableRequestType{
	SCATTERPLOTTABLEREQUESTTYPE_TABLE,
	SCATTERPLOTTABLEREQUESTTYPE_DATA_PROJECTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScatterplotTableRequestType) GetAllowedValues() []ScatterplotTableRequestType {
	return allowedScatterplotTableRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScatterplotTableRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScatterplotTableRequestType(value)
	return nil
}

// NewScatterplotTableRequestTypeFromValue returns a pointer to a valid ScatterplotTableRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScatterplotTableRequestTypeFromValue(v string) (*ScatterplotTableRequestType, error) {
	ev := ScatterplotTableRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScatterplotTableRequestType: valid values are %v", v, allowedScatterplotTableRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScatterplotTableRequestType) IsValid() bool {
	for _, existing := range allowedScatterplotTableRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScatterplotTableRequestType value.
func (v ScatterplotTableRequestType) Ptr() *ScatterplotTableRequestType {
	return &v
}
