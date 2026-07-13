// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetReportScheduleResponseAttributes The configuration and derived state of a report schedule for a published dataset.
type DatasetReportScheduleResponseAttributes struct {
	// The identifier of the notebook cell that published the dataset, or `null` if not set.
	CellId datadog.NullableString `json:"cell_id"`
	// The identifier of the dataset, or `null` if not set.
	DatasetId datadog.NullableString `json:"dataset_id"`
	// The description of the report.
	Description string `json:"description"`
	// The maximum number of rows included in the attached CSV file, or `null` if not set.
	FileRowLimit datadog.NullableInt64 `json:"file_row_limit"`
	// The maximum number of rows included inline in the email body, or `null` if not set.
	InlineRowLimit datadog.NullableInt64 `json:"inline_row_limit"`
	// The Unix timestamp, in milliseconds, of the next scheduled delivery, or
	// `null` if none is scheduled.
	NextRecurrence datadog.NullableInt64 `json:"next_recurrence"`
	// The identifier of the notebook containing the dataset cell, or `null` if not set.
	NotebookId datadog.NullableInt64 `json:"notebook_id"`
	// The recipients of the report (email addresses, Slack channel references, or
	// Microsoft Teams channel references).
	Recipients []string `json:"recipients"`
	// The identifier of the widget containing the dataset.
	ResourceId string `json:"resource_id"`
	// The type of resource targeted by a dataset report schedule.
	ResourceType DatasetReportScheduleResourceType `json:"resource_type"`
	// The recurrence rule for the schedule, expressed as an iCalendar `RRULE` string.
	Rrule string `json:"rrule"`
	// Whether the schedule is currently delivering reports (`active`) or paused (`inactive`).
	Status ReportScheduleStatus `json:"status"`
	// The relative timeframe of data included in the report.
	Timeframe string `json:"timeframe"`
	// The IANA time zone identifier the recurrence rule is evaluated in.
	Timezone string `json:"timezone"`
	// The title of the report.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatasetReportScheduleResponseAttributes instantiates a new DatasetReportScheduleResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatasetReportScheduleResponseAttributes(cellId datadog.NullableString, datasetId datadog.NullableString, description string, fileRowLimit datadog.NullableInt64, inlineRowLimit datadog.NullableInt64, nextRecurrence datadog.NullableInt64, notebookId datadog.NullableInt64, recipients []string, resourceId string, resourceType DatasetReportScheduleResourceType, rrule string, status ReportScheduleStatus, timeframe string, timezone string, title string) *DatasetReportScheduleResponseAttributes {
	this := DatasetReportScheduleResponseAttributes{}
	this.CellId = cellId
	this.DatasetId = datasetId
	this.Description = description
	this.FileRowLimit = fileRowLimit
	this.InlineRowLimit = inlineRowLimit
	this.NextRecurrence = nextRecurrence
	this.NotebookId = notebookId
	this.Recipients = recipients
	this.ResourceId = resourceId
	this.ResourceType = resourceType
	this.Rrule = rrule
	this.Status = status
	this.Timeframe = timeframe
	this.Timezone = timezone
	this.Title = title
	return &this
}

// NewDatasetReportScheduleResponseAttributesWithDefaults instantiates a new DatasetReportScheduleResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatasetReportScheduleResponseAttributesWithDefaults() *DatasetReportScheduleResponseAttributes {
	this := DatasetReportScheduleResponseAttributes{}
	return &this
}

// GetCellId returns the CellId field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *DatasetReportScheduleResponseAttributes) GetCellId() string {
	if o == nil || o.CellId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CellId.Get()
}

// GetCellIdOk returns a tuple with the CellId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DatasetReportScheduleResponseAttributes) GetCellIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CellId.Get(), o.CellId.IsSet()
}

// SetCellId sets field value.
func (o *DatasetReportScheduleResponseAttributes) SetCellId(v string) {
	o.CellId.Set(&v)
}

// GetDatasetId returns the DatasetId field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *DatasetReportScheduleResponseAttributes) GetDatasetId() string {
	if o == nil || o.DatasetId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DatasetId.Get()
}

// GetDatasetIdOk returns a tuple with the DatasetId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DatasetReportScheduleResponseAttributes) GetDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatasetId.Get(), o.DatasetId.IsSet()
}

// SetDatasetId sets field value.
func (o *DatasetReportScheduleResponseAttributes) SetDatasetId(v string) {
	o.DatasetId.Set(&v)
}

// GetDescription returns the Description field value.
func (o *DatasetReportScheduleResponseAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DatasetReportScheduleResponseAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *DatasetReportScheduleResponseAttributes) SetDescription(v string) {
	o.Description = v
}

// GetFileRowLimit returns the FileRowLimit field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *DatasetReportScheduleResponseAttributes) GetFileRowLimit() int64 {
	if o == nil || o.FileRowLimit.Get() == nil {
		var ret int64
		return ret
	}
	return *o.FileRowLimit.Get()
}

// GetFileRowLimitOk returns a tuple with the FileRowLimit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DatasetReportScheduleResponseAttributes) GetFileRowLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileRowLimit.Get(), o.FileRowLimit.IsSet()
}

// SetFileRowLimit sets field value.
func (o *DatasetReportScheduleResponseAttributes) SetFileRowLimit(v int64) {
	o.FileRowLimit.Set(&v)
}

// GetInlineRowLimit returns the InlineRowLimit field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *DatasetReportScheduleResponseAttributes) GetInlineRowLimit() int64 {
	if o == nil || o.InlineRowLimit.Get() == nil {
		var ret int64
		return ret
	}
	return *o.InlineRowLimit.Get()
}

// GetInlineRowLimitOk returns a tuple with the InlineRowLimit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DatasetReportScheduleResponseAttributes) GetInlineRowLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InlineRowLimit.Get(), o.InlineRowLimit.IsSet()
}

// SetInlineRowLimit sets field value.
func (o *DatasetReportScheduleResponseAttributes) SetInlineRowLimit(v int64) {
	o.InlineRowLimit.Set(&v)
}

// GetNextRecurrence returns the NextRecurrence field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *DatasetReportScheduleResponseAttributes) GetNextRecurrence() int64 {
	if o == nil || o.NextRecurrence.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NextRecurrence.Get()
}

// GetNextRecurrenceOk returns a tuple with the NextRecurrence field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DatasetReportScheduleResponseAttributes) GetNextRecurrenceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextRecurrence.Get(), o.NextRecurrence.IsSet()
}

// SetNextRecurrence sets field value.
func (o *DatasetReportScheduleResponseAttributes) SetNextRecurrence(v int64) {
	o.NextRecurrence.Set(&v)
}

// GetNotebookId returns the NotebookId field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *DatasetReportScheduleResponseAttributes) GetNotebookId() int64 {
	if o == nil || o.NotebookId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NotebookId.Get()
}

// GetNotebookIdOk returns a tuple with the NotebookId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DatasetReportScheduleResponseAttributes) GetNotebookIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotebookId.Get(), o.NotebookId.IsSet()
}

// SetNotebookId sets field value.
func (o *DatasetReportScheduleResponseAttributes) SetNotebookId(v int64) {
	o.NotebookId.Set(&v)
}

// GetRecipients returns the Recipients field value.
func (o *DatasetReportScheduleResponseAttributes) GetRecipients() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *DatasetReportScheduleResponseAttributes) GetRecipientsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value.
func (o *DatasetReportScheduleResponseAttributes) SetRecipients(v []string) {
	o.Recipients = v
}

// GetResourceId returns the ResourceId field value.
func (o *DatasetReportScheduleResponseAttributes) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *DatasetReportScheduleResponseAttributes) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value.
func (o *DatasetReportScheduleResponseAttributes) SetResourceId(v string) {
	o.ResourceId = v
}

// GetResourceType returns the ResourceType field value.
func (o *DatasetReportScheduleResponseAttributes) GetResourceType() DatasetReportScheduleResourceType {
	if o == nil {
		var ret DatasetReportScheduleResourceType
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *DatasetReportScheduleResponseAttributes) GetResourceTypeOk() (*DatasetReportScheduleResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value.
func (o *DatasetReportScheduleResponseAttributes) SetResourceType(v DatasetReportScheduleResourceType) {
	o.ResourceType = v
}

// GetRrule returns the Rrule field value.
func (o *DatasetReportScheduleResponseAttributes) GetRrule() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Rrule
}

// GetRruleOk returns a tuple with the Rrule field value
// and a boolean to check if the value has been set.
func (o *DatasetReportScheduleResponseAttributes) GetRruleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rrule, true
}

// SetRrule sets field value.
func (o *DatasetReportScheduleResponseAttributes) SetRrule(v string) {
	o.Rrule = v
}

// GetStatus returns the Status field value.
func (o *DatasetReportScheduleResponseAttributes) GetStatus() ReportScheduleStatus {
	if o == nil {
		var ret ReportScheduleStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DatasetReportScheduleResponseAttributes) GetStatusOk() (*ReportScheduleStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *DatasetReportScheduleResponseAttributes) SetStatus(v ReportScheduleStatus) {
	o.Status = v
}

// GetTimeframe returns the Timeframe field value.
func (o *DatasetReportScheduleResponseAttributes) GetTimeframe() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value
// and a boolean to check if the value has been set.
func (o *DatasetReportScheduleResponseAttributes) GetTimeframeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeframe, true
}

// SetTimeframe sets field value.
func (o *DatasetReportScheduleResponseAttributes) SetTimeframe(v string) {
	o.Timeframe = v
}

// GetTimezone returns the Timezone field value.
func (o *DatasetReportScheduleResponseAttributes) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *DatasetReportScheduleResponseAttributes) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value.
func (o *DatasetReportScheduleResponseAttributes) SetTimezone(v string) {
	o.Timezone = v
}

// GetTitle returns the Title field value.
func (o *DatasetReportScheduleResponseAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DatasetReportScheduleResponseAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *DatasetReportScheduleResponseAttributes) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatasetReportScheduleResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cell_id"] = o.CellId.Get()
	toSerialize["dataset_id"] = o.DatasetId.Get()
	toSerialize["description"] = o.Description
	toSerialize["file_row_limit"] = o.FileRowLimit.Get()
	toSerialize["inline_row_limit"] = o.InlineRowLimit.Get()
	toSerialize["next_recurrence"] = o.NextRecurrence.Get()
	toSerialize["notebook_id"] = o.NotebookId.Get()
	toSerialize["recipients"] = o.Recipients
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["resource_type"] = o.ResourceType
	toSerialize["rrule"] = o.Rrule
	toSerialize["status"] = o.Status
	toSerialize["timeframe"] = o.Timeframe
	toSerialize["timezone"] = o.Timezone
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatasetReportScheduleResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CellId         datadog.NullableString             `json:"cell_id"`
		DatasetId      datadog.NullableString             `json:"dataset_id"`
		Description    *string                            `json:"description"`
		FileRowLimit   datadog.NullableInt64              `json:"file_row_limit"`
		InlineRowLimit datadog.NullableInt64              `json:"inline_row_limit"`
		NextRecurrence datadog.NullableInt64              `json:"next_recurrence"`
		NotebookId     datadog.NullableInt64              `json:"notebook_id"`
		Recipients     *[]string                          `json:"recipients"`
		ResourceId     *string                            `json:"resource_id"`
		ResourceType   *DatasetReportScheduleResourceType `json:"resource_type"`
		Rrule          *string                            `json:"rrule"`
		Status         *ReportScheduleStatus              `json:"status"`
		Timeframe      *string                            `json:"timeframe"`
		Timezone       *string                            `json:"timezone"`
		Title          *string                            `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.CellId.IsSet() {
		return fmt.Errorf("required field cell_id missing")
	}
	if !all.DatasetId.IsSet() {
		return fmt.Errorf("required field dataset_id missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if !all.FileRowLimit.IsSet() {
		return fmt.Errorf("required field file_row_limit missing")
	}
	if !all.InlineRowLimit.IsSet() {
		return fmt.Errorf("required field inline_row_limit missing")
	}
	if !all.NextRecurrence.IsSet() {
		return fmt.Errorf("required field next_recurrence missing")
	}
	if !all.NotebookId.IsSet() {
		return fmt.Errorf("required field notebook_id missing")
	}
	if all.Recipients == nil {
		return fmt.Errorf("required field recipients missing")
	}
	if all.ResourceId == nil {
		return fmt.Errorf("required field resource_id missing")
	}
	if all.ResourceType == nil {
		return fmt.Errorf("required field resource_type missing")
	}
	if all.Rrule == nil {
		return fmt.Errorf("required field rrule missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Timeframe == nil {
		return fmt.Errorf("required field timeframe missing")
	}
	if all.Timezone == nil {
		return fmt.Errorf("required field timezone missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cell_id", "dataset_id", "description", "file_row_limit", "inline_row_limit", "next_recurrence", "notebook_id", "recipients", "resource_id", "resource_type", "rrule", "status", "timeframe", "timezone", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CellId = all.CellId
	o.DatasetId = all.DatasetId
	o.Description = *all.Description
	o.FileRowLimit = all.FileRowLimit
	o.InlineRowLimit = all.InlineRowLimit
	o.NextRecurrence = all.NextRecurrence
	o.NotebookId = all.NotebookId
	o.Recipients = *all.Recipients
	o.ResourceId = *all.ResourceId
	if !all.ResourceType.IsValid() {
		hasInvalidField = true
	} else {
		o.ResourceType = *all.ResourceType
	}
	o.Rrule = *all.Rrule
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Timeframe = *all.Timeframe
	o.Timezone = *all.Timezone
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
