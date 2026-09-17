// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PRCoverageSummaryRequestType JSON:API type for PR coverage summary request. The value must always be `ci_app_coverage_pr_summary_request`.
type PRCoverageSummaryRequestType string

// List of PRCoverageSummaryRequestType.
const (
	PRCOVERAGESUMMARYREQUESTTYPE_CI_APP_COVERAGE_PR_SUMMARY_REQUEST PRCoverageSummaryRequestType = "ci_app_coverage_pr_summary_request"
)

var allowedPRCoverageSummaryRequestTypeEnumValues = []PRCoverageSummaryRequestType{
	PRCOVERAGESUMMARYREQUESTTYPE_CI_APP_COVERAGE_PR_SUMMARY_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PRCoverageSummaryRequestType) GetAllowedValues() []PRCoverageSummaryRequestType {
	return allowedPRCoverageSummaryRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PRCoverageSummaryRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PRCoverageSummaryRequestType(value)
	return nil
}

// NewPRCoverageSummaryRequestTypeFromValue returns a pointer to a valid PRCoverageSummaryRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPRCoverageSummaryRequestTypeFromValue(v string) (*PRCoverageSummaryRequestType, error) {
	ev := PRCoverageSummaryRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PRCoverageSummaryRequestType: valid values are %v", v, allowedPRCoverageSummaryRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PRCoverageSummaryRequestType) IsValid() bool {
	for _, existing := range allowedPRCoverageSummaryRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PRCoverageSummaryRequestType value.
func (v PRCoverageSummaryRequestType) Ptr() *PRCoverageSummaryRequestType {
	return &v
}
