// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CILogItem A CI job log line.
type CILogItem struct {
	// Comma-separated `key:value` tags. A job can have up to 256 tags, including repeated keys.
	Ddtags *string `json:"ddtags,omitempty"`
	// The job event's `resource.id`, sent through the CI Visibility pipeline API.
	JobId string `json:"job_id"`
	// The line number in the job log. Use 0 or 1 for the first line.
	LineNumber *int64 `json:"line_number,omitempty"`
	// The non-empty log line message.
	Message string `json:"message"`
	// The `resource.unique_id` of the pipeline event, which must also match the job event's
	// `resource.pipeline_unique_id`.
	PipelineUniqueId string `json:"pipeline_unique_id"`
	// The provider name sent with the pipeline event. It defaults to `custom` when omitted and, when provided,
	// must be non-empty and cannot contain a comma.
	ProviderName *string `json:"provider_name,omitempty"`
	// The provider-defined section containing this log line, used to display collapsible groups of lines in the CI
	// job log view.
	SectionName *string `json:"section_name,omitempty"`
	// The status of this log line. Any string is accepted. Datadog maps non-empty values to a standard log status.
	// See [status mapping](https://docs.datadoghq.com/logs/log_configuration/processors/log_status_remapper/).
	Status *string `json:"status,omitempty"`
	// The log line time in RFC 3339 format with an explicit timezone. If omitted, the intake time is used. It can
	// be at most 18 hours in the past or 12 hours in the future.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{}         `json:"-"`
	AdditionalProperties map[string]CILogAttributeValue `json:"-"`
}

// NewCILogItem instantiates a new CILogItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCILogItem(jobId string, message string, pipelineUniqueId string) *CILogItem {
	this := CILogItem{}
	this.JobId = jobId
	this.Message = message
	this.PipelineUniqueId = pipelineUniqueId
	return &this
}

// NewCILogItemWithDefaults instantiates a new CILogItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCILogItemWithDefaults() *CILogItem {
	this := CILogItem{}
	return &this
}

// GetDdtags returns the Ddtags field value if set, zero value otherwise.
func (o *CILogItem) GetDdtags() string {
	if o == nil || o.Ddtags == nil {
		var ret string
		return ret
	}
	return *o.Ddtags
}

// GetDdtagsOk returns a tuple with the Ddtags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CILogItem) GetDdtagsOk() (*string, bool) {
	if o == nil || o.Ddtags == nil {
		return nil, false
	}
	return o.Ddtags, true
}

// HasDdtags returns a boolean if a field has been set.
func (o *CILogItem) HasDdtags() bool {
	return o != nil && o.Ddtags != nil
}

// SetDdtags gets a reference to the given string and assigns it to the Ddtags field.
func (o *CILogItem) SetDdtags(v string) {
	o.Ddtags = &v
}

// GetJobId returns the JobId field value.
func (o *CILogItem) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *CILogItem) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value.
func (o *CILogItem) SetJobId(v string) {
	o.JobId = v
}

// GetLineNumber returns the LineNumber field value if set, zero value otherwise.
func (o *CILogItem) GetLineNumber() int64 {
	if o == nil || o.LineNumber == nil {
		var ret int64
		return ret
	}
	return *o.LineNumber
}

// GetLineNumberOk returns a tuple with the LineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CILogItem) GetLineNumberOk() (*int64, bool) {
	if o == nil || o.LineNumber == nil {
		return nil, false
	}
	return o.LineNumber, true
}

// HasLineNumber returns a boolean if a field has been set.
func (o *CILogItem) HasLineNumber() bool {
	return o != nil && o.LineNumber != nil
}

// SetLineNumber gets a reference to the given int64 and assigns it to the LineNumber field.
func (o *CILogItem) SetLineNumber(v int64) {
	o.LineNumber = &v
}

// GetMessage returns the Message field value.
func (o *CILogItem) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CILogItem) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *CILogItem) SetMessage(v string) {
	o.Message = v
}

// GetPipelineUniqueId returns the PipelineUniqueId field value.
func (o *CILogItem) GetPipelineUniqueId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PipelineUniqueId
}

// GetPipelineUniqueIdOk returns a tuple with the PipelineUniqueId field value
// and a boolean to check if the value has been set.
func (o *CILogItem) GetPipelineUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PipelineUniqueId, true
}

// SetPipelineUniqueId sets field value.
func (o *CILogItem) SetPipelineUniqueId(v string) {
	o.PipelineUniqueId = v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *CILogItem) GetProviderName() string {
	if o == nil || o.ProviderName == nil {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CILogItem) GetProviderNameOk() (*string, bool) {
	if o == nil || o.ProviderName == nil {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *CILogItem) HasProviderName() bool {
	return o != nil && o.ProviderName != nil
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *CILogItem) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetSectionName returns the SectionName field value if set, zero value otherwise.
func (o *CILogItem) GetSectionName() string {
	if o == nil || o.SectionName == nil {
		var ret string
		return ret
	}
	return *o.SectionName
}

// GetSectionNameOk returns a tuple with the SectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CILogItem) GetSectionNameOk() (*string, bool) {
	if o == nil || o.SectionName == nil {
		return nil, false
	}
	return o.SectionName, true
}

// HasSectionName returns a boolean if a field has been set.
func (o *CILogItem) HasSectionName() bool {
	return o != nil && o.SectionName != nil
}

// SetSectionName gets a reference to the given string and assigns it to the SectionName field.
func (o *CILogItem) SetSectionName(v string) {
	o.SectionName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CILogItem) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CILogItem) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CILogItem) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CILogItem) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *CILogItem) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CILogItem) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *CILogItem) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *CILogItem) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CILogItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Ddtags != nil {
		toSerialize["ddtags"] = o.Ddtags
	}
	toSerialize["job_id"] = o.JobId
	if o.LineNumber != nil {
		toSerialize["line_number"] = o.LineNumber
	}
	toSerialize["message"] = o.Message
	toSerialize["pipeline_unique_id"] = o.PipelineUniqueId
	if o.ProviderName != nil {
		toSerialize["provider_name"] = o.ProviderName
	}
	if o.SectionName != nil {
		toSerialize["section_name"] = o.SectionName
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Timestamp != nil {
		if o.Timestamp.Nanosecond() == 0 {
			toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CILogItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ddtags           *string    `json:"ddtags,omitempty"`
		JobId            *string    `json:"job_id"`
		LineNumber       *int64     `json:"line_number,omitempty"`
		Message          *string    `json:"message"`
		PipelineUniqueId *string    `json:"pipeline_unique_id"`
		ProviderName     *string    `json:"provider_name,omitempty"`
		SectionName      *string    `json:"section_name,omitempty"`
		Status           *string    `json:"status,omitempty"`
		Timestamp        *time.Time `json:"timestamp,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.JobId == nil {
		return fmt.Errorf("required field job_id missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.PipelineUniqueId == nil {
		return fmt.Errorf("required field pipeline_unique_id missing")
	}
	additionalProperties := make(map[string]CILogAttributeValue)
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ddtags", "job_id", "line_number", "message", "pipeline_unique_id", "provider_name", "section_name", "status", "timestamp"})
	} else {
		return err
	}
	o.Ddtags = all.Ddtags
	o.JobId = *all.JobId
	o.LineNumber = all.LineNumber
	o.Message = *all.Message
	o.PipelineUniqueId = *all.PipelineUniqueId
	o.ProviderName = all.ProviderName
	o.SectionName = all.SectionName
	o.Status = all.Status
	o.Timestamp = all.Timestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
