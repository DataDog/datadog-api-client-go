// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpanKind Test-only generated schema for the dashboard versioning proof of concept.
type SpanKind string

// List of SpanKind.
const (
	SPANKIND_CONSUMER SpanKind = "consumer"
	SPANKIND_SERVER   SpanKind = "server"
	SPANKIND_CLIENT   SpanKind = "client"
	SPANKIND_PRODUCER SpanKind = "producer"
	SPANKIND_INTERNAL SpanKind = "internal"
)

var allowedSpanKindEnumValues = []SpanKind{
	SPANKIND_CONSUMER,
	SPANKIND_SERVER,
	SPANKIND_CLIENT,
	SPANKIND_PRODUCER,
	SPANKIND_INTERNAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SpanKind) GetAllowedValues() []SpanKind {
	return allowedSpanKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SpanKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SpanKind(value)
	return nil
}

// NewSpanKindFromValue returns a pointer to a valid SpanKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSpanKindFromValue(v string) (*SpanKind, error) {
	ev := SpanKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SpanKind: valid values are %v", v, allowedSpanKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SpanKind) IsValid() bool {
	for _, existing := range allowedSpanKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SpanKind value.
func (v SpanKind) Ptr() *SpanKind {
	return &v
}
