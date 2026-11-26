// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AuditStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type AuditStreamQueryDataSource string

// List of AuditStreamQueryDataSource.
const (
	AUDITSTREAMQUERYDATASOURCE_AUDIT_STREAM AuditStreamQueryDataSource = "audit_stream"
)

var allowedAuditStreamQueryDataSourceEnumValues = []AuditStreamQueryDataSource{
	AUDITSTREAMQUERYDATASOURCE_AUDIT_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AuditStreamQueryDataSource) GetAllowedValues() []AuditStreamQueryDataSource {
	return allowedAuditStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AuditStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AuditStreamQueryDataSource(value)
	return nil
}

// NewAuditStreamQueryDataSourceFromValue returns a pointer to a valid AuditStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAuditStreamQueryDataSourceFromValue(v string) (*AuditStreamQueryDataSource, error) {
	ev := AuditStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AuditStreamQueryDataSource: valid values are %v", v, allowedAuditStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AuditStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedAuditStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuditStreamQueryDataSource value.
func (v AuditStreamQueryDataSource) Ptr() *AuditStreamQueryDataSource {
	return &v
}
