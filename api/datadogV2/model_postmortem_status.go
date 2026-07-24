// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PostmortemStatus The status of the postmortem.
type PostmortemStatus string

// List of PostmortemStatus.
const (
	POSTMORTEMSTATUS_DRAFT     PostmortemStatus = "draft"
	POSTMORTEMSTATUS_IN_REVIEW PostmortemStatus = "in_review"
	POSTMORTEMSTATUS_COMPLETED PostmortemStatus = "completed"
)

var allowedPostmortemStatusEnumValues = []PostmortemStatus{
	POSTMORTEMSTATUS_DRAFT,
	POSTMORTEMSTATUS_IN_REVIEW,
	POSTMORTEMSTATUS_COMPLETED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PostmortemStatus) GetAllowedValues() []PostmortemStatus {
	return allowedPostmortemStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PostmortemStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PostmortemStatus(value)
	return nil
}

// NewPostmortemStatusFromValue returns a pointer to a valid PostmortemStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPostmortemStatusFromValue(v string) (*PostmortemStatus, error) {
	ev := PostmortemStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PostmortemStatus: valid values are %v", v, allowedPostmortemStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PostmortemStatus) IsValid() bool {
	for _, existing := range allowedPostmortemStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PostmortemStatus value.
func (v PostmortemStatus) Ptr() *PostmortemStatus {
	return &v
}
