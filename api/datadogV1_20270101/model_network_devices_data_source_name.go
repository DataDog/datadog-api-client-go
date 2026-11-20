// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NetworkDevicesDataSourceName Test-only generated schema for the dashboard versioning proof of concept.
type NetworkDevicesDataSourceName string

// List of NetworkDevicesDataSourceName.
const (
	NETWORKDEVICESDATASOURCENAME_NETWORK_DEVICES NetworkDevicesDataSourceName = "network_devices"
)

var allowedNetworkDevicesDataSourceNameEnumValues = []NetworkDevicesDataSourceName{
	NETWORKDEVICESDATASOURCENAME_NETWORK_DEVICES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NetworkDevicesDataSourceName) GetAllowedValues() []NetworkDevicesDataSourceName {
	return allowedNetworkDevicesDataSourceNameEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NetworkDevicesDataSourceName) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NetworkDevicesDataSourceName(value)
	return nil
}

// NewNetworkDevicesDataSourceNameFromValue returns a pointer to a valid NetworkDevicesDataSourceName
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNetworkDevicesDataSourceNameFromValue(v string) (*NetworkDevicesDataSourceName, error) {
	ev := NetworkDevicesDataSourceName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NetworkDevicesDataSourceName: valid values are %v", v, allowedNetworkDevicesDataSourceNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NetworkDevicesDataSourceName) IsValid() bool {
	for _, existing := range allowedNetworkDevicesDataSourceNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkDevicesDataSourceName value.
func (v NetworkDevicesDataSourceName) Ptr() *NetworkDevicesDataSourceName {
	return &v
}
