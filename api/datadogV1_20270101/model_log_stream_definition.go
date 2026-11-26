// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogStreamDefinition Test-only generated schema for the dashboard versioning proof of concept.
type LogStreamDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Columns []string `json:"columns,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CrossOrgUuids []string `json:"cross_org_uuids,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Indexes []string `json:"indexes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	MessageDisplay *MessageDisplay `json:"message_display,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query *string `json:"query,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowDateColumn *bool `json:"show_date_column,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowMessageColumn *bool `json:"show_message_column,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *SortPair `json:"sort,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Storage *StorageType `json:"storage,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *WidgetTime `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type LogStreamDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewLogStreamDefinition instantiates a new LogStreamDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogStreamDefinition(typeVar LogStreamDefinitionType) *LogStreamDefinition {
	this := LogStreamDefinition{}
	this.Type = typeVar
	return &this
}

// NewLogStreamDefinitionWithDefaults instantiates a new LogStreamDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogStreamDefinitionWithDefaults() *LogStreamDefinition {
	this := LogStreamDefinition{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *LogStreamDefinition) GetColumns() []string {
	if o == nil || o.Columns == nil {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamDefinition) GetColumnsOk() (*[]string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *LogStreamDefinition) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *LogStreamDefinition) SetColumns(v []string) {
	o.Columns = v
}

// GetCrossOrgUuids returns the CrossOrgUuids field value if set, zero value otherwise.
func (o *LogStreamDefinition) GetCrossOrgUuids() []string {
	if o == nil || o.CrossOrgUuids == nil {
		var ret []string
		return ret
	}
	return o.CrossOrgUuids
}

// GetCrossOrgUuidsOk returns a tuple with the CrossOrgUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamDefinition) GetCrossOrgUuidsOk() (*[]string, bool) {
	if o == nil || o.CrossOrgUuids == nil {
		return nil, false
	}
	return &o.CrossOrgUuids, true
}

// HasCrossOrgUuids returns a boolean if a field has been set.
func (o *LogStreamDefinition) HasCrossOrgUuids() bool {
	return o != nil && o.CrossOrgUuids != nil
}

// SetCrossOrgUuids gets a reference to the given []string and assigns it to the CrossOrgUuids field.
func (o *LogStreamDefinition) SetCrossOrgUuids(v []string) {
	o.CrossOrgUuids = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LogStreamDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LogStreamDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LogStreamDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *LogStreamDefinition) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamDefinition) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *LogStreamDefinition) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *LogStreamDefinition) SetIndexes(v []string) {
	o.Indexes = v
}

// GetMessageDisplay returns the MessageDisplay field value if set, zero value otherwise.
func (o *LogStreamDefinition) GetMessageDisplay() MessageDisplay {
	if o == nil || o.MessageDisplay == nil {
		var ret MessageDisplay
		return ret
	}
	return *o.MessageDisplay
}

// GetMessageDisplayOk returns a tuple with the MessageDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamDefinition) GetMessageDisplayOk() (*MessageDisplay, bool) {
	if o == nil || o.MessageDisplay == nil {
		return nil, false
	}
	return o.MessageDisplay, true
}

// HasMessageDisplay returns a boolean if a field has been set.
func (o *LogStreamDefinition) HasMessageDisplay() bool {
	return o != nil && o.MessageDisplay != nil
}

// SetMessageDisplay gets a reference to the given MessageDisplay and assigns it to the MessageDisplay field.
func (o *LogStreamDefinition) SetMessageDisplay(v MessageDisplay) {
	o.MessageDisplay = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *LogStreamDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *LogStreamDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *LogStreamDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *LogStreamDefinition) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamDefinition) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *LogStreamDefinition) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *LogStreamDefinition) SetQuery(v string) {
	o.Query = &v
}

// GetShowDateColumn returns the ShowDateColumn field value if set, zero value otherwise.
func (o *LogStreamDefinition) GetShowDateColumn() bool {
	if o == nil || o.ShowDateColumn == nil {
		var ret bool
		return ret
	}
	return *o.ShowDateColumn
}

// GetShowDateColumnOk returns a tuple with the ShowDateColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamDefinition) GetShowDateColumnOk() (*bool, bool) {
	if o == nil || o.ShowDateColumn == nil {
		return nil, false
	}
	return o.ShowDateColumn, true
}

// HasShowDateColumn returns a boolean if a field has been set.
func (o *LogStreamDefinition) HasShowDateColumn() bool {
	return o != nil && o.ShowDateColumn != nil
}

// SetShowDateColumn gets a reference to the given bool and assigns it to the ShowDateColumn field.
func (o *LogStreamDefinition) SetShowDateColumn(v bool) {
	o.ShowDateColumn = &v
}

// GetShowMessageColumn returns the ShowMessageColumn field value if set, zero value otherwise.
func (o *LogStreamDefinition) GetShowMessageColumn() bool {
	if o == nil || o.ShowMessageColumn == nil {
		var ret bool
		return ret
	}
	return *o.ShowMessageColumn
}

// GetShowMessageColumnOk returns a tuple with the ShowMessageColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamDefinition) GetShowMessageColumnOk() (*bool, bool) {
	if o == nil || o.ShowMessageColumn == nil {
		return nil, false
	}
	return o.ShowMessageColumn, true
}

// HasShowMessageColumn returns a boolean if a field has been set.
func (o *LogStreamDefinition) HasShowMessageColumn() bool {
	return o != nil && o.ShowMessageColumn != nil
}

// SetShowMessageColumn gets a reference to the given bool and assigns it to the ShowMessageColumn field.
func (o *LogStreamDefinition) SetShowMessageColumn(v bool) {
	o.ShowMessageColumn = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *LogStreamDefinition) GetSort() SortPair {
	if o == nil || o.Sort == nil {
		var ret SortPair
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamDefinition) GetSortOk() (*SortPair, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *LogStreamDefinition) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given SortPair and assigns it to the Sort field.
func (o *LogStreamDefinition) SetSort(v SortPair) {
	o.Sort = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *LogStreamDefinition) GetStorage() StorageType {
	if o == nil || o.Storage == nil {
		var ret StorageType
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamDefinition) GetStorageOk() (*StorageType, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *LogStreamDefinition) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given StorageType and assigns it to the Storage field.
func (o *LogStreamDefinition) SetStorage(v StorageType) {
	o.Storage = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *LogStreamDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *LogStreamDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *LogStreamDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *LogStreamDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *LogStreamDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *LogStreamDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *LogStreamDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *LogStreamDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *LogStreamDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *LogStreamDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *LogStreamDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *LogStreamDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *LogStreamDefinition) GetType() LogStreamDefinitionType {
	if o == nil {
		var ret LogStreamDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogStreamDefinition) GetTypeOk() (*LogStreamDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogStreamDefinition) SetType(v LogStreamDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogStreamDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.CrossOrgUuids != nil {
		toSerialize["cross_org_uuids"] = o.CrossOrgUuids
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	if o.MessageDisplay != nil {
		toSerialize["message_display"] = o.MessageDisplay
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.ShowDateColumn != nil {
		toSerialize["show_date_column"] = o.ShowDateColumn
	}
	if o.ShowMessageColumn != nil {
		toSerialize["show_message_column"] = o.ShowMessageColumn
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.TitleAlign != nil {
		toSerialize["title_align"] = o.TitleAlign
	}
	if o.TitleSize != nil {
		toSerialize["title_size"] = o.TitleSize
	}
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogStreamDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns           []string                 `json:"columns,omitempty"`
		CrossOrgUuids     []string                 `json:"cross_org_uuids,omitempty"`
		Description       *string                  `json:"description,omitempty"`
		Indexes           []string                 `json:"indexes,omitempty"`
		MessageDisplay    *MessageDisplay          `json:"message_display,omitempty"`
		Notes             *string                  `json:"notes,omitempty"`
		Query             *string                  `json:"query,omitempty"`
		ShowDateColumn    *bool                    `json:"show_date_column,omitempty"`
		ShowMessageColumn *bool                    `json:"show_message_column,omitempty"`
		Sort              *SortPair                `json:"sort,omitempty"`
		Storage           *StorageType             `json:"storage,omitempty"`
		Time              *WidgetTime              `json:"time,omitempty"`
		Title             *string                  `json:"title,omitempty"`
		TitleAlign        *WidgetTextAlign         `json:"title_align,omitempty"`
		TitleSize         *string                  `json:"title_size,omitempty"`
		Type              *LogStreamDefinitionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Columns = all.Columns
	o.CrossOrgUuids = all.CrossOrgUuids
	o.Description = all.Description
	o.Indexes = all.Indexes
	if all.MessageDisplay != nil && !all.MessageDisplay.IsValid() {
		hasInvalidField = true
	} else {
		o.MessageDisplay = all.MessageDisplay
	}
	o.Notes = all.Notes
	o.Query = all.Query
	o.ShowDateColumn = all.ShowDateColumn
	o.ShowMessageColumn = all.ShowMessageColumn
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort
	if all.Storage != nil && !all.Storage.IsValid() {
		hasInvalidField = true
	} else {
		o.Storage = all.Storage
	}
	o.Time = all.Time
	o.Title = all.Title
	if all.TitleAlign != nil && !all.TitleAlign.IsValid() {
		hasInvalidField = true
	} else {
		o.TitleAlign = all.TitleAlign
	}
	o.TitleSize = all.TitleSize
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
