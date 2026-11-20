// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventPlatformDatasetDefinition Test-only generated schema for the dashboard versioning proof of concept.
type EventPlatformDatasetDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Columns []DatasetColumn `json:"columns"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource AdvancedEventPlatformSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Indexes []string `json:"indexes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Search *EventPlatformDatasetDefinitionSearch `json:"search,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Storage *StorageType `json:"storage,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TimeWindow *TimeWindow `json:"time_window,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEventPlatformDatasetDefinition instantiates a new EventPlatformDatasetDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventPlatformDatasetDefinition(columns []DatasetColumn, dataSource AdvancedEventPlatformSource, name string) *EventPlatformDatasetDefinition {
	this := EventPlatformDatasetDefinition{}
	this.Columns = columns
	this.DataSource = dataSource
	this.Name = name
	return &this
}

// NewEventPlatformDatasetDefinitionWithDefaults instantiates a new EventPlatformDatasetDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventPlatformDatasetDefinitionWithDefaults() *EventPlatformDatasetDefinition {
	this := EventPlatformDatasetDefinition{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *EventPlatformDatasetDefinition) GetColumns() []DatasetColumn {
	if o == nil {
		var ret []DatasetColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *EventPlatformDatasetDefinition) GetColumnsOk() (*[]DatasetColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *EventPlatformDatasetDefinition) SetColumns(v []DatasetColumn) {
	o.Columns = v
}

// GetDataSource returns the DataSource field value.
func (o *EventPlatformDatasetDefinition) GetDataSource() AdvancedEventPlatformSource {
	if o == nil {
		var ret AdvancedEventPlatformSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *EventPlatformDatasetDefinition) GetDataSourceOk() (*AdvancedEventPlatformSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *EventPlatformDatasetDefinition) SetDataSource(v AdvancedEventPlatformSource) {
	o.DataSource = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *EventPlatformDatasetDefinition) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventPlatformDatasetDefinition) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *EventPlatformDatasetDefinition) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *EventPlatformDatasetDefinition) SetIndexes(v []string) {
	o.Indexes = v
}

// GetName returns the Name field value.
func (o *EventPlatformDatasetDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventPlatformDatasetDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EventPlatformDatasetDefinition) SetName(v string) {
	o.Name = v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *EventPlatformDatasetDefinition) GetSearch() EventPlatformDatasetDefinitionSearch {
	if o == nil || o.Search == nil {
		var ret EventPlatformDatasetDefinitionSearch
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventPlatformDatasetDefinition) GetSearchOk() (*EventPlatformDatasetDefinitionSearch, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *EventPlatformDatasetDefinition) HasSearch() bool {
	return o != nil && o.Search != nil
}

// SetSearch gets a reference to the given EventPlatformDatasetDefinitionSearch and assigns it to the Search field.
func (o *EventPlatformDatasetDefinition) SetSearch(v EventPlatformDatasetDefinitionSearch) {
	o.Search = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *EventPlatformDatasetDefinition) GetStorage() StorageType {
	if o == nil || o.Storage == nil {
		var ret StorageType
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventPlatformDatasetDefinition) GetStorageOk() (*StorageType, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *EventPlatformDatasetDefinition) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given StorageType and assigns it to the Storage field.
func (o *EventPlatformDatasetDefinition) SetStorage(v StorageType) {
	o.Storage = &v
}

// GetTimeWindow returns the TimeWindow field value if set, zero value otherwise.
func (o *EventPlatformDatasetDefinition) GetTimeWindow() TimeWindow {
	if o == nil || o.TimeWindow == nil {
		var ret TimeWindow
		return ret
	}
	return *o.TimeWindow
}

// GetTimeWindowOk returns a tuple with the TimeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventPlatformDatasetDefinition) GetTimeWindowOk() (*TimeWindow, bool) {
	if o == nil || o.TimeWindow == nil {
		return nil, false
	}
	return o.TimeWindow, true
}

// HasTimeWindow returns a boolean if a field has been set.
func (o *EventPlatformDatasetDefinition) HasTimeWindow() bool {
	return o != nil && o.TimeWindow != nil
}

// SetTimeWindow gets a reference to the given TimeWindow and assigns it to the TimeWindow field.
func (o *EventPlatformDatasetDefinition) SetTimeWindow(v TimeWindow) {
	o.TimeWindow = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventPlatformDatasetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["columns"] = o.Columns
	toSerialize["data_source"] = o.DataSource
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	toSerialize["name"] = o.Name
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.TimeWindow != nil {
		toSerialize["time_window"] = o.TimeWindow
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventPlatformDatasetDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns    *[]DatasetColumn                      `json:"columns"`
		DataSource *AdvancedEventPlatformSource          `json:"data_source"`
		Indexes    []string                              `json:"indexes,omitempty"`
		Name       *string                               `json:"name"`
		Search     *EventPlatformDatasetDefinitionSearch `json:"search,omitempty"`
		Storage    *StorageType                          `json:"storage,omitempty"`
		TimeWindow *TimeWindow                           `json:"time_window,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Columns == nil {
		return fmt.Errorf("required field columns missing")
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}

	hasInvalidField := false
	o.Columns = *all.Columns
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Indexes = all.Indexes
	o.Name = *all.Name
	if all.Search != nil && all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = all.Search
	if all.Storage != nil && !all.Storage.IsValid() {
		hasInvalidField = true
	} else {
		o.Storage = all.Storage
	}
	if all.TimeWindow != nil && all.TimeWindow.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeWindow = all.TimeWindow

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
