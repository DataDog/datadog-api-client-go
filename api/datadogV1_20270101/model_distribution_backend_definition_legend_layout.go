// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionBackendDefinitionLegendLayout Test-only generated schema for the dashboard versioning proof of concept.
type DistributionBackendDefinitionLegendLayout string

// List of DistributionBackendDefinitionLegendLayout.
const (
	DISTRIBUTIONBACKENDDEFINITIONLEGENDLAYOUT_VERTICAL   DistributionBackendDefinitionLegendLayout = "vertical"
	DISTRIBUTIONBACKENDDEFINITIONLEGENDLAYOUT_HORIZONTAL DistributionBackendDefinitionLegendLayout = "horizontal"
	DISTRIBUTIONBACKENDDEFINITIONLEGENDLAYOUT_AUTO       DistributionBackendDefinitionLegendLayout = "auto"
)

var allowedDistributionBackendDefinitionLegendLayoutEnumValues = []DistributionBackendDefinitionLegendLayout{
	DISTRIBUTIONBACKENDDEFINITIONLEGENDLAYOUT_VERTICAL,
	DISTRIBUTIONBACKENDDEFINITIONLEGENDLAYOUT_HORIZONTAL,
	DISTRIBUTIONBACKENDDEFINITIONLEGENDLAYOUT_AUTO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DistributionBackendDefinitionLegendLayout) GetAllowedValues() []DistributionBackendDefinitionLegendLayout {
	return allowedDistributionBackendDefinitionLegendLayoutEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DistributionBackendDefinitionLegendLayout) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DistributionBackendDefinitionLegendLayout(value)
	return nil
}

// NewDistributionBackendDefinitionLegendLayoutFromValue returns a pointer to a valid DistributionBackendDefinitionLegendLayout
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDistributionBackendDefinitionLegendLayoutFromValue(v string) (*DistributionBackendDefinitionLegendLayout, error) {
	ev := DistributionBackendDefinitionLegendLayout(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DistributionBackendDefinitionLegendLayout: valid values are %v", v, allowedDistributionBackendDefinitionLegendLayoutEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DistributionBackendDefinitionLegendLayout) IsValid() bool {
	for _, existing := range allowedDistributionBackendDefinitionLegendLayoutEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DistributionBackendDefinitionLegendLayout value.
func (v DistributionBackendDefinitionLegendLayout) Ptr() *DistributionBackendDefinitionLegendLayout {
	return &v
}
