// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagDistributionChannel The distribution channel for the feature flag.
type FeatureFlagDistributionChannel string

// List of FeatureFlagDistributionChannel.
const (
	FEATUREFLAGDISTRIBUTIONCHANNEL_ALL    FeatureFlagDistributionChannel = "ALL"
	FEATUREFLAGDISTRIBUTIONCHANNEL_CLIENT FeatureFlagDistributionChannel = "CLIENT"
	FEATUREFLAGDISTRIBUTIONCHANNEL_SERVER FeatureFlagDistributionChannel = "SERVER"
)

var allowedFeatureFlagDistributionChannelEnumValues = []FeatureFlagDistributionChannel{
	FEATUREFLAGDISTRIBUTIONCHANNEL_ALL,
	FEATUREFLAGDISTRIBUTIONCHANNEL_CLIENT,
	FEATUREFLAGDISTRIBUTIONCHANNEL_SERVER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FeatureFlagDistributionChannel) GetAllowedValues() []FeatureFlagDistributionChannel {
	return allowedFeatureFlagDistributionChannelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FeatureFlagDistributionChannel) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FeatureFlagDistributionChannel(value)
	return nil
}

// NewFeatureFlagDistributionChannelFromValue returns a pointer to a valid FeatureFlagDistributionChannel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFeatureFlagDistributionChannelFromValue(v string) (*FeatureFlagDistributionChannel, error) {
	ev := FeatureFlagDistributionChannel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FeatureFlagDistributionChannel: valid values are %v", v, allowedFeatureFlagDistributionChannelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FeatureFlagDistributionChannel) IsValid() bool {
	for _, existing := range allowedFeatureFlagDistributionChannelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureFlagDistributionChannel value.
func (v FeatureFlagDistributionChannel) Ptr() *FeatureFlagDistributionChannel {
	return &v
}
