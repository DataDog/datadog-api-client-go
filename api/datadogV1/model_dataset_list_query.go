// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetListQuery Query that lists the rows of a published dataset (a DDSQL query) without aggregation.
type DatasetListQuery struct {
	// Identifies this as a published-dataset list query.
	DataSource DatasetListQueryDataSourceType `json:"data_source"`
	// ID of the published dataset to query.
	DatasetId string `json:"dataset_id"`
	// Product page that published the dataset queried by a `DatasetListQuery`. `ddsql_query` is the only provider currently supported for host map widgets.
	DatasetProvider PublishedDatasetProvider `json:"dataset_provider"`
	// Filter applied to the dataset's rows, using events-style search syntax.
	Filter *string `json:"filter,omitempty"`
	// Maximum number of rows to return from the dataset query.
	Limit *int64 `json:"limit,omitempty"`
	// Sort configuration for a `DatasetListQuery`.
	Sort *DatasetListQuerySort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatasetListQuery instantiates a new DatasetListQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatasetListQuery(dataSource DatasetListQueryDataSourceType, datasetId string, datasetProvider PublishedDatasetProvider) *DatasetListQuery {
	this := DatasetListQuery{}
	this.DataSource = dataSource
	this.DatasetId = datasetId
	this.DatasetProvider = datasetProvider
	return &this
}

// NewDatasetListQueryWithDefaults instantiates a new DatasetListQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatasetListQueryWithDefaults() *DatasetListQuery {
	this := DatasetListQuery{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *DatasetListQuery) GetDataSource() DatasetListQueryDataSourceType {
	if o == nil {
		var ret DatasetListQueryDataSourceType
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *DatasetListQuery) GetDataSourceOk() (*DatasetListQueryDataSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *DatasetListQuery) SetDataSource(v DatasetListQueryDataSourceType) {
	o.DataSource = v
}

// GetDatasetId returns the DatasetId field value.
func (o *DatasetListQuery) GetDatasetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value
// and a boolean to check if the value has been set.
func (o *DatasetListQuery) GetDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetId, true
}

// SetDatasetId sets field value.
func (o *DatasetListQuery) SetDatasetId(v string) {
	o.DatasetId = v
}

// GetDatasetProvider returns the DatasetProvider field value.
func (o *DatasetListQuery) GetDatasetProvider() PublishedDatasetProvider {
	if o == nil {
		var ret PublishedDatasetProvider
		return ret
	}
	return o.DatasetProvider
}

// GetDatasetProviderOk returns a tuple with the DatasetProvider field value
// and a boolean to check if the value has been set.
func (o *DatasetListQuery) GetDatasetProviderOk() (*PublishedDatasetProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetProvider, true
}

// SetDatasetProvider sets field value.
func (o *DatasetListQuery) SetDatasetProvider(v PublishedDatasetProvider) {
	o.DatasetProvider = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *DatasetListQuery) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetListQuery) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *DatasetListQuery) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *DatasetListQuery) SetFilter(v string) {
	o.Filter = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *DatasetListQuery) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetListQuery) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *DatasetListQuery) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *DatasetListQuery) SetLimit(v int64) {
	o.Limit = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *DatasetListQuery) GetSort() DatasetListQuerySort {
	if o == nil || o.Sort == nil {
		var ret DatasetListQuerySort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetListQuery) GetSortOk() (*DatasetListQuerySort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *DatasetListQuery) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given DatasetListQuerySort and assigns it to the Sort field.
func (o *DatasetListQuery) SetSort(v DatasetListQuerySort) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatasetListQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["dataset_id"] = o.DatasetId
	toSerialize["dataset_provider"] = o.DatasetProvider
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatasetListQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource      *DatasetListQueryDataSourceType `json:"data_source"`
		DatasetId       *string                         `json:"dataset_id"`
		DatasetProvider *PublishedDatasetProvider       `json:"dataset_provider"`
		Filter          *string                         `json:"filter,omitempty"`
		Limit           *int64                          `json:"limit,omitempty"`
		Sort            *DatasetListQuerySort           `json:"sort,omitempty"`
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
	if all.DatasetProvider == nil {
		return fmt.Errorf("required field dataset_provider missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_source", "dataset_id", "dataset_provider", "filter", "limit", "sort"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.DatasetId = *all.DatasetId
	if !all.DatasetProvider.IsValid() {
		hasInvalidField = true
	} else {
		o.DatasetProvider = *all.DatasetProvider
	}
	o.Filter = all.Filter
	o.Limit = all.Limit
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
