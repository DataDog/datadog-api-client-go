// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineClientIdentityMode What happens to a client that no rule allows. `enforce` rejects the TLS handshake. `audit` accepts the client and reports that it would have been rejected.
type ObservabilityPipelineClientIdentityMode string

// List of ObservabilityPipelineClientIdentityMode.
const (
	OBSERVABILITYPIPELINECLIENTIDENTITYMODE_ENFORCE ObservabilityPipelineClientIdentityMode = "enforce"
	OBSERVABILITYPIPELINECLIENTIDENTITYMODE_AUDIT   ObservabilityPipelineClientIdentityMode = "audit"
)

var allowedObservabilityPipelineClientIdentityModeEnumValues = []ObservabilityPipelineClientIdentityMode{
	OBSERVABILITYPIPELINECLIENTIDENTITYMODE_ENFORCE,
	OBSERVABILITYPIPELINECLIENTIDENTITYMODE_AUDIT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineClientIdentityMode) GetAllowedValues() []ObservabilityPipelineClientIdentityMode {
	return allowedObservabilityPipelineClientIdentityModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineClientIdentityMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineClientIdentityMode(value)
	return nil
}

// NewObservabilityPipelineClientIdentityModeFromValue returns a pointer to a valid ObservabilityPipelineClientIdentityMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineClientIdentityModeFromValue(v string) (*ObservabilityPipelineClientIdentityMode, error) {
	ev := ObservabilityPipelineClientIdentityMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineClientIdentityMode: valid values are %v", v, allowedObservabilityPipelineClientIdentityModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineClientIdentityMode) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineClientIdentityModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineClientIdentityMode value.
func (v ObservabilityPipelineClientIdentityMode) Ptr() *ObservabilityPipelineClientIdentityMode {
	return &v
}
