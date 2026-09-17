// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateFeatureFlagStalenessStatus The staleness status for the feature flag at creation.
type CreateFeatureFlagStalenessStatus string

// List of CreateFeatureFlagStalenessStatus.
const (
	CREATEFEATUREFLAGSTALENESSSTATUS_ACTIVE    CreateFeatureFlagStalenessStatus = "ACTIVE"
	CREATEFEATUREFLAGSTALENESSSTATUS_PERMANENT CreateFeatureFlagStalenessStatus = "PERMANENT"
)

var allowedCreateFeatureFlagStalenessStatusEnumValues = []CreateFeatureFlagStalenessStatus{
	CREATEFEATUREFLAGSTALENESSSTATUS_ACTIVE,
	CREATEFEATUREFLAGSTALENESSSTATUS_PERMANENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateFeatureFlagStalenessStatus) GetAllowedValues() []CreateFeatureFlagStalenessStatus {
	return allowedCreateFeatureFlagStalenessStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateFeatureFlagStalenessStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateFeatureFlagStalenessStatus(value)
	return nil
}

// NewCreateFeatureFlagStalenessStatusFromValue returns a pointer to a valid CreateFeatureFlagStalenessStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateFeatureFlagStalenessStatusFromValue(v string) (*CreateFeatureFlagStalenessStatus, error) {
	ev := CreateFeatureFlagStalenessStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateFeatureFlagStalenessStatus: valid values are %v", v, allowedCreateFeatureFlagStalenessStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateFeatureFlagStalenessStatus) IsValid() bool {
	for _, existing := range allowedCreateFeatureFlagStalenessStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateFeatureFlagStalenessStatus value.
func (v CreateFeatureFlagStalenessStatus) Ptr() *CreateFeatureFlagStalenessStatus {
	return &v
}
