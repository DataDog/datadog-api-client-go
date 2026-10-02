// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure Datadog source and query that supply values for the metric.
type ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure struct {
	// Name of the Datadog event field used by this measure.
	ColumnName *string `json:"column_name,omitempty"`
	// Data type of the source column.
	ColumnType *string `json:"column_type,omitempty"`
	// Conditions used to select the metric's source data.
	Filters interface{} `json:"filters,omitempty"`
	// Display name of the Datadog measure.
	Name *string `json:"name,omitempty"`
	// Query used to retrieve the Datadog measure.
	Query *string `json:"query,omitempty"`
	// Filter applied to the Datadog source definition.
	SourceDefinitionFilter interface{} `json:"source_definition_filter,omitempty"`
	// Subtype of the Datadog data source.
	SourceSubtype *string `json:"source_subtype,omitempty"`
	// Type of Datadog data source used for the measure.
	SourceType *string `json:"source_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure instantiates a new ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure() *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure {
	this := ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure{}
	return &this
}

// NewExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasureWithDefaults instantiates a new ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasureWithDefaults() *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure {
	this := ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure{}
	return &this
}

// GetColumnName returns the ColumnName field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) GetColumnName() string {
	if o == nil || o.ColumnName == nil {
		var ret string
		return ret
	}
	return *o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) GetColumnNameOk() (*string, bool) {
	if o == nil || o.ColumnName == nil {
		return nil, false
	}
	return o.ColumnName, true
}

// HasColumnName returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) HasColumnName() bool {
	return o != nil && o.ColumnName != nil
}

// SetColumnName gets a reference to the given string and assigns it to the ColumnName field.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) SetColumnName(v string) {
	o.ColumnName = &v
}

// GetColumnType returns the ColumnType field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) GetColumnType() string {
	if o == nil || o.ColumnType == nil {
		var ret string
		return ret
	}
	return *o.ColumnType
}

// GetColumnTypeOk returns a tuple with the ColumnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) GetColumnTypeOk() (*string, bool) {
	if o == nil || o.ColumnType == nil {
		return nil, false
	}
	return o.ColumnType, true
}

// HasColumnType returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) HasColumnType() bool {
	return o != nil && o.ColumnType != nil
}

// SetColumnType gets a reference to the given string and assigns it to the ColumnType field.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) SetColumnType(v string) {
	o.ColumnType = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) GetFilters() interface{} {
	if o == nil || o.Filters == nil {
		var ret interface{}
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) GetFiltersOk() (*interface{}, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return &o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) HasFilters() bool {
	return o != nil && o.Filters != nil
}

// SetFilters gets a reference to the given interface{} and assigns it to the Filters field.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) SetFilters(v interface{}) {
	o.Filters = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) SetQuery(v string) {
	o.Query = &v
}

// GetSourceDefinitionFilter returns the SourceDefinitionFilter field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) GetSourceDefinitionFilter() interface{} {
	if o == nil || o.SourceDefinitionFilter == nil {
		var ret interface{}
		return ret
	}
	return o.SourceDefinitionFilter
}

// GetSourceDefinitionFilterOk returns a tuple with the SourceDefinitionFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) GetSourceDefinitionFilterOk() (*interface{}, bool) {
	if o == nil || o.SourceDefinitionFilter == nil {
		return nil, false
	}
	return &o.SourceDefinitionFilter, true
}

// HasSourceDefinitionFilter returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) HasSourceDefinitionFilter() bool {
	return o != nil && o.SourceDefinitionFilter != nil
}

// SetSourceDefinitionFilter gets a reference to the given interface{} and assigns it to the SourceDefinitionFilter field.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) SetSourceDefinitionFilter(v interface{}) {
	o.SourceDefinitionFilter = v
}

// GetSourceSubtype returns the SourceSubtype field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) GetSourceSubtype() string {
	if o == nil || o.SourceSubtype == nil {
		var ret string
		return ret
	}
	return *o.SourceSubtype
}

// GetSourceSubtypeOk returns a tuple with the SourceSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) GetSourceSubtypeOk() (*string, bool) {
	if o == nil || o.SourceSubtype == nil {
		return nil, false
	}
	return o.SourceSubtype, true
}

// HasSourceSubtype returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) HasSourceSubtype() bool {
	return o != nil && o.SourceSubtype != nil
}

// SetSourceSubtype gets a reference to the given string and assigns it to the SourceSubtype field.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) SetSourceSubtype(v string) {
	o.SourceSubtype = &v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) GetSourceType() string {
	if o == nil || o.SourceType == nil {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) GetSourceTypeOk() (*string, bool) {
	if o == nil || o.SourceType == nil {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) HasSourceType() bool {
	return o != nil && o.SourceType != nil
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) SetSourceType(v string) {
	o.SourceType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ColumnName != nil {
		toSerialize["column_name"] = o.ColumnName
	}
	if o.ColumnType != nil {
		toSerialize["column_type"] = o.ColumnType
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.SourceDefinitionFilter != nil {
		toSerialize["source_definition_filter"] = o.SourceDefinitionFilter
	}
	if o.SourceSubtype != nil {
		toSerialize["source_subtype"] = o.SourceSubtype
	}
	if o.SourceType != nil {
		toSerialize["source_type"] = o.SourceType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsMetricV2DTODataAttributesPercentileAggregationDatadogMetricMeasure) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColumnName             *string     `json:"column_name,omitempty"`
		ColumnType             *string     `json:"column_type,omitempty"`
		Filters                interface{} `json:"filters,omitempty"`
		Name                   *string     `json:"name,omitempty"`
		Query                  *string     `json:"query,omitempty"`
		SourceDefinitionFilter interface{} `json:"source_definition_filter,omitempty"`
		SourceSubtype          *string     `json:"source_subtype,omitempty"`
		SourceType             *string     `json:"source_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column_name", "column_type", "filters", "name", "query", "source_definition_filter", "source_subtype", "source_type"})
	} else {
		return err
	}
	o.ColumnName = all.ColumnName
	o.ColumnType = all.ColumnType
	o.Filters = all.Filters
	o.Name = all.Name
	o.Query = all.Query
	o.SourceDefinitionFilter = all.SourceDefinitionFilter
	o.SourceSubtype = all.SourceSubtype
	o.SourceType = all.SourceType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
