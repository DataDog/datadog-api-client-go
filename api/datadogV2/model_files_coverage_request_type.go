// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FilesCoverageRequestType JSON:API type for files coverage request. The value must always be `ci_app_coverage_files_request`.
type FilesCoverageRequestType string

// List of FilesCoverageRequestType.
const (
	FILESCOVERAGEREQUESTTYPE_CI_APP_COVERAGE_FILES_REQUEST FilesCoverageRequestType = "ci_app_coverage_files_request"
)

var allowedFilesCoverageRequestTypeEnumValues = []FilesCoverageRequestType{
	FILESCOVERAGEREQUESTTYPE_CI_APP_COVERAGE_FILES_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FilesCoverageRequestType) GetAllowedValues() []FilesCoverageRequestType {
	return allowedFilesCoverageRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FilesCoverageRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FilesCoverageRequestType(value)
	return nil
}

// NewFilesCoverageRequestTypeFromValue returns a pointer to a valid FilesCoverageRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFilesCoverageRequestTypeFromValue(v string) (*FilesCoverageRequestType, error) {
	ev := FilesCoverageRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FilesCoverageRequestType: valid values are %v", v, allowedFilesCoverageRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FilesCoverageRequestType) IsValid() bool {
	for _, existing := range allowedFilesCoverageRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FilesCoverageRequestType value.
func (v FilesCoverageRequestType) Ptr() *FilesCoverageRequestType {
	return &v
}
