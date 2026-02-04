// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormSubmissionType Type for form submissions.
type FormSubmissionType string

// List of FormSubmissionType.
const (
	FORMSUBMISSIONTYPE_FORM_SUBMISSIONS FormSubmissionType = "form_submissions"
)

var allowedFormSubmissionTypeEnumValues = []FormSubmissionType{
	FORMSUBMISSIONTYPE_FORM_SUBMISSIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormSubmissionType) GetAllowedValues() []FormSubmissionType {
	return allowedFormSubmissionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormSubmissionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormSubmissionType(value)
	return nil
}

// NewFormSubmissionTypeFromValue returns a pointer to a valid FormSubmissionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormSubmissionTypeFromValue(v string) (*FormSubmissionType, error) {
	ev := FormSubmissionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormSubmissionType: valid values are %v", v, allowedFormSubmissionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormSubmissionType) IsValid() bool {
	for _, existing := range allowedFormSubmissionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormSubmissionType value.
func (v FormSubmissionType) Ptr() *FormSubmissionType {
	return &v
}
