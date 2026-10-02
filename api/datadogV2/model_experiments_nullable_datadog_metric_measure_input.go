// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsNullableDatadogMetricMeasureInput Optional Datadog measure. Use null when the other measure is selected.
type ExperimentsNullableDatadogMetricMeasureInput struct {
	// Name of the Datadog event field used by this measure.
	ColumnName *string `json:"column_name,omitempty"`
	// Data type of the source column.
	ColumnType string `json:"column_type"`
	// Conditions used to select the metric's source data.
	Filters interface{} `json:"filters,omitempty"`
	// Display name of the Datadog measure.
	Name string `json:"name"`
	// Query used to retrieve the Datadog measure.
	Query *string `json:"query,omitempty"`
	// Filter applied to the Datadog source definition.
	SourceDefinitionFilter interface{} `json:"source_definition_filter,omitempty"`
	// Subtype of the Datadog data source.
	SourceSubtype string `json:"source_subtype"`
	// Type of Datadog data source used for the measure.
	SourceType string `json:"source_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsNullableDatadogMetricMeasureInput instantiates a new ExperimentsNullableDatadogMetricMeasureInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsNullableDatadogMetricMeasureInput(columnType string, name string, sourceSubtype string, sourceType string) *ExperimentsNullableDatadogMetricMeasureInput {
	this := ExperimentsNullableDatadogMetricMeasureInput{}
	this.ColumnType = columnType
	this.Name = name
	this.SourceSubtype = sourceSubtype
	this.SourceType = sourceType
	return &this
}

// NewExperimentsNullableDatadogMetricMeasureInputWithDefaults instantiates a new ExperimentsNullableDatadogMetricMeasureInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsNullableDatadogMetricMeasureInputWithDefaults() *ExperimentsNullableDatadogMetricMeasureInput {
	this := ExperimentsNullableDatadogMetricMeasureInput{}
	return &this
}

// GetColumnName returns the ColumnName field value if set, zero value otherwise.
func (o *ExperimentsNullableDatadogMetricMeasureInput) GetColumnName() string {
	if o == nil || o.ColumnName == nil {
		var ret string
		return ret
	}
	return *o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsNullableDatadogMetricMeasureInput) GetColumnNameOk() (*string, bool) {
	if o == nil || o.ColumnName == nil {
		return nil, false
	}
	return o.ColumnName, true
}

// HasColumnName returns a boolean if a field has been set.
func (o *ExperimentsNullableDatadogMetricMeasureInput) HasColumnName() bool {
	return o != nil && o.ColumnName != nil
}

// SetColumnName gets a reference to the given string and assigns it to the ColumnName field.
func (o *ExperimentsNullableDatadogMetricMeasureInput) SetColumnName(v string) {
	o.ColumnName = &v
}

// GetColumnType returns the ColumnType field value.
func (o *ExperimentsNullableDatadogMetricMeasureInput) GetColumnType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ColumnType
}

// GetColumnTypeOk returns a tuple with the ColumnType field value
// and a boolean to check if the value has been set.
func (o *ExperimentsNullableDatadogMetricMeasureInput) GetColumnTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnType, true
}

// SetColumnType sets field value.
func (o *ExperimentsNullableDatadogMetricMeasureInput) SetColumnType(v string) {
	o.ColumnType = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ExperimentsNullableDatadogMetricMeasureInput) GetFilters() interface{} {
	if o == nil || o.Filters == nil {
		var ret interface{}
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsNullableDatadogMetricMeasureInput) GetFiltersOk() (*interface{}, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return &o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ExperimentsNullableDatadogMetricMeasureInput) HasFilters() bool {
	return o != nil && o.Filters != nil
}

// SetFilters gets a reference to the given interface{} and assigns it to the Filters field.
func (o *ExperimentsNullableDatadogMetricMeasureInput) SetFilters(v interface{}) {
	o.Filters = v
}

// GetName returns the Name field value.
func (o *ExperimentsNullableDatadogMetricMeasureInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExperimentsNullableDatadogMetricMeasureInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExperimentsNullableDatadogMetricMeasureInput) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ExperimentsNullableDatadogMetricMeasureInput) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsNullableDatadogMetricMeasureInput) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ExperimentsNullableDatadogMetricMeasureInput) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ExperimentsNullableDatadogMetricMeasureInput) SetQuery(v string) {
	o.Query = &v
}

// GetSourceDefinitionFilter returns the SourceDefinitionFilter field value if set, zero value otherwise.
func (o *ExperimentsNullableDatadogMetricMeasureInput) GetSourceDefinitionFilter() interface{} {
	if o == nil || o.SourceDefinitionFilter == nil {
		var ret interface{}
		return ret
	}
	return o.SourceDefinitionFilter
}

// GetSourceDefinitionFilterOk returns a tuple with the SourceDefinitionFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsNullableDatadogMetricMeasureInput) GetSourceDefinitionFilterOk() (*interface{}, bool) {
	if o == nil || o.SourceDefinitionFilter == nil {
		return nil, false
	}
	return &o.SourceDefinitionFilter, true
}

// HasSourceDefinitionFilter returns a boolean if a field has been set.
func (o *ExperimentsNullableDatadogMetricMeasureInput) HasSourceDefinitionFilter() bool {
	return o != nil && o.SourceDefinitionFilter != nil
}

// SetSourceDefinitionFilter gets a reference to the given interface{} and assigns it to the SourceDefinitionFilter field.
func (o *ExperimentsNullableDatadogMetricMeasureInput) SetSourceDefinitionFilter(v interface{}) {
	o.SourceDefinitionFilter = v
}

// GetSourceSubtype returns the SourceSubtype field value.
func (o *ExperimentsNullableDatadogMetricMeasureInput) GetSourceSubtype() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SourceSubtype
}

// GetSourceSubtypeOk returns a tuple with the SourceSubtype field value
// and a boolean to check if the value has been set.
func (o *ExperimentsNullableDatadogMetricMeasureInput) GetSourceSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceSubtype, true
}

// SetSourceSubtype sets field value.
func (o *ExperimentsNullableDatadogMetricMeasureInput) SetSourceSubtype(v string) {
	o.SourceSubtype = v
}

// GetSourceType returns the SourceType field value.
func (o *ExperimentsNullableDatadogMetricMeasureInput) GetSourceType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *ExperimentsNullableDatadogMetricMeasureInput) GetSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value.
func (o *ExperimentsNullableDatadogMetricMeasureInput) SetSourceType(v string) {
	o.SourceType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsNullableDatadogMetricMeasureInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ColumnName != nil {
		toSerialize["column_name"] = o.ColumnName
	}
	toSerialize["column_type"] = o.ColumnType
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	toSerialize["name"] = o.Name
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.SourceDefinitionFilter != nil {
		toSerialize["source_definition_filter"] = o.SourceDefinitionFilter
	}
	toSerialize["source_subtype"] = o.SourceSubtype
	toSerialize["source_type"] = o.SourceType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsNullableDatadogMetricMeasureInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColumnName             *string     `json:"column_name,omitempty"`
		ColumnType             *string     `json:"column_type"`
		Filters                interface{} `json:"filters,omitempty"`
		Name                   *string     `json:"name"`
		Query                  *string     `json:"query,omitempty"`
		SourceDefinitionFilter interface{} `json:"source_definition_filter,omitempty"`
		SourceSubtype          *string     `json:"source_subtype"`
		SourceType             *string     `json:"source_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ColumnType == nil {
		return fmt.Errorf("required field column_type missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.SourceSubtype == nil {
		return fmt.Errorf("required field source_subtype missing")
	}
	if all.SourceType == nil {
		return fmt.Errorf("required field source_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column_name", "column_type", "filters", "name", "query", "source_definition_filter", "source_subtype", "source_type"})
	} else {
		return err
	}
	o.ColumnName = all.ColumnName
	o.ColumnType = *all.ColumnType
	o.Filters = all.Filters
	o.Name = *all.Name
	o.Query = all.Query
	o.SourceDefinitionFilter = all.SourceDefinitionFilter
	o.SourceSubtype = *all.SourceSubtype
	o.SourceType = *all.SourceType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableExperimentsNullableDatadogMetricMeasureInput handles when a null is used for ExperimentsNullableDatadogMetricMeasureInput.
type NullableExperimentsNullableDatadogMetricMeasureInput struct {
	value *ExperimentsNullableDatadogMetricMeasureInput
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsNullableDatadogMetricMeasureInput) Get() *ExperimentsNullableDatadogMetricMeasureInput {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsNullableDatadogMetricMeasureInput) Set(val *ExperimentsNullableDatadogMetricMeasureInput) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsNullableDatadogMetricMeasureInput) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableExperimentsNullableDatadogMetricMeasureInput) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsNullableDatadogMetricMeasureInput initializes the struct as if Set has been called.
func NewNullableExperimentsNullableDatadogMetricMeasureInput(val *ExperimentsNullableDatadogMetricMeasureInput) *NullableExperimentsNullableDatadogMetricMeasureInput {
	return &NullableExperimentsNullableDatadogMetricMeasureInput{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsNullableDatadogMetricMeasureInput) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsNullableDatadogMetricMeasureInput) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
