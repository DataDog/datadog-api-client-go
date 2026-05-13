// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PROutputCiStatus The aggregate CI check status for the pull request.
type PROutputCiStatus string

// List of PROutputCiStatus.
const (
	PROUTPUTCISTATUS_PENDING    PROutputCiStatus = "PENDING"
	PROUTPUTCISTATUS_FAILED     PROutputCiStatus = "FAILED"
	PROUTPUTCISTATUS_SUCCESSFUL PROutputCiStatus = "SUCCESSFUL"
)

var allowedPROutputCiStatusEnumValues = []PROutputCiStatus{
	PROUTPUTCISTATUS_PENDING,
	PROUTPUTCISTATUS_FAILED,
	PROUTPUTCISTATUS_SUCCESSFUL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PROutputCiStatus) GetAllowedValues() []PROutputCiStatus {
	return allowedPROutputCiStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PROutputCiStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PROutputCiStatus(value)
	return nil
}

// NewPROutputCiStatusFromValue returns a pointer to a valid PROutputCiStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPROutputCiStatusFromValue(v string) (*PROutputCiStatus, error) {
	ev := PROutputCiStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PROutputCiStatus: valid values are %v", v, allowedPROutputCiStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PROutputCiStatus) IsValid() bool {
	for _, existing := range allowedPROutputCiStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PROutputCiStatus value.
func (v PROutputCiStatus) Ptr() *PROutputCiStatus {
	return &v
}
