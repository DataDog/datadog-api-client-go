// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueStreamColumnName Test-only generated schema for the dashboard versioning proof of concept.
type IssueStreamColumnName string

// List of IssueStreamColumnName.
const (
	ISSUESTREAMCOLUMNNAME_SERVICE           IssueStreamColumnName = "service"
	ISSUESTREAMCOLUMNNAME_COUNT             IssueStreamColumnName = "count"
	ISSUESTREAMCOLUMNNAME_VOLUME            IssueStreamColumnName = "volume"
	ISSUESTREAMCOLUMNNAME_RESOURCES         IssueStreamColumnName = "resources"
	ISSUESTREAMCOLUMNNAME_ERROR_TYPE        IssueStreamColumnName = "error_type"
	ISSUESTREAMCOLUMNNAME_ERROR_MESSAGE     IssueStreamColumnName = "error_message"
	ISSUESTREAMCOLUMNNAME_FILE_PATH         IssueStreamColumnName = "file_path"
	ISSUESTREAMCOLUMNNAME_LAST_SEEN         IssueStreamColumnName = "last_seen"
	ISSUESTREAMCOLUMNNAME_FIRST_SEEN        IssueStreamColumnName = "first_seen"
	ISSUESTREAMCOLUMNNAME_LANGUAGE          IssueStreamColumnName = "language"
	ISSUESTREAMCOLUMNNAME_PLATFORM          IssueStreamColumnName = "platform"
	ISSUESTREAMCOLUMNNAME_DETAILS           IssueStreamColumnName = "details"
	ISSUESTREAMCOLUMNNAME_IMPACTED_SESSIONS IssueStreamColumnName = "impacted_sessions"
	ISSUESTREAMCOLUMNNAME_IMPACTED_USERS    IssueStreamColumnName = "impacted_users"
)

var allowedIssueStreamColumnNameEnumValues = []IssueStreamColumnName{
	ISSUESTREAMCOLUMNNAME_SERVICE,
	ISSUESTREAMCOLUMNNAME_COUNT,
	ISSUESTREAMCOLUMNNAME_VOLUME,
	ISSUESTREAMCOLUMNNAME_RESOURCES,
	ISSUESTREAMCOLUMNNAME_ERROR_TYPE,
	ISSUESTREAMCOLUMNNAME_ERROR_MESSAGE,
	ISSUESTREAMCOLUMNNAME_FILE_PATH,
	ISSUESTREAMCOLUMNNAME_LAST_SEEN,
	ISSUESTREAMCOLUMNNAME_FIRST_SEEN,
	ISSUESTREAMCOLUMNNAME_LANGUAGE,
	ISSUESTREAMCOLUMNNAME_PLATFORM,
	ISSUESTREAMCOLUMNNAME_DETAILS,
	ISSUESTREAMCOLUMNNAME_IMPACTED_SESSIONS,
	ISSUESTREAMCOLUMNNAME_IMPACTED_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueStreamColumnName) GetAllowedValues() []IssueStreamColumnName {
	return allowedIssueStreamColumnNameEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueStreamColumnName) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueStreamColumnName(value)
	return nil
}

// NewIssueStreamColumnNameFromValue returns a pointer to a valid IssueStreamColumnName
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueStreamColumnNameFromValue(v string) (*IssueStreamColumnName, error) {
	ev := IssueStreamColumnName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueStreamColumnName: valid values are %v", v, allowedIssueStreamColumnNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueStreamColumnName) IsValid() bool {
	for _, existing := range allowedIssueStreamColumnNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueStreamColumnName value.
func (v IssueStreamColumnName) Ptr() *IssueStreamColumnName {
	return &v
}
