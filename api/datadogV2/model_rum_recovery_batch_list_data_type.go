// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRecoveryBatchListDataType JSON:API resource type.
type RUMRecoveryBatchListDataType string

// List of RUMRecoveryBatchListDataType.
const (
	RUMRECOVERYBATCHLISTDATATYPE_RUM_BATCH_RECOVERY_REQUEST RUMRecoveryBatchListDataType = "rum_batch_recovery_request"
)

var allowedRUMRecoveryBatchListDataTypeEnumValues = []RUMRecoveryBatchListDataType{
	RUMRECOVERYBATCHLISTDATATYPE_RUM_BATCH_RECOVERY_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMRecoveryBatchListDataType) GetAllowedValues() []RUMRecoveryBatchListDataType {
	return allowedRUMRecoveryBatchListDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMRecoveryBatchListDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMRecoveryBatchListDataType(value)
	return nil
}

// NewRUMRecoveryBatchListDataTypeFromValue returns a pointer to a valid RUMRecoveryBatchListDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMRecoveryBatchListDataTypeFromValue(v string) (*RUMRecoveryBatchListDataType, error) {
	ev := RUMRecoveryBatchListDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMRecoveryBatchListDataType: valid values are %v", v, allowedRUMRecoveryBatchListDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMRecoveryBatchListDataType) IsValid() bool {
	for _, existing := range allowedRUMRecoveryBatchListDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMRecoveryBatchListDataType value.
func (v RUMRecoveryBatchListDataType) Ptr() *RUMRecoveryBatchListDataType {
	return &v
}
