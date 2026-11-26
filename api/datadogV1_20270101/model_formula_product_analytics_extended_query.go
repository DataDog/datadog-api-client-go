// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

// FormulaProductAnalyticsExtendedQuery Test-only generated schema for the dashboard versioning proof of concept.
type FormulaProductAnalyticsExtendedQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	AudienceFilters *AudienceFilters `json:"audience_filters,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute ProductAnalyticsCompute `json:"compute"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaProductAnalyticsExtendedQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []ProductAnalyticsGroupBy `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Indexes *datadogV1.FormulaProductAnalyticsExtendedQuery20270101Indexes `json:"indexes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query BaseProductAnalyticsQuery `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaProductAnalyticsExtendedQuery instantiates a new FormulaProductAnalyticsExtendedQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaProductAnalyticsExtendedQuery(compute ProductAnalyticsCompute, dataSource FormulaProductAnalyticsExtendedQueryDataSource, name string, query BaseProductAnalyticsQuery) *FormulaProductAnalyticsExtendedQuery {
	this := FormulaProductAnalyticsExtendedQuery{}
	this.Compute = compute
	this.DataSource = dataSource
	this.Name = name
	this.Query = query
	return &this
}

// NewFormulaProductAnalyticsExtendedQueryWithDefaults instantiates a new FormulaProductAnalyticsExtendedQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaProductAnalyticsExtendedQueryWithDefaults() *FormulaProductAnalyticsExtendedQuery {
	this := FormulaProductAnalyticsExtendedQuery{}
	return &this
}

// GetAudienceFilters returns the AudienceFilters field value if set, zero value otherwise.
func (o *FormulaProductAnalyticsExtendedQuery) GetAudienceFilters() AudienceFilters {
	if o == nil || o.AudienceFilters == nil {
		var ret AudienceFilters
		return ret
	}
	return *o.AudienceFilters
}

// GetAudienceFiltersOk returns a tuple with the AudienceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaProductAnalyticsExtendedQuery) GetAudienceFiltersOk() (*AudienceFilters, bool) {
	if o == nil || o.AudienceFilters == nil {
		return nil, false
	}
	return o.AudienceFilters, true
}

// HasAudienceFilters returns a boolean if a field has been set.
func (o *FormulaProductAnalyticsExtendedQuery) HasAudienceFilters() bool {
	return o != nil && o.AudienceFilters != nil
}

// SetAudienceFilters gets a reference to the given AudienceFilters and assigns it to the AudienceFilters field.
func (o *FormulaProductAnalyticsExtendedQuery) SetAudienceFilters(v AudienceFilters) {
	o.AudienceFilters = &v
}

// GetCompute returns the Compute field value.
func (o *FormulaProductAnalyticsExtendedQuery) GetCompute() ProductAnalyticsCompute {
	if o == nil {
		var ret ProductAnalyticsCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *FormulaProductAnalyticsExtendedQuery) GetComputeOk() (*ProductAnalyticsCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *FormulaProductAnalyticsExtendedQuery) SetCompute(v ProductAnalyticsCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaProductAnalyticsExtendedQuery) GetDataSource() FormulaProductAnalyticsExtendedQueryDataSource {
	if o == nil {
		var ret FormulaProductAnalyticsExtendedQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaProductAnalyticsExtendedQuery) GetDataSourceOk() (*FormulaProductAnalyticsExtendedQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaProductAnalyticsExtendedQuery) SetDataSource(v FormulaProductAnalyticsExtendedQueryDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *FormulaProductAnalyticsExtendedQuery) GetGroupBy() []ProductAnalyticsGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []ProductAnalyticsGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaProductAnalyticsExtendedQuery) GetGroupByOk() (*[]ProductAnalyticsGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *FormulaProductAnalyticsExtendedQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []ProductAnalyticsGroupBy and assigns it to the GroupBy field.
func (o *FormulaProductAnalyticsExtendedQuery) SetGroupBy(v []ProductAnalyticsGroupBy) {
	o.GroupBy = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *FormulaProductAnalyticsExtendedQuery) GetIndexes() datadogV1.FormulaProductAnalyticsExtendedQuery20270101Indexes {
	if o == nil || o.Indexes == nil {
		var ret datadogV1.FormulaProductAnalyticsExtendedQuery20270101Indexes
		return ret
	}
	return *o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaProductAnalyticsExtendedQuery) GetIndexesOk() (*datadogV1.FormulaProductAnalyticsExtendedQuery20270101Indexes, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *FormulaProductAnalyticsExtendedQuery) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given datadogV1.FormulaProductAnalyticsExtendedQuery20270101Indexes and assigns it to the Indexes field.
func (o *FormulaProductAnalyticsExtendedQuery) SetIndexes(v datadogV1.FormulaProductAnalyticsExtendedQuery20270101Indexes) {
	o.Indexes = &v
}

// GetName returns the Name field value.
func (o *FormulaProductAnalyticsExtendedQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaProductAnalyticsExtendedQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaProductAnalyticsExtendedQuery) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value.
func (o *FormulaProductAnalyticsExtendedQuery) GetQuery() BaseProductAnalyticsQuery {
	if o == nil {
		var ret BaseProductAnalyticsQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *FormulaProductAnalyticsExtendedQuery) GetQueryOk() (*BaseProductAnalyticsQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *FormulaProductAnalyticsExtendedQuery) SetQuery(v BaseProductAnalyticsQuery) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaProductAnalyticsExtendedQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AudienceFilters != nil {
		toSerialize["audience_filters"] = o.AudienceFilters
	}
	toSerialize["compute"] = o.Compute
	toSerialize["data_source"] = o.DataSource
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaProductAnalyticsExtendedQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AudienceFilters *AudienceFilters                                               `json:"audience_filters,omitempty"`
		Compute         *ProductAnalyticsCompute                                       `json:"compute"`
		DataSource      *FormulaProductAnalyticsExtendedQueryDataSource                `json:"data_source"`
		GroupBy         []ProductAnalyticsGroupBy                                      `json:"group_by,omitempty"`
		Indexes         *datadogV1.FormulaProductAnalyticsExtendedQuery20270101Indexes `json:"indexes,omitempty"`
		Name            *string                                                        `json:"name"`
		Query           *BaseProductAnalyticsQuery                                     `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Compute == nil {
		return fmt.Errorf("required field compute missing")
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}

	hasInvalidField := false
	if all.AudienceFilters != nil && all.AudienceFilters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AudienceFilters = all.AudienceFilters
	if all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = *all.Compute
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.GroupBy = all.GroupBy
	o.Indexes = all.Indexes
	o.Name = *all.Name
	o.Query = *all.Query

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
