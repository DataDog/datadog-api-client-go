// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyQueryDataStreamsOrServiceMap Query to service-based topology data sources like the service map or data streams.
type TopologyQueryDataStreamsOrServiceMap struct {
	// Name of the data source
	DataSource TopologyQueryDataStreamsOrServiceMapDataSource `json:"data_source"`
	// Your environment and primary tag (or * if enabled for your account).
	Filters []string `json:"filters"`
	// A search string for filtering services, used in `data_streams` queries only. When set, this replaces the `service` field
	QueryString *string `json:"query_string,omitempty"`
	// Name of the service
	Service *string `json:"service,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTopologyQueryDataStreamsOrServiceMap instantiates a new TopologyQueryDataStreamsOrServiceMap object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopologyQueryDataStreamsOrServiceMap(dataSource TopologyQueryDataStreamsOrServiceMapDataSource, filters []string) *TopologyQueryDataStreamsOrServiceMap {
	this := TopologyQueryDataStreamsOrServiceMap{}
	this.DataSource = dataSource
	this.Filters = filters
	return &this
}

// NewTopologyQueryDataStreamsOrServiceMapWithDefaults instantiates a new TopologyQueryDataStreamsOrServiceMap object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopologyQueryDataStreamsOrServiceMapWithDefaults() *TopologyQueryDataStreamsOrServiceMap {
	this := TopologyQueryDataStreamsOrServiceMap{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *TopologyQueryDataStreamsOrServiceMap) GetDataSource() TopologyQueryDataStreamsOrServiceMapDataSource {
	if o == nil {
		var ret TopologyQueryDataStreamsOrServiceMapDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TopologyQueryDataStreamsOrServiceMap) GetDataSourceOk() (*TopologyQueryDataStreamsOrServiceMapDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *TopologyQueryDataStreamsOrServiceMap) SetDataSource(v TopologyQueryDataStreamsOrServiceMapDataSource) {
	o.DataSource = v
}

// GetFilters returns the Filters field value.
func (o *TopologyQueryDataStreamsOrServiceMap) GetFilters() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *TopologyQueryDataStreamsOrServiceMap) GetFiltersOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value.
func (o *TopologyQueryDataStreamsOrServiceMap) SetFilters(v []string) {
	o.Filters = v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *TopologyQueryDataStreamsOrServiceMap) GetQueryString() string {
	if o == nil || o.QueryString == nil {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyQueryDataStreamsOrServiceMap) GetQueryStringOk() (*string, bool) {
	if o == nil || o.QueryString == nil {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *TopologyQueryDataStreamsOrServiceMap) HasQueryString() bool {
	return o != nil && o.QueryString != nil
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *TopologyQueryDataStreamsOrServiceMap) SetQueryString(v string) {
	o.QueryString = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *TopologyQueryDataStreamsOrServiceMap) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyQueryDataStreamsOrServiceMap) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *TopologyQueryDataStreamsOrServiceMap) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *TopologyQueryDataStreamsOrServiceMap) SetService(v string) {
	o.Service = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TopologyQueryDataStreamsOrServiceMap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["filters"] = o.Filters
	if o.QueryString != nil {
		toSerialize["query_string"] = o.QueryString
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TopologyQueryDataStreamsOrServiceMap) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource  *TopologyQueryDataStreamsOrServiceMapDataSource `json:"data_source"`
		Filters     *[]string                                       `json:"filters"`
		QueryString *string                                         `json:"query_string,omitempty"`
		Service     *string                                         `json:"service,omitempty"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_source", "filters", "query_string", "service"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Filters = *all.Filters
	o.QueryString = all.QueryString
	o.Service = all.Service

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
