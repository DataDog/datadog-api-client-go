// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoveryBatchRequestType Query-based batch recovery.
type RumRecoveryBatchRequestType string

// List of RumRecoveryBatchRequestType.
const (
	RUMRECOVERYBATCHREQUESTTYPE_BATCH RumRecoveryBatchRequestType = "batch"
)

var allowedRumRecoveryBatchRequestTypeEnumValues = []RumRecoveryBatchRequestType{
	RUMRECOVERYBATCHREQUESTTYPE_BATCH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRecoveryBatchRequestType) GetAllowedValues() []RumRecoveryBatchRequestType {
	return allowedRumRecoveryBatchRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRecoveryBatchRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRecoveryBatchRequestType(value)
	return nil
}

// NewRumRecoveryBatchRequestTypeFromValue returns a pointer to a valid RumRecoveryBatchRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRecoveryBatchRequestTypeFromValue(v string) (*RumRecoveryBatchRequestType, error) {
	ev := RumRecoveryBatchRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRecoveryBatchRequestType: valid values are %v", v, allowedRumRecoveryBatchRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRecoveryBatchRequestType) IsValid() bool {
	for _, existing := range allowedRumRecoveryBatchRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRecoveryBatchRequestType value.
func (v RumRecoveryBatchRequestType) Ptr() *RumRecoveryBatchRequestType {
	return &v
}
