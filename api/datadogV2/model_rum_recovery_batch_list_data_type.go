// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoveryBatchListDataType JSON:API resource type.
type RumRecoveryBatchListDataType string

// List of RumRecoveryBatchListDataType.
const (
	RUMRECOVERYBATCHLISTDATATYPE_RUM_BATCH_RECOVERY_REQUEST RumRecoveryBatchListDataType = "rum_batch_recovery_request"
)

var allowedRumRecoveryBatchListDataTypeEnumValues = []RumRecoveryBatchListDataType{
	RUMRECOVERYBATCHLISTDATATYPE_RUM_BATCH_RECOVERY_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRecoveryBatchListDataType) GetAllowedValues() []RumRecoveryBatchListDataType {
	return allowedRumRecoveryBatchListDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRecoveryBatchListDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRecoveryBatchListDataType(value)
	return nil
}

// NewRumRecoveryBatchListDataTypeFromValue returns a pointer to a valid RumRecoveryBatchListDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRecoveryBatchListDataTypeFromValue(v string) (*RumRecoveryBatchListDataType, error) {
	ev := RumRecoveryBatchListDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRecoveryBatchListDataType: valid values are %v", v, allowedRumRecoveryBatchListDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRecoveryBatchListDataType) IsValid() bool {
	for _, existing := range allowedRumRecoveryBatchListDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRecoveryBatchListDataType value.
func (v RumRecoveryBatchListDataType) Ptr() *RumRecoveryBatchListDataType {
	return &v
}
