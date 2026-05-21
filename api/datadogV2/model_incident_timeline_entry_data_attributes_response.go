// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTimelineEntryDataAttributesResponse Attributes of a timeline entry.
type IncidentTimelineEntryDataAttributesResponse struct {
	// The type of a timeline cell.
	CellType IncidentTimelineCellType `json:"cell_type"`
	// The content of a timeline entry.
	Content IncidentTimelineEntryContent `json:"content"`
	// Timestamp when the entry was created.
	Created time.Time `json:"created"`
	// The display time for this timeline entry.
	DisplayTime time.Time `json:"display_time"`
	// Whether this timeline entry is marked as important.
	Important bool `json:"important"`
	// The incident identifier.
	IncidentId string `json:"incident_id"`
	// Timestamp when the entry was last modified.
	Modified time.Time `json:"modified"`
	// UUID of the parent timeline entry.
	ParentUuid datadog.NullableString `json:"parent_uuid,omitempty"`
	// The source of a timeline cell.
	Source IncidentTimelineCellSource `json:"source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTimelineEntryDataAttributesResponse instantiates a new IncidentTimelineEntryDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTimelineEntryDataAttributesResponse(cellType IncidentTimelineCellType, content IncidentTimelineEntryContent, created time.Time, displayTime time.Time, important bool, incidentId string, modified time.Time, source IncidentTimelineCellSource) *IncidentTimelineEntryDataAttributesResponse {
	this := IncidentTimelineEntryDataAttributesResponse{}
	this.CellType = cellType
	this.Content = content
	this.Created = created
	this.DisplayTime = displayTime
	this.Important = important
	this.IncidentId = incidentId
	this.Modified = modified
	this.Source = source
	return &this
}

// NewIncidentTimelineEntryDataAttributesResponseWithDefaults instantiates a new IncidentTimelineEntryDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTimelineEntryDataAttributesResponseWithDefaults() *IncidentTimelineEntryDataAttributesResponse {
	this := IncidentTimelineEntryDataAttributesResponse{}
	return &this
}

// GetCellType returns the CellType field value.
func (o *IncidentTimelineEntryDataAttributesResponse) GetCellType() IncidentTimelineCellType {
	if o == nil {
		var ret IncidentTimelineCellType
		return ret
	}
	return o.CellType
}

// GetCellTypeOk returns a tuple with the CellType field value
// and a boolean to check if the value has been set.
func (o *IncidentTimelineEntryDataAttributesResponse) GetCellTypeOk() (*IncidentTimelineCellType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CellType, true
}

// SetCellType sets field value.
func (o *IncidentTimelineEntryDataAttributesResponse) SetCellType(v IncidentTimelineCellType) {
	o.CellType = v
}

// GetContent returns the Content field value.
func (o *IncidentTimelineEntryDataAttributesResponse) GetContent() IncidentTimelineEntryContent {
	if o == nil {
		var ret IncidentTimelineEntryContent
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *IncidentTimelineEntryDataAttributesResponse) GetContentOk() (*IncidentTimelineEntryContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *IncidentTimelineEntryDataAttributesResponse) SetContent(v IncidentTimelineEntryContent) {
	o.Content = v
}

// GetCreated returns the Created field value.
func (o *IncidentTimelineEntryDataAttributesResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *IncidentTimelineEntryDataAttributesResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value.
func (o *IncidentTimelineEntryDataAttributesResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetDisplayTime returns the DisplayTime field value.
func (o *IncidentTimelineEntryDataAttributesResponse) GetDisplayTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.DisplayTime
}

// GetDisplayTimeOk returns a tuple with the DisplayTime field value
// and a boolean to check if the value has been set.
func (o *IncidentTimelineEntryDataAttributesResponse) GetDisplayTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayTime, true
}

// SetDisplayTime sets field value.
func (o *IncidentTimelineEntryDataAttributesResponse) SetDisplayTime(v time.Time) {
	o.DisplayTime = v
}

// GetImportant returns the Important field value.
func (o *IncidentTimelineEntryDataAttributesResponse) GetImportant() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Important
}

// GetImportantOk returns a tuple with the Important field value
// and a boolean to check if the value has been set.
func (o *IncidentTimelineEntryDataAttributesResponse) GetImportantOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Important, true
}

// SetImportant sets field value.
func (o *IncidentTimelineEntryDataAttributesResponse) SetImportant(v bool) {
	o.Important = v
}

// GetIncidentId returns the IncidentId field value.
func (o *IncidentTimelineEntryDataAttributesResponse) GetIncidentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IncidentId
}

// GetIncidentIdOk returns a tuple with the IncidentId field value
// and a boolean to check if the value has been set.
func (o *IncidentTimelineEntryDataAttributesResponse) GetIncidentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncidentId, true
}

// SetIncidentId sets field value.
func (o *IncidentTimelineEntryDataAttributesResponse) SetIncidentId(v string) {
	o.IncidentId = v
}

// GetModified returns the Modified field value.
func (o *IncidentTimelineEntryDataAttributesResponse) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *IncidentTimelineEntryDataAttributesResponse) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value.
func (o *IncidentTimelineEntryDataAttributesResponse) SetModified(v time.Time) {
	o.Modified = v
}

// GetParentUuid returns the ParentUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentTimelineEntryDataAttributesResponse) GetParentUuid() string {
	if o == nil || o.ParentUuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentUuid.Get()
}

// GetParentUuidOk returns a tuple with the ParentUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentTimelineEntryDataAttributesResponse) GetParentUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentUuid.Get(), o.ParentUuid.IsSet()
}

// HasParentUuid returns a boolean if a field has been set.
func (o *IncidentTimelineEntryDataAttributesResponse) HasParentUuid() bool {
	return o != nil && o.ParentUuid.IsSet()
}

// SetParentUuid gets a reference to the given datadog.NullableString and assigns it to the ParentUuid field.
func (o *IncidentTimelineEntryDataAttributesResponse) SetParentUuid(v string) {
	o.ParentUuid.Set(&v)
}

// SetParentUuidNil sets the value for ParentUuid to be an explicit nil.
func (o *IncidentTimelineEntryDataAttributesResponse) SetParentUuidNil() {
	o.ParentUuid.Set(nil)
}

// UnsetParentUuid ensures that no value is present for ParentUuid, not even an explicit nil.
func (o *IncidentTimelineEntryDataAttributesResponse) UnsetParentUuid() {
	o.ParentUuid.Unset()
}

// GetSource returns the Source field value.
func (o *IncidentTimelineEntryDataAttributesResponse) GetSource() IncidentTimelineCellSource {
	if o == nil {
		var ret IncidentTimelineCellSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *IncidentTimelineEntryDataAttributesResponse) GetSourceOk() (*IncidentTimelineCellSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *IncidentTimelineEntryDataAttributesResponse) SetSource(v IncidentTimelineCellSource) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTimelineEntryDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cell_type"] = o.CellType
	toSerialize["content"] = o.Content
	if o.Created.Nanosecond() == 0 {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.DisplayTime.Nanosecond() == 0 {
		toSerialize["display_time"] = o.DisplayTime.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["display_time"] = o.DisplayTime.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["important"] = o.Important
	toSerialize["incident_id"] = o.IncidentId
	if o.Modified.Nanosecond() == 0 {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ParentUuid.IsSet() {
		toSerialize["parent_uuid"] = o.ParentUuid.Get()
	}
	toSerialize["source"] = o.Source

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTimelineEntryDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CellType    *IncidentTimelineCellType     `json:"cell_type"`
		Content     *IncidentTimelineEntryContent `json:"content"`
		Created     *time.Time                    `json:"created"`
		DisplayTime *time.Time                    `json:"display_time"`
		Important   *bool                         `json:"important"`
		IncidentId  *string                       `json:"incident_id"`
		Modified    *time.Time                    `json:"modified"`
		ParentUuid  datadog.NullableString        `json:"parent_uuid,omitempty"`
		Source      *IncidentTimelineCellSource   `json:"source"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CellType == nil {
		return fmt.Errorf("required field cell_type missing")
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	if all.Created == nil {
		return fmt.Errorf("required field created missing")
	}
	if all.DisplayTime == nil {
		return fmt.Errorf("required field display_time missing")
	}
	if all.Important == nil {
		return fmt.Errorf("required field important missing")
	}
	if all.IncidentId == nil {
		return fmt.Errorf("required field incident_id missing")
	}
	if all.Modified == nil {
		return fmt.Errorf("required field modified missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cell_type", "content", "created", "display_time", "important", "incident_id", "modified", "parent_uuid", "source"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.CellType.IsValid() {
		hasInvalidField = true
	} else {
		o.CellType = *all.CellType
	}
	if all.Content.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Content = *all.Content
	o.Created = *all.Created
	o.DisplayTime = *all.DisplayTime
	o.Important = *all.Important
	o.IncidentId = *all.IncidentId
	o.Modified = *all.Modified
	o.ParentUuid = all.ParentUuid
	if !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = *all.Source
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
