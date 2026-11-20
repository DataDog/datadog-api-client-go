// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PublishedAnalysis Test-only generated schema for the dashboard versioning proof of concept.
type PublishedAnalysis struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Columns []DatasetColumn `json:"columns,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource PublishedAnalysisDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DatasetId string `json:"dataset_id"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryFilter *string `json:"query_filter,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TimeWindow *TimeWindow `json:"time_window,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewPublishedAnalysis instantiates a new PublishedAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPublishedAnalysis(dataSource PublishedAnalysisDataSource, datasetId string, name string) *PublishedAnalysis {
	this := PublishedAnalysis{}
	this.DataSource = dataSource
	this.DatasetId = datasetId
	this.Name = name
	return &this
}

// NewPublishedAnalysisWithDefaults instantiates a new PublishedAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPublishedAnalysisWithDefaults() *PublishedAnalysis {
	this := PublishedAnalysis{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *PublishedAnalysis) GetColumns() []DatasetColumn {
	if o == nil || o.Columns == nil {
		var ret []DatasetColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishedAnalysis) GetColumnsOk() (*[]DatasetColumn, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *PublishedAnalysis) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []DatasetColumn and assigns it to the Columns field.
func (o *PublishedAnalysis) SetColumns(v []DatasetColumn) {
	o.Columns = v
}

// GetDataSource returns the DataSource field value.
func (o *PublishedAnalysis) GetDataSource() PublishedAnalysisDataSource {
	if o == nil {
		var ret PublishedAnalysisDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *PublishedAnalysis) GetDataSourceOk() (*PublishedAnalysisDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *PublishedAnalysis) SetDataSource(v PublishedAnalysisDataSource) {
	o.DataSource = v
}

// GetDatasetId returns the DatasetId field value.
func (o *PublishedAnalysis) GetDatasetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value
// and a boolean to check if the value has been set.
func (o *PublishedAnalysis) GetDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetId, true
}

// SetDatasetId sets field value.
func (o *PublishedAnalysis) SetDatasetId(v string) {
	o.DatasetId = v
}

// GetName returns the Name field value.
func (o *PublishedAnalysis) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PublishedAnalysis) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *PublishedAnalysis) SetName(v string) {
	o.Name = v
}

// GetQueryFilter returns the QueryFilter field value if set, zero value otherwise.
func (o *PublishedAnalysis) GetQueryFilter() string {
	if o == nil || o.QueryFilter == nil {
		var ret string
		return ret
	}
	return *o.QueryFilter
}

// GetQueryFilterOk returns a tuple with the QueryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishedAnalysis) GetQueryFilterOk() (*string, bool) {
	if o == nil || o.QueryFilter == nil {
		return nil, false
	}
	return o.QueryFilter, true
}

// HasQueryFilter returns a boolean if a field has been set.
func (o *PublishedAnalysis) HasQueryFilter() bool {
	return o != nil && o.QueryFilter != nil
}

// SetQueryFilter gets a reference to the given string and assigns it to the QueryFilter field.
func (o *PublishedAnalysis) SetQueryFilter(v string) {
	o.QueryFilter = &v
}

// GetTimeWindow returns the TimeWindow field value if set, zero value otherwise.
func (o *PublishedAnalysis) GetTimeWindow() TimeWindow {
	if o == nil || o.TimeWindow == nil {
		var ret TimeWindow
		return ret
	}
	return *o.TimeWindow
}

// GetTimeWindowOk returns a tuple with the TimeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishedAnalysis) GetTimeWindowOk() (*TimeWindow, bool) {
	if o == nil || o.TimeWindow == nil {
		return nil, false
	}
	return o.TimeWindow, true
}

// HasTimeWindow returns a boolean if a field has been set.
func (o *PublishedAnalysis) HasTimeWindow() bool {
	return o != nil && o.TimeWindow != nil
}

// SetTimeWindow gets a reference to the given TimeWindow and assigns it to the TimeWindow field.
func (o *PublishedAnalysis) SetTimeWindow(v TimeWindow) {
	o.TimeWindow = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PublishedAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["dataset_id"] = o.DatasetId
	toSerialize["name"] = o.Name
	if o.QueryFilter != nil {
		toSerialize["query_filter"] = o.QueryFilter
	}
	if o.TimeWindow != nil {
		toSerialize["time_window"] = o.TimeWindow
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PublishedAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns     []DatasetColumn              `json:"columns,omitempty"`
		DataSource  *PublishedAnalysisDataSource `json:"data_source"`
		DatasetId   *string                      `json:"dataset_id"`
		Name        *string                      `json:"name"`
		QueryFilter *string                      `json:"query_filter,omitempty"`
		TimeWindow  *TimeWindow                  `json:"time_window,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.DatasetId == nil {
		return fmt.Errorf("required field dataset_id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}

	hasInvalidField := false
	o.Columns = all.Columns
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.DatasetId = *all.DatasetId
	o.Name = *all.Name
	o.QueryFilter = all.QueryFilter
	if all.TimeWindow != nil && all.TimeWindow.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeWindow = all.TimeWindow

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
