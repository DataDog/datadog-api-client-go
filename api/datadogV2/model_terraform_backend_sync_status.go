// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TerraformBackendSyncStatus Most recent synchronization outcome, or pending if no outcome has been recorded.
type TerraformBackendSyncStatus string

// List of TerraformBackendSyncStatus.
const (
	TERRAFORMBACKENDSYNCSTATUS_PENDING TerraformBackendSyncStatus = "pending"
	TERRAFORMBACKENDSYNCSTATUS_SUCCESS TerraformBackendSyncStatus = "success"
	TERRAFORMBACKENDSYNCSTATUS_FAILURE TerraformBackendSyncStatus = "failure"
)

var allowedTerraformBackendSyncStatusEnumValues = []TerraformBackendSyncStatus{
	TERRAFORMBACKENDSYNCSTATUS_PENDING,
	TERRAFORMBACKENDSYNCSTATUS_SUCCESS,
	TERRAFORMBACKENDSYNCSTATUS_FAILURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TerraformBackendSyncStatus) GetAllowedValues() []TerraformBackendSyncStatus {
	return allowedTerraformBackendSyncStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TerraformBackendSyncStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TerraformBackendSyncStatus(value)
	return nil
}

// NewTerraformBackendSyncStatusFromValue returns a pointer to a valid TerraformBackendSyncStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTerraformBackendSyncStatusFromValue(v string) (*TerraformBackendSyncStatus, error) {
	ev := TerraformBackendSyncStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TerraformBackendSyncStatus: valid values are %v", v, allowedTerraformBackendSyncStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TerraformBackendSyncStatus) IsValid() bool {
	for _, existing := range allowedTerraformBackendSyncStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TerraformBackendSyncStatus value.
func (v TerraformBackendSyncStatus) Ptr() *TerraformBackendSyncStatus {
	return &v
}
