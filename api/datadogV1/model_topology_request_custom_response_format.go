// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyRequestCustomResponseFormat Response format of a custom topology map request.
type TopologyRequestCustomResponseFormat string

// List of TopologyRequestCustomResponseFormat.
const (
	TOPOLOGYREQUESTCUSTOMRESPONSEFORMAT_SCALAR TopologyRequestCustomResponseFormat = "scalar"
)

var allowedTopologyRequestCustomResponseFormatEnumValues = []TopologyRequestCustomResponseFormat{
	TOPOLOGYREQUESTCUSTOMRESPONSEFORMAT_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TopologyRequestCustomResponseFormat) GetAllowedValues() []TopologyRequestCustomResponseFormat {
	return allowedTopologyRequestCustomResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TopologyRequestCustomResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TopologyRequestCustomResponseFormat(value)
	return nil
}

// NewTopologyRequestCustomResponseFormatFromValue returns a pointer to a valid TopologyRequestCustomResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTopologyRequestCustomResponseFormatFromValue(v string) (*TopologyRequestCustomResponseFormat, error) {
	ev := TopologyRequestCustomResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TopologyRequestCustomResponseFormat: valid values are %v", v, allowedTopologyRequestCustomResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TopologyRequestCustomResponseFormat) IsValid() bool {
	for _, existing := range allowedTopologyRequestCustomResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TopologyRequestCustomResponseFormat value.
func (v TopologyRequestCustomResponseFormat) Ptr() *TopologyRequestCustomResponseFormat {
	return &v
}
