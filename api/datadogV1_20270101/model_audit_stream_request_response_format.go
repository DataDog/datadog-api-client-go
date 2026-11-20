// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AuditStreamRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type AuditStreamRequestResponseFormat string

// List of AuditStreamRequestResponseFormat.
const (
	AUDITSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST AuditStreamRequestResponseFormat = "event_list"
)

var allowedAuditStreamRequestResponseFormatEnumValues = []AuditStreamRequestResponseFormat{
	AUDITSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AuditStreamRequestResponseFormat) GetAllowedValues() []AuditStreamRequestResponseFormat {
	return allowedAuditStreamRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AuditStreamRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AuditStreamRequestResponseFormat(value)
	return nil
}

// NewAuditStreamRequestResponseFormatFromValue returns a pointer to a valid AuditStreamRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAuditStreamRequestResponseFormatFromValue(v string) (*AuditStreamRequestResponseFormat, error) {
	ev := AuditStreamRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AuditStreamRequestResponseFormat: valid values are %v", v, allowedAuditStreamRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AuditStreamRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedAuditStreamRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuditStreamRequestResponseFormat value.
func (v AuditStreamRequestResponseFormat) Ptr() *AuditStreamRequestResponseFormat {
	return &v
}
