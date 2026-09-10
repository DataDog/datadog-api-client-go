// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FilesCoverageResponseType JSON:API type for files coverage response. The value must always be `ci_app_coverage_files`.
type FilesCoverageResponseType string

// List of FilesCoverageResponseType.
const (
	FILESCOVERAGERESPONSETYPE_CI_APP_COVERAGE_FILES FilesCoverageResponseType = "ci_app_coverage_files"
)

var allowedFilesCoverageResponseTypeEnumValues = []FilesCoverageResponseType{
	FILESCOVERAGERESPONSETYPE_CI_APP_COVERAGE_FILES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FilesCoverageResponseType) GetAllowedValues() []FilesCoverageResponseType {
	return allowedFilesCoverageResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FilesCoverageResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FilesCoverageResponseType(value)
	return nil
}

// NewFilesCoverageResponseTypeFromValue returns a pointer to a valid FilesCoverageResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFilesCoverageResponseTypeFromValue(v string) (*FilesCoverageResponseType, error) {
	ev := FilesCoverageResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FilesCoverageResponseType: valid values are %v", v, allowedFilesCoverageResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FilesCoverageResponseType) IsValid() bool {
	for _, existing := range allowedFilesCoverageResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FilesCoverageResponseType value.
func (v FilesCoverageResponseType) Ptr() *FilesCoverageResponseType {
	return &v
}
