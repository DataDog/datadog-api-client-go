// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyQueryServiceMap Query to the service map topology data source.
type TopologyQueryServiceMap struct {
	// Name of the data source.
	DataSource TopologyQueryServiceMapDataSource `json:"data_source"`
	// Your environment and primary tag (or * if enabled for your account).
	Filters []string `json:"filters"`
	// A search string for filtering services. When set, this replaces the `service` field.
	QueryString *string `json:"query_string,omitempty"`
	// (deprecated) Name of the service. Leave this empty and use query_string instead.
	Service string `json:"service"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTopologyQueryServiceMap instantiates a new TopologyQueryServiceMap object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopologyQueryServiceMap(dataSource TopologyQueryServiceMapDataSource, filters []string, service string) *TopologyQueryServiceMap {
	this := TopologyQueryServiceMap{}
	this.DataSource = dataSource
	this.Filters = filters
	this.Service = service
	return &this
}

// NewTopologyQueryServiceMapWithDefaults instantiates a new TopologyQueryServiceMap object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopologyQueryServiceMapWithDefaults() *TopologyQueryServiceMap {
	this := TopologyQueryServiceMap{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *TopologyQueryServiceMap) GetDataSource() TopologyQueryServiceMapDataSource {
	if o == nil {
		var ret TopologyQueryServiceMapDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TopologyQueryServiceMap) GetDataSourceOk() (*TopologyQueryServiceMapDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *TopologyQueryServiceMap) SetDataSource(v TopologyQueryServiceMapDataSource) {
	o.DataSource = v
}

// GetFilters returns the Filters field value.
func (o *TopologyQueryServiceMap) GetFilters() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *TopologyQueryServiceMap) GetFiltersOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value.
func (o *TopologyQueryServiceMap) SetFilters(v []string) {
	o.Filters = v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *TopologyQueryServiceMap) GetQueryString() string {
	if o == nil || o.QueryString == nil {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyQueryServiceMap) GetQueryStringOk() (*string, bool) {
	if o == nil || o.QueryString == nil {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *TopologyQueryServiceMap) HasQueryString() bool {
	return o != nil && o.QueryString != nil
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *TopologyQueryServiceMap) SetQueryString(v string) {
	o.QueryString = &v
}

// GetService returns the Service field value.
func (o *TopologyQueryServiceMap) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *TopologyQueryServiceMap) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *TopologyQueryServiceMap) SetService(v string) {
	o.Service = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TopologyQueryServiceMap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["filters"] = o.Filters
	if o.QueryString != nil {
		toSerialize["query_string"] = o.QueryString
	}
	toSerialize["service"] = o.Service
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TopologyQueryServiceMap) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource  *TopologyQueryServiceMapDataSource `json:"data_source"`
		Filters     *[]string                          `json:"filters"`
		QueryString *string                            `json:"query_string,omitempty"`
		Service     *string                            `json:"service"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Filters == nil {
		return fmt.Errorf("required field filters missing")
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}

	hasInvalidField := false
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Filters = *all.Filters
	o.QueryString = all.QueryString
	o.Service = *all.Service

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
