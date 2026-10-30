// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagChangeEventProperty The part of the feature flag the change applies to.
type FeatureFlagChangeEventProperty string

// List of FeatureFlagChangeEventProperty.
const (
	FEATUREFLAGCHANGEEVENTPROPERTY_FLAG                 FeatureFlagChangeEventProperty = "FLAG"
	FEATUREFLAGCHANGEEVENTPROPERTY_FLAG_NAME            FeatureFlagChangeEventProperty = "FLAG_NAME"
	FEATUREFLAGCHANGEEVENTPROPERTY_FLAG_DESCRIPTION     FeatureFlagChangeEventProperty = "FLAG_DESCRIPTION"
	FEATUREFLAGCHANGEEVENTPROPERTY_JSON_SCHEMA          FeatureFlagChangeEventProperty = "JSON_SCHEMA"
	FEATUREFLAGCHANGEEVENTPROPERTY_DISTRIBUTION_CHANNEL FeatureFlagChangeEventProperty = "DISTRIBUTION_CHANNEL"
	FEATUREFLAGCHANGEEVENTPROPERTY_REQUIRE_APPROVAL     FeatureFlagChangeEventProperty = "REQUIRE_APPROVAL"
	FEATUREFLAGCHANGEEVENTPROPERTY_VARIANT              FeatureFlagChangeEventProperty = "VARIANT"
	FEATUREFLAGCHANGEEVENTPROPERTY_VARIANT_NAME         FeatureFlagChangeEventProperty = "VARIANT_NAME"
	FEATUREFLAGCHANGEEVENTPROPERTY_VARIANT_VALUE        FeatureFlagChangeEventProperty = "VARIANT_VALUE"
	FEATUREFLAGCHANGEEVENTPROPERTY_ALLOCATIONS          FeatureFlagChangeEventProperty = "ALLOCATIONS"
	FEATUREFLAGCHANGEEVENTPROPERTY_ROLLOUT              FeatureFlagChangeEventProperty = "ROLLOUT"
	FEATUREFLAGCHANGEEVENTPROPERTY_ENVIRONMENT_STATUS   FeatureFlagChangeEventProperty = "ENVIRONMENT_STATUS"
	FEATUREFLAGCHANGEEVENTPROPERTY_DEFAULT_VARIANT      FeatureFlagChangeEventProperty = "DEFAULT_VARIANT"
	FEATUREFLAGCHANGEEVENTPROPERTY_OVERRIDE_VARIANT     FeatureFlagChangeEventProperty = "OVERRIDE_VARIANT"
	FEATUREFLAGCHANGEEVENTPROPERTY_STALENESS_STATUS     FeatureFlagChangeEventProperty = "STALENESS_STATUS"
	FEATUREFLAGCHANGEEVENTPROPERTY_SAVED_FILTER         FeatureFlagChangeEventProperty = "SAVED_FILTER"
)

var allowedFeatureFlagChangeEventPropertyEnumValues = []FeatureFlagChangeEventProperty{
	FEATUREFLAGCHANGEEVENTPROPERTY_FLAG,
	FEATUREFLAGCHANGEEVENTPROPERTY_FLAG_NAME,
	FEATUREFLAGCHANGEEVENTPROPERTY_FLAG_DESCRIPTION,
	FEATUREFLAGCHANGEEVENTPROPERTY_JSON_SCHEMA,
	FEATUREFLAGCHANGEEVENTPROPERTY_DISTRIBUTION_CHANNEL,
	FEATUREFLAGCHANGEEVENTPROPERTY_REQUIRE_APPROVAL,
	FEATUREFLAGCHANGEEVENTPROPERTY_VARIANT,
	FEATUREFLAGCHANGEEVENTPROPERTY_VARIANT_NAME,
	FEATUREFLAGCHANGEEVENTPROPERTY_VARIANT_VALUE,
	FEATUREFLAGCHANGEEVENTPROPERTY_ALLOCATIONS,
	FEATUREFLAGCHANGEEVENTPROPERTY_ROLLOUT,
	FEATUREFLAGCHANGEEVENTPROPERTY_ENVIRONMENT_STATUS,
	FEATUREFLAGCHANGEEVENTPROPERTY_DEFAULT_VARIANT,
	FEATUREFLAGCHANGEEVENTPROPERTY_OVERRIDE_VARIANT,
	FEATUREFLAGCHANGEEVENTPROPERTY_STALENESS_STATUS,
	FEATUREFLAGCHANGEEVENTPROPERTY_SAVED_FILTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FeatureFlagChangeEventProperty) GetAllowedValues() []FeatureFlagChangeEventProperty {
	return allowedFeatureFlagChangeEventPropertyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FeatureFlagChangeEventProperty) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FeatureFlagChangeEventProperty(value)
	return nil
}

// NewFeatureFlagChangeEventPropertyFromValue returns a pointer to a valid FeatureFlagChangeEventProperty
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFeatureFlagChangeEventPropertyFromValue(v string) (*FeatureFlagChangeEventProperty, error) {
	ev := FeatureFlagChangeEventProperty(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FeatureFlagChangeEventProperty: valid values are %v", v, allowedFeatureFlagChangeEventPropertyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FeatureFlagChangeEventProperty) IsValid() bool {
	for _, existing := range allowedFeatureFlagChangeEventPropertyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureFlagChangeEventProperty value.
func (v FeatureFlagChangeEventProperty) Ptr() *FeatureFlagChangeEventProperty {
	return &v
}
