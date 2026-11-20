// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyDataset Test-only generated schema for the dashboard versioning proof of concept.
type ProductAnalyticsJourneyDataset struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	AudienceFilters *AudienceFilters `json:"audience_filters,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Columns []DatasetColumn `json:"columns"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConversionTimeMs datadog.NullableFloat64 `json:"conversion_time_ms,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource ProductAnalyticsJourneyDatasetDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Entity ProductAnalyticsJourneyDatasetEntity `json:"entity"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Expression string `json:"expression"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NodeObjects map[string]ProductAnalyticsEventQuery `json:"node_objects"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Search *ProductAnalyticsJourneyDatasetSearch `json:"search,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TimeWindow *TimeWindow `json:"time_window,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneyDataset instantiates a new ProductAnalyticsJourneyDataset object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneyDataset(columns []DatasetColumn, dataSource ProductAnalyticsJourneyDatasetDataSource, entity ProductAnalyticsJourneyDatasetEntity, expression string, name string, nodeObjects map[string]ProductAnalyticsEventQuery) *ProductAnalyticsJourneyDataset {
	this := ProductAnalyticsJourneyDataset{}
	this.Columns = columns
	this.DataSource = dataSource
	this.Entity = entity
	this.Expression = expression
	this.Name = name
	this.NodeObjects = nodeObjects
	return &this
}

// NewProductAnalyticsJourneyDatasetWithDefaults instantiates a new ProductAnalyticsJourneyDataset object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneyDatasetWithDefaults() *ProductAnalyticsJourneyDataset {
	this := ProductAnalyticsJourneyDataset{}
	return &this
}

// GetAudienceFilters returns the AudienceFilters field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyDataset) GetAudienceFilters() AudienceFilters {
	if o == nil || o.AudienceFilters == nil {
		var ret AudienceFilters
		return ret
	}
	return *o.AudienceFilters
}

// GetAudienceFiltersOk returns a tuple with the AudienceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDataset) GetAudienceFiltersOk() (*AudienceFilters, bool) {
	if o == nil || o.AudienceFilters == nil {
		return nil, false
	}
	return o.AudienceFilters, true
}

// HasAudienceFilters returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyDataset) HasAudienceFilters() bool {
	return o != nil && o.AudienceFilters != nil
}

// SetAudienceFilters gets a reference to the given AudienceFilters and assigns it to the AudienceFilters field.
func (o *ProductAnalyticsJourneyDataset) SetAudienceFilters(v AudienceFilters) {
	o.AudienceFilters = &v
}

// GetColumns returns the Columns field value.
func (o *ProductAnalyticsJourneyDataset) GetColumns() []DatasetColumn {
	if o == nil {
		var ret []DatasetColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDataset) GetColumnsOk() (*[]DatasetColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *ProductAnalyticsJourneyDataset) SetColumns(v []DatasetColumn) {
	o.Columns = v
}

// GetConversionTimeMs returns the ConversionTimeMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsJourneyDataset) GetConversionTimeMs() float64 {
	if o == nil || o.ConversionTimeMs.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ConversionTimeMs.Get()
}

// GetConversionTimeMsOk returns a tuple with the ConversionTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsJourneyDataset) GetConversionTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConversionTimeMs.Get(), o.ConversionTimeMs.IsSet()
}

// HasConversionTimeMs returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyDataset) HasConversionTimeMs() bool {
	return o != nil && o.ConversionTimeMs.IsSet()
}

// SetConversionTimeMs gets a reference to the given datadog.NullableFloat64 and assigns it to the ConversionTimeMs field.
func (o *ProductAnalyticsJourneyDataset) SetConversionTimeMs(v float64) {
	o.ConversionTimeMs.Set(&v)
}

// SetConversionTimeMsNil sets the value for ConversionTimeMs to be an explicit nil.
func (o *ProductAnalyticsJourneyDataset) SetConversionTimeMsNil() {
	o.ConversionTimeMs.Set(nil)
}

// UnsetConversionTimeMs ensures that no value is present for ConversionTimeMs, not even an explicit nil.
func (o *ProductAnalyticsJourneyDataset) UnsetConversionTimeMs() {
	o.ConversionTimeMs.Unset()
}

// GetDataSource returns the DataSource field value.
func (o *ProductAnalyticsJourneyDataset) GetDataSource() ProductAnalyticsJourneyDatasetDataSource {
	if o == nil {
		var ret ProductAnalyticsJourneyDatasetDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDataset) GetDataSourceOk() (*ProductAnalyticsJourneyDatasetDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *ProductAnalyticsJourneyDataset) SetDataSource(v ProductAnalyticsJourneyDatasetDataSource) {
	o.DataSource = v
}

// GetEntity returns the Entity field value.
func (o *ProductAnalyticsJourneyDataset) GetEntity() ProductAnalyticsJourneyDatasetEntity {
	if o == nil {
		var ret ProductAnalyticsJourneyDatasetEntity
		return ret
	}
	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDataset) GetEntityOk() (*ProductAnalyticsJourneyDatasetEntity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entity, true
}

// SetEntity sets field value.
func (o *ProductAnalyticsJourneyDataset) SetEntity(v ProductAnalyticsJourneyDatasetEntity) {
	o.Entity = v
}

// GetExpression returns the Expression field value.
func (o *ProductAnalyticsJourneyDataset) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDataset) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value.
func (o *ProductAnalyticsJourneyDataset) SetExpression(v string) {
	o.Expression = v
}

// GetName returns the Name field value.
func (o *ProductAnalyticsJourneyDataset) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDataset) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ProductAnalyticsJourneyDataset) SetName(v string) {
	o.Name = v
}

// GetNodeObjects returns the NodeObjects field value.
func (o *ProductAnalyticsJourneyDataset) GetNodeObjects() map[string]ProductAnalyticsEventQuery {
	if o == nil {
		var ret map[string]ProductAnalyticsEventQuery
		return ret
	}
	return o.NodeObjects
}

// GetNodeObjectsOk returns a tuple with the NodeObjects field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDataset) GetNodeObjectsOk() (*map[string]ProductAnalyticsEventQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeObjects, true
}

// SetNodeObjects sets field value.
func (o *ProductAnalyticsJourneyDataset) SetNodeObjects(v map[string]ProductAnalyticsEventQuery) {
	o.NodeObjects = v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyDataset) GetSearch() ProductAnalyticsJourneyDatasetSearch {
	if o == nil || o.Search == nil {
		var ret ProductAnalyticsJourneyDatasetSearch
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDataset) GetSearchOk() (*ProductAnalyticsJourneyDatasetSearch, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyDataset) HasSearch() bool {
	return o != nil && o.Search != nil
}

// SetSearch gets a reference to the given ProductAnalyticsJourneyDatasetSearch and assigns it to the Search field.
func (o *ProductAnalyticsJourneyDataset) SetSearch(v ProductAnalyticsJourneyDatasetSearch) {
	o.Search = &v
}

// GetTimeWindow returns the TimeWindow field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyDataset) GetTimeWindow() TimeWindow {
	if o == nil || o.TimeWindow == nil {
		var ret TimeWindow
		return ret
	}
	return *o.TimeWindow
}

// GetTimeWindowOk returns a tuple with the TimeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDataset) GetTimeWindowOk() (*TimeWindow, bool) {
	if o == nil || o.TimeWindow == nil {
		return nil, false
	}
	return o.TimeWindow, true
}

// HasTimeWindow returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyDataset) HasTimeWindow() bool {
	return o != nil && o.TimeWindow != nil
}

// SetTimeWindow gets a reference to the given TimeWindow and assigns it to the TimeWindow field.
func (o *ProductAnalyticsJourneyDataset) SetTimeWindow(v TimeWindow) {
	o.TimeWindow = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneyDataset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AudienceFilters != nil {
		toSerialize["audience_filters"] = o.AudienceFilters
	}
	toSerialize["columns"] = o.Columns
	if o.ConversionTimeMs.IsSet() {
		toSerialize["conversion_time_ms"] = o.ConversionTimeMs.Get()
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["entity"] = o.Entity
	toSerialize["expression"] = o.Expression
	toSerialize["name"] = o.Name
	toSerialize["node_objects"] = o.NodeObjects
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}
	if o.TimeWindow != nil {
		toSerialize["time_window"] = o.TimeWindow
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJourneyDataset) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AudienceFilters  *AudienceFilters                          `json:"audience_filters,omitempty"`
		Columns          *[]DatasetColumn                          `json:"columns"`
		ConversionTimeMs datadog.NullableFloat64                   `json:"conversion_time_ms,omitempty"`
		DataSource       *ProductAnalyticsJourneyDatasetDataSource `json:"data_source"`
		Entity           *ProductAnalyticsJourneyDatasetEntity     `json:"entity"`
		Expression       *string                                   `json:"expression"`
		Name             *string                                   `json:"name"`
		NodeObjects      *map[string]ProductAnalyticsEventQuery    `json:"node_objects"`
		Search           *ProductAnalyticsJourneyDatasetSearch     `json:"search,omitempty"`
		TimeWindow       *TimeWindow                               `json:"time_window,omitempty"`
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
	if all.Entity == nil {
		return fmt.Errorf("required field entity missing")
	}
	if all.Expression == nil {
		return fmt.Errorf("required field expression missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.NodeObjects == nil {
		return fmt.Errorf("required field node_objects missing")
	}

	hasInvalidField := false
	if all.AudienceFilters != nil && all.AudienceFilters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AudienceFilters = all.AudienceFilters
	o.Columns = *all.Columns
	o.ConversionTimeMs = all.ConversionTimeMs
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	if !all.Entity.IsValid() {
		hasInvalidField = true
	} else {
		o.Entity = *all.Entity
	}
	o.Expression = *all.Expression
	o.Name = *all.Name
	o.NodeObjects = *all.NodeObjects
	if all.Search != nil && all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = all.Search
	if all.TimeWindow != nil && all.TimeWindow.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeWindow = all.TimeWindow

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
