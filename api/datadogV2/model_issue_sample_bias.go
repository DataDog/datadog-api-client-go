// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueSampleBias Criterion used to pick a sample. `latest` picks the most recent error. `git_metadata` picks the most
// recent error with source code metadata, whose stack trace is parsed. `trace` picks the most recent error
// with an APM trace. `session_replay` picks the most recent error with a RUM session replay.
// `exception_replay` picks the most recent error with an Exception Replay snapshot.
type IssueSampleBias string

// List of IssueSampleBias.
const (
	ISSUESAMPLEBIAS_LATEST           IssueSampleBias = "latest"
	ISSUESAMPLEBIAS_GIT_METADATA     IssueSampleBias = "git_metadata"
	ISSUESAMPLEBIAS_TRACE            IssueSampleBias = "trace"
	ISSUESAMPLEBIAS_SESSION_REPLAY   IssueSampleBias = "session_replay"
	ISSUESAMPLEBIAS_EXCEPTION_REPLAY IssueSampleBias = "exception_replay"
)

var allowedIssueSampleBiasEnumValues = []IssueSampleBias{
	ISSUESAMPLEBIAS_LATEST,
	ISSUESAMPLEBIAS_GIT_METADATA,
	ISSUESAMPLEBIAS_TRACE,
	ISSUESAMPLEBIAS_SESSION_REPLAY,
	ISSUESAMPLEBIAS_EXCEPTION_REPLAY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueSampleBias) GetAllowedValues() []IssueSampleBias {
	return allowedIssueSampleBiasEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueSampleBias) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueSampleBias(value)
	return nil
}

// NewIssueSampleBiasFromValue returns a pointer to a valid IssueSampleBias
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueSampleBiasFromValue(v string) (*IssueSampleBias, error) {
	ev := IssueSampleBias(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueSampleBias: valid values are %v", v, allowedIssueSampleBiasEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueSampleBias) IsValid() bool {
	for _, existing := range allowedIssueSampleBiasEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueSampleBias value.
func (v IssueSampleBias) Ptr() *IssueSampleBias {
	return &v
}
