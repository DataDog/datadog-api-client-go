// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsStopBanditV2RequestDataType Stop bandit request resource type.
type ExperimentsStopBanditV2RequestDataType string

// List of ExperimentsStopBanditV2RequestDataType.
const (
	EXPERIMENTSSTOPBANDITV2REQUESTDATATYPE_STOP_BANDIT_REQUEST ExperimentsStopBanditV2RequestDataType = "stop-bandit-request"
)

var allowedExperimentsStopBanditV2RequestDataTypeEnumValues = []ExperimentsStopBanditV2RequestDataType{
	EXPERIMENTSSTOPBANDITV2REQUESTDATATYPE_STOP_BANDIT_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsStopBanditV2RequestDataType) GetAllowedValues() []ExperimentsStopBanditV2RequestDataType {
	return allowedExperimentsStopBanditV2RequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsStopBanditV2RequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsStopBanditV2RequestDataType(value)
	return nil
}

// NewExperimentsStopBanditV2RequestDataTypeFromValue returns a pointer to a valid ExperimentsStopBanditV2RequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsStopBanditV2RequestDataTypeFromValue(v string) (*ExperimentsStopBanditV2RequestDataType, error) {
	ev := ExperimentsStopBanditV2RequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsStopBanditV2RequestDataType: valid values are %v", v, allowedExperimentsStopBanditV2RequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsStopBanditV2RequestDataType) IsValid() bool {
	for _, existing := range allowedExperimentsStopBanditV2RequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsStopBanditV2RequestDataType value.
func (v ExperimentsStopBanditV2RequestDataType) Ptr() *ExperimentsStopBanditV2RequestDataType {
	return &v
}
