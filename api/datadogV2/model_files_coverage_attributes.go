// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FilesCoverageAttributes Attributes of the per-file code coverage response.
type FilesCoverageAttributes struct {
	// The SHA of the base commit used for comparison (for example, the merge base for a PR).
	BaseCommitSha *string `json:"base_commit_sha,omitempty"`
	// Unix timestamp (milliseconds) of the coverage event.
	EventTimestamp *int64 `json:"event_timestamp,omitempty"`
	// Map of file paths to per-file coverage line data.
	Files map[string]FileCoverageLines `json:"files,omitempty"`
	// The SHA of the head commit for which coverage was evaluated.
	HeadCommitSha *string `json:"head_commit_sha,omitempty"`
	// Number of coverage reports evaluated.
	ReportCount *int64 `json:"report_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFilesCoverageAttributes instantiates a new FilesCoverageAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFilesCoverageAttributes() *FilesCoverageAttributes {
	this := FilesCoverageAttributes{}
	return &this
}

// NewFilesCoverageAttributesWithDefaults instantiates a new FilesCoverageAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFilesCoverageAttributesWithDefaults() *FilesCoverageAttributes {
	this := FilesCoverageAttributes{}
	return &this
}

// GetBaseCommitSha returns the BaseCommitSha field value if set, zero value otherwise.
func (o *FilesCoverageAttributes) GetBaseCommitSha() string {
	if o == nil || o.BaseCommitSha == nil {
		var ret string
		return ret
	}
	return *o.BaseCommitSha
}

// GetBaseCommitShaOk returns a tuple with the BaseCommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesCoverageAttributes) GetBaseCommitShaOk() (*string, bool) {
	if o == nil || o.BaseCommitSha == nil {
		return nil, false
	}
	return o.BaseCommitSha, true
}

// HasBaseCommitSha returns a boolean if a field has been set.
func (o *FilesCoverageAttributes) HasBaseCommitSha() bool {
	return o != nil && o.BaseCommitSha != nil
}

// SetBaseCommitSha gets a reference to the given string and assigns it to the BaseCommitSha field.
func (o *FilesCoverageAttributes) SetBaseCommitSha(v string) {
	o.BaseCommitSha = &v
}

// GetEventTimestamp returns the EventTimestamp field value if set, zero value otherwise.
func (o *FilesCoverageAttributes) GetEventTimestamp() int64 {
	if o == nil || o.EventTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.EventTimestamp
}

// GetEventTimestampOk returns a tuple with the EventTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesCoverageAttributes) GetEventTimestampOk() (*int64, bool) {
	if o == nil || o.EventTimestamp == nil {
		return nil, false
	}
	return o.EventTimestamp, true
}

// HasEventTimestamp returns a boolean if a field has been set.
func (o *FilesCoverageAttributes) HasEventTimestamp() bool {
	return o != nil && o.EventTimestamp != nil
}

// SetEventTimestamp gets a reference to the given int64 and assigns it to the EventTimestamp field.
func (o *FilesCoverageAttributes) SetEventTimestamp(v int64) {
	o.EventTimestamp = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *FilesCoverageAttributes) GetFiles() map[string]FileCoverageLines {
	if o == nil || o.Files == nil {
		var ret map[string]FileCoverageLines
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesCoverageAttributes) GetFilesOk() (*map[string]FileCoverageLines, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return &o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *FilesCoverageAttributes) HasFiles() bool {
	return o != nil && o.Files != nil
}

// SetFiles gets a reference to the given map[string]FileCoverageLines and assigns it to the Files field.
func (o *FilesCoverageAttributes) SetFiles(v map[string]FileCoverageLines) {
	o.Files = v
}

// GetHeadCommitSha returns the HeadCommitSha field value if set, zero value otherwise.
func (o *FilesCoverageAttributes) GetHeadCommitSha() string {
	if o == nil || o.HeadCommitSha == nil {
		var ret string
		return ret
	}
	return *o.HeadCommitSha
}

// GetHeadCommitShaOk returns a tuple with the HeadCommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesCoverageAttributes) GetHeadCommitShaOk() (*string, bool) {
	if o == nil || o.HeadCommitSha == nil {
		return nil, false
	}
	return o.HeadCommitSha, true
}

// HasHeadCommitSha returns a boolean if a field has been set.
func (o *FilesCoverageAttributes) HasHeadCommitSha() bool {
	return o != nil && o.HeadCommitSha != nil
}

// SetHeadCommitSha gets a reference to the given string and assigns it to the HeadCommitSha field.
func (o *FilesCoverageAttributes) SetHeadCommitSha(v string) {
	o.HeadCommitSha = &v
}

// GetReportCount returns the ReportCount field value if set, zero value otherwise.
func (o *FilesCoverageAttributes) GetReportCount() int64 {
	if o == nil || o.ReportCount == nil {
		var ret int64
		return ret
	}
	return *o.ReportCount
}

// GetReportCountOk returns a tuple with the ReportCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesCoverageAttributes) GetReportCountOk() (*int64, bool) {
	if o == nil || o.ReportCount == nil {
		return nil, false
	}
	return o.ReportCount, true
}

// HasReportCount returns a boolean if a field has been set.
func (o *FilesCoverageAttributes) HasReportCount() bool {
	return o != nil && o.ReportCount != nil
}

// SetReportCount gets a reference to the given int64 and assigns it to the ReportCount field.
func (o *FilesCoverageAttributes) SetReportCount(v int64) {
	o.ReportCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FilesCoverageAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BaseCommitSha != nil {
		toSerialize["base_commit_sha"] = o.BaseCommitSha
	}
	if o.EventTimestamp != nil {
		toSerialize["event_timestamp"] = o.EventTimestamp
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.HeadCommitSha != nil {
		toSerialize["head_commit_sha"] = o.HeadCommitSha
	}
	if o.ReportCount != nil {
		toSerialize["report_count"] = o.ReportCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FilesCoverageAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BaseCommitSha  *string                      `json:"base_commit_sha,omitempty"`
		EventTimestamp *int64                       `json:"event_timestamp,omitempty"`
		Files          map[string]FileCoverageLines `json:"files,omitempty"`
		HeadCommitSha  *string                      `json:"head_commit_sha,omitempty"`
		ReportCount    *int64                       `json:"report_count,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"base_commit_sha", "event_timestamp", "files", "head_commit_sha", "report_count"})
	} else {
		return err
	}
	o.BaseCommitSha = all.BaseCommitSha
	o.EventTimestamp = all.EventTimestamp
	o.Files = all.Files
	o.HeadCommitSha = all.HeadCommitSha
	o.ReportCount = all.ReportCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
