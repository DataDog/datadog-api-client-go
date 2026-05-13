// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PROutputStatus The current status of the pull request.
type PROutputStatus string

// List of PROutputStatus.
const (
	PROUTPUTSTATUS_PENDING PROutputStatus = "PENDING"
	PROUTPUTSTATUS_DRAFT   PROutputStatus = "DRAFT"
	PROUTPUTSTATUS_READY   PROutputStatus = "READY"
	PROUTPUTSTATUS_MERGED  PROutputStatus = "MERGED"
	PROUTPUTSTATUS_CLOSED  PROutputStatus = "CLOSED"
)

var allowedPROutputStatusEnumValues = []PROutputStatus{
	PROUTPUTSTATUS_PENDING,
	PROUTPUTSTATUS_DRAFT,
	PROUTPUTSTATUS_READY,
	PROUTPUTSTATUS_MERGED,
	PROUTPUTSTATUS_CLOSED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PROutputStatus) GetAllowedValues() []PROutputStatus {
	return allowedPROutputStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PROutputStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PROutputStatus(value)
	return nil
}

// NewPROutputStatusFromValue returns a pointer to a valid PROutputStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPROutputStatusFromValue(v string) (*PROutputStatus, error) {
	ev := PROutputStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PROutputStatus: valid values are %v", v, allowedPROutputStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PROutputStatus) IsValid() bool {
	for _, existing := range allowedPROutputStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PROutputStatus value.
func (v PROutputStatus) Ptr() *PROutputStatus {
	return &v
}
