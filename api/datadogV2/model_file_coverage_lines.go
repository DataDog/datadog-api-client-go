// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FileCoverageLines Per-file line coverage data including executable, covered, and added lines.
type FileCoverageLines struct {
	// Line numbers that were added in the specified scope (for example, in a PR diff).
	AddedLines []int64 `json:"added_lines,omitempty"`
	// Line numbers that were covered by tests.
	CoveredLines []int64 `json:"covered_lines,omitempty"`
	// Line numbers that are executable (can be covered).
	ExecutableLines []int64 `json:"executable_lines,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFileCoverageLines instantiates a new FileCoverageLines object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFileCoverageLines() *FileCoverageLines {
	this := FileCoverageLines{}
	return &this
}

// NewFileCoverageLinesWithDefaults instantiates a new FileCoverageLines object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFileCoverageLinesWithDefaults() *FileCoverageLines {
	this := FileCoverageLines{}
	return &this
}

// GetAddedLines returns the AddedLines field value if set, zero value otherwise.
func (o *FileCoverageLines) GetAddedLines() []int64 {
	if o == nil || o.AddedLines == nil {
		var ret []int64
		return ret
	}
	return o.AddedLines
}

// GetAddedLinesOk returns a tuple with the AddedLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCoverageLines) GetAddedLinesOk() (*[]int64, bool) {
	if o == nil || o.AddedLines == nil {
		return nil, false
	}
	return &o.AddedLines, true
}

// HasAddedLines returns a boolean if a field has been set.
func (o *FileCoverageLines) HasAddedLines() bool {
	return o != nil && o.AddedLines != nil
}

// SetAddedLines gets a reference to the given []int64 and assigns it to the AddedLines field.
func (o *FileCoverageLines) SetAddedLines(v []int64) {
	o.AddedLines = v
}

// GetCoveredLines returns the CoveredLines field value if set, zero value otherwise.
func (o *FileCoverageLines) GetCoveredLines() []int64 {
	if o == nil || o.CoveredLines == nil {
		var ret []int64
		return ret
	}
	return o.CoveredLines
}

// GetCoveredLinesOk returns a tuple with the CoveredLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCoverageLines) GetCoveredLinesOk() (*[]int64, bool) {
	if o == nil || o.CoveredLines == nil {
		return nil, false
	}
	return &o.CoveredLines, true
}

// HasCoveredLines returns a boolean if a field has been set.
func (o *FileCoverageLines) HasCoveredLines() bool {
	return o != nil && o.CoveredLines != nil
}

// SetCoveredLines gets a reference to the given []int64 and assigns it to the CoveredLines field.
func (o *FileCoverageLines) SetCoveredLines(v []int64) {
	o.CoveredLines = v
}

// GetExecutableLines returns the ExecutableLines field value if set, zero value otherwise.
func (o *FileCoverageLines) GetExecutableLines() []int64 {
	if o == nil || o.ExecutableLines == nil {
		var ret []int64
		return ret
	}
	return o.ExecutableLines
}

// GetExecutableLinesOk returns a tuple with the ExecutableLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCoverageLines) GetExecutableLinesOk() (*[]int64, bool) {
	if o == nil || o.ExecutableLines == nil {
		return nil, false
	}
	return &o.ExecutableLines, true
}

// HasExecutableLines returns a boolean if a field has been set.
func (o *FileCoverageLines) HasExecutableLines() bool {
	return o != nil && o.ExecutableLines != nil
}

// SetExecutableLines gets a reference to the given []int64 and assigns it to the ExecutableLines field.
func (o *FileCoverageLines) SetExecutableLines(v []int64) {
	o.ExecutableLines = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FileCoverageLines) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AddedLines != nil {
		toSerialize["added_lines"] = o.AddedLines
	}
	if o.CoveredLines != nil {
		toSerialize["covered_lines"] = o.CoveredLines
	}
	if o.ExecutableLines != nil {
		toSerialize["executable_lines"] = o.ExecutableLines
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FileCoverageLines) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AddedLines      []int64 `json:"added_lines,omitempty"`
		CoveredLines    []int64 `json:"covered_lines,omitempty"`
		ExecutableLines []int64 `json:"executable_lines,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"added_lines", "covered_lines", "executable_lines"})
	} else {
		return err
	}
	o.AddedLines = all.AddedLines
	o.CoveredLines = all.CoveredLines
	o.ExecutableLines = all.ExecutableLines

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
