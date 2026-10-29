// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateStatusPageRequestDataAttributesVisualizationLayout The layout of components on the status page. Defaults to `list`. `grid` requires `visualization_type` to be `component_name_only`; changing `visualization_type` to any other value resets the layout to `list`.
type CreateStatusPageRequestDataAttributesVisualizationLayout string

// List of CreateStatusPageRequestDataAttributesVisualizationLayout.
const (
	CREATESTATUSPAGEREQUESTDATAATTRIBUTESVISUALIZATIONLAYOUT_LIST CreateStatusPageRequestDataAttributesVisualizationLayout = "list"
	CREATESTATUSPAGEREQUESTDATAATTRIBUTESVISUALIZATIONLAYOUT_GRID CreateStatusPageRequestDataAttributesVisualizationLayout = "grid"
)

var allowedCreateStatusPageRequestDataAttributesVisualizationLayoutEnumValues = []CreateStatusPageRequestDataAttributesVisualizationLayout{
	CREATESTATUSPAGEREQUESTDATAATTRIBUTESVISUALIZATIONLAYOUT_LIST,
	CREATESTATUSPAGEREQUESTDATAATTRIBUTESVISUALIZATIONLAYOUT_GRID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateStatusPageRequestDataAttributesVisualizationLayout) GetAllowedValues() []CreateStatusPageRequestDataAttributesVisualizationLayout {
	return allowedCreateStatusPageRequestDataAttributesVisualizationLayoutEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateStatusPageRequestDataAttributesVisualizationLayout) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateStatusPageRequestDataAttributesVisualizationLayout(value)
	return nil
}

// NewCreateStatusPageRequestDataAttributesVisualizationLayoutFromValue returns a pointer to a valid CreateStatusPageRequestDataAttributesVisualizationLayout
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateStatusPageRequestDataAttributesVisualizationLayoutFromValue(v string) (*CreateStatusPageRequestDataAttributesVisualizationLayout, error) {
	ev := CreateStatusPageRequestDataAttributesVisualizationLayout(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateStatusPageRequestDataAttributesVisualizationLayout: valid values are %v", v, allowedCreateStatusPageRequestDataAttributesVisualizationLayoutEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateStatusPageRequestDataAttributesVisualizationLayout) IsValid() bool {
	for _, existing := range allowedCreateStatusPageRequestDataAttributesVisualizationLayoutEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateStatusPageRequestDataAttributesVisualizationLayout value.
func (v CreateStatusPageRequestDataAttributesVisualizationLayout) Ptr() *CreateStatusPageRequestDataAttributesVisualizationLayout {
	return &v
}
