// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagBucketingMode How subjects are assigned to variants. `DATADOG` is the native scheme. `EPPO` reproduces the assignments that Eppo made for a flag migrated from Eppo, so its subjects keep their variants. Set when the flag is created and never changed.
type FeatureFlagBucketingMode string

// List of FeatureFlagBucketingMode.
const (
	FEATUREFLAGBUCKETINGMODE_DATADOG FeatureFlagBucketingMode = "DATADOG"
	FEATUREFLAGBUCKETINGMODE_EPPO    FeatureFlagBucketingMode = "EPPO"
)

var allowedFeatureFlagBucketingModeEnumValues = []FeatureFlagBucketingMode{
	FEATUREFLAGBUCKETINGMODE_DATADOG,
	FEATUREFLAGBUCKETINGMODE_EPPO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FeatureFlagBucketingMode) GetAllowedValues() []FeatureFlagBucketingMode {
	return allowedFeatureFlagBucketingModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FeatureFlagBucketingMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FeatureFlagBucketingMode(value)
	return nil
}

// NewFeatureFlagBucketingModeFromValue returns a pointer to a valid FeatureFlagBucketingMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFeatureFlagBucketingModeFromValue(v string) (*FeatureFlagBucketingMode, error) {
	ev := FeatureFlagBucketingMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FeatureFlagBucketingMode: valid values are %v", v, allowedFeatureFlagBucketingModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FeatureFlagBucketingMode) IsValid() bool {
	for _, existing := range allowedFeatureFlagBucketingModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureFlagBucketingMode value.
func (v FeatureFlagBucketingMode) Ptr() *FeatureFlagBucketingMode {
	return &v
}
