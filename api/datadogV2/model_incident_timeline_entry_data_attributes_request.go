// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTimelineEntryDataAttributesRequest Attributes for creating or updating a timeline entry.
type IncidentTimelineEntryDataAttributesRequest struct {
	// The type of a timeline cell.
	CellType IncidentTimelineCellType `json:"cell_type"`
	// The content of a timeline entry.
	Content IncidentTimelineEntryContent `json:"content"`
	// The display time for this timeline entry.
	DisplayTime *time.Time `json:"display_time,omitempty"`
	// Whether this timeline entry is marked as important.
	Important *bool `json:"important,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTimelineEntryDataAttributesRequest instantiates a new IncidentTimelineEntryDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTimelineEntryDataAttributesRequest(cellType IncidentTimelineCellType, content IncidentTimelineEntryContent) *IncidentTimelineEntryDataAttributesRequest {
	this := IncidentTimelineEntryDataAttributesRequest{}
	this.CellType = cellType
	this.Content = content
	return &this
}

// NewIncidentTimelineEntryDataAttributesRequestWithDefaults instantiates a new IncidentTimelineEntryDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTimelineEntryDataAttributesRequestWithDefaults() *IncidentTimelineEntryDataAttributesRequest {
	this := IncidentTimelineEntryDataAttributesRequest{}
	return &this
}

// GetCellType returns the CellType field value.
func (o *IncidentTimelineEntryDataAttributesRequest) GetCellType() IncidentTimelineCellType {
	if o == nil {
		var ret IncidentTimelineCellType
		return ret
	}
	return o.CellType
}

// GetCellTypeOk returns a tuple with the CellType field value
// and a boolean to check if the value has been set.
func (o *IncidentTimelineEntryDataAttributesRequest) GetCellTypeOk() (*IncidentTimelineCellType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CellType, true
}

// SetCellType sets field value.
func (o *IncidentTimelineEntryDataAttributesRequest) SetCellType(v IncidentTimelineCellType) {
	o.CellType = v
}

// GetContent returns the Content field value.
func (o *IncidentTimelineEntryDataAttributesRequest) GetContent() IncidentTimelineEntryContent {
	if o == nil {
		var ret IncidentTimelineEntryContent
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *IncidentTimelineEntryDataAttributesRequest) GetContentOk() (*IncidentTimelineEntryContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *IncidentTimelineEntryDataAttributesRequest) SetContent(v IncidentTimelineEntryContent) {
	o.Content = v
}

// GetDisplayTime returns the DisplayTime field value if set, zero value otherwise.
func (o *IncidentTimelineEntryDataAttributesRequest) GetDisplayTime() time.Time {
	if o == nil || o.DisplayTime == nil {
		var ret time.Time
		return ret
	}
	return *o.DisplayTime
}

// GetDisplayTimeOk returns a tuple with the DisplayTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTimelineEntryDataAttributesRequest) GetDisplayTimeOk() (*time.Time, bool) {
	if o == nil || o.DisplayTime == nil {
		return nil, false
	}
	return o.DisplayTime, true
}

// HasDisplayTime returns a boolean if a field has been set.
func (o *IncidentTimelineEntryDataAttributesRequest) HasDisplayTime() bool {
	return o != nil && o.DisplayTime != nil
}

// SetDisplayTime gets a reference to the given time.Time and assigns it to the DisplayTime field.
func (o *IncidentTimelineEntryDataAttributesRequest) SetDisplayTime(v time.Time) {
	o.DisplayTime = &v
}

// GetImportant returns the Important field value if set, zero value otherwise.
func (o *IncidentTimelineEntryDataAttributesRequest) GetImportant() bool {
	if o == nil || o.Important == nil {
		var ret bool
		return ret
	}
	return *o.Important
}

// GetImportantOk returns a tuple with the Important field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTimelineEntryDataAttributesRequest) GetImportantOk() (*bool, bool) {
	if o == nil || o.Important == nil {
		return nil, false
	}
	return o.Important, true
}

// HasImportant returns a boolean if a field has been set.
func (o *IncidentTimelineEntryDataAttributesRequest) HasImportant() bool {
	return o != nil && o.Important != nil
}

// SetImportant gets a reference to the given bool and assigns it to the Important field.
func (o *IncidentTimelineEntryDataAttributesRequest) SetImportant(v bool) {
	o.Important = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTimelineEntryDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cell_type"] = o.CellType
	toSerialize["content"] = o.Content
	if o.DisplayTime != nil {
		if o.DisplayTime.Nanosecond() == 0 {
			toSerialize["display_time"] = o.DisplayTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["display_time"] = o.DisplayTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Important != nil {
		toSerialize["important"] = o.Important
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTimelineEntryDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CellType    *IncidentTimelineCellType     `json:"cell_type"`
		Content     *IncidentTimelineEntryContent `json:"content"`
		DisplayTime *time.Time                    `json:"display_time,omitempty"`
		Important   *bool                         `json:"important,omitempty"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cell_type", "content", "display_time", "important"})
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
	o.DisplayTime = all.DisplayTime
	o.Important = all.Important

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
