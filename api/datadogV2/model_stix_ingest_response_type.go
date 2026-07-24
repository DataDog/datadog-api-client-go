// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// STIXIngestResponseType The STIX ingestion resource type.
type STIXIngestResponseType string

// List of STIXIngestResponseType.
const (
	STIXINGESTRESPONSETYPE_THREAT_INTEL_STIX_INGEST STIXIngestResponseType = "threat-intel-stix-ingest"
)

var allowedSTIXIngestResponseTypeEnumValues = []STIXIngestResponseType{
	STIXINGESTRESPONSETYPE_THREAT_INTEL_STIX_INGEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *STIXIngestResponseType) GetAllowedValues() []STIXIngestResponseType {
	return allowedSTIXIngestResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *STIXIngestResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = STIXIngestResponseType(value)
	return nil
}

// NewSTIXIngestResponseTypeFromValue returns a pointer to a valid STIXIngestResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSTIXIngestResponseTypeFromValue(v string) (*STIXIngestResponseType, error) {
	ev := STIXIngestResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for STIXIngestResponseType: valid values are %v", v, allowedSTIXIngestResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v STIXIngestResponseType) IsValid() bool {
	for _, existing := range allowedSTIXIngestResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to STIXIngestResponseType value.
func (v STIXIngestResponseType) Ptr() *STIXIngestResponseType {
	return &v
}
