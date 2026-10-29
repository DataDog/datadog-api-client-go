// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RerouteOnCallPageRequestDataType The type of resource used when rerouting an On-Call Page.
type RerouteOnCallPageRequestDataType string

// List of RerouteOnCallPageRequestDataType.
const (
	REROUTEONCALLPAGEREQUESTDATATYPE_REROUTE RerouteOnCallPageRequestDataType = "reroute"
)

var allowedRerouteOnCallPageRequestDataTypeEnumValues = []RerouteOnCallPageRequestDataType{
	REROUTEONCALLPAGEREQUESTDATATYPE_REROUTE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RerouteOnCallPageRequestDataType) GetAllowedValues() []RerouteOnCallPageRequestDataType {
	return allowedRerouteOnCallPageRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RerouteOnCallPageRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RerouteOnCallPageRequestDataType(value)
	return nil
}

// NewRerouteOnCallPageRequestDataTypeFromValue returns a pointer to a valid RerouteOnCallPageRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRerouteOnCallPageRequestDataTypeFromValue(v string) (*RerouteOnCallPageRequestDataType, error) {
	ev := RerouteOnCallPageRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RerouteOnCallPageRequestDataType: valid values are %v", v, allowedRerouteOnCallPageRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RerouteOnCallPageRequestDataType) IsValid() bool {
	for _, existing := range allowedRerouteOnCallPageRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RerouteOnCallPageRequestDataType value.
func (v RerouteOnCallPageRequestDataType) Ptr() *RerouteOnCallPageRequestDataType {
	return &v
}
