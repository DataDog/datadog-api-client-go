// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationMetricQuery Metrics query referenced by a formula.
type TimeseriesAnomalyInvestigationMetricQuery struct {
	// Optional scalar aggregator accepted for request compatibility. This field is ignored for timeseries queries.
	Aggregator *string `json:"aggregator,omitempty"`
	// Optional organization UUID used for a cross-organization query. Each query accepts at most one UUID; use separate queries for separate organizations. Influential-tag analysis is currently unsupported for cross-organization queries, but anomaly detection still runs.
	CrossOrgUuids []string `json:"cross_org_uuids,omitempty"`
	// Data source for an anomaly investigation query.
	DataSource TimeseriesAnomalyInvestigationDataSource `json:"data_source"`
	// Name used to reference this query from formulas.
	Name string `json:"name"`
	// Datadog metrics query expression.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationMetricQuery instantiates a new TimeseriesAnomalyInvestigationMetricQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationMetricQuery(dataSource TimeseriesAnomalyInvestigationDataSource, name string, query string) *TimeseriesAnomalyInvestigationMetricQuery {
	this := TimeseriesAnomalyInvestigationMetricQuery{}
	this.DataSource = dataSource
	this.Name = name
	this.Query = query
	return &this
}

// NewTimeseriesAnomalyInvestigationMetricQueryWithDefaults instantiates a new TimeseriesAnomalyInvestigationMetricQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationMetricQueryWithDefaults() *TimeseriesAnomalyInvestigationMetricQuery {
	this := TimeseriesAnomalyInvestigationMetricQuery{}
	return &this
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *TimeseriesAnomalyInvestigationMetricQuery) GetAggregator() string {
	if o == nil || o.Aggregator == nil {
		var ret string
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationMetricQuery) GetAggregatorOk() (*string, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *TimeseriesAnomalyInvestigationMetricQuery) HasAggregator() bool {
	return o != nil && o.Aggregator != nil
}

// SetAggregator gets a reference to the given string and assigns it to the Aggregator field.
func (o *TimeseriesAnomalyInvestigationMetricQuery) SetAggregator(v string) {
	o.Aggregator = &v
}

// GetCrossOrgUuids returns the CrossOrgUuids field value if set, zero value otherwise.
func (o *TimeseriesAnomalyInvestigationMetricQuery) GetCrossOrgUuids() []string {
	if o == nil || o.CrossOrgUuids == nil {
		var ret []string
		return ret
	}
	return o.CrossOrgUuids
}

// GetCrossOrgUuidsOk returns a tuple with the CrossOrgUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationMetricQuery) GetCrossOrgUuidsOk() (*[]string, bool) {
	if o == nil || o.CrossOrgUuids == nil {
		return nil, false
	}
	return &o.CrossOrgUuids, true
}

// HasCrossOrgUuids returns a boolean if a field has been set.
func (o *TimeseriesAnomalyInvestigationMetricQuery) HasCrossOrgUuids() bool {
	return o != nil && o.CrossOrgUuids != nil
}

// SetCrossOrgUuids gets a reference to the given []string and assigns it to the CrossOrgUuids field.
func (o *TimeseriesAnomalyInvestigationMetricQuery) SetCrossOrgUuids(v []string) {
	o.CrossOrgUuids = v
}

// GetDataSource returns the DataSource field value.
func (o *TimeseriesAnomalyInvestigationMetricQuery) GetDataSource() TimeseriesAnomalyInvestigationDataSource {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationMetricQuery) GetDataSourceOk() (*TimeseriesAnomalyInvestigationDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *TimeseriesAnomalyInvestigationMetricQuery) SetDataSource(v TimeseriesAnomalyInvestigationDataSource) {
	o.DataSource = v
}

// GetName returns the Name field value.
func (o *TimeseriesAnomalyInvestigationMetricQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationMetricQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *TimeseriesAnomalyInvestigationMetricQuery) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value.
func (o *TimeseriesAnomalyInvestigationMetricQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationMetricQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *TimeseriesAnomalyInvestigationMetricQuery) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationMetricQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aggregator != nil {
		toSerialize["aggregator"] = o.Aggregator
	}
	if o.CrossOrgUuids != nil {
		toSerialize["cross_org_uuids"] = o.CrossOrgUuids
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesAnomalyInvestigationMetricQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregator    *string                                   `json:"aggregator,omitempty"`
		CrossOrgUuids []string                                  `json:"cross_org_uuids,omitempty"`
		DataSource    *TimeseriesAnomalyInvestigationDataSource `json:"data_source"`
		Name          *string                                   `json:"name"`
		Query         *string                                   `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregator", "cross_org_uuids", "data_source", "name", "query"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Aggregator = all.Aggregator
	o.CrossOrgUuids = all.CrossOrgUuids
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Name = *all.Name
	o.Query = *all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
